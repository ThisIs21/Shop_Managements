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
import axios from "axios";

// State untuk menyimpan data dashboard
const dashboardData = ref({
  purchase_orders_to_process: 0,
  returns_to_review: 0,
  products_in_stock: 0,
});
const loading = ref(true);

// Fungsi untuk mengambil data dari backend
const fetchDashboardData = async () => {
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      console.error(
        "Token autentikasi tidak ditemukan. Silakan login kembali."
      );
      return;
    }

    loading.value = true;

    // Ambil jumlah purchase orders dengan status SUBMITTED
    const purchaseResponse = await axios.get(
      "http://localhost:8081/api/v1/pembelian/purchases",
      {
        params: { status: "SUBMITTED", size: 1 },
        headers: { Authorization: `Bearer ${token}` },
      }
    );
    dashboardData.value.purchase_orders_to_process =
      purchaseResponse.data.data.total || 0;

    // Ambil jumlah purchase returns dengan status PENDING
    const returnResponse = await axios.get(
      "http://localhost:8081/api/v1/pembelian/purchase-returns",
      {
        params: { status: "PENDING", size: 1 },
        headers: { Authorization: `Bearer ${token}` },
      }
    );
    dashboardData.value.returns_to_review = returnResponse.data.total || 0;

    // Ambil data produk dari pembelian untuk menghitung stok (karena /api/v1/products mungkin tidak sesuai)
    const allPurchasesResponse = await axios.get(
      "http://localhost:8081/api/v1/pembelian/purchases",
      {
        params: { size: 100 }, // Ambil lebih banyak data untuk menghitung stok
        headers: { Authorization: `Bearer ${token}` },
      }
    );
    const allItems = allPurchasesResponse.data.data.data.flatMap(
      (purchase) => purchase.Items
    );
    const uniqueProductsInStock = new Set(
      allItems
        .filter((item) => item.Product.stock > 0)
        .map((item) => item.Product.id)
    );
    dashboardData.value.products_in_stock = uniqueProductsInStock.size;
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
