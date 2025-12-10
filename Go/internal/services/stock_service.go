package services

import (
	"app-penjualan/internal/models"
	"app-penjualan/internal/repositories"
	"app-penjualan/internal/dto"
	

	"gorm.io/gorm"
)

type StockService struct {
	db *gorm.DB
}

func NewStockService(db *gorm.DB) *StockService {
	return &StockService{db: db}
}

func (s *StockService) DoStockOpname(uid uint, req dto.CreateOpnameReqService) (*models.StockOpname, error) {
	op := &models.StockOpname{UserID: uid, Date: req.Date, Note: req.Note}

	err := s.db.Transaction(func(tx *gorm.DB) error {
		prodServiceTx := NewProductService(tx)

		for _, it := range req.Items {
			var p models.Product
			if err := tx.Preload("Unit").Preload("Category").First(&p, it.ProductID).Error; err != nil {
				return err
			}

			op.Items = append(op.Items, models.StockOpnameItem{
				StockOpnameID: op.ID,
				ProductID:     it.ProductID,
				Product:       p,
				QtySystem:     p.Stock,
				QtyPhysical:   it.QtyPhysical,
			})

			if delta := it.QtyPhysical - p.Stock; delta != 0 {
				if err := prodServiceTx.AdjustStock(it.ProductID, delta); err != nil {
					return err
				}
			}
		}

		repo := repositories.NewStockRepo(tx)
		if err := repo.Create(op); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return op, nil
}

func (s *StockService) ListStockOpnames(limit int) ([]models.StockOpname, error) {
    var opnames []models.StockOpname
    if err := s.db.Model(&models.StockOpname{}).
        Preload("User").
        Preload("Items").
        Preload("Items.Product.Category").
        Preload("Items.Product.Unit").
        Limit(limit).
        Order("date DESC").
        Find(&opnames).Error; err != nil {
        return nil, err
    }
    return opnames, nil
}
