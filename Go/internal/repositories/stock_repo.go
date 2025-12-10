// internal/repositories/stock_repo.go (Hapus import strconv jika tidak dipakai)
package repositories

import (
	"gorm.io/gorm"
	"app-penjualan/internal/models"
	"time" // Keep untuk parsing di List/GetByID jika butuh
)

type StockRepo struct{ db *gorm.DB }
func NewStockRepo(db *gorm.DB) *StockRepo { return &StockRepo{db} }

func (r *StockRepo) Create(op *models.StockOpname) error { return r.db.Create(op).Error }

// List: Ambil list stock opnames dengan filter (page, size, search q, start_date, end_date)
func (r *StockRepo) List(page, size int, q, startDateStr, endDateStr string) ([]models.StockOpname, int64, error) {
    db := r.db.Model(&models.StockOpname{}).
        Preload("User").
        Preload("Items.Product.Category").
        Preload("Items.Product.Unit")

    // Filter tanggal
    if startDateStr != "" {
        parsedStart, err := time.Parse("2006-01-02", startDateStr)
        if err == nil {
            db = db.Where("date >= ?", parsedStart)
        }
    }
    if endDateStr != "" {
        parsedEnd, err := time.Parse("2006-01-02", endDateStr)
        if err == nil {
            // Sampai akhir hari
            endOfDay := parsedEnd.Add(24*time.Hour - time.Second)
            db = db.Where("date <= ?", endOfDay)
        }
    }

    // Filter search (ID atau nama user)
    if q != "" {
        db = db.Joins("LEFT JOIN users ON users.id = stock_opnames.user_id").
            Where("stock_opnames.id::text LIKE ? OR users.name LIKE ?", "%"+q+"%", "%"+q+"%")
    }

    var total int64
    if err := db.Count(&total).Error; err != nil {
        return nil, 0, err
    }

    var opnames []models.StockOpname
    offset := (page - 1) * size
    if err := db.Order("date DESC").Offset(offset).Limit(size).Find(&opnames).Error; err != nil {
        return nil, total, err
    }

    return opnames, total, nil
}

// GetByID: Ambil detail stock opname dengan preload lengkap
func (r *StockRepo) GetByID(id uint) (*models.StockOpname, error) {
    var opname models.StockOpname
    err := r.db.
        Preload("User").
        Preload("Items.Product.Category").
        Preload("Items.Product.Unit").
        First(&opname, id).Error
    if err != nil {
        return nil, err
    }
    return &opname, nil
}