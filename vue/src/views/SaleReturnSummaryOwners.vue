<template>
  <OwnerLayouts>
    <div class="flex flex-1 flex-col py-5 p-8">
      <div class="layout-content-container flex flex-col flex-1">
        <div class="flex items-center gap-4 p-4">
          <RouterLink
            to="/report-owners"
            class="flex items-center justify-center h-10 w-10 text-[#111418] rounded-full hover:bg-gray-100 transition-colors"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              fill="currentColor"
            >
              <path
                d="M165.66,202.34a8,8,0,0,1-11.32,11.32l-80-80a8,8,0,0,1,0-11.32l80-80a8,8,0,0,1,11.32,11.32L91.31,128Z"
              />
            </svg>
          </RouterLink>
          <div class="flex min-w-72 flex-col gap-1">
            <p
              class="text-[#111418] tracking-light text-[32px] font-bold leading-tight"
            >
              Sale Return Summary
            </p>
            <p class="text-[#60758a] text-sm font-normal leading-normal">
              View and export sale return data.
            </p>
          </div>
        </div>

        <div class="flex flex-wrap items-center justify-between gap-3 p-4">
          <label class="flex flex-col min-w-40 h-12 w-full max-w-64">
            <div class="flex w-full flex-1 items-stretch rounded-lg h-full">
              <div
                class="text-[#60758a] flex border-none bg-[#f0f2f5] items-center justify-center pl-4 rounded-l-lg border-r-0"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="24"
                  height="24"
                  fill="currentColor"
                >
                  <path
                    d="M229.66,218.34l-50.07-50.06a88.11,88.11,0,1,0-11.31,11.31l50.06,50.07a8,8,0,0,0,11.32-11.32ZM40,112a72,72,0,1,1,72,72A72.08,72.08,0,0,1,40,112Z"
                  />
                </svg>
              </div>
              <input
                v-model="search"
                placeholder="Search"
                class="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border-none bg-[#f0f2f5] h-full px-4 placeholder:text-[#60758a]"
              />
            </div>
          </label>

          <div class="flex gap-3 flex-wrap">
            <button
              class="flex h-8 items-center justify-center gap-x-2 rounded-lg bg-[#f0f2f5] pl-4 pr-2"
            >
              <p class="text-[#111418] text-sm font-medium leading-normal">
                {{ selectedRange }}
              </p>
            </button>
          </div>
        </div>

        <div class="px-4 py-3 @container">
          <div
            class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
          >
            <table class="flex-1">
              <thead>
                <tr class="bg-white">
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium"
                  >
                    Period
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium"
                  >
                    Return Count
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium"
                  >
                    Total Return Value
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr v-if="loading">
                  <td colspan="3" class="text-center py-4">Loading...</td>
                </tr>
                <tr
                  v-for="(item, index) in filteredReturns"
                  :key="index"
                  class="border-t border-[#dbe0e6]"
                >
                  <td class="px-4 py-2 text-[#111418] text-sm">
                    {{ item.date }}
                  </td>
                  <td class="px-4 py-2 text-[#60758a] text-sm">
                    {{ item.count }}
                  </td>
                  <td class="px-4 py-2 text-[#60758a] text-sm">
                    Rp {{ formatRupiah(item.total) }}
                  </td>
                </tr>
                <tr v-if="!loading && filteredReturns.length === 0">
                  <td colspan="3" class="text-center py-4">
                    No Data Available
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div class="flex justify-stretch">
          <div class="flex flex-1 gap-3 flex-wrap px-4 py-3 justify-end">
            <button
              @click="exportCSV"
              class="flex items-center justify-center rounded-lg h-10 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-bold"
            >
              Export CSV
            </button>
            <button
              @click="exportPDF"
              class="flex items-center justify-center rounded-lg h-10 px-4 bg-[#0d80f2] text-white text-sm font-bold"
            >
              Export PDF
            </button>
          </div>
        </div>
      </div>
    </div>
  </OwnerLayouts>
</template>

<script setup>
// **PERBAIKAN 1: Sesuaikan path impor jika perlu.** // Asumsi 'OwnerLayouts.vue' ada di "../components/"
import OwnerLayouts from "../components/OwnerLayouts.vue";
import { RouterLink, useRouter } from "vue-router";
import { ref, onMounted, computed } from "vue";
import axios from "axios";

const api = axios.create({ baseURL: "http://localhost:8081/api/v1" });
const savedToken = localStorage.getItem("token");
const router = useRouter();

if (savedToken) {
  api.defaults.headers.common["Authorization"] = `Bearer ${savedToken}`;
} else {
  router.push("/login");
}

const tableData = ref([]);
const loading = ref(false);
const search = ref("");
const selectedRange = ref("Daily");

const fetchReturns = async () => {
  loading.value = true;
  try {
    const res = await api.get("/owner/sale-returns");

    // **PERBAIKAN 2: Akses data yang benar dari respons backend.**
    // Asumsi backend Gin menggunakan {"data": [array]}
    const returnsArray = res.data.data || res.data;

    if (Array.isArray(returnsArray)) {
      tableData.value = returnsArray.map((item) => ({
        // Sesuaikan properti item di bawah ini agar sesuai dengan struct SaleReturn di backend Anda.
        // Misalnya: SaleReturn memiliki field CreatedAt dan TotalValue
        date: item.date || item.created_at || "N/A",
        // Asumsi count dihitung sebagai 1 per item, atau ada field di backend (misal: return_count)
        count: item.return_count || 1,
        total: item.amount || item.total_return_value || item.total || 0,
      }));
    } else {
      console.error("Respons API tidak mengembalikan array data:", res.data);
      tableData.value = []; // Set ke array kosong agar tidak terjadi error map
    }
  } catch (err) {
    console.error("Failed to fetch sale returns", err);
    if (err.response?.status === 401) {
      router.push("/login");
    }
    tableData.value = [];
  }
  loading.value = false;
};

const filteredReturns = computed(() =>
  tableData.value.filter((r) =>
    r.date.toLowerCase().includes(search.value.toLowerCase())
  )
);

const formatRupiah = (value) => new Intl.NumberFormat("id-ID").format(value);

const exportCSV = () => {
  let csvContent = "Period,Return Count,Total Return Value\n";
  tableData.value.forEach((r) => {
    csvContent += `${r.date},${r.count},${r.total}\n`;
  });
  const blob = new Blob([csvContent], { type: "text/csv;charset=utf-8;" });
  const link = document.createElement("a");
  link.href = URL.createObjectURL(blob);
  link.setAttribute("download", "sale_return_summary.csv");
  link.click();
};

const exportPDF = () => window.print();

onMounted(fetchReturns);
</script>

<style scoped>
@container (max-width: 120px) {
  .table-column-120 {
    display: none;
  }
}
@container (max-width: 240px) {
  .table-column-240 {
    display: none;
  }
}
@container (max-width: 360px) {
  .table-column-360 {
    display: none;
  }
}
</style>
