<template>
  <KepalaGudangLayouts>
    <!-- Header -->
    <div class="flex flex-wrap justify-between gap-3 p-4 mb-2">
      <div>
        <p
          class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
        >
          Warehouse Manager Dashboard
        </p>
        <p class="text-[#60758a] text-sm font-normal leading-normal mt-1">
          Welcome back — ringkasan aktivitas gudang terbaru.
        </p>
      </div>
    </div>

    <div class="p-8 bg-gray-50 min-h-screen">
      <div v-if="loading" class="text-center text-[#60758a] py-12">
        <p class="text-lg">Memuat data...</p>
      </div>

      <div v-else>
        <div
          v-if="errorMessage"
          class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded mb-6"
        >
          {{ errorMessage }}
        </div>

        <!-- KPI Cards -->
        <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-8">
          <div
            class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6] bg-gradient-to-br from-blue-50 to-blue-100"
          >
            <p class="text-[#111418] text-base font-medium leading-normal">
              Orders Pending
            </p>
            <p class="text-[#111418] text-3xl font-bold leading-tight">
              {{ dashboard.purchase_orders_pending || 0 }}
            </p>
            <p class="text-blue-600 text-xs font-semibold">Awaiting approval</p>
          </div>

          <div
            class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6] bg-gradient-to-br from-red-50 to-red-100"
          >
            <p class="text-[#111418] text-base font-medium leading-normal">
              Low Stock (≤5)
            </p>
            <p class="text-[#111418] text-3xl font-bold leading-tight">
              {{ dashboard.products_low_stock || 0 }}
            </p>
            <p class="text-red-600 text-xs font-semibold">Need reorder</p>
          </div>

          <div
            class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6] bg-gradient-to-br from-green-50 to-green-100"
          >
            <p class="text-[#111418] text-base font-medium leading-normal">
              In Stock
            </p>
            <p class="text-[#111418] text-3xl font-bold leading-tight">
              {{ dashboard.products_in_stock || 0 }}
            </p>
            <p class="text-green-600 text-xs font-semibold">Available items</p>
          </div>

          <div
            class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6] bg-gradient-to-br from-purple-50 to-purple-100"
          >
            <p class="text-[#111418] text-base font-medium leading-normal">
              Total Products
            </p>
            <p class="text-[#111418] text-3xl font-bold leading-tight">
              {{ dashboard.total_products || 0 }}
            </p>
            <p class="text-purple-600 text-xs font-semibold">In inventory</p>
          </div>
        </div>

        <!-- Main Content Grid -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
          <!-- Left Column (2 cols) -->
          <div class="lg:col-span-2 space-y-6">
            <!-- Pending Purchase Orders -->
            <div class="rounded-lg border border-[#dbe0e6] bg-white">
              <div class="border-b border-[#dbe0e6] p-6">
                <h3 class="text-[#111418] text-lg font-bold leading-tight">
                  Pending Purchase Orders
                </h3>
              </div>
              <div class="divide-y divide-[#dbe0e6]">
                <div
                  v-if="pendingPurchases.length === 0"
                  class="p-6 text-center text-[#60758a] text-sm"
                >
                  No pending orders
                </div>
                <div
                  v-for="purchase in pendingPurchases"
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
                      <p class="text-[#60758a] text-xs">
                        Items: {{ purchase.item_count || 0 }}
                      </p>
                    </div>
                    <span
                      class="px-3 py-1 rounded bg-blue-100 text-blue-800 text-xs font-semibold"
                    >
                      {{ purchase.status || "SUBMITTED" }}
                    </span>
                  </div>
                </div>
              </div>
            </div>

            <!-- Low Stock Products -->
            <div class="rounded-lg border border-[#dbe0e6] bg-white">
              <div class="border-b border-[#dbe0e6] p-6">
                <h3 class="text-[#111418] text-lg font-bold leading-tight">
                  Products Needing Reorder (≤5 units)
                </h3>
              </div>
              <div class="divide-y divide-[#dbe0e6]">
                <div
                  v-if="lowStockProducts.length === 0"
                  class="p-6 text-center text-[#60758a] text-sm"
                >
                  All products have healthy stock levels
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
                        {{ product.category || "No category" }}
                      </p>
                    </div>
                    <div class="text-right">
                      <div
                        class="w-16 bg-red-100 rounded px-2 py-1 text-center"
                      >
                        <p class="text-[#111418] font-bold text-sm">
                          {{ product.stock }}
                        </p>
                        <p class="text-[#60758a] text-xs">units</p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <!-- Right Column (1 col) -->
          <div class="space-y-6">
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
                  No supplier data
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
                    <div class="flex-1 min-w-0">
                      <p class="text-[#111418] font-semibold text-sm truncate">
                        {{ supplier.name }}
                      </p>
                      <p class="text-[#60758a] text-xs">
                        {{ supplier.order_count }} orders
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Stock Summary -->
            <div class="rounded-lg border border-[#dbe0e6] bg-white p-6">
              <h3 class="text-[#111418] text-lg font-bold leading-tight mb-4">
                Stock Summary
              </h3>
              <div class="space-y-3">
                <div>
                  <div class="flex justify-between mb-2">
                    <p class="text-[#111418] text-sm font-medium">
                      Stock Health
                    </p>
                    <p class="text-[#111418] font-bold text-sm">
                      {{ stockHealthPercent }}%
                    </p>
                  </div>
                  <div class="w-full bg-gray-200 rounded-full h-3">
                    <div
                      class="bg-green-500 h-3 rounded-full transition-all"
                      :style="{ width: stockHealthPercent + '%' }"
                    ></div>
                  </div>
                </div>
                <div
                  class="text-xs text-[#60758a] space-y-1 pt-2 border-t border-gray-200"
                >
                  <p>
                    {{ dashboard.products_in_stock }} /
                    {{ dashboard.total_products }} items in stock
                  </p>
                  <p>{{ dashboard.products_low_stock }} items need reorder</p>
                </div>
              </div>
            </div>

            <!-- Quick Actions -->
            <div
              class="rounded-lg border border-[#dbe0e6] bg-gradient-to-br from-blue-50 to-blue-100 p-6"
            >
              <h3 class="text-[#111418] text-lg font-bold leading-tight mb-4">
                Quick Actions
              </h3>
              <div class="space-y-2">
                <button
                  class="w-full text-left p-3 rounded border border-blue-300 hover:border-blue-500 hover:bg-blue-200 transition text-sm font-medium text-[#111418]"
                  @click="goToPendingOrders"
                >
                  📋 Review Orders
                </button>
                <button
                  class="w-full text-left p-3 rounded border border-blue-300 hover:border-blue-500 hover:bg-blue-200 transition text-sm font-medium text-[#111418]"
                  @click="goToLowStockProducts"
                >
                  📦 Reorder Items
                </button>
              </div>
            </div>

            <!-- Info Box -->
            <div class="rounded-lg border border-[#dbe0e6] bg-blue-50 p-6">
              <h4
                class="font-semibold text-blue-900 mb-3 flex items-center gap-2"
              >
                <span>ℹ️</span> Tips
              </h4>
              <ul class="text-blue-800 text-xs space-y-2">
                <li>✓ Check pending orders daily</li>
                <li>✓ Reorder low stock items promptly</li>
                <li>✓ Maintain 70%+ stock health</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>
  </KepalaGudangLayouts>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import KepalaGudangLayouts from "@/components/KepalaGudangLayouts.vue";
