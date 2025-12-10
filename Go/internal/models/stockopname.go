package models

import "time"

type StockOpname struct {
    ID     uint              `gorm:"primaryKey" json:"id"`
    UserID uint              `json:"userId" gorm:"column:user_id"`
    User   User              `json:"user"`
    Date   time.Time         `json:"date"`
    Note   string            `gorm:"size:255" json:"notes"`
    Items  []StockOpnameItem `json:"items"`
    Audit
}

type StockOpnameItem struct {
    ID            uint    `gorm:"primaryKey" json:"id"`
    StockOpnameID uint    `json:"stockOpnameId" gorm:"column:stock_opname_id"`
    ProductID     uint    `json:"productId" gorm:"column:product_id"`
    Product       Product `json:"product"`
    QtySystem     int     `json:"systemStock" gorm:"column:qty_system"`     // Fix: Tambah JSON tag jika frontend expect camelCase
    QtyPhysical   int     `json:"physicalStock" gorm:"column:qty_physical"` // Fix: Tambah JSON tag
}