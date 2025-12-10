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
            <p
              class="text-[#111418] tracking-light text-[32px] font-bold leading-tight"
            >
              Purchase Return Summary
            </p>
            <p class="text-[#60758a] text-sm font-normal leading-normal">
              View and export purchase return data.
            </p>
          </div>
        </div>

        <!-- Tabs Daily/Monthly/Yearly -->
        <div class="flex flex-col">
          <div class="flex border-b border-[#dbe0e6] px-4 gap-8">
            <a
              class="flex flex-col items-center justify-center pb-[13px] pt-4 cursor-pointer"
              :class="{
                'border-b-[3px] border-b-[#111418] text-[#111418]':
                  selectedPeriod === 'daily',
                'border-b-[3px] border-b-transparent text-[#60758a]':
                  selectedPeriod !== 'daily',
              }"
              @click="changePeriod('daily')"
            >
              <p class="text-sm font-bold">Daily</p>
            </a>

            <a
              class="flex flex-col items-center justify-center pb-[13px] pt-4 cursor-pointer"
              :class="{
                'border-b-[3px] border-b-[#111418] text-[#111418]':
                  selectedPeriod === 'monthly',
                'border-b-[3px] border-b-transparent text-[#60758a]':
                  selectedPeriod !== 'monthly',
              }"
              @click="changePeriod('monthly')"
            >
              <p class="text-sm font-bold">Monthly</p>
            </a>

            <a
              class="flex flex-col items-center justify-center pb-[13px] pt-4 cursor-pointer"
              :class="{
                'border-b-[3px] border-b-[#111418] text-[#111418]':
                  selectedPeriod === 'yearly',
                'border-b-[3px] border-b-transparent text-[#60758a]':
                  selectedPeriod !== 'yearly',
              }"
              @click="changePeriod('yearly')"
            >
              <p class="text-sm font-bold">Yearly</p>
            </a>
          </div>
        </div>

        <!-- Table -->
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
                <tr
                  v-for="row in rows"
                  :key="row.period"
                  class="border-t border-t-[#dbe0e6]"
                >
                  <td class="h-[72px] px-4 py-2 text-[#111418] text-sm">
                    {{ row.period }}
                  </td>
                  <td class="h-[72px] px-4 py-2 text-[#60758a] text-sm">
                    {{ row.trx }}
                  </td>
                  <td class="h-[72px] px-4 py-2 text-[#60758a] text-sm">
                    Rp {{ formatNumber(row.total) }}
                  </td>
                </tr>
                <tr v-if="rows.length === 0">
                  <td colspan="3" class="text-center text-gray-500 py-6">
                    No data available
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- Export -->
        <div class="flex justify-stretch">
          <div class="flex flex-1 gap-3 flex-wrap px-4 py-3 justify-end">
            <button
              @click="exportCSV"
              class="flex min-w-[84px] cursor-pointer items-center justify-center rounded-lg h-10 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-bold"
            >
              Export to CSV
            </button>
            <button
              @click="exportPDF"
              class="flex min-w-[84px] cursor-pointer items-center justify-center rounded-lg h-10 px-4 bg-[#0d80f2] text-white text-sm font-bold"
            >
              Export to PDF
            </button>
          </div>
        </div>
      </div>
    </div>
  </OwnerLayouts>
</template>

<script setup>
import { ref, onMounted } from "vue";
import OwnerLayouts from "../components/OwnerLayouts.vue";
import { RouterLink } from "vue-router";
import axios from "axios";
import jsPDF from "jspdf";
import autoTable from "jspdf-autotable";

const selectedPeriod = ref("monthly");
const rows = ref([]);
const token = localStorage.getItem("token");

// Fetch data dari API
const fetchData = async () => {
  try {
    const res = await axios.get("http://localhost:8081/api/v1/owner/reports", {
      params: {
        type: "purchase-returns",
        by:
          selectedPeriod.value === "daily"
            ? "day"
            : selectedPeriod.value === "monthly"
            ? "month"
            : "year",
      },
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    rows.value = res.data.data || [];
  } catch (err) {
    console.error("Failed to fetch data:", err);
  }
};

const changePeriod = (period) => {
  selectedPeriod.value = period;
  fetchData();
};

// Export CSV
const exportCSV = () => {
  if (!rows.value.length) return;
  let csvContent = "data:text/csv;charset=utf-8,";
  csvContent += "Period,Return Count,Total Return Value\n";
  rows.value.forEach((row) => {
    csvContent += `${row.period},${row.trx},${row.total}\n`;
  });
  const encodedUri = encodeURI(csvContent);
  const link = document.createElement("a");
  link.setAttribute("href", encodedUri);
  link.setAttribute("download", "purchase_return_summary.csv");
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
};

// Export PDF
const exportPDF = () => {
  const doc = new jsPDF();
  doc.text("Purchase Return Summary", 14, 16);
  autoTable(doc, {
    head: [["Period", "Return Count", "Total Return Value"]],
    body: rows.value.map((r) => [r.period, r.trx, r.total]),
  });
  doc.save("purchase_return_summary.pdf");
};

const formatNumber = (num) => {
  return new Intl.NumberFormat("id-ID").format(num);
};

onMounted(() => {
  fetchData();
});
</script>
