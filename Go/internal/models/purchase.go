package models

import "time"

type PurchaseStatus string

const (
    PurchaseDraft     PurchaseStatus = "DRAFT"
    PurchaseSubmitted PurchaseStatus = "SUBMITTED"
    PurchaseApproved  PurchaseStatus = "APPROVED"
    PurchaseReceived  PurchaseStatus = "RECEIVED"
    PurchaseRejected  PurchaseStatus = "REJECTED"
)

type Purchase struct {
    ID          uint           `gorm:"primaryKey;column:id"`
    SupplierID  uint           `gorm:"not null;column:supplier_id"`
    Supplier    Supplier
    UserID      uint           `gorm:"column:user_id"`
    User        User
    ApprovedBy  *uint          `gorm:"column:approved_by"`
    Approver    *User          `gorm:"foreignKey:ApprovedBy"`
    ApprovedAt   *time.Time     `gorm:"column:approved_at"`
    
	// ✅ Tambahan Kolom Penerimaan Barang (Goods Receipt)
    
	ReceivedBy   *uint          `gorm:"column:received_by"`
	Receiver     *User          `gorm:"foreignKey:ReceivedBy"`
	ReceivedAt   *time.Time     `gorm:"column:received_at"`
    Total       float64        `gorm:"not null;default:0;column:total"`
    Status      PurchaseStatus `gorm:"size:16;not null;default:'DRAFT';column:status"`
    Date        time.Time      `gorm:"not null;column:date"`
    Items       []PurchaseItem `gorm:"foreignKey:PurchaseID"`
    Audit
}

type PurchaseItem struct {
    ID          uint    `gorm:"primaryKey;column:id"`
    PurchaseID  uint    `gorm:"column:purchase_id"`
    ProductID   uint    `gorm:"column:product_id"`
    Product     Product
    Qty         int     `gorm:"column:qty"`
    Price       float64 `gorm:"column:price"`
    Subtotal    float64 `gorm:"column:subtotal"`
} 