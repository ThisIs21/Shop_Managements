package services

import (
	"fmt"

	"app-penjualan/internal/utils"
	"gorm.io/gorm"
)

type ReportService struct{ DB *gorm.DB }

func NewReportService(db *gorm.DB) *ReportService {
	return &ReportService{DB: db}
}

// ======================= Helper =======================
func applyRange(q *gorm.DB, dr utils.DateRange, field string) *gorm.DB {
	if dr.From != nil {
		q = q.Where(field+" >= ?", *dr.From)
	}
	if dr.To != nil {
		q = q.Where(field+" <= ?", *dr.To)
	}
	return q
}

// ======================= Reports =======================

// SalesSummary
func (s *ReportService) SalesSummary(by string, dr utils.DateRange, month *int) ([]map[string]any, error) {
	col := "DATE(s.date)"
	if by == "month" {
		col = "DATE_FORMAT(s.date, '%Y-%m')"
	}
	q := s.DB.Table("sales s").
		Select(col+" as period, COUNT(*) trx, SUM(s.total) total").
		Group(col).Order(col)

	q = applyRange(q, dr, "s.date")

	if by == "month" && month != nil {
		q = q.Where("MONTH(s.date) = ? AND YEAR(s.date) = YEAR(NOW())", *month)
	}

	var rows []map[string]any
	return rows, q.Find(&rows).Error
}

// PurchaseSummary
func (s *ReportService) PurchaseSummary(by string, dr utils.DateRange, month *int) ([]map[string]any, error) {
	col := "DATE(p.date)"
	if by == "month" {
		col = "DATE_FORMAT(p.date, '%Y-%m')"
	}
	q := s.DB.Table("purchases p").
		Select(col+" as period, COUNT(*) trx, SUM(p.total) total").
		Group(col).Order(col)

	q = applyRange(q, dr, "p.date")

	if by == "month" && month != nil {
		q = q.Where("MONTH(p.date) = ? AND YEAR(p.date) = YEAR(NOW())", *month)
	}

	var rows []map[string]any
	return rows, q.Find(&rows).Error
}

// SaleReturnSummary
func (s *ReportService) SaleReturnSummary(by string, dr utils.DateRange, month *int) ([]map[string]any, error) {
	col := "DATE(sr.date)"
	if by == "month" {
		col = "DATE_FORMAT(sr.date, '%Y-%m')"
	}
	q := s.DB.Table("sale_returns sr").
		Select(col+" as period, COUNT(*) trx").
		Group(col).Order(col)

	q = applyRange(q, dr, "sr.date")

	if by == "month" && month != nil {
		q = q.Where("MONTH(sr.date) = ? AND YEAR(sr.date) = YEAR(NOW())", *month)
	}

	var rows []map[string]any
	return rows, q.Find(&rows).Error
}

// PurchaseReturnSummary
func (s *ReportService) PurchaseReturnSummary(by string, dr utils.DateRange, month *int) ([]map[string]any, error) {
	col := "DATE(pr.date)"
	if by == "month" {
		col = "DATE_FORMAT(pr.date, '%Y-%m')"
	}
	q := s.DB.Table("purchase_returns pr").
		Select(col+" as period, COUNT(*) trx").
		Group(col).Order(col)

	q = applyRange(q, dr, "pr.date")

	if by == "month" && month != nil {
		q = q.Where("MONTH(pr.date) = ? AND YEAR(pr.date) = YEAR(NOW())", *month)
	}

	var rows []map[string]any
	return rows, q.Find(&rows).Error
}

// StockOpnameSummary
func (s *ReportService) StockOpnameSummary(by string, dr utils.DateRange, month *int) ([]map[string]any, error) {
	col := "DATE(so.date)"
	if by == "month" {
		col = "DATE_FORMAT(so.date, '%Y-%m')"
	}
	q := s.DB.Table("stock_opnames so").
		Select(col+" as period, COUNT(*) opname").
		Group(col).Order(col)

	q = applyRange(q, dr, "so.date")

	if by == "month" && month != nil {
		q = q.Where("MONTH(so.date) = ? AND YEAR(so.date) = YEAR(NOW())", *month)
	}

	var rows []map[string]any
	return rows, q.Find(&rows).Error
}

// StockSnapshot
func (s *ReportService) StockSnapshot() ([]map[string]any, error) {
	q := s.DB.Table("products").
		Select("id, name, stock, cost_price, sell_price")

	var rows []map[string]any
	return rows, q.Find(&rows).Error
}

