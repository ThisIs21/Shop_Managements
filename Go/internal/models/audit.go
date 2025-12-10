package models

import (
    "time"
    "gorm.io/gorm" // Sebaiknya sertakan gorm untuk DeletedAt
)

type Audit struct {
    // Tambahkan `json:"-"` pada field yang hanya diatur oleh server/GORM
    CreatedAt time.Time `json:"-"` 
    UpdatedAt time.Time `json:"-"` 
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // ganti *time.Time ke gorm.DeletedAt jika pakai fitur Soft Delete
}

type CorrectionType string
const (
	CorrectionPurchase CorrectionType = "PURCHASE"
	CorrectionSale     CorrectionType = "SALE"
	CorrectionPRet     CorrectionType = "PURCHASE_RETURN"
	CorrectionSRet     CorrectionType = "SALE_RETURN"
)

type Correction struct {
	ID         uint           `gorm:"primaryKey"`
	RefTable   string         // nama tabel target (purchases/sales/dll)
	RefID      uint           // id record yang dikoreksi
	BeforeJSON string         `gorm:"type:json"`
	AfterJSON  string         `gorm:"type:json"`
	DeltaStock int            // perubahan stok (positif/negatif)
	Reason     string
	UserID     uint
	CreatedAt  time.Time
}