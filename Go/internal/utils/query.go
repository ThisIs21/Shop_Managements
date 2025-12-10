package utils

import (
	"fmt"
	"log" // Opsional: Tambah untuk logging error jika diperlukan
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DateRange mendefinisikan rentang waktu menggunakan pointer ke time.Time.
type DateRange struct {
	From *time.Time
	To   *time.Time
}

// Global variable untuk lokasi waktu Asia/Jakarta (WIB)
var WIB = func() *time.Location {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Printf("Warning: Failed to load WIB location, using UTC: %v", err)
		return time.UTC // Fallback ke UTC jika error
	}
	return loc
}()

// ParseDateRange mengambil start_date dan end_date dari gin.Context dan mengonversinya ke DateRange dengan zona WIB
// Return: DateRange dan error jika parsing gagal.
func ParseDateRange(c *gin.Context) (DateRange, error) {
	var dr DateRange

	startDate := c.Query("start_date")
	if startDate != "" {
		// Parsing: YYYY-MM-DD
		from, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			return dr, fmt.Errorf("invalid start_date format, use YYYY-MM-DD: %v", err)
		}
		
		// Set waktu awal hari (00:00:00) di zona WIB
		from = time.Date(from.Year(), from.Month(), from.Day(), 0, 0, 0, 0, WIB)
		dr.From = &from
	}

	endDate := c.Query("end_date")
	if endDate != "" {
		// Parsing: YYYY-MM-DD
		to, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			return dr, fmt.Errorf("invalid end_date format, use YYYY-MM-DD: %v", err)
		}
		
		// Set waktu akhir hari (23:59:59) di zona WIB
		// Menggunakan AddDate(0, 0, 1) untuk pindah ke tanggal berikutnya, lalu dikurangi 1 nanosecond (lebih presisi)
		to = time.Date(to.Year(), to.Month(), to.Day(), 0, 0, 0, 0, WIB).AddDate(0, 0, 1).Add(-time.Nanosecond)
		dr.To = &to
	}

	return dr, nil
}

// ApplyRange adalah helper GORM untuk menambahkan filter WHERE berdasarkan DateRange.
// Fungsi ini penting agar HistoryService dan ReportService tidak perlu mendefinisikannya lagi.
// Parameter: q (query GORM), dr (DateRange), field (nama kolom tanggal di DB, e.g., "date" atau "created_at")
func ApplyRange(q *gorm.DB, dr DateRange, field string) *gorm.DB {
	// Pastikan field adalah nama kolom tanggal di database (misal: "created_at" atau "date")
	if dr.From != nil {
		q = q.Where(field+" >= ?", *dr.From)
	}
	if dr.To != nil {
		q = q.Where(field+" <= ?", *dr.To)
	}
	return q
}

// Paginate adalah GORM Scope untuk menerapkan paginasi (diambil dari solusi sebelumnya)
// Return: func(db *gorm.DB) *gorm.DB yang bisa dipakai sebagai Scope.
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}