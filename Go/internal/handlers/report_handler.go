package handlers

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"app-penjualan/internal/services"
	"app-penjualan/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

type ReportHandler struct {
	Svc *services.ReportService
}

func NewReportHandler(s *services.ReportService) *ReportHandler {
	return &ReportHandler{Svc: s}
}

// helper export CSV
func writeCSV(c *gin.Context, header []string, rows [][]string, filename string) {
	// ... (Kode writeCSV tidak berubah) ...
	buf := &bytes.Buffer{}
	w := csv.NewWriter(buf)
	_ = w.Write(header)
	_ = w.WriteAll(rows)
	w.Flush()
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.String(http.StatusOK, buf.String())
}

// helper export PDF
func writePDF(c *gin.Context, title string, header []string, rows [][]string, filename string) {
	// ... (Kode writePDF tidak berubah) ...
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(40, 10, title)
	pdf.Ln(12)
	pdf.SetFont("Arial", "", 10)
	for _, h := range header {
		pdf.CellFormat(40, 8, h, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)
	for _, r := range rows {
		for _, col := range r {
			pdf.CellFormat(40, 8, col, "1", 0, "L", false, 0, "")
		}
		pdf.Ln(-1)
	}
	var buf bytes.Buffer
	_ = pdf.Output(&buf)
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "application/pdf", buf.Bytes())
}

func toS(v any) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%v", v)
}

func (h *ReportHandler) GenerateReport(c *gin.Context) {
	reportType := c.DefaultQuery("type", "")
	if reportType == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "report type is required"})
		return
	}

	by := c.DefaultQuery("by", "day")
	dr, err := utils.ParseDateRange(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	monthStr := c.Query("month")
	var month *int
	if monthStr != "" && by == "month" {
		m, err := strconv.Atoi(monthStr)
		if err != nil || m < 1 || m > 12 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid month (1-12)"})
			return
		}
		month = &m
	}

	var data []map[string]any
	var err2 error
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Tentukan apakah kita perlu data DETAIL
	needsDetail := c.DefaultQuery("detail", "false") == "true"

	switch reportType {
	case "sales":
		data, err2 = h.Svc.SalesSummary(by, dr, month)
	case "purchases":
		data, err2 = h.Svc.PurchaseSummary(by, dr, month)
	case "sale-returns":
		data, err2 = h.Svc.SaleReturnSummary(by, dr, month)
	case "purchase-returns":
		data, err2 = h.Svc.PurchaseReturnSummary(by, dr, month)
	case "stock-opnames":
		// 💡 MODIFIKASI DIMULAI DI SINI:
		// Cek query parameter 'detail'. Jika true, panggil service Detail.
		if needsDetail {
			// Asumsi Anda sudah membuat fungsi di ReportService untuk detail.
			// Nama fungsi harus diganti sesuai yang Anda buat (misalnya StockOpnameDetail)
			data, err2 = h.Svc.StockOpnameDetail(dr) // Contoh pemanggilan dengan DateRange
		} else {
			// Jika tidak ada parameter 'detail' atau 'detail=false', gunakan ringkasan default.
			data, err2 = h.Svc.StockOpnameSummary(by, dr, month)
		}
		// 💡 MODIFIKASI BERAKHIR DI SINI.

	case "stock-snapshot":
		data, err2 = h.Svc.StockSnapshot()
	case "inventory-value":
		data, err2 = h.Svc.InventoryValue()
	case "top-products":
		data, err2 = h.Svc.TopProducts(by, dr, month, limit)
	case "bottom-products":
		data, err2 = h.Svc.BottomProducts(by, dr, month, limit)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid report type"})
		return
	}

	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}

	// Bagian export CSV dan PDF harus dimodifikasi total untuk menangani struktur data detail!
	// Karena struktur data detail sangat berbeda dari ringkasan, bagian ini berpotensi error/rusak.
	// Jika data adalah DETAIL Stok Opname, Anda perlu header dan logic rows yang berbeda.
	// Untuk saat ini, kita biarkan logic export default berjalan, dan hanya fokus pada JSON response.

	// ... (Kode export CSV dan PDF tidak berubah, karena ini hanya fokus pada JSON data) ...
	// HATI-HATI: Jika data yang dikembalikan adalah detail, logic di bawah ini untuk CSV/PDF
	// ('period', 'trx', 'total') tidak akan cocok dan bisa menyebabkan crash atau data salah.

	export := c.Query("export")
	if export == "csv" {
		var header []string
		var rows [][]string
		if strings.Contains(reportType, "products") {
			header = []string{"product_name", "total_sold"}
			for _, m := range data {
				rows = append(rows, []string{toS(m["product_name"]), toS(m["total_sold"])})
			}
		} else {
			header = []string{"period", "trx", "total"}
			for _, m := range data {
				rows = append(rows, []string{toS(m["period"]), toS(m["trx"]), toS(m["total"])})
			}
		}
		filename := strings.ReplaceAll(reportType, "-", "_") + ".csv"
		writeCSV(c, header, rows, filename)
		return
	}

	if export == "pdf" {
		var header []string
		var rows [][]string
		title := strings.Title(strings.ReplaceAll(reportType, "-", " ")) + " Report"
		if strings.Contains(reportType, "products") {
			header = []string{"Product Name", "Total Sold"}
			for _, m := range data {
				rows = append(rows, []string{toS(m["product_name"]), toS(m["total_sold"])})
			}
		} else {
			header = []string{"Period", "Trx", "Total"}
			for _, m := range data {
				rows = append(rows, []string{toS(m["period"]), toS(m["trx"]), toS(m["total"])})
			}
		}
		filename := strings.ReplaceAll(reportType, "-", "_") + ".pdf"
		writePDF(c, title, header, rows, filename)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": data})
}

