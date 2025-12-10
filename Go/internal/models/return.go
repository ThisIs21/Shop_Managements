package models

import (
	"time"

)
// PurchaseReturn represents a purchase return transaction
type PurchaseReturn struct {
	ID            uint              `gorm:"primaryKey"`
	PurchaseID    uint              `gorm:"not null"`
	UserID        uint              `gorm:"not null"` // Dibuat oleh
	Total         float64           `gorm:"type:decimal(10,2)"`
	Date          time.Time         `gorm:"not null"`
	ApprovedBy    *uint             `gorm:"default:null"`
	Approver      *User             `gorm:"foreignKey:ApprovedBy"`
	ApprovedAt    *time.Time        `gorm:"default:null"`
	// *** PERUBAHAN DI SINI: Tambahkan 'QC_DEFECT' ke enum ***
	Status        string            `gorm:"type:enum('PENDING','APPROVED','REJECTED','QC_DEFECT');default:'PENDING'"`
	Items         []PurchaseReturnItem `gorm:"foreignKey:PurchaseReturnID"`
	Audit
}

// PurchaseReturnItem represents an item in a purchase return
type PurchaseReturnItem struct {
    ID            uint      `gorm:"primaryKey"`
    PurchaseReturnID uint   `gorm:"not null"`
    ProductID     uint      `gorm:"not null"`
    Product       Product   `gorm:"foreignKey:ProductID"`
    Qty           int       `gorm:"not null"`
    Price         float64   `gorm:"type:decimal(10,2);not null"`
    Subtotal      float64   `gorm:"type:decimal(10,2);not null"`
    Audit
}

// SaleReturn represents a sales return transaction
type SaleReturn struct {
    ID         uint            `gorm:"primaryKey" json:"id"`
    SaleID     uint            `gorm:"not null" json:"sale_id"`
    UserID     uint            `gorm:"not null" json:"user_id"`
    Sale       Sale            `gorm:"foreignKey:SaleID" json:"sale"`
    User       User            `gorm:"foreignKey:UserID" json:"user"`
    Total      float64         `gorm:"type:decimal(10,2)" json:"total"`
    Date       time.Time       `gorm:"not null" json:"date"`
    ApprovedBy *uint           `json:"approved_by"`
    ApprovedAt *time.Time      `json:"approved_at"`
    Status     string          `gorm:"type:enum('PENDING','APPROVED','REJECTED');default:'PENDING'" json:"status"`
    Notes      string          `gorm:"type:text" json:"notes"`
    Items      []SaleReturnItem `gorm:"foreignKey:SaleReturnID" json:"items"`
    Audit
}

// SaleReturnItem represents an item in a sales return
type SaleReturnItem struct {
    ID           uint      `gorm:"primaryKey" json:"id"`
    SaleReturnID uint      `gorm:"not null" json:"sale_return_id"`
    ProductID    uint      `gorm:"not null" json:"product_id"`
    Product      Product   `gorm:"foreignKey:ProductID" json:"product"`
    Qty          int       `gorm:"not null" json:"qty"`
    Price        float64   `gorm:"type:decimal(10,2)" json:"price"`
    Subtotal     float64   `gorm:"type:decimal(10,2)" json:"subtotal"`
    Reason       string    `gorm:"type:text" json:"reason"`
    Audit
}

