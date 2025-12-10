<template>
  <CashierLayouts>
    <div>
      <div class="flex flex-wrap justify-between gap-3 p-4">
        <p
          class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
        >
          Retur Penjualan
        </p>
      </div>

      <!-- Filter Tanggal untuk History Penjualan -->
      <div class="p-4 bg-[#f0f2f5] rounded-lg">
        <div class="flex flex-wrap gap-4 items-end">
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1"
              >Dari Tanggal</label
            >
            <input
              type="date"
              v-model="fromDate"
              class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <div>
            <label class="block text-sm font-medium text-gray-700 mb-1"
              >Sampai Tanggal</label
            >
            <input
              type="date"
              v-model="toDate"
              class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            />
          </div>
          <button
            @click="fetchSalesHistory"
            :disabled="loadingHistory"
            class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 disabled:opacity-50"
          >
            {{ loadingHistory ? "Loading..." : "Filter" }}
          </button>
          <button
            @click="resetFilter"
            class="px-4 py-2 bg-gray-500 text-white rounded-md hover:bg-gray-600"
          >
            Reset (10 Terakhir)
          </button>
        </div>
        <p v-if="fromDate || toDate" class="text-sm text-gray-600 mt-2">
          Filter aktif: {{ fromDate || "N/A" }} s/d {{ toDate || "N/A" }}
        </p>
      </div>

      <hr class="my-4" />

      <div class="p-4">
        <p
          class="text-[#111418] tracking-light text-2xl font-bold leading-tight"
        >
          History Penjualan (untuk Retur) ({{
            fromDate || toDate ? "Filtered" : "10 Terakhir"
          }})
        </p>
        <div
          class="mt-4 overflow-x-auto rounded-lg border border-[#dbe0e6] bg-white"
        >
          <table class="min-w-full">
            <thead>
              <tr class="bg-gray-100">
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Tanggal
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Nomor Transaksi
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Total
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Aksi
                </th>
              </tr>
            </thead>
            <tbody>
              <tr v-if="loadingHistory">
                <td colspan="4" class="text-center py-4">Loading history...</td>
              </tr>
              <tr v-else-if="salesHistory.length === 0">
                <td colspan="4" class="text-center py-4 text-gray-500">
                  Tidak ada data penjualan.
                </td>
              </tr>
              <tr
                v-else
                v-for="sale in salesHistory"
                :key="sale.ID"
                class="border-t border-t-[#dbe0e6]"
              >
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ formatDate(sale.Date) || "N/A" }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ sale.transaction_number || `TRX-${sale.ID || "N/A"}` }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ formatCurrency(sale.Total) || "Rp 0" }}
                </td>
                <td
                  class="px-6 py-4 whitespace-nowrap text-sm text-green-500 hover:text-green-700 cursor-pointer"
                >
                  <button @click="openReturnModal(sale.ID)" class="underline">
                    ADD
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Daftar Retur yang Dibuat Kasir dengan Filter -->
      <div class="p-4 mt-8">
        <p
          class="text-[#111418] tracking-light text-2xl font-bold leading-tight"
        >
          Daftar Retur Saya
        </p>
        <div class="p-4 bg-[#f0f2f5] rounded-lg mt-4">
          <div class="flex flex-wrap gap-4 items-end">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1"
                >Dari Tanggal</label
              >
              <input
                type="date"
                v-model="returnFromDate"
                class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1"
                >Sampai Tanggal</label
              >
              <input
                type="date"
                v-model="returnToDate"
                class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              />
            </div>
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-1"
                >Status</label
              >
              <select
                v-model="returnStatusFilter"
                class="px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
              >
                <option value="">Semua</option>
                <option value="PENDING">Pending</option>
                <option value="APPROVED">Approved</option>
                <option value="REJECTED">Rejected</option>
              </select>
            </div>
            <button
              @click="fetchSaleReturns"
              :disabled="loadingReturns"
              class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 disabled:opacity-50"
            >
              {{ loadingReturns ? "Loading..." : "Filter" }}
            </button>
            <button
              @click="resetReturnFilter"
              class="px-4 py-2 bg-gray-500 text-white rounded-md hover:bg-gray-600"
            >
              Reset (10 Terakhir)
            </button>
          </div>
          <p
            v-if="returnFromDate || returnToDate || returnStatusFilter"
            class="text-sm text-gray-600 mt-2"
          >
            Filter aktif: {{ returnFromDate || "N/A" }} s/d
            {{ returnToDate || "N/A" }} (Status:
            {{ returnStatusFilter || "Semua" }})
          </p>
        </div>
        <div
          class="mt-4 overflow-x-auto rounded-lg border border-[#dbe0e6] bg-white"
        >
          <table class="min-w-full">
            <thead>
              <tr class="bg-gray-100">
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Return No.
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Nomor Transaksi
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Tanggal
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Total
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Status
                </th>
              </tr>
            </thead>
            <tbody>
              <tr v-if="loadingReturns">
                <td colspan="5" class="text-center py-4">Loading returns...</td>
              </tr>
              <tr v-else-if="saleReturns.length === 0">
                <td colspan="5" class="text-center py-4 text-gray-500">
                  Tidak ada retur yang dibuat.
                </td>
              </tr>
              <tr
                v-else
                v-for="ret in saleReturns"
                :key="ret.id"
                class="border-t border-t-[#dbe0e6]"
              >
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  RTN-{{ formatReturnNumber(ret.id) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{
                    ret.sale?.transaction_number ||
                    `TRX-${ret.sale_id || "N/A"}`
                  }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ formatDate(ret.date) || "N/A" }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ formatCurrency(ret.total) || "Rp 0" }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <span
                    :class="getStatusClass(ret.status)"
                    class="px-2 py-1 rounded text-xs"
                  >
                    {{ getStatusText(ret.status) }}
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Modal untuk Buat Retur -->
      <div
        v-if="showModal"
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
      >
        <div
          class="bg-white rounded-lg p-6 max-w-4xl w-full max-h-[90vh] overflow-y-auto"
        >
          <h2 class="text-xl font-bold mb-4">Buat Retur Penjualan</h2>

          <div v-if="selectedSale" class="mb-4 p-4 bg-gray-100 rounded">
            <p>
              <strong>Nomor Transaksi:</strong>
              {{ selectedSale.transaction_number || `TRX-${selectedSale.ID}` }}
            </p>
            <p><strong>Tanggal:</strong> {{ formatDate(selectedSale.Date) }}</p>
            <p>
              <strong>Total:</strong> {{ formatCurrency(selectedSale.Total) }}
            </p>
          </div>

          <div class="mb-4">
            <h3 class="text-lg font-semibold mb-2">Items untuk Diretur</h3>
            <div
              v-for="(retItem, index) in returnForm.items"
              :key="index"
              class="border-b py-3 last:border-b-0"
            >
              <div class="flex flex-wrap gap-4 items-start">
                <div class="flex-1">
                  <p class="font-medium">
                    {{ selectedSale?.Items[index]?.Product?.name || "N/A" }}
                  </p>
                  <p class="text-sm text-gray-600">
                    Qty Original: {{ selectedSale?.Items[index]?.Qty || 0 }} |
                    Harga:
                    {{ formatCurrency(selectedSale?.Items[index]?.Price) }}
                  </p>
                </div>
                <div class="w-32">
                  <label class="block text-sm font-medium text-gray-700 mb-1"
                    >Qty Retur</label
                  >
                  <input
                    type="number"
                    v-model.number="retItem.qty"
                    min="0"
                    :max="selectedSale?.Items[index]?.Qty || 0"
                    class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  />
                </div>
                <div class="flex-1">
                  <label class="block text-sm font-medium text-gray-700 mb-1"
                    >Alasan</label
                  >
                  <textarea
                    v-model="retItem.reason"
                    placeholder="Masukkan alasan retur item ini"
                    rows="2"
                    class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  ></textarea>
                </div>
              </div>
            </div>
          </div>

          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-1"
              >Catatan Tambahan</label
            >
            <textarea
              v-model="returnForm.notes"
              placeholder="Catatan opsional untuk retur ini"
              rows="3"
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
            ></textarea>
          </div>

          <div class="flex justify-end gap-2">
            <button
              @click="closeModal"
              class="px-4 py-2 bg-gray-500 text-white rounded-md hover:bg-gray-600"
            >
              Batal
            </button>
            <button
              @click="submitReturn"
              :disabled="submitting || !hasValidItems"
              class="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 disabled:opacity-50"
            >
              {{ submitting ? "Menyimpan..." : "Submit Retur" }}
            </button>
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