import api from "../services/api.js";

// State untuk dashboard
const loading = ref(true);
const errorMessage = ref("");
const dashboard = ref({
  purchase_orders_pending: 0,
  products_low_stock: 0,
  products_in_stock: 0,
  total_products: 0,
});

const pendingPurchases = ref([]);
const lowStockProducts = ref([]);
const topSuppliers = ref([]);

// Helper untuk mengambil total dari berbagai struktur respons
function getTotalFromResponse(resp) {
  if (!resp) return 0;
  const d = resp.data ?? resp;
  return (
    d.total ?? d?.data?.total ?? d?.meta?.total ?? d?.data?.meta?.total ?? 0
  );
}

// Ambil array dari respons
function getDataArray(resp) {
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

// Stock health percent
const stockHealthPercent = computed(() => {
  const total = dashboard.value.total_products;
  if (total === 0) return 0;
  const inStock = dashboard.value.products_in_stock;
  return Math.round((inStock / total) * 100);
});

// Navigation helpers
const goToPendingOrders = () => {
  window.location.href = "/#/purchase-orders"; // Adjust route as needed
};

const goToLowStockProducts = () => {
  window.location.href = "/#/products"; // Adjust route as needed
};

// Fungsi untuk mengambil data dashboard
const fetchDashboard = async () => {
  loading.value = true;
  errorMessage.value = "";
  try {
    // 1) Purchase orders pending approval (SUBMITTED)
    const poResp = await api.get("/pembelian/purchases", {
      params: { status: "SUBMITTED", size: 200 },
    });
    const purchases = getDataArray(poResp);
    dashboard.value.purchase_orders_pending = getTotalFromResponse(poResp);

    // Process pending purchases untuk display
    pendingPurchases.value = purchases.slice(0, 5).map((p) => ({
      id: p.id || p.ID,
      status: p.status || "SUBMITTED",
      created_at: p.created_at || p.CreatedAt,
      supplier: p.supplier || p.Supplier || {},
      item_count: (p.Items || p.items || p.PurchaseItems || []).length,
    }));

    // 2) Products list untuk low stock dan in-stock counts
    let productsResp;
    let products = [];

    const productEndpoints = [
      "/products",
      "/pembelian/products",
      "/owner/products",
    ];

    for (const endpoint of productEndpoints) {
      try {
        productsResp = await api.get(endpoint, { params: { size: 500 } });
        products = getDataArray(productsResp);
        if (products.length > 0) {
          console.log(`✓ Products fetched from ${endpoint}:`, products.length);
          break;
        }
      } catch (err) {
        console.log(`✗ Failed to fetch from ${endpoint}`);
      }
    }

    // Fallback: extract dari purchases jika produk endpoint gagal
    if (products.length === 0 && purchases.length > 0) {
      console.log("Extracting products from purchase orders...");
      const productMap = {};
      purchases.forEach((p) => {
        const items = p.Items || p.items || p.PurchaseItems || [];
        items.forEach((item) => {
          const prod = item.Product || item.product || item.ProductDto || {};
          const id = prod.id || prod.ID;
          if (id && !productMap[id]) {
            productMap[id] = {
              id,
              name: prod.name || prod.Name || "Unknown",
              stock: Number(prod.stock || prod.Stock || 0),
              category: prod.category?.name || prod.Category?.name || "Unknown",
            };
          }
        });
      });
      products = Object.values(productMap);
      console.log(`✓ Extracted ${products.length} products from purchases`);
    }

    // Threshold untuk low stock
    const LOW_STOCK_THRESHOLD = 5;

    const productMap = {};
    let lowCount = 0;
    let inStockCount = 0;

    products.forEach((p) => {
      const prod = p.Product ?? p.product ?? p;
      const id = prod.id ?? prod.ID;
      const stock = Number(
        prod.stock ?? prod.Stock ?? prod.qty ?? prod.quantity ?? 0
      );
      const name = prod.name ?? prod.Name ?? "Unknown";
      const category = prod.category?.name ?? prod.Category?.name ?? "Unknown";

      if (id && !Number.isNaN(stock)) {
        if (!productMap[id]) {
          productMap[id] = {
            id,
            name,
            stock,
            category,
          };

          if (stock <= LOW_STOCK_THRESHOLD) lowCount += 1;
          if (stock > 0) inStockCount += 1;
        }
      }
    });

    dashboard.value.products_low_stock = lowCount;
    dashboard.value.products_in_stock = inStockCount;
    dashboard.value.total_products = Object.keys(productMap).length;

    // Process low stock products untuk display
    lowStockProducts.value = Object.values(productMap)
      .filter((p) => Number(p.stock) <= LOW_STOCK_THRESHOLD)
      .sort((a, b) => Number(a.stock) - Number(b.stock))
      .slice(0, 5);

    console.log("Dashboard summary:", dashboard.value);

    // 3) Top suppliers dari purchases
    const supplierMap = {};
    purchases.forEach((p) => {
      const supplier = p.supplier || p.Supplier || {};
      const supplierId = supplier.id || supplier.ID;
      if (supplierId) {
        if (!supplierMap[supplierId]) {
          supplierMap[supplierId] = {
            id: supplierId,
            name: supplier.name || supplier.Name || "Unknown",
            order_count: 0,
          };
        }
        supplierMap[supplierId].order_count++;
      }
    });

    topSuppliers.value = Object.values(supplierMap)
      .sort((a, b) => b.order_count - a.order_count)
      .slice(0, 5);
  } catch (error) {
    console.error("Gagal fetch dashboard Kepala Gudang:", error);
    if (error && error.response) {
      const status = error.response.status;
      const body = error.response.data;
      let bodyText = "";
      try {
        bodyText = typeof body === "string" ? body : JSON.stringify(body);
      } catch (e) {
        bodyText = String(body);
      }
      errorMessage.value = `Request failed with status code ${status}: ${bodyText}`;
    } else {
      errorMessage.value = error?.message || String(error);
    }
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  fetchDashboard();
});
</script>
