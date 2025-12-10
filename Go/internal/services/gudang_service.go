package services

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"app-penjualan/internal/dto"
	"app-penjualan/internal/models"
	"app-penjualan/internal/utils"
)

// GudangService menangani seluruh proses QC, GR, dan retur gudang
type GudangService struct {
	db   *gorm.DB
	prod *ProductService
}

func NewGudangService(db *gorm.DB) *GudangService {
	return &GudangService{db: db, prod: NewProductService(db)}
}

//
// ================== 1️⃣ QC + Good Receipt ==================
//

func (s *GudangService) RecordQcAndReceive(purchaseID uint, req dto.CreateQcReportRequest, userID uint) error {
	var purchase models.Purchase
	if err := s.db.Preload("Items").First(&purchase, purchaseID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("purchase order not found")
		}
		return err
	}

	if purchase.Status != models.PurchaseApproved {
		return errors.New("purchase order is not APPROVED; cannot proceed with QC/Receipt")
	}

	// buat map item untuk validasi
	poItemMap := make(map[uint]struct {
		Qty       int
		Price     float64
		ProductID uint
	})
	for _, item := range purchase.Items {
		poItemMap[item.ID] = struct {
			Qty       int
			Price     float64
			ProductID uint
		}{Qty: item.Qty, Price: item.Price, ProductID: item.ProductID}
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var (
		hasDefects        bool
		totalDefectValue  float64
		defectItems       []models.PurchaseReturnItem
		now               = time.Now()
	)

	for _, reqItem := range req.Items {
		poData, ok := poItemMap[reqItem.PurchaseItemID]
		if !ok {
			tx.Rollback()
			return fmt.Errorf("item %d not found in purchase order", reqItem.PurchaseItemID)
		}

		totalReceived := reqItem.QtyGood + reqItem.QtyDamaged
		if totalReceived > poData.Qty {
			tx.Rollback()
			return fmt.Errorf("received qty (%d) exceeds ordered (%d) for item %d",
				totalReceived, poData.Qty, reqItem.PurchaseItemID)
		}

		// tambah stok barang bagus
		if reqItem.QtyGood > 0 {
			if err := tx.Model(&models.Product{}).
				Where("id = ?", poData.ProductID).
				Update("stock", gorm.Expr("stock + ?", reqItem.QtyGood)).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to update stock for product %d: %w", poData.ProductID, err)
			}
		}

		// barang cacat dicatat ke PurchaseReturn
		if reqItem.QtyDamaged > 0 {
			hasDefects = true
			subtotal := float64(reqItem.QtyDamaged) * poData.Price
			totalDefectValue += subtotal

			defectItems = append(defectItems, models.PurchaseReturnItem{
				ProductID: poData.ProductID,
				Qty:       reqItem.QtyDamaged,
				Price:     poData.Price,
				Subtotal:  subtotal,
			})
		}
	}

	if hasDefects {
		qcReturn := &models.PurchaseReturn{
			PurchaseID: purchaseID,
			UserID:     userID,
			Date:       now,
			Status:     "QC_DEFECT",
			Total:      totalDefectValue,
			Items:      defectItems,
		}
		if err := tx.Create(qcReturn).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to create QC defect return: %w", err)
		}
	}

	// update status PO jadi RECEIVED
	if err := tx.Model(&models.Purchase{}).
		Where("id = ?", purchaseID).
		Updates(map[string]interface{}{
			"status":      models.PurchaseReceived,
			"received_by": userID,
			"received_at": &now,
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

//
// ================== 2️⃣ List Purchase Returns ==================
//

func (s *GudangService) ListApprovedPurchaseReturns(page, size int) ([]models.PurchaseReturn, int64, error) {
	var returns []models.PurchaseReturn
	var total int64

	query := s.db.Model(&models.PurchaseReturn{}).
		Preload("Purchase").
		Preload("Items.Product").
		Where("status IN ?", []string{"APPROVED", "QC_DEFECT"}).
		Order("created_at DESC")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Gunakan pagination dari query.go
	if err := query.Scopes(Paginate(page, size)).Find(&returns).Error; err != nil {
		return nil, 0, err
	}

	return returns, total, nil
}

//
// ================== 3️⃣ List Approved Purchases ==================
//

func (s *GudangService) ListApprovedPurchases(page, size int) ([]models.Purchase, int64, error) {
	var purchases []models.Purchase
	var total int64

	query := s.db.Model(&models.Purchase{}).
		Preload("Supplier").
		Preload("Items.Product").
		Where("status = ?", models.PurchaseApproved).
		Order("created_at DESC")

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Scopes(Paginate(page, size)).Find(&purchases).Error; err != nil {
		return nil, 0, err
	}

	return purchases, total, nil
}

// ListApprovedSaleReturns
// Menampilkan daftar SaleReturn dengan status "APPROVED" yang siap diproses oleh Gudang
func (s *GudangService) ListApprovedSaleReturns(page, size int, dr utils.DateRange) ([]models.SaleReturn, int64, error) {
	var returns []models.SaleReturn
	var total int64

	q := s.db.Model(&models.SaleReturn{}).
		Preload("User").
		Preload("Items.Product").
		Where("status = ?", "APPROVED")

	// Terapkan filter tanggal jika ada
	q = utils.ApplyRange(q, dr, "date")

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := q.Scopes(utils.Paginate(page, size)).Find(&returns).Error; err != nil {
		return nil, 0, err
	}

	return returns, total, nil
}
//
// ================== 4️⃣ QC Sale Return ==================
//

func (s *GudangService) RecordQcSaleReturn(saleReturnID uint, req dto.CreateQcReportRequest, userID uint) error {
	var saleReturn models.SaleReturn
	if err := s.db.Preload("Items").Preload("Items.Product").First(&saleReturn, saleReturnID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("sale return not found")
		}
		return err
	}

	if saleReturn.Status != "APPROVED" {
		return errors.New("sale return is not APPROVED")
	}

	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	now := time.Now()

	for _, reqItem := range req.Items {
		found := false
		for _, sItem := range saleReturn.Items {
			if uint(reqItem.PurchaseItemID) == sItem.ID || uint(reqItem.PurchaseItemID) == sItem.ProductID {
				found = true

				if reqItem.QtyGood+reqItem.QtyDamaged > sItem.Qty {
					tx.Rollback()
					return fmt.Errorf("received qty exceeds returned qty for item %d", sItem.ID)
				}

				if reqItem.QtyGood > 0 {
					if err := tx.Model(&models.Product{}).
						Where("id = ?", sItem.ProductID).
						Update("stock", gorm.Expr("stock + ?", reqItem.QtyGood)).Error; err != nil {
						tx.Rollback()
						return err
					}
				}
			}
		}
		if !found {
			tx.Rollback()
			return fmt.Errorf("no matching sale return item found for id %d", reqItem.PurchaseItemID)
		}
	}

	if err := tx.Model(&models.SaleReturn{}).
		Where("id = ?", saleReturnID).
		Updates(map[string]interface{}{
			"status": "RECEIVED",
			"date":   now,
			"notes":  fmt.Sprintf("Processed by Gudang UID=%d at %s", userID, now.Format(time.RFC3339)),
		}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
