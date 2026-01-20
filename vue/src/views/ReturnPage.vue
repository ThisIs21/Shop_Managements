<template>
  <CashierLayouts>
    <div class="bg-gray-50 min-h-screen">
      <!-- Header -->
      <div class="p-8 bg-white border-b border-[#dbe0e6]">
        <h1 class="text-[#111418] text-[32px] font-bold leading-tight">
          Sales Returns Management 📦
        </h1>
        <p class="text-[#60758a] text-sm font-normal leading-normal mt-2">
          Manage product returns and track return status
        </p>
      </div>

      <div class="p-8">
        <!-- Tab Navigation -->
        <div class="bg-white rounded-lg border border-[#dbe0e6] mb-8 shadow-sm">
          <div class="flex border-b border-[#dbe0e6]">
            <button
              @click="activeTab = 'sales'"
              :class="[
                'flex-1 py-4 px-6 font-semibold text-center transition',
                activeTab === 'sales'
                  ? 'bg-blue-600 text-white border-b-2 border-blue-600'
                  : 'bg-gray-50 text-[#60758a] hover:bg-gray-100',
              ]"
            >
              <span class="text-lg">📋</span> Sales History
            </button>
            <button
              @click="activeTab = 'returns'"
              :class="[
                'flex-1 py-4 px-6 font-semibold text-center transition',
                activeTab === 'returns'
                  ? 'bg-blue-600 text-white border-b-2 border-blue-600'
                  : 'bg-gray-50 text-[#60758a] hover:bg-gray-100',
              ]"
            >
              <span class="text-lg">📤</span> My Returns
            </button>
          </div>
        </div>

        <!-- Tab 1: Sales History -->
        <div v-if="activeTab === 'sales'" class="animate-fadeIn">
          <!-- Filter Section -->
          <div
            class="bg-white rounded-lg border border-[#dbe0e6] p-6 mb-6 shadow-sm"
          >
            <h3 class="text-[#111418] text-lg font-semibold mb-4">
              Filter Sales
            </h3>
            <div class="flex flex-wrap gap-4 items-end">
              <div>
                <label class="block text-sm font-medium text-[#111418] mb-2"
                  >From Date</label
                >
                <input
                  type="date"
                  v-model="fromDate"
                  class="px-4 py-2 border border-[#dbe0e6] rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-[#111418] mb-2"
                  >To Date</label
                >
                <input
                  type="date"
                  v-model="toDate"
                  class="px-4 py-2 border border-[#dbe0e6] rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
              <button
                @click="fetchSalesHistory"
                :disabled="loadingHistory"
                class="px-6 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg font-medium disabled:opacity-50 transition"
              >
                {{ loadingHistory ? "Loading..." : "Apply" }}
              </button>
              <button
                @click="resetFilter"
                class="px-6 py-2 bg-gray-300 hover:bg-gray-400 text-gray-800 rounded-lg font-medium transition"
              >
                Reset
              </button>
            </div>
            <p v-if="fromDate || toDate" class="text-sm text-[#60758a] mt-3">
              📅 Showing: {{ fromDate || "Start" }} to {{ toDate || "End" }}
            </p>
          </div>

          <!-- Sales Table -->
          <div
            class="bg-white rounded-lg border border-[#dbe0e6] overflow-hidden shadow-sm"
          >
            <div class="border-b border-[#dbe0e6] p-6">
              <h3 class="text-[#111418] text-lg font-bold">
                Available Sales
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
                      Transaction
                    </th>
                    <th
                      class="px-6 py-3 text-left text-xs font-semibold text-[#111418] uppercase"
                    >
                      Customer
                    </th>
                    <th
                      class="px-6 py-3 text-right text-xs font-semibold text-[#111418] uppercase"
                    >
                      Amount
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
                    <td colspan="5" class="text-center py-8 text-[#60758a]">
                      Loading transactions...
                    </td>
                  </tr>
                  <tr
                    v-else-if="salesHistory.length === 0"
                    class="border-b border-[#dbe0e6]"
                  >
                    <td colspan="5" class="text-center py-8 text-[#60758a]">
                      No sales found
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
                      class="px-6 py-4 text-sm text-right text-[#111418] font-bold"
                    >
                      {{ formatCurrency(sale.Total) }}
                    </td>
                    <td class="px-6 py-4 text-center">
                      <button
                        @click="openReturnModal(sale.ID)"
                        class="px-4 py-2 bg-green-100 hover:bg-green-200 text-green-700 rounded-lg text-sm font-semibold transition"
                      >
                        Create Return
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        <!-- Tab 2: Returns List -->
        <div v-if="activeTab === 'returns'" class="animate-fadeIn">
          <!-- Filter Section -->
          <div
            class="bg-white rounded-lg border border-[#dbe0e6] p-6 mb-6 shadow-sm"
          >
            <h3 class="text-[#111418] text-lg font-semibold mb-4">
              Filter Returns
            </h3>
            <div class="flex flex-wrap gap-4 items-end">
              <div>
                <label class="block text-sm font-medium text-[#111418] mb-2"
                  >From Date</label
                >
                <input
                  type="date"
                  v-model="returnFromDate"
                  class="px-4 py-2 border border-[#dbe0e6] rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-[#111418] mb-2"
                  >To Date</label
                >
                <input
                  type="date"
                  v-model="returnToDate"
                  class="px-4 py-2 border border-[#dbe0e6] rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-[#111418] mb-2"
                  >Status</label
                >
                <select
                  v-model="returnStatusFilter"
                  class="px-4 py-2 border border-[#dbe0e6] rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="">All Status</option>
                  <option value="PENDING">Pending</option>
                  <option value="APPROVED">Approved</option>
                  <option value="REJECTED">Rejected</option>
                </select>
              </div>
              <button
                @click="fetchSaleReturns"
                :disabled="loadingReturns"
                class="px-6 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg font-medium disabled:opacity-50 transition"
              >
                {{ loadingReturns ? "Loading..." : "Apply" }}
              </button>
              <button
                @click="resetReturnFilter"
                class="px-6 py-2 bg-gray-300 hover:bg-gray-400 text-gray-800 rounded-lg font-medium transition"
              >
                Reset
              </button>
            </div>
          </div>

          <!-- Returns Table -->
          <div
            class="bg-white rounded-lg border border-[#dbe0e6] overflow-hidden shadow-sm"
          >
            <div class="border-b border-[#dbe0e6] p-6">
              <h3 class="text-[#111418] text-lg font-bold">
                My Returns List
                <span class="text-[#60758a] text-sm font-normal ml-2">
                  ({{
                    returnFromDate || returnToDate || returnStatusFilter
                      ? "Filtered"
                      : "Latest 10"
                  }})
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
                      Return No.
                    </th>
                    <th
                      class="px-6 py-3 text-left text-xs font-semibold text-[#111418] uppercase"
                    >
                      Transaction
                    </th>
                    <th
                      class="px-6 py-3 text-left text-xs font-semibold text-[#111418] uppercase"
                    >
                      Date
                    </th>
                    <th
                      class="px-6 py-3 text-right text-xs font-semibold text-[#111418] uppercase"
                    >
                      Amount
                    </th>
                    <th
                      class="px-6 py-3 text-center text-xs font-semibold text-[#111418] uppercase"
                    >
                      Status
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-if="loadingReturns" class="border-b border-[#dbe0e6]">
                    <td colspan="5" class="text-center py-8 text-[#60758a]">
                      Loading returns...
                    </td>
                  </tr>
                  <tr
                    v-else-if="saleReturns.length === 0"
                    class="border-b border-[#dbe0e6]"
                  >
                    <td colspan="5" class="text-center py-8 text-[#60758a]">
                      No returns created yet
                    </td>
                  </tr>
                  <tr
                    v-else
                    v-for="ret in saleReturns"
                    :key="ret.id"
                    class="border-b border-[#dbe0e6] hover:bg-gray-50 transition"
                  >
                    <td class="px-6 py-4 text-sm text-[#111418] font-medium">
                      RTN-{{ formatReturnNumber(ret.id) }}
                    </td>
                    <td class="px-6 py-4 text-sm text-[#111418]">
                      {{ ret.sale?.transaction_number || `TRX-${ret.sale_id}` }}
                    </td>
                    <td class="px-6 py-4 text-sm text-[#60758a]">
                      {{ formatDate(ret.date) }}
                    </td>
                    <td
                      class="px-6 py-4 text-sm text-right text-[#111418] font-bold"
                    >
                      {{ formatCurrency(ret.total) }}
                    </td>
                    <td class="px-6 py-4 text-center">
                      <span
                        :class="[
                          'px-3 py-1 rounded-lg text-xs font-semibold',
                          getStatusClass(ret.status),
                        ]"
                      >
                        {{ getStatusText(ret.status) }}
                      </span>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>

      <!-- Create Return Modal -->
      <div
        v-if="showModal"
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
      >
        <div
          class="bg-white rounded-lg max-w-4xl w-full max-h-[90vh] overflow-y-auto shadow-xl"
        >
          <!-- Modal Header -->
          <div class="border-b border-[#dbe0e6] p-6 sticky top-0 bg-white">
            <div class="flex justify-between items-center">
              <h2 class="text-[#111418] text-2xl font-bold">
                Create Sales Return
              </h2>
              <button
                @click="closeModal"
                class="text-[#60758a] hover:text-[#111418] text-3xl font-light"
              >
                ×
              </button>
            </div>
          </div>

          <!-- Modal Body -->
          <div class="p-6 space-y-6">
            <!-- Sale Info -->
            <div
              v-if="selectedSale"
              class="bg-gradient-to-br from-blue-50 to-blue-100 p-6 rounded-lg"
            >
              <h3 class="text-[#111418] font-semibold mb-4">
                Original Transaction
              </h3>
              <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
                <div>
                  <p class="text-[#60758a] text-xs font-medium">
                    Transaction #
                  </p>
                  <p class="text-[#111418] font-bold">
                    {{
                      selectedSale.transaction_number ||
                      `TRX-${selectedSale.ID}`
                    }}
                  </p>
                </div>
                <div>
                  <p class="text-[#60758a] text-xs font-medium">Date</p>
                  <p class="text-[#111418] font-bold">
                    {{ formatDate(selectedSale.Date) }}
                  </p>
                </div>
                <div>
                  <p class="text-[#60758a] text-xs font-medium">Customer</p>
                  <p class="text-[#111418] font-bold">
                    {{ selectedSale.Customer?.name || "Walk-in" }}
                  </p>
                </div>
                <div>
                  <p class="text-[#60758a] text-xs font-medium">Total</p>
                  <p class="text-[#111418] font-bold">
                    {{ formatCurrency(selectedSale.Total) }}
                  </p>
                </div>
              </div>
            </div>

            <!-- Return Items -->
            <div>
              <h3 class="text-[#111418] text-lg font-semibold mb-4">
                Items to Return
              </h3>
              <div class="space-y-4">
                <div
                  v-for="(retItem, index) in returnForm.items"
                  :key="index"
                  class="border border-[#dbe0e6] rounded-lg p-4 bg-gray-50"
                >
                  <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
                    <!-- Product Info -->
                    <div class="md:col-span-1">
                      <p class="text-[#60758a] text-xs font-medium">Product</p>
                      <p class="text-[#111418] font-semibold text-sm">
                        {{ selectedSale?.Items[index]?.Product?.name || "N/A" }}
                      </p>
                      <p class="text-[#60758a] text-xs mt-1">
                        Original:
                        {{ selectedSale?.Items[index]?.Qty || 0 }} units
                      </p>
                    </div>

                    <!-- Quantity Input -->
                    <div>
                      <label
                        class="block text-[#60758a] text-xs font-medium mb-2"
                        >Return Qty</label
                      >
                      <input
                        type="number"
                        v-model.number="retItem.qty"
                        min="0"
                        :max="selectedSale?.Items[index]?.Qty || 0"
                        class="w-full px-3 py-2 border border-[#dbe0e6] rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                      />
                    </div>

                    <!-- Reason Textarea -->
                    <div class="md:col-span-2">
                      <label
                        class="block text-[#60758a] text-xs font-medium mb-2"
                        >Reason</label
                      >
                      <textarea
                        v-model="retItem.reason"
                        placeholder="Why is this item being returned?"
                        rows="2"
                        class="w-full px-3 py-2 border border-[#dbe0e6] rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 text-sm"
                      ></textarea>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- Additional Notes -->
            <div>
              <label class="block text-[#111418] text-sm font-semibold mb-2"
                >Additional Notes</label
              >
              <textarea
                v-model="returnForm.notes"
                placeholder="Any additional information about this return..."
                rows="3"
                class="w-full px-4 py-3 border border-[#dbe0e6] rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              ></textarea>
            </div>

            <!-- Actions -->
            <div class="flex justify-end gap-3 pt-4 border-t border-[#dbe0e6]">
              <button
                @click="closeModal"
                class="px-6 py-2 bg-gray-300 hover:bg-gray-400 text-gray-800 rounded-lg font-medium transition"
              >
                Cancel
              </button>
              <button
                @click="submitReturn"
                :disabled="submitting || !hasValidItems"
                class="px-6 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg font-medium disabled:opacity-50 transition"
              >
                {{ submitting ? "Creating..." : "Create Return" }}
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </CashierLayouts>
</template>

