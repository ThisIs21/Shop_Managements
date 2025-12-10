package repositories

import (
    "strconv"
    "gorm.io/gorm"
    "app-penjualan/internal/models"
)

type ProductRepo struct{ DB *gorm.DB } // Ubah db jadi DB

func NewProductRepo(db *gorm.DB) *ProductRepo { return &ProductRepo{DB: db} }

func (r *ProductRepo) FindAll() ([]models.Product, error) {
    var rows []models.Product
    err := r.DB.Preload("Category").Preload("Unit").Preload("Supplier").Find(&rows).Error
    return rows, err
}

func (r *ProductRepo) FindByID(id uint) (*models.Product, error) {
    var m models.Product
    err := r.DB.Preload("Category").Preload("Unit").Preload("Supplier").First(&m, id).Error
    return &m, err
}

func (r *ProductRepo) Create(m *models.Product) error { return r.DB.Create(m).Error }

func (r *ProductRepo) Update(m *models.Product) error { return r.DB.Save(m).Error }

func (r *ProductRepo) Delete(id uint) error { return r.DB.Delete(&models.Product{}, id).Error }

func (r *ProductRepo) AddStock(id uint, amount int) error {
    return r.DB.Model(&models.Product{}).Where("id = ?", id).Update("stock", gorm.Expr("stock + ?", amount)).Error
}

func (r *ProductRepo) ReduceStock(id uint, amount int) error {
    return r.DB.Model(&models.Product{}).
        Where("id = ? AND stock >= ?", id, amount).
        Update("stock", gorm.Expr("stock - ?", amount)).Error
}

// FIX: Tambahkan Preload Unit dan Category ke FindFiltered
func (r *ProductRepo) FindFiltered(search, category, minPrice, maxPrice string) ([]models.Product, error) {
    // Pastikan Preload Unit dan Category di sini
    db := r.DB.Model(&models.Product{}).
        Preload("Category").
        Preload("Unit"). 
        Preload("Supplier") 

    if search != "" {
        db = db.Where("name LIKE ?", "%"+search+"%")
    }
    if category != "" {
        db = db.Where("category_id = ?", category)
    }
    if minPrice != "" {
        if v, err := strconv.Atoi(minPrice); err == nil {
            db = db.Where("sell_price >= ?", v)
        }
    }
    if maxPrice != "" {
        if v, err := strconv.Atoi(maxPrice); err == nil {
            db = db.Where("sell_price <= ?", v)
        }
    }
    var rows []models.Product
    if err := db.Find(&rows).Error; err != nil {
        return nil, err
    }
    return rows, nil
}