<template>
  <CashierLayouts>
    <div class="bg-gray-50 min-h-screen">
      <!-- Header -->
      <div class="p-8 bg-white border-b border-[#dbe0e6]">
        <h1
          class="text-[#111418] tracking-light text-[32px] font-bold leading-tight"
        >
          Welcome back, {{ userName }}! 👋
        </h1>
        <p class="text-[#60758a] text-sm font-normal leading-normal mt-2">
          Here's your sales performance for today and recent transactions.
        </p>
      </div>

      <div class="p-8">
        <!-- KPI Cards -->
        <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-8">
          <div
            class="rounded-lg p-6 border border-[#dbe0e6] bg-gradient-to-br from-blue-50 to-blue-100"
          >
            <p class="text-[#60758a] text-sm font-medium">Total Sales</p>
            <p class="text-[#111418] text-3xl font-bold leading-tight mt-2">
              {{ formatCurrency(totalSales) }}
            </p>
            <p class="text-blue-600 text-xs font-semibold mt-2">
              {{ salesCount }} transactions
            </p>
          </div>

          <div
            class="rounded-lg p-6 border border-[#dbe0e6] bg-gradient-to-br from-green-50 to-green-100"
          >
            <p class="text-[#60758a] text-sm font-medium">Avg Transaction</p>
            <p class="text-[#111418] text-3xl font-bold leading-tight mt-2">
              {{ formatCurrency(avgTransaction) }}
            </p>
            <p class="text-green-600 text-xs font-semibold mt-2">Per sale</p>
          </div>

          <div
            class="rounded-lg p-6 border border-[#dbe0e6] bg-gradient-to-br from-purple-50 to-purple-100"
          >
            <p class="text-[#60758a] text-sm font-medium">Highest Sale</p>
            <p class="text-[#111418] text-3xl font-bold leading-tight mt-2">
              {{ formatCurrency(highestSale) }}
            </p>
            <p class="text-purple-600 text-xs font-semibold mt-2">
              Today's peak
            </p>
          </div>

          <div
            class="rounded-lg p-6 border border-[#dbe0e6] bg-gradient-to-br from-amber-50 to-amber-100"
          >
            <p class="text-[#60758a] text-sm font-medium">Total Items</p>
            <p class="text-[#111418] text-3xl font-bold leading-tight mt-2">
              {{ totalItems }}
            </p>
            <p class="text-amber-600 text-xs font-semibold mt-2">Sold</p>
          </div>
        </div>

        <!-- Filter Section -->
        <div class="bg-white rounded-lg border border-[#dbe0e6] p-6 mb-8">
          <h3 class="text-[#111418] text-lg font-semibold mb-4">
            Filter Sales History
          </h3>
          <div class="flex flex-wrap gap-4 items-end">
            <div>
              <label class="block text-sm font-medium text-[#111418] mb-2"
                >From Date</label
              >
              <input
                type="date"
                v-model="fromDate"
                class="px-3 py-2 border border-[#dbe0e6] rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-[#111418] mb-2"
                >To Date</label
              >
              <input
                type="date"
                v-model="toDate"
                class="px-3 py-2 border border-[#dbe0e6] rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
            <button
              @click="fetchDashboardData"
              :disabled="loadingHistory"
              class="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg font-medium disabled:opacity-50 transition"
            >
              {{ loadingHistory ? "Loading..." : "Apply Filter" }}
            </button>
            <button
              @click="resetFilter"
              class="px-4 py-2 bg-gray-300 hover:bg-gray-400 text-gray-800 rounded-lg font-medium transition"
            >
              Reset
            </button>
          </div>
          <p v-if="fromDate || toDate" class="text-sm text-[#60758a] mt-3">
            📅 Showing: {{ fromDate || "Start" }} to {{ toDate || "End" }}
          </p>
        </div>

        <!-- Sales History Table -->
        <div
          class="rounded-lg border border-[#dbe0e6] bg-white overflow-hidden shadow-sm"
        >
          <div class="border-b border-[#dbe0e6] p-6">
            <h3 class="text-[#111418] text-lg font-bold leading-tight">
              Sales History
              <span class="text-[#60758a] text-sm font-normal ml-2">
                ({{ fromDate || toDate ? "Filtered" : "Latest 10" }})
              </span>
            </h3>
          </div>

          <div class="overflow-x-auto">
            <table class="w-full">
              <thead>
                <tr class="bg-gray-100 border-b border-[#dbe0e6]">
                  <th
                    class="px-6 py-3 text-left text-xs font-semibold text-[#111418] uppercase"
                  >
                    Date
                  </th>
                  <th
                    class="px-6 py-3 text-left text-xs font-semibold text-[#111418] uppercase"
                  >
                    Transaction #
                  </th>
                  <th
                    class="px-6 py-3 text-left text-xs font-semibold text-[#111418] uppercase"
                  >
                    Customer
                  </th>
                  <th
                    class="px-6 py-3 text-right text-xs font-semibold text-[#111418] uppercase"
                  >
                    Items
                  </th>
                  <th
                    class="px-6 py-3 text-right text-xs font-semibold text-[#111418] uppercase"
                  >
                    Total
                  </th>
                  <th
                    class="px-6 py-3 text-center text-xs font-semibold text-[#111418] uppercase"
                  >
                    Action
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr v-if="loadingHistory" class="border-b border-[#dbe0e6]">
                  <td colspan="6" class="text-center py-8 text-[#60758a]">
                    Loading transactions...
                  </td>
                </tr>
                <tr
                  v-else-if="salesHistory.length === 0"
                  class="border-b border-[#dbe0e6]"
                >
                  <td colspan="6" class="text-center py-8 text-[#60758a]">
                    No sales data found
                  </td>
                </tr>
                <tr
                  v-else
                  v-for="sale in salesHistory"
                  :key="sale.ID"
                  class="border-b border-[#dbe0e6] hover:bg-gray-50 transition"
                >
                  <td class="px-6 py-4 text-sm text-[#111418]">
                    {{ formatDate(sale.Date) }}
                  </td>
                  <td class="px-6 py-4 text-sm text-[#111418] font-medium">
                    {{ sale.transaction_number || `TRX-${sale.ID}` }}
                  </td>
                  <td class="px-6 py-4 text-sm text-[#60758a]">
                    {{ sale.Customer?.name || "Walk-in" }}
                  </td>
                  <td
                    class="px-6 py-4 text-sm text-right text-[#111418] font-medium"
                  >
                    {{ (sale.Items || []).length }} items
                  </td>
                  <td
                    class="px-6 py-4 text-sm text-right text-[#111418] font-bold"
                  >
                    {{ formatCurrency(sale.Total) }}
                  </td>
                  <td class="px-6 py-4 text-center">
                    <button
                      @click="viewDetails(sale.ID)"
                      class="px-3 py-1 bg-blue-100 hover:bg-blue-200 text-blue-700 rounded-lg text-sm font-medium transition"
                    >
                      View
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- Details Modal -->
      <div
        v-if="showDetailsModal"
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
      >
        <div
          class="bg-white rounded-lg max-w-2xl w-full max-h-[90vh] overflow-y-auto shadow-xl"
        >
          <!-- Modal Header -->
          <div
            class="border-b border-[#dbe0e6] p-6 flex justify-between items-center sticky top-0 bg-white"
          >
            <h2 class="text-[#111418] text-xl font-bold">
              Transaction Details
            </h2>
            <button
              @click="closeDetailsModal"
              class="text-[#60758a] hover:text-[#111418] text-2xl font-light"
            >
              ×
            </button>
          </div>

          <!-- Modal Body -->
          <div v-if="selectedSale" class="p-6 space-y-6">
            <!-- Transaction Info -->
            <div class="grid grid-cols-2 gap-6">
              <div>
                <p class="text-[#60758a] text-sm font-medium">Transaction #</p>
                <p class="text-[#111418] text-lg font-bold mt-1">
                  {{
                    selectedSale.transaction_number || `TRX-${selectedSale.ID}`
                  }}
                </p>
              </div>
              <div>
                <p class="text-[#60758a] text-sm font-medium">Date</p>
                <p class="text-[#111418] text-lg font-bold mt-1">
                  {{ formatDate(selectedSale.Date) }}
                </p>
              </div>
              <div>
                <p class="text-[#60758a] text-sm font-medium">Customer</p>
                <p class="text-[#111418] text-lg font-bold mt-1">
                  {{ selectedSale.Customer?.name || "Walk-in Customer" }}
                </p>
              </div>
              <div>
                <p class="text-[#60758a] text-sm font-medium">Payment Method</p>
                <p class="text-[#111418] text-lg font-bold mt-1">
                  {{ selectedSale.payment_method || "Cash" }}
                </p>
              </div>
            </div>

            <!-- Items Table -->
            <div>
              <h3 class="text-[#111418] text-lg font-bold mb-4">
                Items Purchased
              </h3>
              <div class="overflow-x-auto rounded-lg border border-[#dbe0e6]">
                <table class="w-full text-sm">
                  <thead>
                    <tr class="bg-gray-100 border-b border-[#dbe0e6]">
                      <th
                        class="px-4 py-3 text-left text-[#111418] font-semibold"
                      >
                        Product
                      </th>
                      <th
                        class="px-4 py-3 text-right text-[#111418] font-semibold"
                      >
                        Qty
                      </th>
                      <th
                        class="px-4 py-3 text-right text-[#111418] font-semibold"
                      >
                        Price
                      </th>
                      <th
                        class="px-4 py-3 text-right text-[#111418] font-semibold"
                      >
                        Subtotal
                      </th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr
                      v-for="(item, idx) in selectedSale.Items || []"
                      :key="idx"
                      class="border-b border-[#dbe0e6] hover:bg-gray-50"
                    >
                      <td class="px-4 py-3 text-[#111418]">
                        {{ item.Product?.name || "Unknown" }}
                      </td>
                      <td
                        class="px-4 py-3 text-right text-[#111418] font-medium"
                      >
                        {{ item.Qty || 0 }}
                      </td>
                      <td class="px-4 py-3 text-right text-[#111418]">
                        {{ formatCurrency(item.Price || 0) }}
                      </td>
                      <td class="px-4 py-3 text-right text-[#111418] font-bold">
                        {{
                          formatCurrency((item.Qty || 0) * (item.Price || 0))
                        }}
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>

            <!-- Summary -->
            <div
              class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-lg p-4 space-y-2"
            >
              <div class="flex justify-between">
                <p class="text-[#60758a] font-medium">Subtotal:</p>
                <p class="text-[#111418] font-bold">
                  {{ formatCurrency(selectedSale.Total) }}
                </p>
              </div>
              <div class="flex justify-between">
                <p class="text-[#60758a] font-medium">Payment:</p>
                <p class="text-[#111418] font-bold">
                  {{ formatCurrency(selectedSale.Paid || 0) }}
                </p>
              </div>
              <div class="border-t border-blue-200 pt-2 flex justify-between">
                <p class="text-[#111418] font-bold">Change:</p>
                <p class="text-[#111418] font-bold text-lg">
                  {{
                    formatCurrency(
                      (selectedSale.Paid || 0) - (selectedSale.Total || 0)
                    )
                  }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </CashierLayouts>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import axios from "axios";
import CashierLayouts from "@/components/CashierLayouts.vue";

const userName = ref("Cashier");
const salesHistory = ref([]);
const loadingHistory = ref(true);
const fromDate = ref("");
const toDate = ref("");
const showDetailsModal = ref(false);
const selectedSale = ref(null);

const api = axios.create({ baseURL: "http://localhost:8081/api/v1" });
const savedToken = localStorage.getItem("token");
if (savedToken) {
  api.defaults.headers.common["Authorization"] = `Bearer ${savedToken}`;
}

// Computed KPI values
const totalSales = computed(() => {
  return salesHistory.value.reduce((sum, sale) => sum + (sale.Total || 0), 0);
});

const salesCount = computed(() => {
  return salesHistory.value.length;
});

const avgTransaction = computed(() => {
  if (salesCount.value === 0) return 0;
  return totalSales.value / salesCount.value;
});

const highestSale = computed(() => {
  if (salesHistory.value.length === 0) return 0;
  return Math.max(...salesHistory.value.map((s) => s.Total || 0));
});

const totalItems = computed(() => {
  return salesHistory.value.reduce((sum, sale) => {
    return sum + (sale.Items || []).length;
  }, 0);
});

const fetchDashboardData = async () => {
  loadingHistory.value = true;
  try {
    const nameFromStorage = localStorage.getItem("userName");
    if (nameFromStorage) {
      userName.value = nameFromStorage;
    }

    const params = {
      page: 1,
      size: 10,
      orderBy: "date DESC",
    };

    if (fromDate.value) {
      params.start_date = fromDate.value;
    }
    if (toDate.value) {
      params.end_date = toDate.value;
    }

    console.log("Request Params:", params);
    const endpoint =
      fromDate.value || toDate.value
        ? "/history/sales"
        : "/history/sales/latest";
    const response = await api.get(endpoint, { params });
    console.log("API Response:", response.data);

    const data = response.data;
    if (data && data.data) {
      salesHistory.value = data.data.map((sale) => ({
        ...sale,
        Paid: sale.Paid || 0,
        Change: sale.Change || (sale.Paid || 0) - (sale.Total || 0),
        Customer: sale.Customer || { name: "N/A" },
        Items: sale.Items || [],
      }));
    } else {
      console.warn("No 'data' property in response:", data);
      salesHistory.value = [];
      alert("No data from server. Check API response.");
    }
  } catch (error) {
    console.error("Error fetch:", error);
    if (error.response) {
      console.error("Status:", error.response.status);
      console.error("Response data:", error.response.data);
      alert(
        `Error fetching data: ${error.response.status} - ${
          error.response.data.error || error.message
        }`
      );
    } else {
      alert("Network error or server unavailable. Check console.");
    }
    salesHistory.value = [];
  } finally {
    loadingHistory.value = false;
  }
};

const resetFilter = () => {
  fromDate.value = "";
  toDate.value = "";
  fetchDashboardData();
};

const formatDate = (dateStr) => {
  if (!dateStr) return "N/A";
  const date = new Date(dateStr);
  if (isNaN(date.getTime())) {
    const sliced = dateStr.split("T")[0];
    return sliced || "Invalid Date";
  }
  return date.toLocaleDateString("id-ID", {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
  });
};

const formatCurrency = (amount) => {
  return amount !== undefined && amount !== null
    ? new Intl.NumberFormat("id-ID", {
        style: "currency",
        currency: "IDR",
        minimumFractionDigits: 0,
      }).format(amount)
    : "Rp 0";
};

const viewDetails = async (saleId) => {
  try {
    const response = await api.get(`/history/sales/${saleId}`);
    selectedSale.value = {
      ...response.data,
      Paid: response.data.Paid || 0,
      Change:
        response.data.Change ||
        (response.data.Paid || 0) - (response.data.Total || 0),
      Customer: response.data.Customer || { name: "N/A" },
      Items: response.data.Items || [],
    };
    console.log("Selected Sale Data:", selectedSale.value);
    showDetailsModal.value = true;
  } catch (error) {
    console.error("Error fetching sale details:", error);
    alert("Failed to load transaction details. Check console.");
  }
};

const closeDetailsModal = () => {
  showDetailsModal.value = false;
  selectedSale.value = null;
};

onMounted(() => {
  fetchDashboardData();
});
</script>

<style scoped>
/* Tailwind styles applied inline */
</style>
