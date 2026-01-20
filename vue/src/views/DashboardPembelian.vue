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
      <div v-else class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-8">
        <div class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6]">
          <p class="text-[#111418] text-base font-medium leading-normal">
            Purchase Orders to Process
          </p>
          <p class="text-[#111418] text-2xl font-bold leading-tight">
            {{ dashboardData.purchase_orders_to_process || 0 }}
          </p>
        </div>
        <div class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6]">
          <p class="text-[#111418] text-base font-medium leading-normal">
            Returns to Review
          </p>
          <p class="text-[#111418] text-2xl font-bold leading-tight">
            {{ dashboardData.returns_to_review || 0 }}
          </p>
        </div>
        <div class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6]">
          <p class="text-[#111418] text-base font-medium leading-normal">
            Products in Stock
          </p>
          <p class="text-[#111418] text-2xl font-bold leading-tight">
            {{ dashboardData.products_in_stock || 0 }}
          </p>
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
});
const loading = ref(true);

// Helper untuk mengambil total dari berbagai bentuk respons API
function getTotalFromResponse(resp) {
  if (!resp) return 0;
  const d = resp.data ?? resp;
  // Cek berbagai struktur yang mungkin dari backend
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

// Fungsi untuk mengambil data dari backend (menggunakan instance `api`)
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

    // Ambil data pembelian (lebih banyak) untuk menghitung produk yang ada di stok
    const allPurchasesResponse = await api.get("/pembelian/purchases", {
      params: { size: 200 },
    });

    const purchases = getPurchasesArray(allPurchasesResponse);

    // Flatten items dengan fallback properti yang berbeda
    const allItems = purchases.flatMap((purchase) => {
      return purchase.Items ?? purchase.items ?? purchase.PurchaseItems ?? [];
    });

    // Hitung produk unik dimana stock > 0 dengan fallback nama properti
    const uniqueProducts = new Set();
    allItems.forEach((item) => {
      const product = item.Product ?? item.product ?? item.ProductDto ?? {};
      const stock = product.stock ?? product.Stock ?? product.qty ?? 0;
      const id = product.id ?? product.ID ?? product.product_id ?? null;
      if (id && Number(stock) > 0) uniqueProducts.add(String(id));
    });

    dashboardData.value.products_in_stock = uniqueProducts.size;
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
