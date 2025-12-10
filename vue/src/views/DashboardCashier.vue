<template>
  <CashierLayouts>
    <div>
      <div class="flex flex-wrap justify-between gap-3 p-4">
        <p
          class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
        >
          Selamat datang, {{ userName }}
        </p>
      </div>

      <!-- Filter Tanggal -->
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
            @click="fetchDashboardData"
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
          History Penjualan ({{
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
                  Print PDF
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Detail
                </th>
              </tr>
            </thead>
            <tbody>
              <tr v-if="loadingHistory">
                <td colspan="5" class="text-center py-4">Loading history...</td>
              </tr>
              <tr v-else-if="salesHistory.length === 0">
                <td colspan="5" class="text-center py-4 text-gray-500">
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
                  class="px-6 py-4 whitespace-nowrap text-sm text-blue-500 hover:text-blue-700 cursor-pointer"
                >
                  <button @click="printPDF(sale.ID)" class="underline">
                    Print PDF
                  </button>
                </td>
                <td
                  class="px-6 py-4 whitespace-nowrap text-sm text-blue-500 hover:text-blue-700 cursor-pointer"
                >
                  <button @click="viewDetails(sale.ID)" class="underline">
                    Lihat Detail
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Modal untuk Lihat Detail -->
      <div
        v-if="showDetailsModal"
        class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
      >
        <div
          class="bg-white rounded-lg p-6 max-w-2xl w-full max-h-[90vh] overflow-y-auto"
        >
          <div class="flex justify-between items-center mb-4">
            <h2 class="text-xl font-bold">Detail Transaksi</h2>
            <button
              @click="closeDetailsModal"
              class="text-gray-500 hover:text-gray-700 text-xl"
            >
              &times;
            </button>
          </div>
          <div v-if="selectedSale" class="space-y-4">
            <p>
              <strong>Nomor Transaksi:</strong>
              {{
                selectedSale.transaction_number ||
                `TRX-${selectedSale.ID || "N/A"}`
              }}
            </p>
            <p>
              <strong>Tanggal:</strong>
              {{ formatDate(selectedSale.Date) || "N/A" }}
            </p>
            <p>
              <strong>Total:</strong>
              {{ formatCurrency(selectedSale.Total) || "Rp 0" }}
            </p>
            <p>
              <strong>Uang Bayar:</strong>
              {{ formatCurrency(selectedSale.Paid || 0) || "Rp 0" }}
            </p>
            <p>
              <strong>Kembalian:</strong>
              {{
                formatCurrency(
                  selectedSale.Change ||
                    (selectedSale.Paid || 0) - (selectedSale.Total || 0)
                ) || "Rp 0"
              }}
            </p>
            <p>
              <strong>Customer:</strong>
              {{ selectedSale.Customer?.name || "N/A" }}
            </p>

            <h3 class="text-lg font-semibold mt-4">Daftar Barang</h3>
            <div class="overflow-x-auto">
              <table class="min-w-full bg-white border border-gray-300">
                <thead>
                  <tr class="bg-gray-100">
                    <th class="px-4 py-2 text-left text-gray-600">
                      Nama Barang
                    </th>
                    <th class="px-4 py-2 text-left text-gray-600">Qty</th>
                    <th class="px-4 py-2 text-left text-gray-600">Harga</th>
                    <th class="px-4 py-2 text-left text-gray-600">Subtotal</th>
                  </tr>
                </thead>
                <tbody>
                  <tr
                    v-for="(item, index) in selectedSale.Items || []"
                    :key="index"
                    class="border-t border-gray-300"
                  >
                    <td class="px-4 py-2 text-gray-700">
                      {{ item.Product?.name || "N/A" }}
                    </td>
                    <td class="px-4 py-2 text-gray-700">{{ item.Qty || 0 }}</td>
                    <td class="px-4 py-2 text-gray-700">
                      {{ formatCurrency(item.Price || 0) }}
                    </td>
                    <td class="px-4 py-2 text-gray-700">
                      {{ formatCurrency((item.Qty || 0) * (item.Price || 0)) }}
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  </CashierLayouts>
</template>

<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";
import CashierLayouts from "@/components/CashierLayouts.vue";

const userName = ref("Kasir");
const salesHistory = ref([]);
const loadingHistory = ref(true);
const fromDate = ref("");
const toDate = ref("");
const showDetailsModal = ref(false); // State untuk modal detail
const selectedSale = ref(null); // Data transaksi yang dipilih

const api = axios.create({ baseURL: "http://localhost:8081/api/v1" });
const savedToken = localStorage.getItem("token");
if (savedToken) {
  api.defaults.headers.common["Authorization"] = `Bearer ${savedToken}`;
}

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
      alert("Tidak ada data dari server. Periksa respons API.");
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
      alert("Kesalahan jaringan atau server tidak tersedia. Cek console.");
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

const printPDF = async (saleId) => {
  try {
    const response = await api.get(`/history/sales/${saleId}/pdf`, {
      responseType: "blob",
    });

    const blob = new Blob([response.data], { type: "application/pdf" });
    const url = window.URL.createObjectURL(blob);
    window.open(url);

    const link = document.createElement("a");
    link.href = url;
    link.setAttribute("download", `struk_${saleId}.pdf`);
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

    window.URL.revokeObjectURL(url);
  } catch (error) {
    console.error("Error printing PDF:", error);
    alert("Gagal mencetak PDF. Cek console untuk detail.");
  }
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
    console.log("Selected Sale Data:", selectedSale.value); // Debugging
    showDetailsModal.value = true;
  } catch (error) {
    console.error("Error fetching sale details:", error);
    alert("Gagal memuat detail transaksi. Cek console untuk detail.");
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
/* Tambahkan style jika diperlukan */
</style>
