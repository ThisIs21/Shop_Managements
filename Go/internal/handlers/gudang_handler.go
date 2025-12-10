package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"app-penjualan/internal/dto"
	"app-penjualan/internal/services"
	"app-penjualan/internal/utils"
	
)

type GudangHandler struct{ svc *services.GudangService }

func NewGudangHandler(s *services.GudangService) *GudangHandler { 
	return &GudangHandler{svc: s} 
}

func (h *GudangHandler) Register(rg *gin.RouterGroup) {
	// QC + Receive untuk Purchase
	rg.POST("/purchases/:id/qc-receive", h.QcAndReceive)
	rg.GET("/purchases", h.ListApprovedPurchases)
	
	// QC + Receive untuk Sale Return (retur dari kasir)
	rg.POST("/sale-returns/:id/qc-receive", h.QcSaleReturn)

	// List Sale Return yang sudah disetujui kepala gudang dan siap diterima
	rg.GET("/sale-returns", h.ListApprovedSaleReturns)

}

// ==================== HANDLERS ====================

// QcAndReceive: proses QC + penerimaan barang PO
func (h *GudangHandler) QcAndReceive(c *gin.Context) {
	purchaseID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid Purchase ID"})
		return
	}

	var req dto.CreateQcReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("uid")
	if err := h.svc.RecordQcAndReceive(uint(purchaseID), req, userID); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "QC report recorded and stock updated"})
}

// QcSaleReturn: proses QC untuk retur penjualan (dari kasir, sudah disetujui kepala gudang)
func (h *GudangHandler) QcSaleReturn(c *gin.Context) {
	saleReturnID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid SaleReturn ID"})
		return
	}

	var req dto.CreateQcReportRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("uid")
	if err := h.svc.RecordQcSaleReturn(uint(saleReturnID), req, userID); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Sale return received and QC recorded"})
}

// ListApprovedSaleReturns: tampilkan data retur dari kasir yang sudah APPROVED kepala gudang
func (h *GudangHandler) ListApprovedSaleReturns(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	dr, _ := utils.ParseDateRange(c)

	returns, total, err := h.svc.ListApprovedSaleReturns(page, size, dr)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"total": total,
		"data":  returns,
		"page":  page,
		"size":  size,
	})
}

// ListApprovedPurchases: tampilkan daftar Purchase Order yang sudah APPROVED, siap diterima gudang
func (h *GudangHandler) ListApprovedPurchases(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	purchases, total, err := h.svc.ListApprovedPurchases(page, size)
	if err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	utils.OK(c, gin.H{
		"total": total,
		"data":  purchases,
		"page":  page,
		"size":  size,
	})
}
