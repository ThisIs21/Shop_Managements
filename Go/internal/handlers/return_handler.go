package handlers

import (
    "net/http"
    "strconv"
    "time"
    "log"
    "errors"

    "github.com/gin-gonic/gin"
    "app-penjualan/internal/middlewares"
    "app-penjualan/internal/services"
    "gorm.io/gorm"
)

type ReturnHandler struct {
    svc *services.ReturnService
}

func NewReturnHandler(db *gorm.DB) *ReturnHandler {
    return &ReturnHandler{svc: services.NewReturnService(db)}
}

/* ===== List Returns ===== */
func (h *ReturnHandler) ListSaleReturns(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
    search := c.DefaultQuery("q", "")
    status := c.DefaultQuery("status", "") // Filter by status: PENDING, APPROVED, REJECTED

    log.Printf("Mengambil daftar sale returns - page: %d, size: %d, search: %s, status: %s", page, size, search, status)
    returns, total, err := h.svc.ListSaleReturns(page, size, search, status)
    if err != nil {
        log.Printf("Gagal mengambil daftar sale returns: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    log.Printf("Mengambil %d sale returns", total)
    c.JSON(http.StatusOK, gin.H{
        "data":  returns,
        "total": total,
        "page":  page,
        "size":  size,
    })
}

func (h *ReturnHandler) ListPurchaseReturns(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
    search := c.DefaultQuery("q", "")
    status := c.DefaultQuery("status", "") // Filter by status: PENDING, APPROVED, REJECTED

    log.Printf("Mengambil daftar purchase returns - page: %d, size: %d, search: %s, status: %s", page, size, search, status)
    returns, total, err := h.svc.ListPurchaseReturns(page, size, search, status)
    if err != nil {
        log.Printf("Gagal mengambil daftar purchase returns: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    log.Printf("Mengambil %d purchase returns", total)
    c.JSON(http.StatusOK, gin.H{
        "data":  returns,
        "total": total,
        "page":  page,
        "size":  size,
    })
}

func (h *ReturnHandler) GetSaleReturnByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Printf("ID retur tidak valid: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid return ID"})
        return
    }
    includeUser := c.DefaultQuery("include_user", "false") == "true"
    ret, err := h.svc.GetSaleReturnByID(uint(id), includeUser)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            log.Printf("Retur ID %d tidak ditemukan", id)
            c.JSON(http.StatusNotFound, gin.H{"error": "Retur tidak ditemukan"})
        } else {
            log.Printf("Gagal mengambil retur ID %d: %v", id, err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses permintaan"})
        }
        return
    }
    log.Printf("Mengambil detail retur ID %d", id)
    c.JSON(http.StatusOK, ret)
}

func (h *ReturnHandler) GetPurchaseReturnByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Printf("ID retur pembelian tidak valid: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid return ID"})
        return
    }
    includeApprover := c.DefaultQuery("include_approver", "false") == "true"
    ret, err := h.svc.GetPurchaseReturnByID(uint(id), includeApprover)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            log.Printf("Retur pembelian ID %d tidak ditemukan", id)
            c.JSON(http.StatusNotFound, gin.H{"error": "Retur tidak ditemukan"})
        } else {
            log.Printf("Gagal mengambil retur pembelian ID %d: %v", id, err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memproses permintaan"})
        }
        return
    }
    log.Printf("Mengambil detail retur pembelian ID %d", id)
    c.JSON(http.StatusOK, ret)
}

