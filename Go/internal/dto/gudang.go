package dto

// --- DTO untuk Hasil QC (Pre-Good Receipt) ---
type QcItemRequest struct {
	PurchaseItemID uint `json:"purchase_item_id" binding:"required"`
	QtyGood        int  `json:"qty_good" binding:"min=0"`         // Kuantitas Lulus QC (Siap Masuk Stok)
	QtyDamaged     int  `json:"qty_damaged" binding:"min=0"`               // Kuantitas Rusak (Akan diretur/dicatat rugi)
	Note           string `json:"note"`                                      // Catatan QC (e.g., "Gores di 3 unit")
}

type CreateQcReportRequest struct {
	Items []QcItemRequest `json:"items" binding:"required,dive"`
}