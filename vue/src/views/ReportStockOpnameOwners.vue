<template>
  <OwnerLayout>
    <div class="p-4">
      <h1 class="text-2xl font-bold mb-4">Laporan Stok Opname</h1>

      <div
        v-if="errorMessage"
        class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative mb-4"
        role="alert"
      >
        <span class="block sm:inline">{{ errorMessage }}</span>
      </div>

      <div class="flex flex-wrap gap-4 mb-6 items-end">
        <label class="flex flex-col">
          <span class="text-sm font-medium mb-1">Dari Tanggal</span>
          <input
            type="date"
            v-model="filters.startDate"
            class="border p-2 rounded w-40"
            @change="validateFilters"
          />
        </label>
        <label class="flex flex-col">
          <span class="text-sm font-medium mb-1">Sampai Tanggal</span>
          <input
            type="date"
            v-model="filters.endDate"
            class="border p-2 rounded w-40"
            @change="validateFilters"
          />
        </label>
        <button
          @click="applyFilter"
          class="h-10 px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700 font-semibold"
          :disabled="loading || !isFilterValid"
        >
          Terapkan Filter
        </button>
        <button
          @click="resetFilter"
          class="h-10 px-4 py-2 bg-gray-200 text-gray-800 rounded hover:bg-gray-300 font-semibold"
          :disabled="loading"
        >
          Reset
        </button>
      </div>

      <div v-if="loading" class="text-center py-8">
        <svg
          class="animate-spin h-6 w-6 text-blue-600 mx-auto"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
        >
          <circle
            class="opacity-25"
            cx="12"
            cy="12"
            r="10"
            stroke="currentColor"
            stroke-width="4"
          ></circle>
          <path
            class="opacity-75"
            fill="currentColor"
            d="M4 12a8 8 0 018-8v8H4z"
          ></path>
        </svg>
        <p class="mt-2 text-gray-600">Memuat data...</p>
      </div>

      <table
        v-else
        class="w-full border-collapse border border-gray-300 shadow-md"
      >
        <thead>
          <tr class="bg-gray-100">
            <th class="border border-gray-300 px-4 py-3 text-left w-[150px]">
              Tanggal Opname
            </th>
            <th class="border border-gray-300 px-4 py-3 text-left">
              Nama Barang
            </th>
            <th class="border border-gray-300 px-4 py-3 text-right w-[150px]">
              Stok Sistem
            </th>
            <th class="border border-gray-300 px-4 py-3 text-right w-[150px]">
              Stok Fisik
            </th>
            <th class="border border-gray-300 px-4 py-3 text-right w-[150px]">
              Selisih
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in filteredData" :key="index">
            <td class="border border-gray-300 px-4 py-2 text-sm">
              {{ item.date }}
            </td>
            <td class="border border-gray-300 px-4 py-2 text-sm">
              {{ item.book_name }} ({{ item.product_code }})
            </td>
            <td class="border border-gray-300 px-4 py-2 text-right text-sm">
              {{ item.system_stock !== undefined ? item.system_stock : "-" }}
            </td>
            <td class="border border-gray-300 px-4 py-2 text-right text-sm">
              {{
                item.physical_stock !== undefined ? item.physical_stock : "-"
              }}
            </td>
            <td
              class="border border-gray-300 px-4 py-2 text-right text-sm font-semibold"
              :class="{
                'text-red-600': item.difference < 0,
                'text-green-600': item.difference > 0,
              }"
            >
              {{ item.difference !== undefined ? item.difference : "-" }}
            </td>
          </tr>
          <tr v-if="filteredData.length === 0">
            <td colspan="5" class="text-center py-6 text-gray-500 text-base">
              Tidak ada data Stok Opname dalam periode ini.
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </OwnerLayout>
</template>

<script setup>
import { ref, computed, onMounted, watch } from "vue";
import axios from "axios";
import OwnerLayout from "../components/OwnerLayouts.vue";

const filters = ref({
  startDate: "",
  endDate: "",
});
const rawData = ref([]);
const flattenedData = ref([]);
const loading = ref(false);
const errorMessage = ref("");
const token = localStorage.getItem("token");
const isFilterValid = ref(false);

// Validasi filter
const validateFilters = () => {
  const start = new Date(filters.value.startDate);
  const end = new Date(filters.value.endDate);
  isFilterValid.value =
    filters.value.startDate && filters.value.endDate && start <= end;
  if (!isFilterValid.value) {
    errorMessage.value =
      "Harap masukkan kedua tanggal yang valid (Dari ≤ Sampai).";
  } else {
    errorMessage.value = "";
  }
};

// Ambil data dari backend
const fetchData = async () => {
  loading.value = true;
  errorMessage.value = "";
  try {
    if (!token) {
      errorMessage.value =
        "Token autentikasi tidak ditemukan. Silakan login kembali.";
      return;
    }

    let url =
      "http://localhost:8081/api/v1/owner/reports?type=stock-opnames&detail=true";

    if (filters.value.startDate && filters.value.endDate) {
      url += `&from=${filters.value.startDate}&to=${filters.value.endDate}`;
    }

    const res = await axios.get(url, {
      headers: { Authorization: `Bearer ${token}` },
    });

    rawData.value = res.data.data || [];
    // Sort data berdasarkan tanggal secara descending (terbaru di atas)
    flattenedData.value = rawData.value
      .map((item) => ({
        date: item.date || "N/A",
        book_name: item.book_name || "",
        product_code: item.product_code || "",
        system_stock: item.system_stock,
        physical_stock: item.physical_stock,
        difference: item.difference,
      }))
      .sort((a, b) => new Date(b.date) - new Date(a.date));

    errorMessage.value = ""; // Reset error jika berhasil
  } catch (err) {
    console.error("Gagal ambil data laporan:", err.response || err);
    if (err.response) {
      const status = err.response.status;
      if (status === 404) {
        errorMessage.value =
          "Endpoint Owner 404. Periksa router Go atau pastikan 'type=stock-opnames' valid.";
      } else if (status === 401) {
        errorMessage.value = "Sesi Anda telah berakhir. Silakan login kembali.";
      } else {
        errorMessage.value = `Gagal mengambil data laporan. Status: ${status}. Detail: ${
          err.response.data?.error || err.message
        }`;
      }
    } else {
      errorMessage.value =
        "Gagal terhubung ke server. Pastikan backend berjalan di port 8081.";
    }
  } finally {
    loading.value = false;
  }
};

onMounted(fetchData);

// Logika filter
function applyFilter() {
  if (!isFilterValid.value) {
    errorMessage.value = "Harap masukkan kedua tanggal yang valid.";
    return;
  }
  fetchData();
}

function resetFilter() {
  filters.value.startDate = "";
  filters.value.endDate = "";
  isFilterValid.value = false;
  errorMessage.value = "";
  fetchData();
}

const filteredData = computed(() => flattenedData.value);

// Watch untuk validasi otomatis saat input berubah
watch(() => filters.value, validateFilters, { immediate: true });
</script>