/* ===== Retur Penjualan ===== */
func (h *ReturnHandler) CreateSaleReturn(c *gin.Context) {
    userID := middlewares.MustUserID(c)
    var dto services.CreateSaleReturnDTO
    if err := c.ShouldBindJSON(&dto); err != nil {
        log.Printf("Gagal bind JSON untuk retur penjualan: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data: " + err.Error()})
        return
    }
    // Parse date dari string ke time.Time
    var date time.Time
    if dto.Date == "" {
        date = time.Now()
    } else {
        // Daftar layout yang didukung
        layouts := []string{
            "02/01/2006",           // DD/MM/YYYY
            "2006-01-02",           // YYYY-MM-DD
            time.RFC3339,           // ISO 8601, e.g., 2025-10-02T03:23:09.009Z
        }
        var err error
        for _, layout := range layouts {
            date, err = time.Parse(layout, dto.Date)
            if err == nil {
                break // Berhenti jika berhasil
            }
        }
        if err != nil {
            log.Printf("Gagal parse tanggal %s dengan semua format: %v", dto.Date, err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use DD/MM/YYYY, YYYY-MM-DD, or ISO 8601 (e.g., 2025-10-02T03:23:09Z)"})
            return
        }
    }
    // Buat DTO baru dengan time.Time diformat ke string YYYY-MM-DD untuk service
    newDTO := services.CreateSaleReturnDTO{
        SaleID: dto.SaleID,
        Date:   date.Format("2006-01-02"), // Selalu kirim YYYY-MM-DD ke service
        Items:  dto.Items,
        Notes:  dto.Notes,
    }
    log.Printf("Membuat retur penjualan untuk user ID %d, sale ID %d, date: %s", userID, dto.SaleID, newDTO.Date)
    ret, err := h.svc.CreateSaleReturn(userID, newDTO)
    if err != nil {
        log.Printf("Gagal membuat retur penjualan: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    log.Printf("Retur penjualan ID %d berhasil dibuat", ret.ID)
    c.JSON(http.StatusCreated, ret)
}

func (h *ReturnHandler) ApproveSaleReturn(c *gin.Context) {
    approver := middlewares.MustUserID(c)
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Printf("ID retur penjualan tidak valid: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid return ID"})
        return
    }
    role, exists := c.Get("role")
    if !exists || role != "KEPALA_GUDANG" {
        log.Printf("Hanya Kepala Gudang yang dapat menyetujui retur penjualan ID %d", id)
        c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak, hanya Kepala Gudang yang dapat menyetujui"})
        return
    }
    log.Printf("Menyetujui retur penjualan ID %d oleh user ID %d", id, approver)
    if err := h.svc.ApproveSaleReturn(uint(id), approver); err != nil {
        log.Printf("Gagal menyetujui retur penjualan ID %d: %v", id, err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    log.Printf("Retur penjualan ID %d disetujui", id)
    c.JSON(http.StatusOK, gin.H{"message": "Retur penjualan disetujui"})
}

func (h *ReturnHandler) RejectSaleReturn(c *gin.Context) {
    approver := middlewares.MustUserID(c)
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Printf("ID retur penjualan tidak valid: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid return ID"})
        return
    }
    role, exists := c.Get("role")
    if !exists || role != "KEPALA_GUDANG" {
        log.Printf("Hanya Kepala Gudang yang dapat menolak retur penjualan ID %d", id)
        c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak, hanya Kepala Gudang yang dapat menolak"})
        return
    }
    var dto struct {
        Reason string `json:"reason" binding:"required"`
    }
    if err := c.ShouldBindJSON(&dto); err != nil {
        log.Printf("Gagal bind JSON untuk penolakan: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Alasan penolakan wajib diisi: " + err.Error()})
        return
    }
    log.Printf("Menolak retur penjualan ID %d oleh user ID %d, alasan: %s", id, approver, dto.Reason)
    if err := h.svc.RejectSaleReturn(uint(id), approver, dto.Reason); err != nil {
        log.Printf("Gagal menolak retur penjualan ID %d: %v", id, err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    log.Printf("Retur penjualan ID %d ditolak", id)
    c.JSON(http.StatusOK, gin.H{"message": "Retur penjualan ditolak"})
}

/* ===== Retur Pembelian ===== */
func (h *ReturnHandler) CreatePurchaseReturn(c *gin.Context) {
    userID := middlewares.MustUserID(c)
    var dto services.CreatePurchaseReturnDTO
    if err := c.ShouldBindJSON(&dto); err != nil {
        log.Printf("Gagal bind JSON untuk retur pembelian: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data: " + err.Error()})
        return
    }
    // Parse date dari string ke time.Time
    var date time.Time
    if dto.Date == "" {
        date = time.Now()
    } else {
        parsedDate, err := time.Parse("2006-01-02", dto.Date)
        if err != nil {
            log.Printf("Gagal parse tanggal %s: %v", dto.Date, err)
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use YYYY-MM-DD"})
            return
        }
        date = parsedDate
    }
    // Buat DTO baru dengan time.Time
    newDTO := services.CreatePurchaseReturnDTO{
        PurchaseID: dto.PurchaseID,
        Date:       date.Format("2006-01-02"), // Simpan sebagai string jika service mengharapkannya
        Items:      dto.Items,
        Notes:      dto.Notes,
    }
    log.Printf("Membuat retur pembelian untuk user ID %d, purchase ID %d", userID, dto.PurchaseID)
    ret, err := h.svc.CreatePurchaseReturn(userID, newDTO)
    if err != nil {
        log.Printf("Gagal membuat retur pembelian: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    log.Printf("Retur pembelian ID %d berhasil dibuat", ret.ID)
    c.JSON(http.StatusCreated, ret)
}

func (h *ReturnHandler) ApprovePurchaseReturn(c *gin.Context) {
    approver := middlewares.MustUserID(c)
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Printf("ID retur pembelian tidak valid: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid return ID"})
        return
    }
    role, exists := c.Get("role")
    if !exists || role != "KEPALA_GUDANG" {
        log.Printf("Hanya Kepala Gudang yang dapat menyetujui retur pembelian ID %d", id)
        c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak, hanya Kepala Gudang yang dapat menyetujui"})
        return
    }
    log.Printf("Menyetujui retur pembelian ID %d oleh user ID %d", id, approver)
    if err := h.svc.ApprovePurchaseReturn(uint(id), approver); err != nil {
        log.Printf("Gagal menyetujui retur pembelian ID %d: %v", id, err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    log.Printf("Retur pembelian ID %d disetujui", id)
    c.JSON(http.StatusOK, gin.H{"message": "Retur pembelian disetujui"})
}

func (h *ReturnHandler) RejectPurchaseReturn(c *gin.Context) {
    approver := middlewares.MustUserID(c)
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        log.Printf("ID retur pembelian tidak valid: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid return ID"})
        return
    }
    role, exists := c.Get("role")
    if !exists || role != "KEPALA_GUDANG" {
        log.Printf("Hanya Kepala Gudang yang dapat menolak retur pembelian ID %d", id)
        c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak, hanya Kepala Gudang yang dapat menolak"})
        return
    }
    var dto struct {
        Reason string `json:"reason" binding:"required"`
    }
    if err := c.ShouldBindJSON(&dto); err != nil {
        log.Printf("Gagal bind JSON untuk penolakan: %v", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Alasan penolakan wajib diisi: " + err.Error()})
        return
    }
    log.Printf("Menolak retur pembelian ID %d oleh user ID %d, alasan: %s", id, approver, dto.Reason)
    if err := h.svc.RejectPurchaseReturn(uint(id), approver, dto.Reason); err != nil {
        log.Printf("Gagal menolak retur pembelian ID %d: %v", id, err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    log.Printf("Retur pembelian ID %d ditolak", id)
    c.JSON(http.StatusOK, gin.H{"message": "Retur pembelian ditolak"})
}