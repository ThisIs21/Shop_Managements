package routes

import (
	"app-penjualan/config"
	"app-penjualan/internal/handlers"
	"app-penjualan/internal/middlewares"
	"app-penjualan/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(r *gin.Engine, db *gorm.DB, cfg *config.AppConfig) {
	// =======================
	// HEALTH CHECK
	// =======================
	r.GET("/api/health", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"db":     "down",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"db":      "connected",
			"message": "API is running 🚀",
		})
	})

	// =======================
	// AUTH (Public)
	// =======================
	authSvc := services.NewAuthService(db, cfg)
	authH := handlers.NewAuthHandler(authSvc)
	r.POST("/api/v1/auth/login", authH.Login)

	// ==========================================================
	// Protected Group (Base: /api/v1)
	// ==========================================================
	auth := r.Group("/api/v1")
	auth.Use(middlewares.JWTAuth(cfg))

	// Inisialisasi Handler yang digunakan di berbagai group
	voucherH := handlers.NewVoucherHandler(services.NewVoucherService(db))
	categoryH := handlers.NewCategoryHandler(services.NewCategoryService(db))
	unitH := handlers.NewUnitHandler(services.NewUnitService(db))
	supplierH := handlers.NewSupplierHandler(services.NewSupplierService(db))
	prodH := handlers.NewProductHandler(services.NewProductService(db))
	returnH := handlers.NewReturnHandler(db)
	reportH := handlers.NewReportHandler(services.NewReportService(db))
	historyH := handlers.NewHistoryHandler(services.NewHistoryService(db))

	// =======================
	// OWNER (Full Access)
	// =======================
	owner := auth.Group("/owner")
	owner.Use(middlewares.RequireRoles("OWNER"))
	{
		// Role & Users
		owner.GET("/roles", authH.ListRoles)
		owner.POST("/users", authH.CreateUser)
		owner.GET("/users", authH.ListUsers)
		owner.PUT("/users/:id", authH.UpdateUser)
		owner.DELETE("/users/:id", authH.DeleteUser)

		// Voucher CRUD
		owner.POST("/vouchers", voucherH.Create)
		// owner.GET("/vouchers", voucherH.List) <-- DIPINDAH ke Shared Routes
		owner.GET("/vouchers/:id", voucherH.Get)
		owner.PUT("/vouchers/:id", voucherH.Update)
		owner.DELETE("/vouchers/:id", voucherH.Delete)

		// Master Data
		owner.POST("/categories", categoryH.Create)
		owner.GET("/categories", categoryH.List)
		owner.GET("/categories/:id", categoryH.Get)
		owner.PUT("/categories/:id", categoryH.Update)
		owner.DELETE("/categories/:id", categoryH.Delete)

		owner.POST("/units", unitH.Create)
		owner.GET("/units", unitH.List)
		owner.GET("/units/:id", unitH.Get)
		owner.PUT("/units/:id", unitH.Update)
		owner.DELETE("/units/:id", unitH.Delete)

		owner.POST("/suppliers", supplierH.Create)
		owner.GET("/suppliers", supplierH.List)
		owner.GET("/suppliers/:id", supplierH.Get)
		owner.PUT("/suppliers/:id", supplierH.Update)
		owner.DELETE("/suppliers/:id", supplierH.Delete)

		// Produk CRUD
		prodH.Register(owner)

		// Semua laporan
		owner.GET("/reports", reportH.GenerateReport)
		owner.GET("/reports/inventory-detail", reportH.InventoryDetail)
		// Dashboard summary for Owner
		owner.GET("/dashboard/summary", reportH.DashboardSummary)

		// Approval Returns
		// Sale Returns Approval (OWNER)
		owner.GET("/sale-returns", returnH.ListSaleReturns)
		owner.GET("/sale-returns/:id", returnH.GetSaleReturnByID)
		owner.POST("/sale-returns/:id/approve", returnH.ApproveSaleReturn)
		owner.POST("/sale-returns/:id/reject", returnH.RejectSaleReturn)

		// Purchase Returns Approval (OWNER)
		owner.GET("/purchase-returns", returnH.ListPurchaseReturns)
		owner.GET("/purchase-returns/:id", returnH.GetPurchaseReturnByID)
		owner.POST("/purchase-returns/:id/approve", returnH.ApprovePurchaseReturn)
		owner.POST("/purchase-returns/:id/reject", returnH.RejectPurchaseReturn)
	}

	// =======================
	// SHARED ROUTES
	// =======================
	// List Vouchers (Fix 404, karena klien memanggil /api/v1/vouchers)
	auth.GET("/vouchers",
		middlewares.RequireRoles("OWNER", "KASIR"),
		voucherH.List)

	auth.GET("/products",
		middlewares.RequireRoles("OWNER", "GUDANG", "KASIR", "PEMBELIAN", "KEPALA_GUDANG"),
		prodH.List)

	auth.GET("/categories",
		middlewares.RequireRoles("OWNER", "KASIR", "PEMBELIAN", "GUDANG", "KEPALA_GUDANG"),
		categoryH.List)

	auth.GET("/suppliers",
		middlewares.RequireRoles("OWNER", "PEMBELIAN", "GUDANG", "KEPALA_GUDANG"),
		supplierH.List)

	auth.GET("/units",
		middlewares.RequireRoles("OWNER", "KASIR", "PEMBELIAN", "GUDANG", "KEPALA_GUDANG"),
		unitH.List)

	// =======================
	// KASIR
	// =======================
	kasir := auth.Group("/kasir")
	kasir.Use(middlewares.RequireRoles("KASIR", "OWNER"))
	{
		saleH := handlers.NewSaleHandler(services.NewSaleService(db))
		saleH.Register(kasir)

		kasir.POST("/sale-returns", returnH.CreateSaleReturn)
		kasir.GET("/sale-returns", returnH.ListSaleReturns)
		kasir.GET("/sale-returns/:id", returnH.GetSaleReturnByID)

		kasir.GET("/reports/sales", reportH.GenerateReport)
	}

	// =======================
	// PEMBELIAN
	// =======================
	pembelian := auth.Group("/pembelian")
	pembelian.Use(middlewares.RequireRoles("PEMBELIAN", "OWNER", "KEPALA_GUDANG"))
	{
		purchaseH := handlers.NewPurchaseHandler(services.NewPurchaseService(db))

		// Purchase Approval (Untuk PEMBELIAN yang juga bisa Approve/Reject)
		pembelian.POST("/purchases", purchaseH.Create)
		pembelian.POST("/purchases/:id/approve", purchaseH.Approve)
		pembelian.POST("/purchases/:id/reject", purchaseH.Reject)

		// History Purchase
		pembelian.GET("/purchases", historyH.Purchases)
		pembelian.GET("/purchases/:id", historyH.GetPurchaseByID)

		// Category
		pembelian.POST("/categories", categoryH.Create)
		pembelian.GET("/categories", categoryH.List)
		pembelian.GET("/categories/:id", categoryH.Get)

		pembelian.GET("/reports/purchases", reportH.GenerateReport)
	}

	// internal/routes/routes.go (Snippet fix untuk grup GUDANG - hapus call langsung stockH.ListStockOpnames(c))
	// =======================
	// GUDANG
	// =======================
	gudang := auth.Group("/gudang")
	gudang.Use(middlewares.RequireRoles("GUDANG", "OWNER"))
	{
		stockH := handlers.NewStockHandler(services.NewStockService(db))
		stockH.Register(gudang) // Ini sudah register GET /stock-opnames/history -> h.List

		prodH.Register(gudang)

		purchaseH := handlers.NewPurchaseHandler(services.NewPurchaseService(db))
		gudang.POST("/purchases/:id/receive", purchaseH.Receive) // ROLE GUDANG: Menerima barang

		gudang.GET("/reports/stocks", reportH.GenerateReport)
		gudang.GET("/reports/stock-histories", reportH.GenerateReport)

		// Retur Pembelian: CREATE/LIST/GET
		gudang.POST("/purchase-returns", returnH.CreatePurchaseReturn) // ✅ ROLE GUDANG: Membuat retur
		gudang.GET("/purchase-returns", returnH.ListPurchaseReturns)
		gudang.GET("/purchase-returns/:id", returnH.GetPurchaseReturnByID)

		gudangService := services.NewGudangService(db)
		gudangHandler := handlers.NewGudangHandler(gudangService)
		gudang.POST("/purchases/:id/qc-receive", gudangHandler.QcAndReceive)

		gudang.GET("/purchases", gudangHandler.ListApprovedPurchases)
		gudang.GET("/purchases/:id", historyH.GetPurchaseByID)

		// HAPUS jika ada: stockH.ListStockOpnames(c) // Ini salah; gunakan route aja
	}

	// =======================
	// KEPALA GUDANG
	// =======================
	kepalaGudang := auth.Group("/kepala-gudang")
	kepalaGudang.Use(middlewares.RequireRoles("KEPALA_GUDANG", "OWNER"))
	{
		purchaseH := handlers.NewPurchaseHandler(services.NewPurchaseService(db))

		// Purchase Approval/Reject
		kepalaGudang.GET("/purchases", historyH.Purchases)
		kepalaGudang.GET("/purchases/:id", historyH.GetPurchaseByID)

		kepalaGudang.POST("/purchases/:id/approve", purchaseH.Approve)
		kepalaGudang.POST("/purchases/:id/reject", purchaseH.Reject)

		// Sale Returns Approval/Reject
		kepalaGudang.GET("/sale-returns", returnH.ListSaleReturns)
		kepalaGudang.GET("/sale-returns/:id", returnH.GetSaleReturnByID)
		kepalaGudang.POST("/sale-returns/:id/approve", returnH.ApproveSaleReturn)
		kepalaGudang.POST("/sale-returns/:id/reject", returnH.RejectSaleReturn)

		// Purchase Returns Approval/Reject
		kepalaGudang.GET("/purchase-returns", returnH.ListPurchaseReturns)
		kepalaGudang.GET("/purchase-returns/:id", returnH.GetPurchaseReturnByID)
		kepalaGudang.POST("/purchase-returns/:id/approve", returnH.ApprovePurchaseReturn) // ✅ ROLE KEPALA GUDANG: Approve
		kepalaGudang.POST("/purchase-returns/:id/reject", returnH.RejectPurchaseReturn)   // ✅ ROLE KEPALA GUDANG: Reject

		kepalaGudang.GET("/reports/stocks", reportH.GenerateReport)
		kepalaGudang.GET("/reports/stock-histories", reportH.GenerateReport)
	}

	// =======================
	// HISTORY
	// =======================
	hist := auth.Group("/history")
	hist.Use(middlewares.RequireRoles("OWNER", "PEMBELIAN", "KASIR", "GUDANG", "KEPALA_GUDANG"))
	{
		hist.GET("/purchases", historyH.Purchases)
		hist.GET("/purchases/:id", historyH.GetPurchaseByID)
		hist.GET("/sales", historyH.Sales)
		hist.GET("/sales/:id", historyH.GetSaleByID)
		hist.GET("/sales/:id/pdf", historyH.GenerateSalePDF)
		hist.GET("/sales/latest", historyH.LatestSales)
		hist.GET("/stock-opnames/:id", historyH.GetStockOpnameByID)

		// History Retur Pembelian dan Penjualan (Akses untuk semua yang berkepentingan)
		hist.GET("/purchase-returns", historyH.PurchaseReturns)
		hist.GET("/sale-returns", historyH.SaleReturns)

		hist.GET("/stock-opnames", historyH.StockOpnames)
	}

	// =======================
	// CORRECTIONS
	// =======================
	corrH := handlers.NewCorrectionHandler(services.NewCorrectionService(db))
	corr := auth.Group("/corrections")
	corr.Use(middlewares.RequireRoles("OWNER", "PEMBELIAN", "KASIR", "GUDANG", "KEPALA_GUDANG"))
	{
		corr.POST("", corrH.Create)
	}
}
