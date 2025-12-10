import { createRouter, createWebHistory } from "vue-router";
import LoginForm from "@/views/LoginForm.vue";
import DashboardCashier from "@/views/DashboardCashier.vue";
import CashierPage from "@/views/CashierPage.vue";
import CompletedTransactions from "@/views/CompletedTransactions.vue";
import DashboardOwner from "@/views/DashboardOwner.vue";
import FinancialReportsPage from "@/views/FinancialReportsPage.vue";
import UsersView from "@/views/UsersView.vue";
import DashboardGudang from "@/views/DashboardGudang.vue";
import DashboardKepalagudang from "@/views/DashboardKepalagudang.vue";
import DashboardPembelian from "@/views/DashboardPembelian.vue";
import ReturnPage from "@/views/ReturnPage.vue";
import ProductList from "@/views/ProductList.vue";
import CreatePurchase from "@/views/CreatePurchase.vue";
import ProductListGudang from "@/views/ProductListGudang.vue";
import ProductListPurchase from "@/views/ProductListPurchase.vue";
import PurchaseReturn from "@/views/PurchaseReturn.vue";
import SalesReturnKepalaGudang from "@/views/SalesReturnKepalaGudang.vue";
import PurchaseOrderKepalaGudang from "@/views/PurchaseOrderKepalaGudang.vue";
import PurchaseReturnKepalaGudang from "@/views/PurchaseReturnKepalaGudang.vue";
import ProductCatalogKepalaGudang from "@/views/ProductCatalogKepalaGudang.vue";
import ManagementUserOwner from "@/views/ManagementUserOwner.vue";
import ManagementVoucherOwner from "@/views/ManagementVoucherOwner.vue";
import ManagementSupplierOwner from "@/views/ManagementSupplierOwner.vue";
import ManagementCategoryOwner from "@/views/ManagementCategoryOwner.vue";
import ManagementUnitOwner from "@/views/ManagementUnitOwner.vue";
import ManageUserOwner from "@/views/ManageUserOwner.vue";
import ManagementProductOwner from "@/views/ManagementProductOwner.vue";
import AddProductOwner from "@/views/AddProductOwner.vue";
import ReportOwners from "@/views/ReportOwners.vue";
import SalesSummaryReportOwners from "@/views/SalesSummaryReportOwners.vue";
import PurchaseSummaryReportOwners from "@/views/PurchaseSummaryReportOwners.vue";
import SaleReturnSummaryOwners from "@/views/SaleReturnSummaryOwners.vue";
import PurchaseReturnSummaryOwners from "@/views/PurchaseReturnSummaryOwners.vue";
import ReportStockOpnameOwners from "@/views/ReportStockOpnameOwners.vue";
import InventoryValueOwners from "@/views/InventoryValueOwners.vue";
import EditUsers from "@/views/EditUsers.vue";
import ListCategoryOwner from "@/views/ListCategoryOwner.vue";
import StockSnapshotOwners from "@/views/StockSnapshotOwners.vue";
import GoodsReceipt from "@/views/GoodsReceipt.vue";
import StockOpnameHistory from "@/views/StockOpnameHistory.vue";

