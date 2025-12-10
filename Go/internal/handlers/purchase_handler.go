package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"app-penjualan/internal/dto"
	"app-penjualan/internal/services"
	"app-penjualan/internal/utils"
	"gorm.io/gorm"
)

type PurchaseHandler struct{ svc *services.PurchaseService }

func NewPurchaseHandler(s *services.PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{svc: s}
}

func (h *PurchaseHandler) Register(rg *gin.RouterGroup) {
	rg.POST("/purchases", h.Create)              // role PEMBELIAN
	rg.POST("/purchases/:id/approve", h.Approve) // role KEPALA_GUDANG
	rg.POST("/purchases/:id/reject", h.Reject)

	// ✅ Tambahan: endpoint GET list purchases
	
}

// ✅ handler baru untuk list purchase
func (h *PurchaseHandler) List(c *gin.Context) {
    pageStr := c.DefaultQuery("page", "1")
    sizeStr := c.DefaultQuery("size", "20")
    search := c.Query("q")
    status := c.DefaultQuery("status", "SUBMITTED")

    // ✅ Tambahan
    from := c.Query("from")
    to := c.Query("to")

    page, err := strconv.Atoi(pageStr)
    if err != nil {
        utils.BadRequest(c, "Invalid page parameter")
        return
    }
    size, err := strconv.Atoi(sizeStr)
    if err != nil {
        utils.BadRequest(c, "Invalid size parameter")
        return
    }
    if page < 1 {
        page = 1
    }
    if size < 1 || size > 100 {
        size = 20
    }

    // ✅ panggil FindAll baru dengan filter tanggal
    out, err := h.svc.FindAll(page, size, search, status, from, to)
    if err != nil {
        utils.ServerError(c, err)
        return
    }
    utils.OK(c, gin.H{"data": out, "success": true})
}

func (h *PurchaseHandler) Create(c *gin.Context) {
	var req dto.CreatePurchaseReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	if err := utils.Validate.Struct(req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}
	uid := c.GetUint("uid")
	out, err := h.svc.Create(uid, req)
	if err != nil {
		utils.ServerError(c, err)
		return
	}
	utils.Created(c, out)
}

func (h *PurchaseHandler) Approve(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequest(c, "Invalid purchase ID")
		return
	}
	out, err := h.svc.Approve(c.GetUint("uid"), uint(id), true)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.NotFound(c, "Purchase not found")
			return
		}
		utils.BadRequest(c, err.Error())
		return
	}
	utils.OK(c, out)
}

func (h *PurchaseHandler) Reject(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequest(c, "Invalid purchase ID")
		return
	}
	out, err := h.svc.Reject(c.GetUint("uid"), uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.NotFound(c, "Purchase not found")
			return
		}
		utils.BadRequest(c, err.Error())
		return
	}
	utils.OK(c, out)
}

func (h *PurchaseHandler) Receive(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.BadRequest(c, "Invalid purchase ID")
		return
	}
	// Panggil service Receive
	out, err := h.svc.Receive(c.GetUint("uid"), uint(id)) 
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.NotFound(c, "Approved Purchase not found")
			return
		}
		utils.ServerError(c, err)
		return
	}
	utils.OK(c, out)
}