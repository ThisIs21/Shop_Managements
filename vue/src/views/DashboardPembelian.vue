<template>
  <PurchasingLayouts>
    <div class="p-8">
      <div class="mb-6">
        <h1
          class="text-[#111418] text-[32px] font-bold leading-tight tracking-[-0.015em]"
        >
          Purchase Dashboard
        </h1>
        <p class="text-[#60758a] text-sm font-normal leading-normal mt-1">
          Welcome back, Sarah! Here's an overview of your purchase activities.
        </p>
      </div>

      <div v-if="loading" class="text-center text-[#60758a]">
        Memuat data...
      </div>
      <div v-else>
        <!-- KPI Cards -->
        <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-8">
          <div
            class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6] bg-gradient-to-br from-blue-50 to-blue-100"
          >
            <p class="text-[#111418] text-base font-medium leading-normal">
              Orders to Process
            </p>
            <p class="text-[#111418] text-3xl font-bold leading-tight">
              {{ dashboardData.purchase_orders_to_process || 0 }}
            </p>
            <p class="text-blue-600 text-xs font-semibold">
              Waiting for approval
            </p>
          </div>
          <div
            class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6] bg-gradient-to-br from-amber-50 to-amber-100"
          >
            <p class="text-[#111418] text-base font-medium leading-normal">
              Returns to Review
            </p>
            <p class="text-[#111418] text-3xl font-bold leading-tight">
              {{ dashboardData.returns_to_review || 0 }}
            </p>
            <p class="text-amber-600 text-xs font-semibold">Pending action</p>
          </div>
          <div
            class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6] bg-gradient-to-br from-green-50 to-green-100"
          >
            <p class="text-[#111418] text-base font-medium leading-normal">
              Products in Stock
            </p>
            <p class="text-[#111418] text-3xl font-bold leading-tight">
              {{ dashboardData.products_in_stock || 0 }}
            </p>
            <p class="text-green-600 text-xs font-semibold">Available items</p>
          </div>
          <div
            class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6] bg-gradient-to-br from-red-50 to-red-100"
          >
            <p class="text-[#111418] text-base font-medium leading-normal">
              Low Stock
            </p>
            <p class="text-[#111418] text-3xl font-bold leading-tight">
              {{ dashboardData.low_stock_products || 0 }}
            </p>
            <p class="text-red-600 text-xs font-semibold">Need reorder</p>
          </div>
        </div>

        <!-- Recent Purchases & Low Stock Products -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
          <!-- Recent Purchases -->
          <div class="rounded-lg border border-[#dbe0e6] bg-white">
            <div class="border-b border-[#dbe0e6] p-6">
              <h3 class="text-[#111418] text-lg font-bold leading-tight">
                Recent Purchases
              </h3>
            </div>
            <div class="divide-y divide-[#dbe0e6]">
              <div
                v-if="recentPurchases.length === 0"
                class="p-6 text-center text-[#60758a] text-sm"
              >
                No recent purchases
              </div>
              <div
                v-for="purchase in recentPurchases"
                :key="purchase.id"
                class="p-4 hover:bg-gray-50 transition"
              >
                <div class="flex justify-between items-start gap-3">
                  <div>
                    <p class="text-[#111418] font-semibold text-sm">
                      Order #{{ purchase.id }}
                    </p>
                    <p class="text-[#60758a] text-xs mt-1">
                      {{ formatDate(purchase.created_at) }}
                    </p>
                    <p class="text-[#60758a] text-xs">
                      Supplier: {{ purchase.supplier?.name || "Unknown" }}
                    </p>
                  </div>
                  <span
                    :class="[
                      'px-2 py-1 rounded text-xs font-semibold',
                      getStatusColor(purchase.status),
                    ]"
                  >
                    {{ purchase.status || "DRAFT" }}
                  </span>
                </div>
              </div>
            </div>
          </div>

          <!-- Low Stock Products -->
          <div class="rounded-lg border border-[#dbe0e6] bg-white">
            <div class="border-b border-[#dbe0e6] p-6">
              <h3 class="text-[#111418] text-lg font-bold leading-tight">
                Low Stock Products
              </h3>
            </div>
            <div class="divide-y divide-[#dbe0e6]">
              <div
                v-if="lowStockProducts.length === 0"
                class="p-6 text-center text-[#60758a] text-sm"
              >
                All products have good stock levels
              </div>
              <div
                v-for="product in lowStockProducts"
                :key="product.id"
                class="p-4 hover:bg-gray-50 transition"
              >
                <div class="flex justify-between items-start gap-3">
                  <div class="flex-1">
                    <p class="text-[#111418] font-semibold text-sm">
                      {{ product.name }}
                    </p>
                    <p class="text-[#60758a] text-xs mt-1">
                      Category: {{ product.category?.name || "Unknown" }}
                    </p>
                  </div>
                  <div class="text-right">
                    <p class="text-[#111418] font-bold text-sm">
                      {{ product.stock }}
                    </p>
                    <p class="text-[#60758a] text-xs">Stock</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Top Suppliers & Purchase Trends -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
          <!-- Top Suppliers -->
          <div class="rounded-lg border border-[#dbe0e6] bg-white">
            <div class="border-b border-[#dbe0e6] p-6">
              <h3 class="text-[#111418] text-lg font-bold leading-tight">
                Top Suppliers
              </h3>
            </div>
            <div class="divide-y divide-[#dbe0e6]">
              <div
                v-if="topSuppliers.length === 0"
                class="p-6 text-center text-[#60758a] text-sm"
              >
                No supplier data available
              </div>
              <div
                v-for="(supplier, idx) in topSuppliers"
                :key="supplier.id"
                class="p-4 hover:bg-gray-50 transition"
              >
                <div class="flex items-center gap-3">
                  <div
                    class="flex-shrink-0 w-8 h-8 rounded-full bg-blue-200 text-blue-800 flex items-center justify-center font-bold text-xs"
                  >
                    {{ idx + 1 }}
                  </div>
                  <div class="flex-1">
                    <p class="text-[#111418] font-semibold text-sm">
                      {{ supplier.name }}
                    </p>
                    <p class="text-[#60758a] text-xs">{{ supplier.contact }}</p>
                  </div>
                  <div class="text-right">
                    <p class="text-[#111418] font-bold">
                      {{ supplier.purchase_count || 0 }}
                    </p>
                    <p class="text-[#60758a] text-xs">Orders</p>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Monthly Statistics -->
          <div class="rounded-lg border border-[#dbe0e6] bg-white">
            <div class="border-b border-[#dbe0e6] p-6">
              <h3 class="text-[#111418] text-lg font-bold leading-tight">
                Monthly Statistics
              </h3>
            </div>
            <div class="p-6">
              <div class="space-y-4">
                <div>
                  <div class="flex justify-between mb-2">
                    <p class="text-[#111418] text-sm font-medium">
                      Total Purchases This Month
                    </p>
                    <p class="text-[#111418] font-bold">
                      {{ dashboardData.monthly_purchases || 0 }}
                    </p>
                  </div>
                  <div class="w-full bg-gray-200 rounded-full h-2">
                    <div
                      class="bg-blue-500 h-2 rounded-full"
                      :style="{
                        width: getProgressWidth(
                          dashboardData.monthly_purchases,
                          100
                        ),
                      }"
                    ></div>
                  </div>
                </div>
                <div>
                  <div class="flex justify-between mb-2">
                    <p class="text-[#111418] text-sm font-medium">
                      Approved Purchases
                    </p>
                    <p class="text-[#111418] font-bold">
                      {{ dashboardData.approved_purchases || 0 }}
                    </p>
                  </div>
                  <div class="w-full bg-gray-200 rounded-full h-2">
                    <div
                      class="bg-green-500 h-2 rounded-full"
                      :style="{
                        width: getProgressWidth(
                          dashboardData.approved_purchases,
                          dashboardData.monthly_purchases || 1
                        ),
                      }"
                    ></div>
                  </div>
                </div>
                <div>
                  <div class="flex justify-between mb-2">
                    <p class="text-[#111418] text-sm font-medium">
                      Pending Purchases
                    </p>
                    <p class="text-[#111418] font-bold">
                      {{ dashboardData.pending_purchases || 0 }}
                    </p>
                  </div>
                  <div class="w-full bg-gray-200 rounded-full h-2">
                    <div
                      class="bg-yellow-500 h-2 rounded-full"
                      :style="{
                        width: getProgressWidth(
                          dashboardData.pending_purchases,
                          dashboardData.monthly_purchases || 1
                        ),
                      }"
                    ></div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </PurchasingLayouts>