// ✅ InventoryDetail
func (h *ReportHandler) InventoryDetail(c *gin.Context) {
	data, err := h.Svc.InventoryDetail()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}

// DashboardSummary returns simple aggregated totals for owner dashboard
func (h *ReportHandler) DashboardSummary(c *gin.Context) {
	// Support query params:
	// - period: today | month | year (default: month)
	// - from, to: explicit YYYY-MM-DD range (overrides period when both provided)
	period := c.Query("period")
	from := c.Query("from")
	to := c.Query("to")

	var totalSales float64
	var totalPurchases float64
	var totalReturns float64

	// Helper to build date condition and args
	buildDateCond := func(base string) (string, []interface{}) {
		if from != "" && to != "" {
			return base + " AND DATE(date) BETWEEN ? AND ?", []interface{}{from, to}
		}
		switch period {
		case "today":
			return base + " AND DATE(date) = CURRENT_DATE()", nil
		case "year":
			return base + " AND YEAR(date) = YEAR(NOW())", nil
		default:
			// default to current month
			return base + " AND MONTH(date) = MONTH(NOW()) AND YEAR(date) = YEAR(NOW())", nil
		}
	}

	// Sales: include all sales (sum of total)
	salesBase := "SELECT COALESCE(SUM(total),0) FROM sales WHERE 1=1"
	salesQuery, salesArgs := buildDateCond(salesBase)
	if err := h.Svc.DB.Raw(salesQuery, salesArgs...).Scan(&totalSales).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Purchases: only count purchases that have been processed by kepala gudang
	// (either approved OR rejected) — i.e. approved_by set or status in ('APPROVED','RECEIVED','REJECTED')
	purchasesBase := "SELECT COALESCE(SUM(total),0) FROM purchases WHERE (approved_by IS NOT NULL OR status IN ('APPROVED','RECEIVED','REJECTED'))"
	purchasesQuery, purchasesArgs := buildDateCond(purchasesBase)
	if err := h.Svc.DB.Raw(purchasesQuery, purchasesArgs...).Scan(&totalPurchases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Returns (sales returns): only include returns that have been approved
	returnsBase := "SELECT COALESCE(SUM(total),0) FROM sale_returns WHERE status = 'APPROVED'"
	returnsQuery, returnsArgs := buildDateCond(returnsBase)
	if err := h.Svc.DB.Raw(returnsQuery, returnsArgs...).Scan(&totalReturns).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": gin.H{
		"total_sales":     totalSales,
		"total_purchases": totalPurchases,
		"total_returns":   totalReturns,
		"period":          period,
		"from":            from,
		"to":              to,
	}})
}
