package handlers

import (
    "net/http"
    "strconv"
    "log"

    "github.com/gin-gonic/gin"
    "app-penjualan/internal/dto"
    "app-penjualan/internal/services"
    "app-penjualan/internal/utils"
)

type SaleHandler struct {
    svc *services.SaleService
}

func NewSaleHandler(s *services.SaleService) *SaleHandler {
    return &SaleHandler{svc: s}
}

func (h *SaleHandler) Register(rg *gin.RouterGroup) {
    rg.POST("/sales", h.Create)
    rg.GET("/sales", h.List)
    rg.GET("/sales/:id", h.GetByID)
}

func (h *SaleHandler) Create(c *gin.Context) {
    var req dto.CreateSaleReq
    if err := c.ShouldBindJSON(&req); err != nil {
        log.Printf("Gagal bind JSON: %v", err)
        utils.BadRequest(c, err.Error())
        return
    }
    if err := utils.Validate.Struct(req); err != nil {
        log.Printf("Validasi gagal: %v", err)
        utils.BadRequest(c, err.Error())
        return
    }

    uid := c.GetUint("uid")
    log.Printf("Membuat transaksi untuk user ID: %d, request: %+v", uid, req)
    sale, err := h.svc.Create(uid, req)
    if err != nil {
        log.Printf("Gagal membuat transaksi: %v", err)
        utils.BadRequest(c, err.Error())
        return
    }

    // Preload details
    sale, err = h.svc.LoadDetails(sale.ID)
    if err != nil {
        log.Printf("Gagal load detail transaksi ID %d: %v", sale.ID, err)
        utils.BadRequest(c, "Gagal load detail transaksi")
        return
    }

    log.Printf("Transaksi ID %d berhasil dibuat", sale.ID)
    utils.Created(c, sale)
}

func (h *SaleHandler) List(c *gin.Context) {
    sales, err := h.svc.List()
    if err != nil {
        log.Printf("Gagal mengambil daftar penjualan: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    log.Printf("Mengambil %d penjualan", len(sales))
    c.JSON(http.StatusOK, gin.H{"data": sales})
}

func (h *SaleHandler) GetByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Printf("ID penjualan tidak valid: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sale ID"})
        return
    }
    sale, err := h.svc.LoadDetails(uint(id))
    if err != nil {
        log.Printf("Gagal mengambil penjualan ID %d: %v", id, err)
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    log.Printf("Mengambil penjualan ID %d", id)
    c.JSON(http.StatusOK, sale)
}