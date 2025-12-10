package services

import (
	"errors"
	"time"

	"app-penjualan/internal/dto"
	"app-penjualan/internal/models"
	"app-penjualan/internal/repositories"
	"gorm.io/gorm"
)

type PurchaseService struct {
	db    *gorm.DB
	prepo *repositories.PurchaseRepo
	prod  *ProductService
}

func NewPurchaseService(db *gorm.DB) *PurchaseService {
	return &PurchaseService{
		db:    db,
		prepo: repositories.NewPurchaseRepo(db),
		prod:  NewProductService(db),
	}
}

// FindAll memuat semua data pembelian beserta relasi (Supplier, User, dan Items dengan Product).
func (s *PurchaseService) FindAll(page, size int, search, status, from, to string) ([]models.Purchase, error) {
	var purchases []models.Purchase
	offset := (page - 1) * size
	query := s.db.Preload("Supplier").Preload("Items.Product").Preload("User")

	if search != "" {
		query = query.Where("id LIKE ? OR supplier_id LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// ✅ Filter tanggal
	if from != "" && to != "" {
		query = query.Where("date BETWEEN ? AND ?", from, to)
	} else if from != "" {
		query = query.Where("date >= ?", from)
	} else if to != "" {
		query = query.Where("date <= ?", to)
	}

	err := query.Offset(offset).Limit(size).Find(&purchases).Error
	if err != nil {
		return nil, err
	}
	return purchases, nil
}

// ✅ Summary rekap pembelian per tanggal
func (s *PurchaseService) Summary(from, to string) ([]map[string]interface{}, error) {
	var results []map[string]interface{}

	query := s.db.Table("purchases").
		Select("DATE(date) as period, COUNT(*) as trx, COALESCE(SUM(total),0) as total").
		Where("total > 0"). // hindari record kosong
		Where("status = ?", models.PurchaseApproved). // hanya approved
		Group("DATE(date)").
		Order("DATE(date) ASC")

	// filter tanggal opsional
	if from != "" && to != "" {
		query = query.Where("date BETWEEN ? AND ?", from, to)
	} else if from != "" {
		query = query.Where("date >= ?", from)
	} else if to != "" {
		query = query.Where("date <= ?", to)
	}

	if err := query.Find(&results).Error; err != nil {
		return nil, err
	}

	return results, nil
}

// Create membuat data pembelian baru.
func (s *PurchaseService) Create(uid uint, req dto.CreatePurchaseReq) (*models.Purchase, error) {
	// Validasi SupplierID
	var supplier models.Supplier
	if err := s.db.First(&supplier, req.SupplierID).Error; err != nil {
		return nil, err
	}

	// Validasi ProductID untuk setiap item
	for _, it := range req.Items {
		var product models.Product
		if err := s.db.First(&product, it.ProductID).Error; err != nil {
			return nil, err
		}
	}

	// Validasi tanggal
	if req.Date.IsZero() {
		return nil, errors.New("invalid date")
	}

	p := &models.Purchase{
		SupplierID: req.SupplierID,
		UserID:     uid,
		Date:       req.Date,
		Status:     models.PurchaseDraft,
	}

	var total float64
	for _, it := range req.Items {
		item := models.PurchaseItem{
			ProductID: it.ProductID,
			Qty:       it.Qty,
			Price:     it.Price,
			Subtotal:  float64(it.Qty) * it.Price,
		}
		p.Items = append(p.Items, item)
		total += item.Subtotal
	}

	p.Total = total
	if req.Submit {
		p.Status = models.PurchaseSubmitted
	}

	if err := s.prepo.Create(p); err != nil {
		return nil, err
	}

	return p, nil
}

// Approve menyetujui atau menolak pembelian.
func (s *PurchaseService) Approve(approverID, purchaseID uint, approve bool) (*models.Purchase, error) {
	po, err := s.prepo.WithItems(purchaseID)
	if err != nil { return nil, err }
	
	if po.Status != models.PurchaseSubmitted {
		return nil, errors.New("purchase can only be approved if status is SUBMITTED")
	}

	if approve {
		now := time.Now()
		po.Status = models.PurchaseApproved
		po.ApprovedBy = &approverID
		po.ApprovedAt = &now
		
		// ❌ HAPUS: Penambahan stok dihapus dari sini!
		// for _, it := range po.Items { _ = s.prod.AdjustStock(it.ProductID, it.Qty) } 
		
	} else {
		po.Status = models.PurchaseRejected
	}

	if err := s.prepo.Update(po); err != nil { return nil, err }

	return po, nil
}

func (s *PurchaseService) Reject(approverID, purchaseID uint) (*models.Purchase, error) {
	po, err := s.prepo.WithItems(purchaseID)
	if err != nil {
		return nil, err
	}

	po.Status = models.PurchaseRejected
	if err := s.prepo.Update(po); err != nil {
		return nil, err
	}
	return po, nil
}

func (s *PurchaseService) GetPurchaseByID(id uint) (*models.Purchase, error) {
	var purchase models.Purchase
	err := s.db.Preload("Supplier").Preload("Items.Product").Preload("User").First(&purchase, id).Error
	if err != nil {
		return nil, err
	}
	return &purchase, nil
} 


// ✅ FUNGSI BARU: Receive (Goods Receipt) - Menambah Stok
func (s *PurchaseService) Receive(receiverID, purchaseID uint) (*models.Purchase, error) {
	po, err := s.prepo.WithItems(purchaseID)
	if err != nil { return nil, err }

	// 1. Validasi Status: Hanya bisa diterima jika statusnya APPROVED
	if po.Status != models.PurchaseApproved {
		return nil, errors.New("purchase can only be received if status is APPROVED")
	}

	// 2. Transaksi Database: Update Status dan Stock
	txErr := s.db.Transaction(func(tx *gorm.DB) error {
		// 2a. Tambah Stok Produk
		for _, item := range po.Items {
			// panggil AdjustStock
			if err := s.prod.AdjustStock(item.ProductID, item.Qty); err != nil {
				return err 
			}
		}

		// 2b. Update Status Purchase menjadi RECEIVED
		po.Status = models.PurchaseReceived
		now := time.Now()
		po.ReceivedBy = &receiverID
		po.ReceivedAt = &now

		if err := tx.Save(po).Error; err != nil { return err }

		return nil
	})

	if txErr != nil { return nil, txErr }

	// Muat ulang Purchase untuk menampilkan data Receiver
	return s.prepo.WithItems(purchaseID)
}