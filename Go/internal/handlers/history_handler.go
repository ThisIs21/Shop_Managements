package handlers

import (
	"errors" // Diperlukan untuk errors.Is
	"log"
	"net/http"
	"strconv"

	"app-penjualan/internal/services" // Sesuaikan dengan path service Anda
	"app-penjualan/internal/utils"    // Sesuaikan dengan path utils Anda
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HistoryHandler struct {
	Svc *services.HistoryService
}

func NewHistoryHandler(s *services.HistoryService) *HistoryHandler {
	return &HistoryHandler{Svc: s}
}

// Register menambahkan endpoint untuk HistoryHandler
func (h *HistoryHandler) Register(rg *gin.RouterGroup) {
	rg.GET("/purchases", h.Purchases)
	rg.GET("/purchases/:id", h.GetPurchaseByID)
	rg.GET("/purchase-returns/history", h.PurchaseReturnsHistory)
	rg.GET("/sale-returns/history", h.SaleReturnsHistory)
	rg.GET("/stock-opnames", h.StockOpnames) // Tambahkan endpoint ini
}

func getPaging(c *gin.Context) (int, int, string) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	search := c.DefaultQuery("q", "")
	log.Printf("Paging params - page: %d, size: %d, search: %s", page, size, search)
	return page, size, search
}

func (h *HistoryHandler) Purchases(c *gin.Context) {
	dr, err := utils.ParseDateRange(c)
	if err != nil {
		log.Printf("Gagal parsing date range: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page, size, search := getPaging(c)
	status := c.DefaultQuery("status", "APPROVED")

	data, total, err := h.Svc.Purchases(dr, page, size, search, status)
	if err != nil {
		log.Printf("Gagal mengambil histori pembelian: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Mengambil %d pembelian dengan status %s, page %d, size %d", total, status, page, size)
	c.JSON(http.StatusOK, gin.H{
		"data":    gin.H{"data": data, "total": total, "page": page, "size": size},
		"success": true,
	})
}

// GetPurchaseByID menangani detail pembelian berdasarkan ID
func (h *HistoryHandler) GetPurchaseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("ID pembelian tidak valid: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid purchase ID"})
		return
	}

	purchase, err := h.Svc.GetPurchaseByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Pembelian ID %d tidak ditemukan", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "Purchase not found"})
			return
		}
		log.Printf("Gagal mengambil pembelian ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Mengambil pembelian ID %d", id)
	c.JSON(http.StatusOK, gin.H{"data": purchase, "success": true})
}

