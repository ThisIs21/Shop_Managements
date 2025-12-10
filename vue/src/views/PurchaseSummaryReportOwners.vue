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
              Purchase Summary Report
            </p>
            <p class="text-[#60758a] text-sm font-normal leading-normal">
              Overview of purchase performance across different periods.
            </p>
          </div>
        </div>

        <!-- Filters -->
        <h3 class="text-[#111418] text-lg font-bold px-4 pb-2 pt-4">Filters</h3>
        <div class="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
          <label class="flex flex-col">
            <span class="text-sm mb-1">From</span>
            <input
              type="date"
              v-model="fromDate"
              class="form-input w-full h-12 rounded-lg border border-[#dbe0e6] p-2"
            />
          </label>
          <label class="flex flex-col">
            <span class="text-sm mb-1">To</span>
            <input
              type="date"
              v-model="toDate"
              class="form-input w-full h-12 rounded-lg border border-[#dbe0e6] p-2"
            />
          </label>
          <button
            @click="fetchReport"
            class="h-12 px-6 bg-[#111418] text-white rounded-lg"
          >
            Apply
          </button>
        </div>

        <!-- Export -->
        <div class="flex justify-stretch">
          <div class="flex flex-1 gap-3 flex-wrap px-4 py-3 justify-start">
            <button
              @click="exportReport('csv')"
              class="flex items-center justify-center h-10 px-4 bg-[#f0f2f5] rounded-lg text-sm font-bold"
            >
              Export to CSV
            </button>
            <button
              @click="exportReport('pdf')"
              class="flex items-center justify-center h-10 px-4 bg-[#f0f2f5] rounded-lg text-sm font-bold"
            >
              Export to PDF
            </button>
          </div>
        </div>

        <!-- Data -->
        <h3 class="text-[#111418] text-lg font-bold px-4 pb-2 pt-4">
          Purchase Data
        </h3>
        <div class="px-4 py-3 @container">
          <div
            class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
          >
            <table class="flex-1">
              <thead>
                <tr class="bg-white">
                  <th class="px-4 py-3 text-left w-[200px] text-sm font-medium">
                    Date
                  </th>
                  <th class="px-4 py-3 text-left w-[150px] text-sm font-medium">
                    Transactions
                  </th>
                  <th class="px-4 py-3 text-left w-[150px] text-sm font-medium">
                    Total
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="(row, idx) in reportData"
                  :key="idx"
                  class="border-t border-[#dbe0e6]"
                >
                  <td class="px-4 py-2 text-sm">{{ row.date }}</td>
                  <td class="px-4 py-2 text-sm text-[#60758a]">
                    {{ row.trx }}
                  </td>
                  <td class="px-4 py-2 text-sm text-[#60758a]">
                    {{ formatCurrency(row.total) }}
                  </td>
                </tr>
                <tr v-if="!reportData.length">
                  <td colspan="3" class="p-4 text-center text-gray-500">
                    No report data available.
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
const fromDate = ref("");
const toDate = ref("");

// Fetch purchase report untuk OWNER
const fetchReport = async () => {
  console.log(">>> Apply clicked", fromDate.value, toDate.value); // 👈 debug log

  const token = localStorage.getItem("token");
  if (!token) {
    alert("Authentication token not found. Please log in again.");
    return;
  }

  try {
    const params = {
      type: "purchases",
    };

    if (fromDate.value) params.from = fromDate.value;
    if (toDate.value) params.to = toDate.value;

    console.log(">>> Params dikirim ke API:", params); // 👈 cek params

    const response = await axios.get(API_URL, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
      params,
    });

    console.log(">>> Raw backend data:", response.data);

    const purchases = response.data.data || [];
    reportData.value = purchases.map((p) => ({
      date: p.period ? new Date(p.period).toLocaleDateString("id-ID") : "-",
      trx: p.trx || 0,
      total: p.total || 0,
    }));
  } catch (error) {
    console.error("Failed to fetch report data:", error);
    alert("Failed to fetch report data: " + error.message);
  }
};

// Export
const exportReport = async (format) => {
  const token = localStorage.getItem("token");
  if (!token) {
    alert("Authentication token not found. Please log in again.");
    return;
  }

  try {
    const params = {
      type: "purchases",
      export: format,
    };
    if (fromDate.value) params.from = fromDate.value;
    if (toDate.value) params.to = toDate.value;

    const response = await axios.get(API_URL, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
      params,
      responseType: "blob",
    });

    const url = window.URL.createObjectURL(new Blob([response.data]));
    const link = document.createElement("a");
    link.href = url;
    link.setAttribute("download", `purchase_report.${format}`);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
  } catch (error) {
    console.error("Export failed:", error);
    alert("Export failed: " + error.message);
  }
};

// Format Rupiah
const formatCurrency = (num) => {
  if (num === null || num === undefined) return "Rp 0";
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
  }).format(num);
};

// Load default report pas halaman dibuka
onMounted(fetchReport);
</script>