const routes = [
  // ======================= LOGIN =======================
  {
    path: "/",
    name: "Login",
    component: LoginForm,
    meta: { public: true },
  },

  // ======================= KASIR =======================
  {
    path: "/dashboard-cashier",
    name: "DashboardCashier",
    component: DashboardCashier,
    meta: { requiresAuth: true, roles: ["KASIR"] },
  },
  {
    path: "/cashier",
    name: "CashierPage",
    component: CashierPage,
    meta: { requiresAuth: true, roles: ["KASIR"] },
  },
  {
    path: "/completed-transactions",
    name: "CompletedTransactions",
    component: CompletedTransactions,
    meta: { requiresAuth: true, roles: ["KASIR"] },
  },

  // ======================= OWNER =======================
  {
    path: "/dashboard-owner",
    name: "DashboardOwner",
    component: DashboardOwner,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/management-user",
    name: "ManagementUserOwner",
    component: ManagementUserOwner,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/management-category",
    name: "ManagementCategoryOwner",
    component: ManagementCategoryOwner,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/management-unit",
    name: "ManagementUnitOwner",
    component: ManagementUnitOwner,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/management-voucher",
    name: "ManagementVoucherOwner",
    component: ManagementVoucherOwner,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/management-supplier",
    name: "ManagementSupplierOwner",
    component: ManagementSupplierOwner,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/management-product",
    name: "ManagementProductOwner",
    component: ManagementProductOwner,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/add-user",
    name: "AddUserOwner",
    component: ManageUserOwner,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/add-product",
    name: "AddProductOwner",
    component: AddProductOwner,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/report-owners", // Rute utama Laporan, biarkan
    name: "ReportOwners",
    component: ReportOwners,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/report-sales-summary", // Diperbaiki dari /reports-sales-summary
    name: "SalesSummaryReportOwners",
    component: SalesSummaryReportOwners,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/report-purchase-summary", // Diperbaiki dari /reports-purchase-summary
    name: "PurchaseSummaryReportOwners",
    component: PurchaseSummaryReportOwners,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/report-sale-return-summary", // Diperbaiki agar sesuai dengan sidebar /reports-returns-sale
    name: "SaleReturnSummaryOwners",
    component: SaleReturnSummaryOwners,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/report-purchase-return-summary", // Diperbaiki agar sesuai dengan sidebar /reports-returns-purchase
    name: "PurchaseReturnSummaryOwners",
    component: PurchaseReturnSummaryOwners,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/report-stock-opname-summary", // Diperbaiki agar sesuai dengan sidebar /report-stock-opname
    name: "ReportStockOpnameOwners",
    component: ReportStockOpnameOwners,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/report-stock-snapshot", // Diperbaiki agar sesuai dengan sidebar /stocksnapshot
    name: "ReportStockSnapshotOwners",
    component: StockSnapshotOwners,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/report-inventory-value", // Diperbaiki agar sesuai dengan sidebar /inventory-value
    name: "InventoryValueOwners",
    component: InventoryValueOwners,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/edit-user/:id",
    name: "EditUsers",
    component: EditUsers,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/list-category",
    name: "ListCategoryOwner",
    component: ListCategoryOwner,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },
  {
    path: "/users",
    name: "UsersView",
    component: UsersView,
    meta: { requiresAuth: true, roles: ["OWNER"] },
  },

  // ======================= GUDANG =======================
  {
    path: "/dashboard-gudang",
    name: "DashboardGudang",
    component: DashboardGudang,
    meta: { requiresAuth: true, roles: ["GUDANG"] },
  },
  {
    path: "/products",
    name: "ProductList",
    component: ProductList,
    meta: { requiresAuth: true, roles: ["KASIR"] },
  },
  {
    path: "/productsgudang",
    name: "ProductListGudang",
    component: ProductListGudang,
    meta: { requiresAuth: true, roles: ["GUDANG"] },
  },
  {
    path: "/return",
    name: "ReturnPage",
    component: ReturnPage,
    meta: { requiresAuth: true, roles: ["GUDANG", "KASIR"] },
  },

  // ======================= KEPALA GUDANG =======================
  {
    path: "/dashboard-kepalagudang",
    name: "DashboardKepalagudang",
    component: DashboardKepalagudang,
    meta: { requiresAuth: true, roles: ["KEPALA_GUDANG"] },
  },
  {
    path: "/purchaseorder-kepalagudang",
    name: "PurchaseOrderKepalaGudang",
    component: PurchaseOrderKepalaGudang,
    meta: { requiresAuth: true, roles: ["KEPALA_GUDANG"] },
  },
  {
    path: "/sales-return",
    name: "SalesReturnKepalaGudang",
    component: SalesReturnKepalaGudang,
    meta: { requiresAuth: true, roles: ["KEPALA_GUDANG"] },
  },
  {
    path: "/purchasereturn-kepalagudang",
    name: "PurchaseReturnKepalaGudang",
    component: PurchaseReturnKepalaGudang,
    meta: { requiresAuth: true, roles: ["KEPALA_GUDANG"] },
  },
  {
    path: "/product-catalog",
    name: "ProductCatalogKepalaGudang",
    component: ProductCatalogKepalaGudang,
    meta: { requiresAuth: true, roles: ["KEPALA_GUDANG"] },
  },

  // ======================= PEMBELIAN =======================
  {
    path: "/dashboard-pembelian",
    name: "DashboardPembelian",
    component: DashboardPembelian,
    meta: { requiresAuth: true, roles: ["PEMBELIAN"] },
  },
  {
    path: "/product-list-purchase",
    name: "ProductListPurchase",
    component: ProductListPurchase,
    meta: { requiresAuth: true, roles: ["PEMBELIAN"] },
  },
  {
    path: "/create-purchase-order",
    name: "CreatePurchase",
    component: CreatePurchase,
    meta: { requiresAuth: true, roles: ["PEMBELIAN"] },
  },
  {
    path: "/purchase-returns",
    name: "PurchaseReturn",
    component: PurchaseReturn,
    meta: { requiresAuth: true, roles: ["GUDANG"] },
  },
  {
    path: "/gudang/goods-receipt",
    name: "GoodsReceipt",
    component: GoodsReceipt, // 🚀 Rute baru untuk Goods Receipt
    meta: { requiresAuth: true, roles: ["GUDANG", "OWNER"] },
  },
  {
    path: "/history/stock-opnames",
    name: "StockOpnameHistory",
    component: StockOpnameHistory, // 🚀 Rute baru untuk Goods Receipt
    meta: { requiresAuth: true, roles: ["GUDANG"] },
  },

  // ======================= UMUM =======================
  {
    path: "/financial-reports",
    name: "FinancialReportsPage",
    component: FinancialReportsPage,
    meta: {
      requiresAuth: true,
      roles: ["OWNER", "KEPALA_GUDANG", "PEMBELIAN"],
    },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// ======================= NAVIGATION GUARD =======================
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem("token");
  const userRole = localStorage.getItem("userRole");

  if (to.meta.public) {
    if (token && userRole) {
      switch (userRole) {
        case "OWNER":
          return next({ name: "DashboardOwner" });
        case "KASIR":
          return next({ name: "DashboardCashier" });
        case "GUDANG":
          return next({ name: "DashboardGudang" });
        case "KEPALA_GUDANG":
          return next({ name: "DashboardKepalagudang" });
        case "PEMBELIAN":
          return next({ name: "DashboardPembelian" });
        default:
          localStorage.removeItem("token");
          localStorage.removeItem("userRole");
          return next({ name: "Login" });
      }
    }
    return next();
  }

  if (to.meta.requiresAuth) {
    if (!token || !userRole) {
      return next({ name: "Login" });
    }

    if (to.meta.roles && !to.meta.roles.includes(userRole)) {
      alert("Anda tidak memiliki akses ke halaman ini.");
      switch (userRole) {
        case "OWNER":
          return next({ name: "DashboardOwner" });
        case "KASIR":
          return next({ name: "DashboardCashier" });
        case "GUDANG":
          return next({ name: "DashboardGudang" });
        case "KEPALA_GUDANG":
          return next({ name: "DashboardKepalagudang" });
        case "PEMBELIAN":
          return next({ name: "DashboardPembelian" });
        default:
          return next({ name: "Login" });
      }
    }
    return next();
  }

  return next();
});

export default router;
