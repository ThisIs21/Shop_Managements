<template>
  <OwnerLayouts>
    <div class="flex flex-1 flex-col py-5 p-8">
      <div class="layout-content-container flex flex-col flex-1">
        <!-- Header -->
        <div class="flex items-center gap-4 p-4">
          <RouterLink
            to="/report-owners"
            class="flex items-center justify-center h-10 w-10 text-[#111418] rounded-full hover:bg-gray-100 transition-colors"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24px"
              height="24px"
              fill="currentColor"
              viewBox="0 0 256 256"
            >
              <path
                d="M165.66,202.34a8,8,0,0,1-11.32,11.32l-80-80a8,8,0,0,1,0-11.32l80-80a8,8,0,0,1,11.32,11.32L91.31,128Z"
              ></path>
            </svg>
          </RouterLink>
          <div class="flex min-w-72 flex-col gap-1">
            <p class="text-[#111418] text-[32px] font-bold leading-tight">
              Sales Summary Report
            </p>
            <p class="text-[#60758a] text-sm font-normal leading-normal">
              Overview of sales performance across different periods.
            </p>
          </div>
        </div>

        <!-- Filters -->
        <h3 class="text-[#111418] text-lg font-bold px-4 pb-2 pt-4">Filters</h3>
        <div class="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
          <label class="flex flex-col min-w-40 flex-1">
            <select
              v-model="filterBy"
              @change="onFilterByChange"
              class="form-input w-full h-14 rounded-lg border border-[#dbe0e6] bg-white p-[15px] text-base"
            >
              <option value="day">Harian</option>
              <option value="month">Bulanan</option>
            </select>
          </label>

          <!-- Date Range untuk Day -->
          <div v-if="filterBy === 'day'" class="flex gap-2">
            <input
              type="date"
              v-model="startDate"
              @change="fetchReport"
              class="form-input w-32 h-14 rounded-lg border border-[#dbe0e6] bg-white p-2"
            />
            <span class="text-sm text-gray-500">s/d</span>
            <input
              type="date"
              v-model="endDate"
              @change="fetchReport"
              class="form-input w-32 h-14 rounded-lg border border-[#dbe0e6] bg-white p-2"
            />
          </div>

          <!-- Month Select untuk Month -->
          <select
            v-if="filterBy === 'month'"
            v-model="selectedMonth"
            @change="fetchReport"
            class="form-input w-48 h-14 rounded-lg border border-[#dbe0e6] bg-white p-[15px] text-base"
          >
            <option value="">Semua Bulan</option>
            <option v-for="m in months" :key="m.value" :value="m.value">
              {{ m.label }}
            </option>
          </select>
        </div>

        <!-- Export -->
        <div class="flex justify-stretch">
          <div class="flex flex-1 gap-3 flex-wrap px-4 py-3 justify-start">
            <button
              @click="exportReport('csv')"
              class="flex items-center justify-center h-10 px-4 bg-[#f0f2f5] rounded-lg text-sm font-bold"
              :disabled="loading"
            >
              Ekspor ke CSV
            </button>
            <button
              @click="exportReport('pdf')"
              class="flex items-center justify-center h-10 px-4 bg-[#f0f2f5] rounded-lg text-sm font-bold"
              :disabled="loading"
            >
              Ekspor ke PDF
            </button>
          </div>
        </div>

        <!-- Loader -->
        <div v-if="loading" class="p-4 text-center text-gray-500">
          Memuat laporan...
        </div>

        <!-- Sales Data -->
        <h3 class="text-[#111418] text-lg font-bold px-4 pb-2 pt-4">
          Data Penjualan
        </h3>
        <div class="px-4 py-3 @container">
          <div
            class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
          >
            <table class="flex-1">
              <thead>
                <tr class="bg-white">
                  <th class="px-4 py-3 text-left w-[400px] text-sm font-medium">
                    Periode
                  </th>
                  <th class="px-4 py-3 text-left w-[400px] text-sm font-medium">
                    Transaksi
                  </th>
                  <th class="px-4 py-3 text-left w-[400px] text-sm font-medium">
                    Total Penjualan
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="row in reportData"
                  :key="row.period"
                  class="border-t border-[#dbe0e6]"
                >
                  <td class="px-4 py-2 text-sm">{{ row.period }}</td>
                  <td class="px-4 py-2 text-sm text-[#60758a]">
                    {{ row.trx }}
                  </td>
                  <td class="px-4 py-2 text-sm text-[#60758a]">
                    {{ formatCurrency(row.total) }}
                  </td>
                </tr>
                <tr v-if="!reportData.length && !loading">
                  <td colspan="3" class="p-4 text-center text-gray-500">
                    Tidak ada data laporan tersedia.
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- Top 10 Best-Selling Products -->
        <h3 class="text-[#111418] text-lg font-bold px-4 pb-2 pt-4">
          10 Produk Terlaris
        </h3>
        <div class="px-4 py-3 @container">
          <div
            class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
          >
            <table class="flex-1">
              <thead>
                <tr class="bg-white">
                  <th class="px-4 py-3 text-left w-[400px] text-sm font-medium">
                    Nama Produk
                  </th>
                  <th class="px-4 py-3 text-left w-[400px] text-sm font-medium">
                    Total Terjual
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="product in topProducts"
                  :key="product.product_name"
                  class="border-t border-[#dbe0e6]"
                >
                  <td class="px-4 py-2 text-sm">{{ product.product_name }}</td>
                  <td class="px-4 py-2 text-sm text-[#60758a]">
                    {{ product.total_sold }}
                  </td>
                </tr>
                <tr v-if="!topProducts.length && !loading">
                  <td colspan="2" class="p-4 text-center text-gray-500">
                    Tidak ada data produk terlaris tersedia.
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- Top 10 Least-Selling Products -->
        <h3 class="text-[#111418] text-lg font-bold px-4 pb-2 pt-4">
          10 Produk Paling Tidak Laku
        </h3>
        <div class="px-4 py-3 @container">
          <div
            class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
          >
            <table class="flex-1">
              <thead>
                <tr class="bg-white">
                  <th class="px-4 py-3 text-left w-[400px] text-sm font-medium">
                    Nama Produk
                  </th>
                  <th class="px-4 py-3 text-left w-[400px] text-sm font-medium">
                    Total Terjual
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="product in bottomProducts"
                  :key="product.product_name"
                  class="border-t border-[#dbe0e6]"
                >
                  <td class="px-4 py-2 text-sm">{{ product.product_name }}</td>
                  <td class="px-4 py-2 text-sm text-[#60758a]">
                    {{ product.total_sold }}
                  </td>
                </tr>
                <tr v-if="!bottomProducts.length && !loading">
                  <td colspan="2" class="p-4 text-center text-gray-500">
                    Tidak ada data produk tidak laku tersedia.
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </OwnerLayouts>
</template>