func (h *HistoryHandler) Sales(c *gin.Context) {
	dr, err := utils.ParseDateRange(c)
	if err != nil {
		log.Printf("Gagal parsing date range: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page, size, search := getPaging(c)
	data, total, totalValue, err := h.Svc.Sales(dr, page, size, search)
	if err != nil {
		log.Printf("Gagal mengambil histori penjualan: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Menemukan %d penjualan dengan total nilai: %f", total, totalValue)
	c.JSON(http.StatusOK, gin.H{
		"data":        data,
		"total":       total,
		"total_value": totalValue,
		"page":        page,
		"size":        size,
	})
}

func (h *HistoryHandler) GetSaleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("ID penjualan tidak valid: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sale ID"})
		return
	}
	sale, err := h.Svc.GetSaleByID(uint(id))
	if err != nil {
		log.Printf("Gagal mengambil penjualan ID %d: %v", id, err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if sale == nil {
		log.Printf("Penjualan ID %d tidak ditemukan", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Sale not found"})
		return
	}
	log.Printf("Mengambil penjualan ID %d", id)
	c.JSON(http.StatusOK, sale)
}

func (h *HistoryHandler) PurchaseReturns(c *gin.Context) {
	dr, err := utils.ParseDateRange(c)
	if err != nil {
		log.Printf("Gagal parsing date range: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page, size, search := getPaging(c)
	data, total, err := h.Svc.PurchaseReturns(dr, page, size, search)
	if err != nil {
		log.Printf("Gagal mengambil histori retur pembelian: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Mengambil %d retur pembelian", total)
	c.JSON(http.StatusOK, gin.H{"data": data, "total": total, "page": page, "size": size})
}

func (h *HistoryHandler) SaleReturns(c *gin.Context) {
	dr, err := utils.ParseDateRange(c)
	if err != nil {
		log.Printf("Gagal parsing date range: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page, size, search := getPaging(c)
	data, total, err := h.Svc.SaleReturns(dr, page, size, search)
	if err != nil {
		log.Printf("Gagal mengambil histori retur penjualan: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Mengambil %d retur penjualan", total)
	c.JSON(http.StatusOK, gin.H{"data": data, "total": total, "page": page, "size": size})
}

func (h *HistoryHandler) StockOpnames(c *gin.Context) {
    dr, err := utils.ParseDateRange(c)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    page, size, search := getPaging(c)

    data, total, err := h.Svc.StockOpnames(dr, page, size, search)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data":  data,   // sudah include user + product + category + unit
        "total": total,
        "page":  page,
        "size":  size,
    })
}

func (h *HistoryHandler) GetStockOpnameByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("ID stock opname tidak valid: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock opname ID"})
		return
	}

	opname, err := h.Svc.GetStockOpnameByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Stock opname ID %d tidak ditemukan", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "Stock opname not found"})
			return
		}
		log.Printf("Gagal mengambil stock opname ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Mengambil stock opname ID %d", id)
	c.JSON(http.StatusOK, gin.H{"data": opname, "success": true})
}

func (h *HistoryHandler) LatestSales(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}
	log.Printf("Mengambil %d penjualan terakhir", limit)
	sales, total, totalValue, err := h.Svc.LatestSales(limit)
	if err != nil {
		log.Printf("Gagal mengambil penjualan terakhir: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("Mengambil %d penjualan terakhir dengan total nilai: %f", total, totalValue)
	c.JSON(http.StatusOK, gin.H{
		"data":        sales,
		"total":       total,
		"total_value": totalValue,
	})
}

func (h *HistoryHandler) GenerateSalePDF(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("ID penjualan tidak valid: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sale ID"})
		return
	}
	pdfBytes, err := h.Svc.GenerateSalePDF(uint(id))
	if err != nil {
		log.Printf("Gagal menghasilkan PDF untuk sale ID %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename=struk_"+strconv.Itoa(id)+".pdf")
	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}

// PurchaseReturnsHistory menangani riwayat retur pembelian
func (h *HistoryHandler) PurchaseReturnsHistory(c *gin.Context) {
	dr, err := utils.ParseDateRange(c)
	if err != nil {
		log.Printf("Gagal parsing date range: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page, size, search := getPaging(c)
	status := c.DefaultQuery("status", "")
	data, total, err := h.Svc.PurchaseReturns(dr, page, size, search)
	if err != nil {
		log.Printf("Gagal mengambil histori retur pembelian: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if status != "" {
		if list, ok := data.([]map[string]any); ok {
			var filteredData []map[string]any
			for _, item := range list {
				if itemStatus, ok := item["status"].(string); ok && itemStatus == status {
					filteredData = append(filteredData, item)
				}
			}
			data = filteredData
			total = int64(len(filteredData))
		}
	}
	log.Printf("Mengambil %d retur pembelian dengan status %s", total, status)
	c.JSON(http.StatusOK, gin.H{"data": data, "total": total, "page": page, "size": size})
}

// SaleReturnsHistory menangani riwayat retur penjualan
func (h *HistoryHandler) SaleReturnsHistory(c *gin.Context) {
	dr, err := utils.ParseDateRange(c)
	if err != nil {
		log.Printf("Gagal parsing date range: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	page, size, search := getPaging(c)
	status := c.DefaultQuery("status", "")
	data, total, err := h.Svc.SaleReturns(dr, page, size, search)
	if err != nil {
		log.Printf("Gagal mengambil histori retur penjualan: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if status != "" {
		if list, ok := data.([]map[string]any); ok {
			var filteredData []map[string]any
			for _, item := range list {
				if itemStatus, ok := item["status"].(string); ok && itemStatus == status {
					filteredData = append(filteredData, item)
				}
			}
			data = filteredData
			total = int64(len(filteredData))
		}
	}
	log.Printf("Mengambil %d retur penjualan dengan status %s", total, status)
	c.JSON(http.StatusOK, gin.H{"data": data, "total": total, "page": page, "size": size})
}