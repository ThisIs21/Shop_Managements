<template>
  <KepalaGudangLayouts>
    <div class="flex flex-wrap justify-between gap-3 p-4">
      <p
        class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
      >
        Warehouse Manager Dashboard
      </p>
    </div>

    <div class="p-4">
      <div class="mb-6">
        <p class="text-[#60758a] text-sm font-normal leading-normal mt-1">
          Welcome back — ringkasan aktivitas gudang terbaru.
        </p>
      </div>

      <div v-if="loading" class="text-center text-[#60758a]">
        Memuat data...
      </div>

      <div v-else>
        <div v-if="errorMessage" class="text-red-600 mb-4">
          {{ errorMessage }}
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-8">
          <div
            class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6]"
          >
            <p class="text-[#111418] text-base font-medium leading-normal">
              Purchase Orders Pending
            </p>
            <p class="text-[#111418] text-2xl font-bold leading-tight">
              {{ dashboard.purchase_orders_pending || 0 }}
            </p>
          </div>

          <div
            class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6]"
          >
            <p class="text-[#111418] text-base font-medium leading-normal">
              Products Low Stock (&le;5)
            </p>
            <p class="text-[#111418] text-2xl font-bold leading-tight">
              {{ dashboard.products_low_stock || 0 }}
            </p>
          </div>

          <div
            class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6]"
          >
            <p class="text-[#111418] text-base font-medium leading-normal">
              Products In Stock
            </p>
            <p class="text-[#111418] text-2xl font-bold leading-tight">
              {{ dashboard.products_in_stock || 0 }}
            </p>
          </div>
        </div>

        <div
          class="flex flex-col gap-2 rounded-lg p-6 bg-[#f0f2f5] min-h-[120px]"
        >
          <p
            class="text-[#111418] tracking-light text-lg font-semibold leading-tight"
          >
            Selamat datang, Kepala Gudang!
          </p>
          <p class="text-[#60758a] text-sm">
            Gunakan menu kiri untuk melihat permintaan dan katalog produk.
          </p>
        </div>
      </div>
    </div>
  </KepalaGudangLayouts>
</template>

<script setup>
import { ref, onMounted } from "vue";
import KepalaGudangLayouts from "@/components/KepalaGudangLayouts.vue";
import api from "../services/api.js";

// State untuk dashboard
const loading = ref(true);
const errorMessage = ref("");
const dashboard = ref({
  purchase_orders_pending: 0,
  products_low_stock: 0,
  products_in_stock: 0,
});

// Helper untuk mengambil total dari berbagai struktur respons
function getTotalFromResponse(resp) {
  if (!resp) return 0;
  const d = resp.data ?? resp;
  return (
    d.total ?? d?.data?.total ?? d?.meta?.total ?? d?.data?.meta?.total ?? 0
  );
}

// Ambil array produk dari respons
function getProductsArray(resp) {
  if (!resp) return [];
  const d = resp.data ?? resp;
  return d?.data ?? d?.items ?? d?.products ?? [];
}

// Fungsi untuk mengambil data dashboard
const fetchDashboard = async () => {
  loading.value = true;
  errorMessage.value = "";
  try {
    // 1) Purchase orders pending approval (SUBMITTED)
    const poResp = await api.get("/pembelian/purchases", {
      params: { status: "SUBMITTED", size: 1 },
    });
    dashboard.value.purchase_orders_pending = getTotalFromResponse(poResp);

    // Purchase returns removed from dashboard (disabled for safety)

    // 3) Products list to compute low stock and in-stock counts
    // try owner/products or /products depending on backend
    let productsResp;
    try {
      productsResp = await api.get("/products", { params: { size: 500 } });
    } catch (e) {
      // fallback to owner/products
      productsResp = await api.get("/owner/products", {
        params: { size: 500 },
      });
    }

    const products = getProductsArray(productsResp) || [];

    // Threshold low stock (tweakable)
    const LOW_STOCK_THRESHOLD = 5;

    let lowCount = 0;
    let inStockCount = 0;
    products.forEach((p) => {
      const prod = p.Product ?? p.product ?? p; // support wrapped objects
      const stock = Number(
        prod.stock ?? prod.Stock ?? prod.qty ?? prod.quantity ?? 0
      );
      if (!Number.isNaN(stock)) {
        if (stock <= LOW_STOCK_THRESHOLD) lowCount += 1;
        if (stock > 0) inStockCount += 1;
      }
    });

    dashboard.value.products_low_stock = lowCount;
    dashboard.value.products_in_stock = inStockCount;
  } catch (error) {
    console.error("Gagal fetch dashboard Kepala Gudang:", error);
    // Build a detailed error message for UI debugging
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
