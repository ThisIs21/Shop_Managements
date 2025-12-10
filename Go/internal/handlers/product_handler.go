package handlers

import (
    "log"
    "strconv"
    "errors"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "app-penjualan/internal/models"
    "app-penjualan/internal/services"
    "app-penjualan/internal/utils"
)

type ProductHandler struct{ svc *services.ProductService }

func NewProductHandler(s *services.ProductService) *ProductHandler {
    return &ProductHandler{svc: s}
}

func (h *ProductHandler) Register(rg *gin.RouterGroup) {
    rg.GET("/products", h.List)
    rg.GET("/products/:id", h.Get)
    rg.POST("/products", h.Create)
    rg.PUT("/products/:id", h.Update)
    rg.DELETE("/products/:id", h.Delete)
}

func (h *ProductHandler) List(c *gin.Context) {
    search := c.Query("search")
    category := c.Query("category")
    minPrice := c.Query("min_price")
    maxPrice := c.Query("max_price")
    out, err := h.svc.ListFiltered(search, category, minPrice, maxPrice)
    if err != nil {
        log.Printf("Error listing products: %v", err)
        utils.ServerError(c, err)
        return
    }
    utils.OK(c, out)
}

func (h *ProductHandler) Get(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil || id <= 0 {
        log.Printf("Invalid product ID: %v", c.Param("id"))
        utils.BadRequest(c, "ID produk tidak valid")
        return
    }
    out, err := h.svc.Get(uint(id))
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            utils.NotFound(c, "Produk tidak ditemukan")
            return
        }
        log.Printf("Error getting product: %v", err)
        utils.ServerError(c, err)
        return
    }
    utils.OK(c, out)
}

func (h *ProductHandler) Create(c *gin.Context) {
    var m models.Product
    if err := c.ShouldBindJSON(&m); err != nil {
        log.Printf("Invalid request body: %v", err)
        utils.BadRequest(c, err.Error())
        return
    }
    if err := h.svc.Create(&m); err != nil {
        log.Printf("Error creating product: %v", err)
        utils.ServerError(c, err)
        return
    }
    utils.Created(c, m)
}

func (h *ProductHandler) Update(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil || id <= 0 {
        log.Printf("Invalid product ID: %v", c.Param("id"))
        utils.BadRequest(c, "ID produk tidak valid")
        return
    }
    var in models.Product
    if err := c.ShouldBindJSON(&in); err != nil {
        log.Printf("Invalid request body: %v", err)
        utils.BadRequest(c, err.Error())
        return
    }
    log.Printf("Updating product ID %d with data: %+v", id, in)
    out, err := h.svc.Update(uint(id), &in)
    if err != nil {
        log.Printf("Error updating product: %v", err)
        if errors.Is(err, gorm.ErrRecordNotFound) {
            utils.NotFound(c, "Produk tidak ditemukan")
            return
        }
        utils.ServerError(c, err)
        return
    }
    utils.OK(c, out)
}

func (h *ProductHandler) Delete(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil || id <= 0 {
        log.Printf("Invalid product ID: %v", c.Param("id"))
        utils.BadRequest(c, "ID produk tidak valid")
        return
    }
    log.Printf("Deleting product ID %d", id)
    if err := h.svc.Delete(uint(id)); err != nil {
        log.Printf("Error deleting product: %v", err)
        if errors.Is(err, gorm.ErrRecordNotFound) {
            utils.NotFound(c, "Produk tidak ditemukan")
            return
        }
        utils.ServerError(c, err)
        return
    }
    utils.OK(c, gin.H{"deleted": id})
}