// InventoryValue
func (s *ReportService) InventoryValue() ([]map[string]any, error) {
	q := s.DB.Table("products").
		Select("SUM(stock*cost_price) as nilai_beli, SUM(stock*sell_price) as nilai_jual")

	var rows []map[string]any
	return rows, q.Find(&rows).Error
}

// InventoryDetail
func (s *ReportService) InventoryDetail() ([]map[string]any, error) {
	q := s.DB.Table("products p").
		Select(`
			p.id,
			p.name as product_name,
			CONCAT('P', p.id) as product_code,
			COALESCE(c.name, '') as category,
			COALESCE(u.name, '') as unit,
			p.stock as stock_quantity,
			CAST(p.cost_price AS DOUBLE) as purchase_price,
			CAST((p.stock * p.cost_price) AS DOUBLE) as total_value
		`).
		Joins("LEFT JOIN categories c ON c.id = p.category_id").
		Joins("LEFT JOIN units u ON u.id = p.unit_id").
		Order("p.name ASC")

	var rows []map[string]any
	err := q.Find(&rows).Error
	if err != nil {
		fmt.Println("GORM Error InventoryDetail:", err)
		return nil, err
	}

	return rows, nil
}
 
// TopProducts
func (s *ReportService) TopProducts(by string, dr utils.DateRange, month *int, limit int) ([]map[string]any, error) {
	q := s.DB.Table("sale_items si").
		Joins("JOIN sales s ON s.id = si.sale_id").
		Joins("JOIN products p ON p.id = si.product_id").
		Select("p.name as product_name, SUM(si.qty) as total_sold").
		Group("p.id, p.name").
		Order("total_sold DESC").
		Limit(limit)

	q = applyRange(q, dr, "s.date")

	if by == "month" && month != nil {
		q = q.Where("MONTH(s.date) = ? AND YEAR(s.date) = YEAR(NOW())", *month)
	}

	var rows []map[string]any
	return rows, q.Find(&rows).Error
}

// BottomProducts
func (s *ReportService) BottomProducts(by string, dr utils.DateRange, month *int, limit int) ([]map[string]any, error) {
	q := s.DB.Table("products p").
		Joins("LEFT JOIN sale_items si ON si.product_id = p.id").
		Joins("LEFT JOIN sales s ON s.id = si.sale_id").
		Select("p.name as product_name, COALESCE(SUM(si.qty), 0) as total_sold").
		Group("p.id, p.name").
		Order("total_sold ASC").
		Limit(limit)

	q = applyRange(q, dr, "s.date")

	if by == "month" && month != nil {
		q = q.Where("MONTH(s.date) = ? AND YEAR(s.date) = YEAR(NOW())", *month)
	}

	var rows []map[string]any
	return rows, q.Find(&rows).Error
}

// ======================= TAMBAHAN KODE DETAIL STOK OPNAME =======================

// StockOpnameDetail (BARU)
// Mengembalikan detail stok opname per barang (item)
func (s *ReportService) StockOpnameDetail(dr utils.DateRange) ([]map[string]any, error) {
    // ✅ PERBAIKAN: Ganti tabel dari 'stock_opname_details' menjadi 'stock_opname_items'
	q := s.DB.Table("stock_opname_items soi"). 
        // ✅ PERBAIKAN: Gunakan alias 'soi'
		Joins("JOIN stock_opnames so ON so.id = soi.stock_opname_id").
		Joins("JOIN products p ON p.id = soi.product_id").
		Select(`
			so.date AS date, 
			p.name AS book_name,
			CONCAT('P', p.id) AS product_code,
            
            -- ✅ PERBAIKAN NAMA KOLOM: Menggunakan qty_system dan qty_physical
			soi.qty_system AS system_stock, 
			soi.qty_physical AS physical_stock, 
            
            -- ✅ PERBAIKAN RUMUS SELISIH
			(soi.qty_physical - soi.qty_system) AS difference 
		`).
		Order("so.date DESC, p.name ASC")

	// Terapkan filter tanggal pada kolom 'so.date'
	q = applyRange(q, dr, "so.date")

	var rows []map[string]any
	err := q.Find(&rows).Error
	if err != nil {
		fmt.Println("GORM Error StockOpnameDetail:", err)
		return nil, err
	}

	return rows, nil
}