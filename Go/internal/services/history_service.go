package services

import (
	"app-penjualan/internal/models"
	"app-penjualan/internal/utils"
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/jung-kurt/gofpdf"
	"gorm.io/gorm"
)

type HistoryService struct {
	DB *gorm.DB
}

func NewHistoryService(db *gorm.DB) *HistoryService {
	return &HistoryService{DB: db}
}

func (s *HistoryService) Purchases(dr utils.DateRange, page, size int, search, status string) (any, int64, error) {
	var purchases []models.Purchase
	var total int64

	query := s.DB.Model(&models.Purchase{}).
		Preload("Supplier").
		Preload("Items.Product").
		Scopes(utils.Paginate(page, size)).
		Order("date DESC")

	if search != "" {
		query = query.Where("CAST(id as CHAR) LIKE ? OR CAST(total as CHAR) LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	query = utils.ApplyRange(query, dr, "date")

	if err := query.Count(&total).Error; err != nil {
		log.Printf("Gagal menghitung total purchases: %v", err)
		return nil, 0, err
	}
	if err := query.Find(&purchases).Error; err != nil {
		log.Printf("Gagal mengambil purchases: %v", err)
		return nil, 0, err
	}

	log.Printf("Mengambil %d purchases dengan status %s, page %d, size %d", total, status, page, size)
	return purchases, total, nil
}

func (s *HistoryService) Sales(dr utils.DateRange, page, size int, search string) ([]models.Sale, int64, float64, error) {
	var sales []models.Sale
	var total int64
	var totalValue float64

	baseQuery := s.DB.Model(&models.Sale{})

	if search != "" {
		baseQuery = baseQuery.
			Joins("left join customers on customers.id = sales.customer_id").
			Joins("left join sale_items on sale_items.sale_id = sales.id").
			Joins("left join products on products.id = sale_items.product_id").
			Where("LOWER(customers.name) LIKE ? OR LOWER(products.name) LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	baseQuery = utils.ApplyRange(baseQuery, dr, "sales.created_at")

	if err := baseQuery.Count(&total).Error; err != nil {
		log.Printf("Gagal menghitung total penjualan: %v", err)
		return nil, 0, 0, err
	}

	totalQuery := s.DB.Model(&models.Sale{})
	if search != "" {
		totalQuery = totalQuery.
			Joins("left join customers on customers.id = sales.customer_id").
			Joins("left join sale_items on sale_items.sale_id = sales.id").
			Joins("left join products on products.id = sale_items.product_id").
			Where("LOWER(customers.name) LIKE ? OR LOWER(products.name) LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	totalQuery = utils.ApplyRange(totalQuery, dr, "sales.created_at")
	var result struct {
		TotalValue float64
	}
	if err := totalQuery.Select("COALESCE(SUM(total), 0) as total_value").Scan(&result).Error; err != nil {
		log.Printf("Gagal menghitung total_value penjualan: %v", err)
		return nil, 0, 0, err
	}
	totalValue = result.TotalValue

	dataQuery := s.DB.Model(&models.Sale{})
	if search != "" {
		dataQuery = dataQuery.
			Joins("left join customers on customers.id = sales.customer_id").
			Joins("left join sale_items on sale_items.sale_id = sales.id").
			Joins("left join products on products.id = sale_items.product_id").
			Where("LOWER(customers.name) LIKE ? OR LOWER(products.name) LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	dataQuery = utils.ApplyRange(dataQuery, dr, "sales.created_at")
	dataQuery = dataQuery.Scopes(utils.Paginate(page, size))
	dataQuery = dataQuery.Preload("Customer").
		Preload("Voucher").
		Preload("Items.Product").
		Order("sales.created_at DESC")

	if err := dataQuery.Find(&sales).Error; err != nil {
		log.Printf("Gagal mengambil data penjualan: %v", err)
		return nil, 0, 0, err
	}

	log.Printf("Mengambil %d penjualan dengan total nilai: %f", total, totalValue)
	return sales, total, totalValue, nil
}

func (s *HistoryService) LatestSales(limit int) ([]models.Sale, int64, float64, error) {
	var sales []models.Sale
	var total int64
	var totalValue float64

	if err := s.DB.Model(&models.Sale{}).Count(&total).Error; err != nil {
		log.Printf("Gagal menghitung total penjualan: %v", err)
		return nil, 0, 0, err
	}

	var result struct {
		TotalValue float64
	}
	if err := s.DB.Model(&models.Sale{}).Select("COALESCE(SUM(total), 0) as total_value").Scan(&result).Error; err != nil {
		log.Printf("Gagal menghitung total_value penjualan: %v", err)
		return nil, 0, 0, err
	}
	totalValue = result.TotalValue

	if err := s.DB.Model(&models.Sale{}).
		Preload("Customer").
		Preload("Voucher").
		Preload("Items.Product").
		Order("sales.created_at DESC").
		Limit(limit).
		Find(&sales).Error; err != nil {
		log.Printf("Gagal mengambil %d penjualan terakhir: %v", limit, err)
		return nil, 0, 0, err
	}

	log.Printf("Mengambil %d penjualan terakhir dengan total nilai: %f", len(sales), totalValue)
	return sales, total, totalValue, nil
}

func (s *HistoryService) GetSaleByID(id uint) (*models.Sale, error) {
	var sale models.Sale
	err := s.DB.
		Preload("Customer").
		Preload("User").
		Preload("Voucher").
		Preload("Items.Product").
		First(&sale, id).Error
	if err != nil {
		log.Printf("Gagal mengambil penjualan ID %d: %v", id, err)
		return nil, err
	}
	return &sale, nil
}

func (s *HistoryService) PurchaseReturns(dr utils.DateRange, page, size int, search string) (any, int64, error) {
	q := s.DB.Table("purchase_returns pr").
		Select(`pr.id, pr.purchase_id, pr.status, pr.date, pr.total, pr.created_at, p.supplier_id`).
		Joins("LEFT JOIN purchases p ON p.id = pr.purchase_id").
		Scopes(utils.Paginate(page, size)).
		Order("pr.date DESC")

	if search != "" {
		q = q.Where("CAST(pr.id as CHAR) LIKE ? OR CAST(p.total as CHAR) LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	q = utils.ApplyRange(q, dr, "pr.date")
	var rows []map[string]any
	var total int64
	if err := q.Count(&total).Error; err != nil {
		log.Printf("Gagal menghitung total purchase returns: %v", err)
		return nil, 0, err
	}
	if err := q.Find(&rows).Error; err != nil {
		log.Printf("Gagal mengambil purchase returns: %v", err)
		return nil, 0, err
	}
	log.Printf("Mengambil %d purchase returns", total)
	return rows, total, nil
}

func (s *HistoryService) SaleReturns(dr utils.DateRange, page, size int, search string) (any, int64, error) {
	q := s.DB.Table("sale_returns sr").
		Select(`sr.id, sr.sale_id, sr.status, sr.date, sr.reason`).
		Scopes(utils.Paginate(page, size)).
		Order("sr.date DESC")

	if search != "" {
		q = q.Where("CAST(sr.id as CHAR) LIKE ? OR sr.reason LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	q = utils.ApplyRange(q, dr, "sr.date")
	var rows []map[string]any
	var total int64
	if err := q.Count(&total).Error; err != nil {
		log.Printf("Gagal menghitung total sale returns: %v", err)
		return nil, 0, err
	}
	if err := q.Find(&rows).Error; err != nil {
		log.Printf("Gagal mengambil sale returns: %v", err)
		return nil, 0, err
	}
	log.Printf("Mengambil %d sale returns", total)
	return rows, total, nil
}

func (s *HistoryService) StockOpnames(dr utils.DateRange, page, size int, search string) ([]models.StockOpname, int64, error) {
	var opnames []models.StockOpname
	var total int64

	query := s.DB.Model(&models.StockOpname{}).
		Preload("User").
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Unit").
		Preload("Items.Product.Category").
		Scopes(utils.Paginate(page, size)).
		Order("date DESC")

	if search != "" {
		query = query.Joins("LEFT JOIN users ON users.id = stock_opnames.user_id").
			Where("CAST(stock_opnames.id AS CHAR) LIKE ? OR LOWER(users.name) LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	query = utils.ApplyRange(query, dr, "date")

	if err := query.Count(&total).Error; err != nil {
		log.Printf("Gagal menghitung total stock opnames: %v", err)
		return nil, 0, err
	}
	if err := query.Find(&opnames).Error; err != nil {
		log.Printf("Gagal mengambil stock opnames: %v", err)
		return nil, 0, err
	}

	for _, op := range opnames {
		log.Printf("StockOpname ID %d, Items count: %d", op.ID, len(op.Items))
		for _, item := range op.Items {
			log.Printf("Item ID %d, Product: %v, QtySystem: %d, QtyPhysical: %d", item.ID, item.Product, item.QtySystem, item.QtyPhysical)
		}
	}

	log.Printf("Mengambil %d stock opnames", total)
	return opnames, total, nil
}

func (s *HistoryService) GetStockOpnameByID(id uint) (*models.StockOpname, error) {
	var opname models.StockOpname
	err := s.DB.
		Preload("User").
		Preload("Items").
		Preload("Items.Product").
		Preload("Items.Product.Unit").
		Preload("Items.Product.Category").
		First(&opname, id).Error
	if err != nil {
		log.Printf("Gagal mengambil stock opname ID %d: %v", id, err)
		return nil, err
	}
	return &opname, nil
}

func (s *HistoryService) GenerateSalePDF(id uint) ([]byte, error) {
	var sale models.Sale
	if err := s.DB.
		Preload("Customer").
		Preload("User").
		Preload("Voucher").
		Preload("Items.Product").
		First(&sale, id).Error; err != nil {
		log.Printf("Gagal mengambil sale ID %d untuk PDF: %v", id, err)
		return nil, err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Struk Pembayaran")
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "ID Transaksi: "+strconv.Itoa(int(sale.ID)))
	pdf.Ln(5)
	pdf.Cell(40, 10, "Tanggal: "+sale.Date.Format("2006-01-02 15:04:05"))
	pdf.Ln(5)
	customerName := "Umum"
	if sale.Customer != nil {
		customerName = sale.Customer.Name
	}
	pdf.Cell(40, 10, "Pelanggan: "+customerName)
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Detail Pembelian")
	pdf.Ln(5)

	pdf.SetFont("Arial", "", 10)
	for _, item := range sale.Items {
		pdf.Cell(80, 8, item.Product.Name)
		pdf.Cell(20, 8, "x"+strconv.Itoa(item.Qty))
		pdf.Cell(40, 8, "Rp "+formatCurrency(item.Subtotal))
		pdf.Ln(5)
	}

	pdf.Ln(5)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, "Subtotal: Rp "+formatCurrency(sale.Subtotal))
	pdf.Ln(5)
	pdf.Cell(40, 10, "Diskon: Rp "+formatCurrency(sale.Discount))
	pdf.Ln(5)
	pdf.Cell(40, 10, "Total: Rp "+formatCurrency(sale.Total))
	pdf.Ln(5)
	pdf.Cell(40, 10, "Dibayar: Rp "+formatCurrency(sale.Paid))
	pdf.Ln(5)
	pdf.Cell(40, 10, "Kembali: Rp "+formatCurrency(sale.Change))
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 10)
	pdf.Cell(40, 10, "Terima kasih telah berbelanja!")
	pdf.Ln(5)

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		log.Printf("Gagal menghasilkan PDF: %v", err)
		return nil, err
	}

	return buf.Bytes(), nil
}

func formatCurrency(amount float64) string {
	return fmt.Sprintf("Rp %.0f", amount)
}

func (s *HistoryService) GetPurchaseByID(id uint) (*models.Purchase, error) {
	var p models.Purchase
	err := s.DB.
		Preload("Supplier").
		Preload("Items.Product.Unit").
		Preload("Items.Product.Category").
		Preload("User").
		Preload("Approver").
		Preload("Receiver").
		First(&p, id).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}