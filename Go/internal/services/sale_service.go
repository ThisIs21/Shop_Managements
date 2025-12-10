package services

import (
    "errors"
    "fmt"
    "log"
    "time"

    "app-penjualan/internal/dto"
    "app-penjualan/internal/models"
    "app-penjualan/internal/repositories"
    "gorm.io/gorm"
)

type SaleService struct {
    db    *gorm.DB
    srepo *repositories.SaleRepo
    prod  *ProductService
}

func NewSaleService(db *gorm.DB) *SaleService {
    return &SaleService{db: db, srepo: repositories.NewSaleRepo(db), prod: NewProductService(db)}
}

func (s *SaleService) Create(uid uint, req dto.CreateSaleReq) (*models.Sale, error) {
    var saleDate time.Time
    if req.Date != nil {
        saleDate = *req.Date
    } else {
        saleDate = time.Now()
    }

    sale := &models.Sale{
        UserID: uid,
        Date:   saleDate,
        Audit:  models.Audit{CreatedAt: time.Now()},
    }

    // Generate TransactionNumber
    var count int64
    if err := s.db.Model(&models.Sale{}).Count(&count).Error; err != nil {
        log.Printf("Gagal menghitung jumlah transaksi: %v", err)
        return nil, err
    }
    sale.TransactionNumber = fmt.Sprintf("TRX-%06d", count+1)

    // Customer baru atau existing
    if req.CustomerName != nil && *req.CustomerName != "" {
        newCustomer := models.Customer{Name: *req.CustomerName}
        if err := s.db.Create(&newCustomer).Error; err != nil {
            log.Printf("Gagal membuat pelanggan baru: %v", err)
            return nil, errors.New("gagal membuat pelanggan baru")
        }
        sale.CustomerID = &newCustomer.ID
    } else if req.CustomerID != nil {
        sale.CustomerID = req.CustomerID
    }

    var subtotal float64
    var saleItems []models.SaleItem

    // Transaksi penuh
    err := s.db.Transaction(func(tx *gorm.DB) error {
        // Simpan item dan kurangi stok
        for _, it := range req.Items {
            var product models.Product
            if err := tx.First(&product, it.ProductID).Error; err != nil {
                log.Printf("Produk ID %d tidak ditemukan: %v", it.ProductID, err)
                return errors.New("produk tidak ditemukan")
            }

            itemSubtotal := float64(it.Qty) * product.SellPrice
            item := models.SaleItem{
                ProductID: it.ProductID,
                Qty:       it.Qty,
                Price:     product.SellPrice,
                Subtotal:  itemSubtotal,
            }

            saleItems = append(saleItems, item)
            subtotal += itemSubtotal

            // Kurangi stok
            if err := s.prod.AdjustStock(it.ProductID, -it.Qty); err != nil {
                log.Printf("Gagal mengurangi stok untuk produk ID %d: %v", it.ProductID, err)
                return errors.New("stok tidak cukup")
            }
        }

        // Hitung diskon
        var discountAmount float64
        if req.VoucherCode != nil {
            var v models.Voucher
            if err := tx.Where("code = ? AND active = ?", *req.VoucherCode, 1).First(&v).Error; err == nil {
                sale.VoucherID = &v.ID
                switch v.Type {
                case models.VoucherPercent:
                    discountAmount = (v.Value / 100.0) * subtotal
                case models.VoucherAmount:
                    discountAmount = v.Value
                }
            } else {
                log.Printf("Voucher %s tidak ditemukan atau tidak aktif", *req.VoucherCode)
            }
        }

        sale.Subtotal = subtotal
        sale.Discount = discountAmount
        sale.Total = subtotal - discountAmount

        // Validasi pembayaran
        if req.PaidAmount == nil || *req.PaidAmount < sale.Total {
            log.Printf("Paid amount tidak cukup: %f < %f", *req.PaidAmount, sale.Total)
            return errors.New("paid_amount harus lebih besar atau sama dengan total")
        }
        sale.Paid = *req.PaidAmount
        sale.Change = *req.PaidAmount - sale.Total
        sale.Items = saleItems

        // Simpan sale
        if err := s.srepo.Create(sale); err != nil {
            log.Printf("Gagal menyimpan sale: %v", err)
            return err
        }

        return nil
    })

    if err != nil {
        return nil, err
    }

    // Muat relasi
    sale, err = s.LoadDetails(sale.ID)
    if err != nil {
        log.Printf("Gagal memuat detail sale ID %d: %v", sale.ID, err)
        return nil, err
    }

    log.Printf("Sale ID %d created successfully with TransactionNumber %s", sale.ID, sale.TransactionNumber)
    return sale, nil
}

func (s *SaleService) LoadDetails(id uint) (*models.Sale, error) {
    var sale models.Sale
    err := s.db.
        Preload("Customer").
        Preload("User").
        Preload("Voucher").
        Preload("Items.Product").
        First(&sale, id).Error
    if err != nil {
        log.Printf("Gagal memuat detail sale ID %d: %v", id, err)
        return nil, err
    }
    return &sale, nil
}

func (s *SaleService) List() ([]models.Sale, error) {
    var sales []models.Sale
    err := s.db.
        Preload("Customer").
        Preload("User").
        Preload("Voucher").
        Preload("Items.Product").
        Order("created_at DESC").
        Find(&sales).Error
    if err != nil {
        log.Printf("Gagal mengambil daftar penjualan: %v", err)
        return nil, errors.New("gagal mengambil daftar penjualan")
    }
    log.Printf("Mengambil %d penjualan", len(sales))
    return sales, nil
}