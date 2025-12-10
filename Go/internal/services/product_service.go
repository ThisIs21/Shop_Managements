package services

import (
    "errors"
    "gorm.io/gorm"
    "app-penjualan/internal/models"
    "app-penjualan/internal/repositories"
)

type ProductService struct {
    Repo *repositories.ProductRepo
}

func NewProductService(db *gorm.DB) *ProductService {
    return &ProductService{Repo: repositories.NewProductRepo(db)}
}

func (s *ProductService) List() ([]models.Product, error) {
    return s.Repo.FindAll()
}

func (s *ProductService) Get(id uint) (*models.Product, error) {
    return s.Repo.FindByID(id)
}

func (s *ProductService) Create(p *models.Product) error {
    return s.Repo.Create(p)
}

func (s *ProductService) Update(id uint, in *models.Product) (*models.Product, error) {
    p, err := s.Repo.FindByID(id)
    if err != nil {
        return nil, err
    }
    // Validasi ID relasi
    if in.CategoryID != 0 {
        var category models.Category
        if err := s.Repo.DB.First(&category, in.CategoryID).Error; err != nil {
            return nil, errors.New("category_id tidak valid")
        }
    }
    if in.UnitID != 0 {
        var unit models.Unit
        if err := s.Repo.DB.First(&unit, in.UnitID).Error; err != nil {
            return nil, errors.New("unit_id tidak valid")
        }
    }
    if in.SupplierID != 0 {
        var supplier models.Supplier
        if err := s.Repo.DB.First(&supplier, in.SupplierID).Error; err != nil {
            return nil, errors.New("supplier_id tidak valid")
        }
    }
    p.Name = in.Name
    p.CategoryID = in.CategoryID
    p.UnitID = in.UnitID
    p.CostPrice = in.CostPrice
    p.SellPrice = in.SellPrice
    p.SupplierID = in.SupplierID
    p.Stock = in.Stock
    if err := s.Repo.Update(p); err != nil {
        return nil, err
    }
    // Preload relasi
    if err := s.Repo.DB.Preload("Category").Preload("Unit").Preload("Supplier").First(p, p.ID).Error; err != nil {
        return nil, errors.New("gagal memuat data relasi")
    }
    return p, nil
}

func (s *ProductService) Delete(id uint) error {
    // Hapus data terkait
    if err := s.Repo.DB.Where("product_id = ?", id).Delete(&models.PurchaseItem{}).Error; err != nil {
        return errors.New("gagal hapus purchase_items: " + err.Error())
    }
    if err := s.Repo.DB.Where("product_id = ?", id).Delete(&models.PurchaseReturnItem{}).Error; err != nil {
        return errors.New("gagal hapus purchase_return_items: " + err.Error())
    }
    if err := s.Repo.DB.Where("product_id = ?", id).Delete(&models.SaleItem{}).Error; err != nil {
        return errors.New("gagal hapus sale_items: " + err.Error())
    }
    if err := s.Repo.DB.Where("product_id = ?", id).Delete(&models.SaleReturnItem{}).Error; err != nil {
        return errors.New("gagal hapus sale_return_items: " + err.Error())
    }
    // Hapus produk
    return s.Repo.Delete(id)
}

func (s *ProductService) AdjustStock(productID uint, delta int) error {
    if delta == 0 {
        return nil
    }
    if delta > 0 {
        return s.Repo.AddStock(productID, delta)
    }
    if err := s.Repo.ReduceStock(productID, -delta); err != nil {
        return errors.New("stok tidak cukup atau produk tidak ditemukan")
    }
    return nil
}

func (s *ProductService) ListFiltered(search, category, minPrice, maxPrice string) ([]models.Product, error) {
    return s.Repo.FindFiltered(search, category, minPrice, maxPrice)
}