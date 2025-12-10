package models

type Category struct {
    ID   uint   `gorm:"primaryKey;column:id" json:"id"`
    Name string `gorm:"size:80;uniqueIndex;not null;column:name" json:"name"`
    Audit
}

type Unit struct {
    ID   uint   `gorm:"primaryKey;column:id" json:"id"`
    Name string `gorm:"size:30;uniqueIndex;not null;column:name" json:"name"`
    Audit
}

type Supplier struct {
    ID      uint   `gorm:"primaryKey;column:id" json:"id"`
    Name    string `gorm:"size:120;not null;column:name" json:"name"`
    Contact string `gorm:"size:120;column:contact" json:"contact"`
    Address string `gorm:"size:255;column:address" json:"address"`
    Audit
}

type Customer struct {
    ID      uint   `gorm:"primaryKey;column:id" json:"id"`
    Name    string `gorm:"size:120;not null;column:name" json:"name"`
    Contact string `gorm:"size:120;column:contact" json:"contact"`
    Audit
}

type VoucherType string

const (
    VoucherPercent VoucherType = "PERCENT"
    VoucherAmount  VoucherType = "AMOUNT"
)

type Voucher struct {
    ID     uint        `gorm:"primaryKey;column:id"`
    Code   string      `gorm:"size:40;uniqueIndex;not null;column:code"`
    Value  float64     `gorm:"not null;column:value"` // percent or amount
    Type   VoucherType `gorm:"size:16;not null;column:type"`
    Active bool        `gorm:"default:true;column:active"`
    Audit
}


type Product struct {
    ID         uint    `gorm:"primaryKey;column:id" json:"id"`
    Name       string  `gorm:"size:120;uniqueIndex;not null;column:name" json:"name" binding:"required"`
    CategoryID uint    `gorm:"not null;column:category_id" json:"category_id" binding:"required"`
    Category   Category `json:"category"`
    UnitID     uint    `gorm:"not null;column:unit_id" json:"unit_id" binding:"required"`
    Unit       Unit    `json:"unit"`
    CostPrice  float64 `gorm:"type:DECIMAL(18,2);not null;column:cost_price" json:"cost_price" binding:"required"`
    SellPrice  float64 `gorm:"type:DECIMAL(18,2);not null;column:sell_price" json:"sell_price" binding:"required"`
    Stock      int     `gorm:"not null;default:0;column:stock" json:"stock"`
    SupplierID uint    `gorm:"not null;column:supplier_id" json:"supplier_id" binding:"required"`
    Supplier   Supplier `json:"supplier"`
    Audit
}


type OrderItem struct {
    ID          uint   `gorm:"primaryKey;column:id" json:"id"`
    OrderID     uint   `gorm:"not null;column:order_id" json:"order_id"`
    ProductID   uint   `gorm:"not null;column:product_id" json:"product_id"`
    // Kolom lain sesuai tabelmu
}