<script setup>
import OwnerLayouts from "../components/OwnerLayouts.vue";
import { RouterLink } from "vue-router";
import axios from "axios";
import { ref, onMounted } from "vue";

const API_URL = "http://localhost:8081/api/v1/owner/reports";

const reportData = ref([]);
const topProducts = ref([]);
const bottomProducts = ref([]);
const filterBy = ref("month");
const startDate = ref("");
const endDate = ref("");
const selectedMonth = ref("");

const months = [
  { label: "Januari", value: "1" },
  { label: "Februari", value: "2" },
  { label: "Maret", value: "3" },
  { label: "April", value: "4" },
  { label: "Mei", value: "5" },
  { label: "Juni", value: "6" },
  { label: "Juli", value: "7" },
  { label: "Agustus", value: "8" },
  { label: "September", value: "9" },
  { label: "Oktober", value: "10" },
  { label: "November", value: "11" },
  { label: "Desember", value: "12" },
];

const loading = ref(false);

const onFilterByChange = () => {
  // Reset filter saat ganti by
  startDate.value = "";
  endDate.value = "";
  selectedMonth.value = "";
  fetchReport();
};

const getParams = () => {
  const params = {
    by: filterBy.value,
  };

  if (filterBy.value === "day" && startDate.value && endDate.value) {
    params.start_date = startDate.value;
    params.end_date = endDate.value;
  }

  if (filterBy.value === "month" && selectedMonth.value) {
    params.month = selectedMonth.value;
  }

  return params;
};

const fetchReport = async () => {
  loading.value = true;

  const token = localStorage.getItem("token");
  if (!token) {
    console.error("Token tidak ditemukan di localStorage");
    loading.value = false;
    return;
  }

  const baseParams = getParams();

  // Fetch sales data
  try {
    const salesRes = await axios.get(API_URL, {
      params: { ...baseParams, type: "sales" },
      headers: { Authorization: `Bearer ${token}` },
    });
    console.log("Sales Response:", salesRes.data);
    reportData.value = salesRes.data.data || [];
  } catch (err) {
    console.error(
      "Gagal mengambil data penjualan:",
      err.response?.data || err.message
    );
    reportData.value = [];
  }

  // Fetch top 10 best-selling products
  try {
    const topRes = await axios.get(API_URL, {
      params: { ...baseParams, type: "top-products", limit: 10 },
      headers: { Authorization: `Bearer ${token}` },
    });
    console.log("Top Products Response:", topRes.data);
    topProducts.value = topRes.data.data || [];
  } catch (err) {
    console.error(
      "Gagal mengambil produk terlaris:",
      err.response?.data || err.message
    );
    topProducts.value = [];
  }

  // Fetch top 10 least-selling products
  try {
    const bottomRes = await axios.get(API_URL, {
      params: { ...baseParams, type: "bottom-products", limit: 10 },
      headers: { Authorization: `Bearer ${token}` },
    });
    console.log("Bottom Products Response:", bottomRes.data);
    bottomProducts.value = bottomRes.data.data || [];
  } catch (err) {
    console.error(
      "Gagal mengambil produk tidak laku:",
      err.response?.data || err.message
    );
    bottomProducts.value = [];
  }

  loading.value = false;
};

const exportReport = async (format) => {
  try {
    loading.value = true;
    const token = localStorage.getItem("token");
    const params = getParams();

    const res = await axios.get(API_URL, {
      params: { ...params, type: "sales", export: format },
      headers: { Authorization: `Bearer ${token}` },
      responseType: "blob",
    });

    const url = window.URL.createObjectURL(new Blob([res.data]));
    const link = document.createElement("a");
    link.href = url;
    link.setAttribute("download", `sales_report.${format}`);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    window.URL.revokeObjectURL(url);
  } catch (err) {
    console.error("Ekspor gagal:", err.response?.data || err.message);
  } finally {
    loading.value = false;
  }
};

const formatCurrency = (num) => {
  if (!num) return "Rp 0";
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
  }).format(num);
};

onMounted(fetchReport);
</script>