<style scoped>
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(5px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.animate-fadeIn {
  animation: fadeIn 0.3s ease-in-out;
}
</style>

<script setup>
import { ref, onMounted, computed } from "vue";
import axios from "axios";
import CashierLayouts from "@/components/CashierLayouts.vue";

const activeTab = ref("sales");
const userName = ref("Kasir");
const salesHistory = ref([]);
const saleReturns = ref([]);
const loadingHistory = ref(true);
const loadingReturns = ref(false);
const fromDate = ref("");
const toDate = ref("");
const returnFromDate = ref("");
const returnToDate = ref("");
const returnStatusFilter = ref("");

// Modal states
const showModal = ref(false);
const selectedSale = ref(null);
const returnForm = ref({
  items: [],
  notes: "",
});
const submitting = ref(false);

// Computed untuk validasi items
const hasValidItems = computed(() => {
  return returnForm.value.items.some(
    (item) => item.qty > 0 && item.reason.trim()
  );
});

const api = axios.create({ baseURL: "http://localhost:8081/api/v1" });
const savedToken = localStorage.getItem("token");
if (savedToken) {
  api.defaults.headers.common["Authorization"] = `Bearer ${savedToken}`;
}

const fetchSalesHistory = async () => {
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
      salesHistory.value = data.data;
    } else {
      salesHistory.value = [];
      alert("Tidak ada data dari server.");
    }
  } catch (error) {
    console.error("Error fetch:", error);
    if (error.response) {
      console.error("Status:", error.response.status);
      console.error("Response data:", error.response.data);
      alert(
        `Error mengambil data: ${error.response.status} - ${
          error.response.data.error || error.message
        }`
      );
    } else {
      alert("Kesalahan jaringan atau server tidak tersedia.");
    }
    salesHistory.value = [];
  } finally {
    loadingHistory.value = false;
  }
};

