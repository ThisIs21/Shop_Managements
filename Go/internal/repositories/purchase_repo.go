package repositories

import (
    "app-penjualan/internal/models"
    "gorm.io/gorm"
)

type PurchaseRepo struct {
    db *gorm.DB
}

func NewPurchaseRepo(db *gorm.DB) *PurchaseRepo { return &PurchaseRepo{db} }

func (r *PurchaseRepo) Create(p *models.Purchase) error {
    return r.db.Create(p).Error
}

func (r *PurchaseRepo) WithItems(id uint) (*models.Purchase, error) {
    var m models.Purchase
    err := r.db.Preload("Supplier").Preload("User").Preload("Approver").
        Preload("Items.Product").First(&m, id).Error
    return &m, err
}

func (r *PurchaseRepo) Update(p *models.Purchase) error {
    return r.db.Save(p).Error
}

// ✅ KODE YANG SUDAH DIPERBAIKI. UBAH FindAll() menjadi FindAllWithRelations()
func (r *PurchaseRepo) FindAllWithRelations() ([]models.Purchase, error) {
    var list []models.Purchase
    err := r.db.
        Preload("Supplier").
        Preload("Items.Product").
        Find(&list).Error
    return list, err
}