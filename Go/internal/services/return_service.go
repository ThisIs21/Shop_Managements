package services

import (
    "errors"
    "log"
    "time"

    "app-penjualan/internal/models"
    "gorm.io/gorm"
)

type ReturnService struct {
    db *gorm.DB
}

func NewReturnService(db *gorm.DB) *ReturnService {
    return &ReturnService{db: db}
}

/* ====== DTO ====== */
type SaleReturnItemDTO struct {
    ProductID uint   `json:"product_id" binding:"required"`
    Qty       int    `json:"qty" binding:"required,gt=0"`
    Reason    string `json:"reason" binding:"required"`
}

type CreateSaleReturnDTO struct {
    SaleID uint                `json:"sale_id" binding:"required"`
    Date   string               `json:"date"` // Ubah dari time.Time ke string
    Items  []SaleReturnItemDTO `json:"items" binding:"required,dive"`
    Notes  string              `json:"notes"`
}

type PurchaseReturnItemDTO struct {
    ProductID uint    `json:"product_id" binding:"required"`
    Qty       int     `json:"qty" binding:"required,gt=0"`
    Price     float64 `json:"price" binding:"required,gte=0"`
}

type CreatePurchaseReturnDTO struct {
    PurchaseID uint                  `json:"purchase_id" binding:"required"`
    Date       string                 `json:"date"` // Ubah dari time.Time ke string
    Items      []PurchaseReturnItemDTO `json:"items" binding:"required,dive"`
    Notes      string                 `json:"notes"`
}

