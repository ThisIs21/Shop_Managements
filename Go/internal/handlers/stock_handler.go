package handlers

import (
	"log"
	"strconv"
	"time"
    "errors"
    "net/http"

	"github.com/gin-gonic/gin"
	"app-penjualan/internal/dto"
	"app-penjualan/internal/models" // Ditambahkan
	"app-penjualan/internal/services"
	"app-penjualan/internal/utils"
	"gorm.io/gorm"
)

type StockHandler struct {
	svc *services.StockService
}

func NewStockHandler(s *services.StockService) *StockHandler {
	if s == nil {
		log.Fatalf("StockService is nil during handler initialization")
	}
	return &StockHandler{svc: s}
}

func (h *StockHandler) Register(rg *gin.RouterGroup) {
	if rg == nil {
		log.Fatalf("RouterGroup is nil during handler registration")
	}
	rg.POST("/stock-opnames", h.Create)    // role GUDANG
	rg.GET("/stock-opnames/history", h.List) // role GUDANG, OWNER
}

func (h *StockHandler) Create(c *gin.Context) {
	var req dto.CreateOpnameReq

	// 1. Binding dan Validasi DTO
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Gagal binding JSON untuk CreateOpname: %v, body: %v", err, c.Request.Body)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}
	if err := utils.Validate.Struct(req); err != nil {
		log.Printf("Validasi gagal untuk CreateOpname: %v, data: %+v", err, req)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed: " + err.Error()})
		return
	}

	// 2. Parsing Tanggal
	parsedDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		log.Printf("Tanggal tidak valid untuk CreateOpname: %v, input: %s", err, req.Date)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Expected YYYY-MM-DD."})
		return
	}

	// 3. Siapkan DTO untuk Service
	serviceReq := dto.CreateOpnameReqService{
		Date:  parsedDate,
		Note:  req.Note,
		Items: req.Items,
	}

	// 4. Ambil UID dari context
	uidInterface, exists := c.Get("uid")
	if !exists {
		log.Printf("UID tidak ditemukan di context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	uid, ok := uidInterface.(uint)
	if !ok {
		log.Printf("UID bukan tipe uint, got type %T with value %v", uidInterface, uidInterface)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID format"})
		return
	}

	// 5. Panggil Service
	out, err := h.svc.DoStockOpname(uid, serviceReq)
	if err != nil {
		log.Printf("Gagal menyimpan stock opname untuk UID %d: %v", uid, err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Related record not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stock opname: " + err.Error()})
		}
		return
	}

	log.Printf("Stock opname berhasil dibuat untuk UID %d, ID: %d", uid, out.ID)
	c.JSON(http.StatusCreated, gin.H{"data": out, "success": true})
}

// List mendukung params (hanya limit untuk kompatibilitas dengan service)
func (h *StockHandler) List(c *gin.Context) {
	// 1. Parsing Parameter
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
		log.Printf("Parameter limit tidak valid, default ke 10: %v", err)
	}

	// 2. Call Service (diasumsikan hanya menerima limit)
	opnames, err := h.svc.ListStockOpnames(limit)
	if err != nil {
		log.Printf("Gagal mengambil daftar stock opnames: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{"data": []models.StockOpname{}, "total": 0, "limit": limit, "success": true})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stock opnames: " + err.Error()})
		}
		return
	}

	// Hitung total secara manual (jika service tidak mengembalikannya)
	total := int64(len(opnames))

	log.Printf("Mengambil %d stock opnames dengan limit %d", total, limit)
	c.JSON(http.StatusOK, gin.H{"data": opnames, "total": total, "limit": limit, "success": true})
}