</template>

<script setup>
import { ref, onMounted } from "vue";
import PurchasingLayouts from "../components/PurchasingLayouts.vue";
import api from "../services/api.js";

// State untuk menyimpan data dashboard
const dashboardData = ref({
  purchase_orders_to_process: 0,
  returns_to_review: 0,
  products_in_stock: 0,
  low_stock_products: 0,
  monthly_purchases: 0,
  approved_purchases: 0,
  pending_purchases: 0,
});
const loading = ref(true);
const recentPurchases = ref([]);
const lowStockProducts = ref([]);
const topSuppliers = ref([]);

// Helper untuk mengambil total dari berbagai bentuk respons API
function getTotalFromResponse(resp) {
  if (!resp) return 0;
  const d = resp.data ?? resp;
  return (
    d.total ?? d?.data?.total ?? d?.meta?.total ?? d?.data?.meta?.total ?? 0
  );
}

// Helper untuk mendapatkan array purchases dari respons
function getPurchasesArray(resp) {
  if (!resp) return [];
  const d = resp.data ?? resp;
  return d?.data?.data ?? d?.data ?? d?.items ?? [];
}

// Format tanggal
const formatDate = (dateStr) => {
  if (!dateStr) return "-";
  try {
    return new Date(dateStr).toLocaleDateString("id-ID", {
      year: "numeric",
      month: "short",
      day: "numeric",
    });
  } catch {
    return dateStr;
  }
};