const fetchSaleReturns = async () => {
  loadingReturns.value = true;
  try {
    const params = {
      page: 1,
      size: 10, // Membatasi ke 10 retur terakhir
      orderBy: "date DESC",
    };

    if (returnFromDate.value) {
      params.start_date = returnFromDate.value;
    }
    if (returnToDate.value) {
      params.end_date = returnToDate.value;
    }
    if (returnStatusFilter.value) {
      params.status = returnStatusFilter.value;
    }

    console.log("Return Params:", params);
    const response = await api.get("/kasir/sale-returns", { params });
    saleReturns.value = response.data.data || [];
    console.log("Fetched returns:", saleReturns.value);
  } catch (error) {
    console.error(
      "Error fetching returns:",
      error.response ? error.response.data : error.message
    );
    alert(
      `Gagal mengambil daftar retur: ${
        error.response?.data?.error || error.message
      }`
    );
    saleReturns.value = [];
  } finally {
    loadingReturns.value = false;
  }
};

const resetFilter = () => {
  fromDate.value = "";
  toDate.value = "";
  fetchSalesHistory();
};

const resetReturnFilter = () => {
  returnFromDate.value = "";
  returnToDate.value = "";
  returnStatusFilter.value = "";
  fetchSaleReturns();
};

const openReturnModal = async (saleId) => {
  try {
    const response = await api.get(`/history/sales/${saleId}`);
    selectedSale.value = response.data;

    returnForm.value.items = selectedSale.value.Items.map((item) => ({
      product_id: item.ProductID,
      qty: 0,
      reason: "",
    }));
    returnForm.value.notes = "";

    showModal.value = true;
  } catch (error) {
    console.error("Error fetching sale:", error);
    alert("Gagal memuat detail transaksi.");
  }
};