/* ====== List Returns ====== */
func (s *ReturnService) ListSaleReturns(page, size int, search, status string) ([]models.SaleReturn, int64, error) {
    var returns []models.SaleReturn
    var total int64

    baseQuery := s.db.Model(&models.SaleReturn{}).
        Preload("User").
        Preload("Sale.Customer").
        Preload("Sale.Items.Product").
        Preload("Items.Product").
        Order("created_at DESC")

    if status != "" {
        baseQuery = baseQuery.Where("status = ?", status)
    }

    if search != "" {
        baseQuery = baseQuery.Joins("JOIN sales ON sales.id = sale_returns.sale_id").
            Joins("LEFT JOIN customers ON customers.id = sales.customer_id").
            Joins("JOIN sale_return_items ON sale_return_items.sale_return_id = sale_returns.id").
            Joins("JOIN products ON products.id = sale_return_items.product_id").
            Where("sale_returns.id::text LIKE ? OR customers.name LIKE ? OR products.name LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
    }

    if err := baseQuery.Count(&total).Error; err != nil {
        log.Printf("Gagal menghitung total sale returns: %v", err)
        return nil, 0, err
    }

    baseQuery = baseQuery.Offset((page - 1) * size).Limit(size)
    if err := baseQuery.Find(&returns).Error; err != nil {
        log.Printf("Gagal mengambil sale returns: %v", err)
        return nil, 0, err
    }

    log.Printf("Mengambil %d sale returns dari total %d", len(returns), total)
    return returns, total, nil
}

func (s *ReturnService) ListPurchaseReturns(page, size int, search, status string) ([]models.PurchaseReturn, int64, error) {
	var returns []models.PurchaseReturn
	var total int64

	baseQuery := s.db.Model(&models.PurchaseReturn{}).
		Preload("Approver"). // Approver di sini adalah Kepala Gudang yang MENYETUJUI retur ke Supplier
		Preload("Items.Product").
		Order("created_at DESC")

	// PENTING: Jika status kosong, tampilkan laporan Barang Rusak hasil QC dan PENDING manual.
	if status == "" {
		// QC_DEFECT adalah status yang dicatat oleh GUDANG saat QC/Penerimaan.
		// PENDING adalah retur manual yang diajukan.
		baseQuery = baseQuery.Where("status = ? OR status = ?", "QC_DEFECT", "PENDING")
	} else {
		// Jika status diisi (misal: status=APPROVED), gunakan filter status.
		baseQuery = baseQuery.Where("status = ?", status)
	}

	if search != "" {
		baseQuery = baseQuery.Joins("JOIN purchase_return_items ON purchase_return_items.purchase_return_id = purchase_returns.id").
			Joins("JOIN products ON products.id = purchase_return_items.product_id").
			Where("purchase_returns.id::text LIKE ? OR products.name LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := baseQuery.Count(&total).Error; err != nil {
		log.Printf("Gagal menghitung total purchase returns: %v", err)
		return nil, 0, err
	}

	baseQuery = baseQuery.Offset((page - 1) * size).Limit(size)
	if err := baseQuery.Find(&returns).Error; err != nil {
		log.Printf("Gagal mengambil purchase returns: %v", err)
		return nil, 0, err
	}

	log.Printf("Mengambil %d purchase returns dari total %d", len(returns), total)
	return returns, total, nil
}

func (s *ReturnService) GetSaleReturnByID(id uint, includeUser bool) (*models.SaleReturn, error) {
    var ret models.SaleReturn
    query := s.db.
        Preload("Sale.Customer").
        Preload("Sale.Items.Product").
        Preload("Items.Product")
    if includeUser {
        query = query.Preload("User")
    }
    err := query.First(&ret, id).Error
    if err != nil {
        log.Printf("Gagal mengambil sale return ID %d: %v", id, err)
        return nil, err
    }
    return &ret, nil
}

func (s *ReturnService) GetPurchaseReturnByID(id uint, includeApprover bool) (*models.PurchaseReturn, error) {
    var ret models.PurchaseReturn
    query := s.db.Preload("Items.Product")
    if includeApprover {
        query = query.Preload("Approver")
    }
    err := query.First(&ret, id).Error
    if err != nil {
        log.Printf("Gagal mengambil purchase return ID %d: %v", id, err)
        return nil, err
    }
    return &ret, nil
}

/* ====== Retur Penjualan ====== */
func (s *ReturnService) CreateSaleReturn(userID uint, dto CreateSaleReturnDTO) (*models.SaleReturn, error) {
    if len(dto.Items) == 0 {
        log.Printf("Items retur penjualan kosong")
        return nil, errors.New("items wajib diisi")
    }

    var sale models.Sale
    if err := s.db.First(&sale, dto.SaleID).Error; err != nil {
        log.Printf("Sale ID %d tidak ditemukan: %v", dto.SaleID, err)
        return nil, errors.New("sale tidak ditemukan")
    }

    // Parse date dari string ke time.Time
    var date time.Time
    if dto.Date == "" {
        date = time.Now()
    } else {
        parsedDate, err := time.Parse("2006-01-02", dto.Date)
        if err != nil {
            log.Printf("Gagal parse tanggal %s: %v", dto.Date, err)
            return nil, errors.New("invalid date format, use YYYY-MM-DD")
        }
        date = parsedDate
    }

    ret := &models.SaleReturn{
        SaleID: dto.SaleID,
        UserID: userID,
        Date:   date,
        Status: "PENDING",
        Notes:  dto.Notes,
    }

    var total float64
    for _, it := range dto.Items {
        var product models.Product
        if err := s.db.First(&product, it.ProductID).Error; err != nil {
            log.Printf("Produk ID %d tidak ditemukan: %v", it.ProductID, err)
            return nil, errors.New("produk tidak ditemukan")
        }
        item := models.SaleReturnItem{
            ProductID: it.ProductID,
            Qty:       it.Qty,
            Price:     product.SellPrice,
            Reason:    it.Reason,
            Subtotal:  float64(it.Qty) * product.SellPrice,
        }
        ret.Items = append(ret.Items, item)
        total += item.Subtotal
    }
    ret.Total = total

    if err := s.db.Create(ret).Error; err != nil {
        log.Printf("Gagal menyimpan retur penjualan: %v", err)
        return nil, err
    }
    log.Printf("Retur penjualan ID %d berhasil dibuat", ret.ID)
    return ret, nil
}

func (s *ReturnService) ApproveSaleReturn(retID uint, approverID uint) error {
    return s.db.Transaction(func(tx *gorm.DB) error {
        var ret models.SaleReturn
        if err := tx.Preload("Items").First(&ret, retID).Error; err != nil {
            log.Printf("Retur penjualan ID %d tidak ditemukan: %v", retID, err)
            return errors.New("retur penjualan tidak ditemukan")
        }
        if ret.Status != "PENDING" {
            log.Printf("Retur penjualan ID %d sudah diproses: %s", retID, ret.Status)
            return errors.New("retur penjualan sudah diproses")
        }

        for _, it := range ret.Items {
            var product models.Product
            if err := tx.First(&product, it.ProductID).Error; err != nil {
                log.Printf("Produk ID %d tidak ditemukan: %v", it.ProductID, err)
                return err
            }
            if err := tx.Model(&models.Product{}).
                Where("id = ?", it.ProductID).
                Update("stock", gorm.Expr("stock + ?", it.Qty)).Error; err != nil {
                log.Printf("Gagal memperbarui stok untuk produk ID %d: %v", it.ProductID, err)
                return err
            }
        }

        now := time.Now()
        ret.Status = "APPROVED"
        ret.ApprovedBy = &approverID
        ret.ApprovedAt = &now

        if err := tx.Save(&ret).Error; err != nil {
            // Rollback stok jika save gagal
            for _, it := range ret.Items {
                tx.Model(&models.Product{}).
                    Where("id = ?", it.ProductID).
                    Update("stock", gorm.Expr("stock - ?", it.Qty))
            }
            log.Printf("Gagal menyimpan retur penjualan ID %d: %v", retID, err)
            return err
        }
        log.Printf("Retur penjualan ID %d disetujui", retID)
        return nil
    })
}

func (s *ReturnService) RejectSaleReturn(retID uint, approverID uint, reason string) error {
    return s.db.Transaction(func(tx *gorm.DB) error {
        var ret models.SaleReturn
        if err := tx.First(&ret, retID).Error; err != nil {
            log.Printf("Retur penjualan ID %d tidak ditemukan: %v", retID, err)
            return errors.New("retur penjualan tidak ditemukan")
        }
        if ret.Status != "PENDING" {
            log.Printf("Retur penjualan ID %d sudah diproses: %s", retID, ret.Status)
            return errors.New("retur penjualan sudah diproses")
        }

        now := time.Now()
        ret.Status = "REJECTED" // Pastikan nilai ini sesuai enum
        ret.ApprovedBy = &approverID
        ret.ApprovedAt = &now
        if ret.Notes == "" {
            ret.Notes = "[REJECTED: " + reason + "]"
        } else {
            ret.Notes += "\n\n[REJECTED: " + reason + "]"
        }

        if err := tx.Save(&ret).Error; err != nil {
            log.Printf("Gagal menyimpan retur penjualan ID %d: %v", retID, err)
            return err
        }
        log.Printf("Retur penjualan ID %d ditolak dengan alasan: %s", retID, reason)
        return nil
    })
}

/* ====== Retur Pembelian ====== */
func (s *ReturnService) CreatePurchaseReturn(userID uint, dto CreatePurchaseReturnDTO) (*models.PurchaseReturn, error) {
    if len(dto.Items) == 0 {
        log.Printf("Items retur pembelian kosong")
        return nil, errors.New("items wajib diisi")
    }

    var pur models.Purchase
    if err := s.db.First(&pur, dto.PurchaseID).Error; err != nil {
        log.Printf("Purchase ID %d tidak ditemukan: %v", dto.PurchaseID, err)
        return nil, errors.New("purchase tidak ditemukan")
    }

    // Parse date dari string ke time.Time
    var date time.Time
    if dto.Date == "" {
        date = time.Now()
    } else {
        parsedDate, err := time.Parse("2006-01-02", dto.Date)
        if err != nil {
            log.Printf("Gagal parse tanggal %s: %v", dto.Date, err)
            return nil, errors.New("invalid date format, use YYYY-MM-DD")
        }
        date = parsedDate
    }

    ret := &models.PurchaseReturn{
        PurchaseID: dto.PurchaseID,
        UserID:     userID,
        Date:       date,
        Status:     "PENDING",
    }

    var total float64
    for _, it := range dto.Items {
        item := models.PurchaseReturnItem{
            ProductID: it.ProductID,
            Qty:       it.Qty,
            Price:     it.Price,
            Subtotal:  float64(it.Qty) * it.Price,
        }
        ret.Items = append(ret.Items, item)
        total += item.Subtotal
    }
    ret.Total = total

    if err := s.db.Create(ret).Error; err != nil {
        log.Printf("Gagal menyimpan retur pembelian: %v", err)
        return nil, err
    }
    log.Printf("Retur pembelian ID %d berhasil dibuat", ret.ID)
    return ret, nil
}

func (s *ReturnService) ApprovePurchaseReturn(retID uint, approverID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var ret models.PurchaseReturn
		if err := tx.Preload("Items").First(&ret, retID).Error; err != nil {
			log.Printf("Retur pembelian ID %d tidak ditemukan: %v", retID, err)
			return errors.New("retur pembelian tidak ditemukan")
		}
		
		// Memperbolehkan PENDING (retur manual) atau QC_DEFECT (laporan barang rusak) untuk diproses
		if ret.Status != "PENDING" && ret.Status != "QC_DEFECT" { 
			log.Printf("Retur pembelian ID %d sudah diproses: %s", retID, ret.Status)
			return errors.New("retur pembelian sudah diproses")
		}

		// *** HAPUS TOTAL LOGIKA PENGURANGAN STOK DI SINI ***
		// Karena: Barang Rusak hasil QC tidak pernah ditambahkan ke stok.
		// Jika barang PENDING berasal dari retur manual yang barangnya sudah ada di stok, 
		// Anda mungkin perlu menyesuaikan alur model, tetapi untuk kasus QC_DEFECT, stok tidak disentuh lagi.

		now := time.Now()
		ret.Status = "APPROVED" // Disetujui untuk dikembalikan ke supplier
		ret.ApprovedBy = &approverID
		ret.ApprovedAt = &now

		if err := tx.Save(&ret).Error; err != nil {
			log.Printf("Gagal menyimpan retur pembelian ID %d: %v", retID, err)
			return err
		}
		log.Printf("Retur pembelian ID %d disetujui", retID)
		return nil
	})
}

func (s *ReturnService) RejectPurchaseReturn(retID uint, approverID uint, reason string) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var ret models.PurchaseReturn
		if err := tx.First(&ret, retID).Error; err != nil {
			log.Printf("Retur pembelian ID %d tidak ditemukan: %v", retID, err)
			return errors.New("retur pembelian tidak ditemukan")
		}
		
		// Memperbolehkan PENDING atau QC_DEFECT untuk diproses
		if ret.Status != "PENDING" && ret.Status != "QC_DEFECT" {
			log.Printf("Retur pembelian ID %d sudah diproses: %s", retID, ret.Status)
			return errors.New("retur pembelian sudah diproses")
		}

		// ... (lanjutan sama)
		now := time.Now()
		ret.Status = "REJECTED"
		ret.ApprovedBy = &approverID
		ret.ApprovedAt = &now

		// ... (lanjutan sama)
		if err := tx.Save(&ret).Error; err != nil {
			log.Printf("Gagal menyimpan retur pembelian ID %d: %v", retID, err)
			return err
		}
		log.Printf("Retur pembelian ID %d ditolak dengan alasan: %s", retID, reason)
		return nil
	})
}