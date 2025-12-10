<template>
  <GudangLayouts>
    <div
      class="container mx-auto px-4 py-8 font-poppins bg-gray-50 min-h-screen"
    >
      <!-- Header Section -->
      <header class="mb-8">
        <h1 class="text-2xl md:text-3xl font-semibold text-gray-900">
          Laporan Retur Pembelian & Defect QC
        </h1>
        <p class="mt-2 text-base text-gray-500">
          Memantau semua pengajuan Retur Manual (Pending) dan Laporan Barang
          Rusak (QC Defect).
        </p>
      </header>

      <!-- Returns and Defects Section -->
      <section class="bg-white rounded-lg shadow-md border border-gray-200">
        <div class="p-6 border-b border-gray-100">
          <h2 class="text-xl font-semibold text-gray-800 flex items-center">
            Daftar Laporan Retur & Defect
            <span
              class="ml-3 text-sm font-normal text-gray-500 bg-gray-100 px-3 py-1 rounded-full"
            >
              Data menampilkan status PENDING dan QC_DEFECT secara default.
            </span>
          </h2>
        </div>

        <!-- Filter Section -->
        <div class="p-6 bg-gray-100 rounded-t-lg border-b border-gray-300">
          <div class="flex flex-wrap gap-4 items-end">
            <div class="flex-1 min-w-[160px]">
              <label class="block text-sm font-medium text-gray-700 mb-1.5"
                >Dari Tanggal</label
              >
              <input
                type="date"
                v-model="returnFromDate"
                class="w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm focus:ring-indigo-500 focus:border-indigo-500 transition duration-200"
              />
            </div>
            <div class="flex-1 min-w-[160px]">
              <label class="block text-sm font-medium text-gray-700 mb-1.5"
                >Sampai Tanggal</label
              >
              <input
                type="date"
                v-model="returnToDate"
                class="w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm focus:ring-indigo-500 focus:border-indigo-500 transition duration-200"
              />
            </div>
            <div class="flex-1 min-w-[160px]">
              <label class="block text-sm font-medium text-gray-700 mb-1.5"
                >Status</label
              >
              <select
                v-model="returnStatusFilter"
                class="w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm focus:ring-indigo-500 focus:border-indigo-500 transition duration-200"
              >
                <option value="">Pending & QC Defect (Default)</option>
                <option value="PENDING">Pending (Retur Manual)</option>
                <option value="QC_DEFECT">QC Defect (Laporan Rusak)</option>
                <option value="APPROVED">Approved (Siap Retur)</option>
                <option value="REJECTED">Rejected</option>
              </select>
            </div>
            <button
              @click="fetchPurchaseReturns"
              :disabled="loadingReturns"
              class="px-5 py-2.5 bg-indigo-600 text-white font-medium rounded-lg shadow-md hover:bg-indigo-700 disabled:opacity-50 transition duration-200"
            >
              {{ loadingReturns ? "Memuat..." : "Terapkan Filter" }}
            </button>
            <button
              @click="resetReturnFilter"
              class="px-5 py-2.5 bg-gray-600 text-white font-medium rounded-lg shadow-md hover:bg-gray-700 transition duration-200"
            >
              Reset
            </button>
            <button
              @click="openSelectPOModal"
              class="px-5 py-2.5 bg-red-600 text-white font-medium rounded-lg shadow-md hover:bg-red-700 transition duration-200"
              title="Membuat retur untuk barang yang sudah ada di gudang"
            >
              + Retur Manual Baru
            </button>
          </div>
          <p
            v-if="returnFromDate || returnToDate || returnStatusFilter"
            class="text-sm text-gray-600 mt-4"
          >
            Filter aktif: <strong>{{ returnFromDate || "N/A" }}</strong> s/d
            <strong>{{ returnToDate || "N/A" }}</strong> (Status:
            <strong>{{ returnStatusFilter || "Default" }}</strong
            >)
          </p>
        </div>

        <!-- Returns Table -->
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th
                  class="px-6 py-3 text-left text-xs font-semibold text-gray-500 uppercase tracking-wider"
                >
                  Return No.
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-semibold text-gray-500 uppercase tracking-wider"
                >
                  Nomor PO Asal
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-semibold text-gray-500 uppercase tracking-wider"
                >
                  Tanggal Retur
                </th>
                <th
                  class="px-6 py-3 text-right text-xs font-semibold text-gray-500 uppercase tracking-wider"
                >
                  Total Retur
                </th>
                <th
                  class="px-6 py-3 text-center text-xs font-semibold text-gray-500 uppercase tracking-wider"
                >
                  Status
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-if="loadingReturns">
                <td colspan="5" class="text-center py-8 text-gray-500">
                  <div class="flex flex-col items-center justify-center">
                    <svg
                      class="animate-spin h-5 w-5 text-indigo-500 mb-2"
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
                        d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                      ></path>
                    </svg>
                    Memuat Daftar Retur...
                  </div>
                </td>
              </tr>
              <tr v-else-if="purchaseReturns.length === 0">
                <td colspan="5" class="text-center py-8 text-gray-500 italic">
                  Tidak ada retur/defect yang ditemukan dalam rentang filter.
                </td>
              </tr>
              <tr
                v-else
                v-for="ret in purchaseReturns"
                :key="ret?.id || Math.random()"
                class="hover:bg-gray-50 transition duration-150"
              >
                <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <span
                    :class="
                      ret?.status === 'QC_DEFECT'
                        ? 'text-red-600 font-bold'
                        : 'text-indigo-600'
                    "
                  >
                    {{ ret?.status === "QC_DEFECT" ? "QC-DEFECT" : "RTN" }}-{{
                      formatReturnNumber(ret?.id)
                    }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
                  {{
                    ret?.purchase_id ? formatPONumber(ret.purchase_id) : "N/A"
                  }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-600">
                  {{ formatDate(ret?.date || ret?.created_at) || "N/A" }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-right">
                  <span class="font-semibold text-gray-800">
                    {{ formatCurrency(ret?.total || ret?.Total) || "Rp 0" }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-center">
                  <span
                    :class="getStatusClass(ret?.status)"
                    class="px-3 py-1 text-xs font-semibold rounded-full"
                  >
                    {{ getStatusText(ret?.status) || "Unknown" }}
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>

      <!-- Select PO Modal -->
      <div
        v-if="showSelectPOModal"
        class="fixed inset-0 bg-gray-900 bg-opacity-60 flex items-center justify-center z-50 p-4 transition-opacity duration-300"
      >
        <div
          class="bg-white rounded-lg shadow-xl p-8 max-w-2xl w-full max-h-[80vh] overflow-y-auto"
        >
          <h2 class="text-xl font-semibold text-red-700 mb-6 border-b pb-3">
            Pilih PO untuk Retur Manual
          </h2>
          <p class="mb-4 text-gray-600 text-sm">
            Pilih PO yang sudah di-<strong>RECEIVE</strong> (stok sudah masuk)
            untuk membuat retur manual.
          </p>

          <div class="flex flex-wrap gap-4 mb-6">
            <div class="flex-1 min-w-[160px]">
              <label class="block text-sm font-medium text-gray-700 mb-1.5"
                >Dari Tanggal</label
              >
              <input
                type="date"
                v-model="fromDate"
                class="w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
            <div class="flex-1 min-w-[160px]">
              <label class="block text-sm font-medium text-gray-700 mb-1.5"
                >Sampai Tanggal</label
              >
              <input
                type="date"
                v-model="toDate"
                class="w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500"
              />
            </div>
            <button
              @click="fetchPurchasesHistory('RECEIVED')"
              :disabled="loadingHistory"
              class="px-4 py-2.5 bg-blue-600 text-white font-medium rounded-lg shadow-md hover:bg-blue-700 disabled:opacity-50"
            >
              Filter
            </button>
            <button
              @click="resetFilter"
              class="px-4 py-2.5 bg-gray-600 text-white font-medium rounded-lg shadow-md hover:bg-gray-700"
            >
              Reset
            </button>
          </div>

          <div class="overflow-x-auto border rounded-lg">
            <table class="min-w-full divide-y divide-gray-200">
              <thead class="bg-gray-50 sticky top-0">
                <tr>
                  <th
                    class="px-6 py-3 text-left text-xs font-semibold text-gray-500 uppercase tracking-wider"
                  >
                    PO No.
                  </th>
                  <th
                    class="px-6 py-3 text-left text-xs font-semibold text-gray-500 uppercase tracking-wider"
                  >
                    Supplier
                  </th>
                  <th
                    class="px-6 py-3 text-center text-xs font-semibold text-gray-500 uppercase tracking-wider"
                  >
                    Aksi
                  </th>
                </tr>
              </thead>
              <tbody class="bg-white divide-y divide-gray-200">
                <tr v-if="loadingHistory">
                  <td
                    colspan="3"
                    class="text-center py-4 text-sm text-gray-500"
                  >
                    Memuat...
                  </td>
                </tr>
                <tr v-else-if="purchasesHistory.length === 0">
                  <td
                    colspan="3"
                    class="text-center py-4 text-sm text-gray-500 italic"
                  >
                    Tidak ada PO RECEIVED.
                  </td>
                </tr>
                <tr
                  v-else
                  v-for="purchase in purchasesHistory"
                  :key="purchase?.ID || purchase?.id"
                >
                  <td class="px-6 py-3 whitespace-nowrap text-sm font-medium">
                    {{ formatPONumber(purchase?.ID || purchase?.id) }}
                  </td>
                  <td class="px-6 py-3 whitespace-nowrap text-sm text-gray-600">
                    {{ purchase?.Supplier?.name || "N/A" }}
                  </td>
                  <td class="px-6 py-3 whitespace-nowrap text-center">
                    <button
                      @click="openReturnModal(purchase?.ID || purchase?.id)"
                      :disabled="
                        (purchase?.Status || purchase?.status) !== 'RECEIVED'
                      "
                      class="px-3 py-1.5 bg-red-500 text-white text-xs font-medium rounded-md hover:bg-red-600 disabled:bg-gray-300 transition duration-200"
                    >
                      Pilih
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <div class="flex justify-end gap-3 pt-4 border-t border-gray-200">
            <button
              @click="closeSelectPOModal"
              class="px-6 py-2.5 bg-gray-500 text-white font-medium rounded-lg shadow-md hover:bg-gray-600 transition duration-200"
            >
              Tutup
            </button>
          </div>
        </div>
      </div>

      <!-- Create Manual Return Modal -->
      <div
        v-if="showModal"
        class="fixed inset-0 bg-gray-900 bg-opacity-60 flex items-center justify-center z-50 p-4 transition-opacity duration-300"
      >
        <div
          class="bg-white rounded-lg shadow-xl p-8 max-w-4xl w-full max-h-[80vh] overflow-y-auto"
        >
          <h2 class="text-xl font-semibold text-red-700 mb-6 border-b pb-3">
            Buat Retur Pembelian Manual
          </h2>

          <div
            v-if="selectedPurchase"
            class="mb-6 p-4 bg-red-50 border-l-4 border-red-400 text-red-800 rounded-md"
          >
            <h3 class="font-semibold text-base mb-2">Detail Pembelian Asal</h3>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
              <p>
                <strong>Nomor PO:</strong>
                <span class="font-mono text-red-900">{{
                  formatPONumber(selectedPurchase.ID || selectedPurchase.id)
                }}</span>
              </p>
              <p>
                <strong>Tanggal PO:</strong>
                <span class="font-medium">{{
                  formatDate(
                    selectedPurchase.Date || selectedPurchase.created_at
                  )
                }}</span>
              </p>
              <p>
                <strong>Status PO:</strong>
                <span class="font-bold uppercase text-green-700">{{
                  selectedPurchase.Status || selectedPurchase.status
                }}</span>
              </p>
              <p>
                <strong>Total PO:</strong>
                <span class="font-medium text-green-700">{{
                  formatCurrency(
                    selectedPurchase.Total || selectedPurchase.total
                  )
                }}</span>
              </p>
            </div>
          </div>

          <div class="mb-6">
            <h3 class="text-lg font-semibold text-gray-700 mb-4">
              Pilih Item yang Akan Dikembalikan
            </h3>
            <div class="space-y-4">
              <div
                v-for="(retItem, index) in returnForm.items"
                :key="index"
                class="p-4 border border-gray-200 rounded-lg bg-white shadow-sm hover:shadow-md transition duration-200"
              >
                <div class="flex flex-wrap gap-4 items-center">
                  <div class="flex-1 min-w-[200px]">
                    <p class="font-semibold text-gray-800 text-base">
                      {{
                        selectedPurchase?.Items[index]?.Product?.name || "N/A"
                      }}
                    </p>
                    <p class="text-sm text-gray-500 mt-1">
                      Qty Dipesan:
                      <span class="font-semibold">{{
                        selectedPurchase?.Items[index]?.Qty || 0
                      }}</span>
                      | Harga Satuan:
                      <span class="font-semibold">{{
                        formatCurrency(selectedPurchase?.Items[index]?.Price)
                      }}</span>
                    </p>
                  </div>
                  <div class="w-28">
                    <label
                      class="block text-xs font-medium text-gray-700 mb-1.5"
                      >Qty Retur</label
                    >
                    <input
                      type="number"
                      v-model.number="retItem.qty"
                      min="0"
                      :max="selectedPurchase?.Items[index]?.Qty || 0"
                      class="w-full px-3 py-2.5 border border-gray-300 rounded-lg shadow-sm focus:ring-red-500 focus:border-red-500 text-center font-semibold"
                    />
                  </div>
                  <div class="flex-1 min-w-[250px]">
                    <label
                      class="block text-xs font-medium text-gray-700 mb-1.5"
                      >Alasan Retur (Wajib jika Qty > 0)</label
                    >
                    <textarea
                      v-model="retItem.reason"
                      placeholder="Contoh: Barang cacat/rusak, salah kirim, dll."
                      rows="2"
                      class="w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm focus:ring-red-500 focus:border-red-500 text-sm"
                    ></textarea>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="mb-6">
            <label class="block text-sm font-semibold text-gray-700 mb-2">
              Catatan Tambahan untuk Keseluruhan Retur (Opsional)
            </label>
            <textarea
              v-model="returnForm.notes"
              placeholder="Catatan opsional untuk proses retur ini."
              rows="3"
              class="w-full px-4 py-2.5 border border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500"
            ></textarea>
          </div>

          <div class="flex justify-end gap-3 pt-4 border-t border-gray-200">
            <button
              @click="closeModal"
              class="px-6 py-2.5 bg-gray-500 text-white font-medium rounded-lg shadow-md hover:bg-gray-600 transition duration-200"
            >
              Batal
            </button>
            <button
              @click="submitReturn"
              :disabled="submitting || !hasValidItems"
              class="px-6 py-2.5 bg-red-600 text-white font-medium rounded-lg shadow-md hover:bg-red-700 disabled:bg-red-300 disabled:cursor-not-allowed transition duration-200"
            >
              <span v-if="submitting">
                <svg
                  class="animate-spin h-5 w-5 mr-3 inline-block"
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
                    d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                  ></path>
                </svg>
                Menyimpan...
              </span>
              <span v-else>
                Submit Retur Manual
                <span
                  class="ml-2 px-2 py-0.5 text-xs bg-white text-red-600 rounded-full"
                >
                  Total Item: {{ validReturnItemsCount }}
                </span>
              </span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </GudangLayouts>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import axios from "axios";
import GudangLayouts from "@/components/GudangLayouts.vue"; // Pastikan path ini benar!

const purchasesHistory = ref([]); // Digunakan di modal Retur Manual
const purchaseReturns = ref([]); // Daftar Retur/Defect
const loadingHistory = ref(true);
const loadingReturns = ref(false);
const fromDate = ref("");
const toDate = ref("");
const returnFromDate = ref("");
const returnToDate = ref("");
const returnStatusFilter = ref(""); // Default: ""

// Modal states
const showSelectPOModal = ref(false); // Modal untuk memilih PO
const showModal = ref(false); // Modal untuk form retur
const selectedPurchase = ref(null);
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

const validReturnItemsCount = computed(() => {
  return returnForm.value.items.filter(
    (item) => item.qty > 0 && item.reason.trim()
  ).length;
});

const api = axios.create({ baseURL: "http://localhost:8081/api/v1" });
const savedToken = localStorage.getItem("token");
if (savedToken) {
  api.defaults.headers.common["Authorization"] = `Bearer ${savedToken}`;
}

// ✅ FUNGSI INI KINI DIPAKAI HANYA UNTUK MEMILIH PO DI MODAL RETUR MANUAL
const fetchPurchasesHistory = async (status = "RECEIVED") => {
  loadingHistory.value = true;
  try {
    const params = {
      page: 1,
      size: 10,
      status: status, // Kunci: Hanya PO yang sudah RECEIVED yang bisa diretur manual
    };

    if (fromDate.value) {
      params.start_date = fromDate.value;
    }
    if (toDate.value) {
      params.end_date = toDate.value;
    }

    // Menggunakan endpoint /gudang/purchases/history yang diasumsikan hanya menampilkan
    // PO yang *bisa* dilihat oleh gudang.
    const response = await api.get("/history/purchases", { params });

    const data = response.data;
    // Asumsi: data PO dari backend dikemas dalam { data: { data: [] } } atau { data: [] }
    let rawData = [];
    if (data && data.data && Array.isArray(data.data.data)) {
      rawData = data.data.data;
    } else if (data && Array.isArray(data.data)) {
      rawData = data.data;
    } else if (data && Array.isArray(data)) {
      rawData = data;
    }

    purchasesHistory.value = rawData.filter((p) => p && (p.ID || p.id));
  } catch (error) {
    console.error("Error fetch purchases for manual return:", error);
    if (error.response) {
      alert(
        `Error memuat PO: ${error.response.status} - ${
          error.response.data.error || error.message
        }`
      );
    } else {
      alert("Kesalahan jaringan atau server tidak tersedia.");
    }
    purchasesHistory.value = [];
  } finally {
    loadingHistory.value = false;
  }
};

// ✅ FUNGSI UTAMA UNTUK MENAMPILKAN DAFTAR RETUR/DEFECT
const fetchPurchaseReturns = async () => {
  loadingReturns.value = true;
  try {
    const params = {
      page: 1,
      size: 10,
    };

    if (returnFromDate.value) {
      params.start_date = returnFromDate.value;
    }
    if (returnToDate.value) {
      params.end_date = returnToDate.value;
    }
    if (returnStatusFilter.value) {
      // Jika status diset, gunakan itu
      params.status = returnStatusFilter.value;
    }
    // Asumsi: Endpoint /gudang/purchase-returns/history akan default menampilkan PENDING & QC_DEFECT jika params.status kosong.

    // Mengganti /history/purchase-returns dengan /gudang/purchase-returns/history untuk konsistensi di service Gudang
    const response = await api.get("/history/purchase-returns", {
      params,
    });
    let rawData = response.data.data || response.data || [];
    purchaseReturns.value = Array.isArray(rawData)
      ? rawData.filter((ret) => ret && ret.id)
      : [];
  } catch (error) {
    console.error(
      "Error fetching returns:",
      error.response ? error.response.data : error.message
    );
    alert(
      `Gagal mengambil daftar retur: ${
        error.response?.data?.error || "Cek otorisasi (role not allowed)."
      }`
    );
    purchaseReturns.value = [];
  } finally {
    loadingReturns.value = false;
  }
};

const openSelectPOModal = () => {
  fetchPurchasesHistory("RECEIVED"); // Ambil PO yang sudah RECEIVED
  showSelectPOModal.value = true;
};

const closeSelectPOModal = () => {
  showSelectPOModal.value = false;
  purchasesHistory.value = [];
};

const resetFilter = () => {
  fromDate.value = "";
  toDate.value = "";
  fetchPurchasesHistory("RECEIVED");
};

const resetReturnFilter = () => {
  returnFromDate.value = "";
  returnToDate.value = "";
  returnStatusFilter.value = "";
  fetchPurchaseReturns();
};

const openReturnModal = async (purchaseId) => {
  closeSelectPOModal(); // Tutup modal pemilihan PO
  if (!purchaseId) {
    alert("ID pembelian tidak valid.");
    return;
  }
  try {
    // Ambil detail PO dari endpoint detail. Asumsi: /gudang/purchases/:id
    const response = await api.get(`/gudang/purchases/${purchaseId}`);
    selectedPurchase.value = response.data.data || response.data;

    if (!selectedPurchase.value) {
      alert("Data pembelian tidak ditemukan.");
      return;
    }

    const items = selectedPurchase.value.Items || [];
    if (!Array.isArray(items)) {
      alert("Items pembelian tidak valid.");
      return;
    }

    // Inisialisasi return form items
    returnForm.value.items = items.map((item) => ({
      // Menyimpan ID item PO untuk referensi backend
      purchase_item_id: item.ID || item.id,
      product_id: item.ProductID || item.product_id,
      qty: 0,
      reason: "", // 'reason' adalah v-model di template
      price: item.Price || item.price,
    }));
    returnForm.value.notes = "";

    showModal.value = true;
  } catch (error) {
    console.error("Error fetching purchase:", error);
    alert(
      `Gagal memuat detail pembelian: ${
        error.response?.data?.error || error.message
      }`
    );
  }
};

const closeModal = () => {
  showModal.value = false;
  selectedPurchase.value = null;
  returnForm.value = { items: [], notes: "" };
};

// 💡 FUNGSI INI ADALAH FOKUS PERBAIKAN
const submitReturn = async () => {
  if (!hasValidItems.value) {
    alert("Pilih minimal satu item dengan qty > 0 dan alasan.");
    return;
  }

  submitting.value = true;
  try {
    const validItems = returnForm.value.items
      .filter((item) => item.qty > 0 && item.reason.trim())
      .map((item) => ({
        // KIRIM DATA YANG LENGKAP KE BACKEND UNTUK KEMUDAHAN PENANGANAN:
        product_id: item.product_id,
        purchase_item_id: item.purchase_item_id, // Penting untuk mencocokkan item PO
        qty: item.qty,
        price: item.price,
        // ASUMSI PERBAIKAN: Menggunakan field 'reason'
        reason: item.reason,
      }));

    const dto = {
      purchase_id: selectedPurchase.value.ID || selectedPurchase.value.id,
      date: new Date().toISOString().split("T")[0],
      items: validItems,
      notes: returnForm.value.notes,
    };

    // Endpoint Retur Manual: /gudang/purchase-returns (harus status PENDING)
    // Asumsi backend DTO menerima field 'items' dengan isian di atas.
    const response = await api.post("/gudang/purchase-returns", dto);

    if (response.status === 201 || response.status === 200) {
      alert("Retur pembelian manual berhasil dibuat (status: PENDING).");
      closeModal();
      fetchPurchaseReturns();
    }
  } catch (error) {
    console.error("Error creating return:", error.response?.data || error);
    // Tampilkan pesan error spesifik dari backend jika ada
    alert(
      `Gagal membuat retur: ${
        error.response?.data?.error ||
        error.response?.data?.message ||
        "Kesalahan tidak diketahui. Cek Console."
      }`
    );
  } finally {
    submitting.value = false;
  }
};

onMounted(() => {
  // Hanya ambil data Retur/Defect saat mounted
  fetchPurchaseReturns();
});

// --- Utility Functions ---

const formatPONumber = (id) => {
  return `PO-${String(id || "").padStart(3, "0")}`;
};

const formatReturnNumber = (id) => {
  return String(id || "").padStart(4, "0"); // Diperpendek agar lebih mudah dibaca
};

const getStatusClass = (status) => {
  switch (status) {
    case "PENDING":
      return "bg-yellow-100 text-yellow-800 border border-yellow-300";
    case "APPROVED":
      return "bg-green-100 text-green-800 border border-green-300";
    case "REJECTED":
      return "bg-red-100 text-red-800 border border-red-300";
    case "QC_DEFECT": // Warna khusus untuk laporan defect
      return "bg-red-200 text-red-900 border border-red-500 font-bold";
    default:
      return "bg-gray-100 text-gray-800 border border-gray-300";
  }
};

const getStatusText = (status) => {
  switch (status) {
    case "PENDING":
      return "Pending (Retur Manual)";
    case "APPROVED":
      return "Approved";
    case "REJECTED":
      return "Rejected";
    case "QC_DEFECT":
      return "Laporan QC Defect";
    default:
      return status || "Unknown";
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
  if (amount === null || amount === undefined) return "Rp 0";
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  }).format(amount);
};
</script>

<style scoped>
/* Import Poppins font */
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700&display=swap");

.font-poppins {
  font-family: "Poppins", sans-serif;
}

/* Custom scrollbar styles */
.max-h-\[80vh\]::-webkit-scrollbar {
  width: 8px;
}
.max-h-\[80vh\]::-webkit-scrollbar-thumb {
  background-color: #cbd5e1;
  border-radius: 4px;
}
.max-h-\[80vh\]::-webkit-scrollbar-track {
  background: #f1f1f1;
}
</style>