const closeModal = () => {
  showModal.value = false;
  selectedSale.value = null;
  returnForm.value = { items: [], notes: "" };
};

const submitReturn = async () => {
  if (!hasValidItems.value) {
    alert("Pilih minimal satu item dengan qty > 0 dan alasan.");
    return;
  }

  submitting.value = true;
  try {
    const validItems = returnForm.value.items.filter(
      (item) => item.qty > 0 && item.reason.trim()
    );

    const dto = {
      sale_id: selectedSale.value.ID,
      date: new Date().toISOString(),
      items: validItems,
      notes: returnForm.value.notes,
    };

    console.log("Return payload:", dto);

    const response = await api.post("/kasir/sale-returns", dto);
    if (response.status === 201) {
      alert("Retur penjualan berhasil dibuat (status: PENDING).");
      closeModal();
      fetchSaleReturns(); // Perbarui daftar retur setelah submit
    }
  } catch (error) {
    console.error("Error creating return:", error);
    alert(
      `Gagal membuat retur: ${error.response?.data?.error || error.message}`
    );
  } finally {
    submitting.value = false;
  }
};

const formatReturnNumber = (id) => {
  return `2025-${String(id).padStart(3, "0")}`;
};

const getStatusClass = (status) => {
  switch (status) {
    case "PENDING":
      return "bg-yellow-100 text-yellow-800";
    case "APPROVED":
      return "bg-green-100 text-green-800";
    case "REJECTED":
      return "bg-red-100 text-red-800";
    default:
      return "bg-gray-100 text-gray-800";
  }
};

const getStatusText = (status) => {
  switch (status) {
    case "PENDING":
      return "Pending";
    case "APPROVED":
      return "Approved";
    case "REJECTED":
      return "Rejected";
    default:
      return status;
  }
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

onMounted(() => {
  fetchSalesHistory();
  fetchSaleReturns();
});
</script>

<style scoped>
/* Style tambahan jika diperlukan */
</style>