// Get status color
const getStatusColor = (status) => {
  const colors = {
    SUBMITTED: "bg-blue-100 text-blue-800",
    APPROVED: "bg-green-100 text-green-800",
    RECEIVED: "bg-emerald-100 text-emerald-800",
    REJECTED: "bg-red-100 text-red-800",
    DRAFT: "bg-gray-100 text-gray-800",
  };
  return colors[status] || "bg-gray-100 text-gray-800";
};

// Get progress width
const getProgressWidth = (value, total) => {
  if (total === 0) return "0%";
  const percent = (value / total) * 100;
  return Math.min(100, Math.max(0, percent)) + "%";
};

// Fungsi untuk mengambil data dari backend
const fetchDashboardData = async () => {
  try {
    loading.value = true;

    // Ambil jumlah purchase orders dengan status SUBMITTED
    const purchaseResponse = await api.get("/pembelian/purchases", {
      params: { status: "SUBMITTED", size: 1 },
    });
    dashboardData.value.purchase_orders_to_process =
      getTotalFromResponse(purchaseResponse) || 0;

    // Ambil jumlah purchase returns dengan status PENDING
    const returnResponse = await api.get("/pembelian/purchase-returns", {
      params: { status: "PENDING", size: 1 },
    });
    dashboardData.value.returns_to_review =
      getTotalFromResponse(returnResponse);

    // Ambil data pembelian untuk berbagai analisis
    const allPurchasesResponse = await api.get("/pembelian/purchases", {
      params: { size: 200 },
    });
    const purchases = getPurchasesArray(allPurchasesResponse);

    // Ambil 5 pembelian terbaru
    recentPurchases.value = purchases.slice(0, 5).map((p) => ({
      id: p.id || p.ID,
      status: p.status || "DRAFT",
      created_at: p.created_at || p.CreatedAt,
      supplier: p.supplier || p.Supplier || {},
    }));

    // Ambil data produk dari berbagai endpoint
    let allProducts = [];

    // Try multiple endpoints
    const endpoints = ["/products", "/pembelian/products", "/owner/products"];

    for (const endpoint of endpoints) {
      try {
        const response = await api.get(endpoint, {
          params: { size: 500 },
        });
        const data = getPurchasesArray(response);
        if (data && data.length > 0) {
          allProducts = data;
          console.log(
            `✓ Products fetched from ${endpoint}:`,
            allProducts.length
          );
          break;
        }
      } catch (err) {
        console.log(`✗ Failed to fetch from ${endpoint}`);
      }
    }

    // Jika semua endpoint gagal, extract dari purchase items
    if (allProducts.length === 0) {
      console.log("Extracting products from purchase items...");
      const productMap = {};
      purchases.forEach((purchase) => {
        const items =
          purchase.Items || purchase.items || purchase.PurchaseItems || [];
        items.forEach((item) => {
          const product = item.Product || item.product || item.ProductDto || {};
          const id = product.id || product.ID;
          if (id && !productMap[id]) {
            productMap[id] = {
              id,
              name: product.name || product.Name || "Unknown",
              stock: Number(product.stock || product.Stock || 0),
              category: product.category || product.Category || {},
            };
          }
        });
      });
      allProducts = Object.values(productMap);
      console.log(`✓ Extracted ${allProducts.length} products from purchases`);
    }

    // Process produk: hitung yang in stock dan low stock
    const productMap = {};
    allProducts.forEach((product) => {
      const id = product.id || product.ID;
      const stock = Number(product.stock || product.Stock || 0);
      const name = product.name || product.Name || "Unknown";
      const category = product.category || product.Category || {};

      if (id) {
        productMap[id] = {
          id,
          name,
          stock,
          category,
        };
      }
    });

    console.log("Product Map:", productMap);

    // Filter: products in stock (stock > 0)
    const productsInStock = Object.values(productMap).filter(
      (p) => Number(p.stock) > 0
    );
    dashboardData.value.products_in_stock = productsInStock.length;
    console.log(
      `Products in stock: ${productsInStock.length}`,
      productsInStock
    );

    // Filter: low stock products (stock < 10 dan stock > 0)
    lowStockProducts.value = Object.values(productMap)
      .filter((p) => {
        const s = Number(p.stock);
        return s > 0 && s < 10;
      })
      .sort((a, b) => Number(a.stock) - Number(b.stock))
      .slice(0, 5);

    dashboardData.value.low_stock_products = lowStockProducts.value.length;
    console.log(
      `Low stock products: ${lowStockProducts.value.length}`,
      lowStockProducts.value
    );

    // Hitung statistik bulanan
    const currentMonth = new Date().getMonth();
    const currentYear = new Date().getFullYear();

    const monthlyPurchases = purchases.filter((p) => {
      const date = new Date(p.created_at || p.CreatedAt);
      return (
        date.getMonth() === currentMonth && date.getFullYear() === currentYear
      );
    });

    dashboardData.value.monthly_purchases = monthlyPurchases.length;
    dashboardData.value.approved_purchases = monthlyPurchases.filter(
      (p) => p.status === "APPROVED" || p.status === "RECEIVED"
    ).length;
    dashboardData.value.pending_purchases = monthlyPurchases.filter(
      (p) => p.status === "SUBMITTED" || p.status === "DRAFT"
    ).length;

    // Hitung top suppliers
    const supplierMap = {};
    purchases.forEach((p) => {
      const supplier = p.supplier || p.Supplier || {};
      const supplierId = supplier.id || supplier.ID;
      if (supplierId) {
        if (!supplierMap[supplierId]) {
          supplierMap[supplierId] = {
            id: supplierId,
            name: supplier.name || supplier.Name || "Unknown",
            contact:
              supplier.contact || supplier.Contact || supplier.phone || "-",
            purchase_count: 0,
          };
        }
        supplierMap[supplierId].purchase_count++;
      }
    });

    topSuppliers.value = Object.values(supplierMap)
      .sort((a, b) => b.purchase_count - a.purchase_count)
      .slice(0, 5);
  } catch (error) {
    console.error("Gagal mengambil data dashboard:", error);
    if (error.response) {
      if (error.response.status === 401) {
        console.error("Sesi Anda telah berakhir. Silakan login kembali.");
      } else if (error.response.status === 404) {
        console.error("Endpoint tidak ditemukan. Periksa konfigurasi server.");
      } else {
        console.error("Terjadi kesalahan saat memuat data:", error.message);
      }
    } else {
      console.error("Koneksi ke server gagal. Periksa jaringan Anda.");
    }
  } finally {
    loading.value = false;
  }
};

// Panggil fungsi saat komponen dimuat
onMounted(fetchDashboardData);
</script>
