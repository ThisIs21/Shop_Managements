<template>
  <div class="flex">
    <GudangSidebar />

    <main class="flex-1 ml-64 p-6">
      <h1 class="text-2xl font-bold mb-6">Penerimaan Barang (Goods Receipt)</h1>

      <div class="flex justify-between items-center mb-4">
        <div class="flex gap-2">
          <label class="text-sm font-medium pt-2">Filter Status:</label>
          <select
            v-model="currentStatus"
            @change="loadPurchases"
            class="border rounded p-2 text-sm"
          >
            <option value="APPROVED">Siap Diterima (APPROVED)</option>
            <option value="RECEIVED">Sudah Diterima (RECEIVED)</option>
            <option value="">Semua</option>
          </select>
        </div>
        <input
          type="text"
          v-model="searchQuery"
          @keyup.enter="loadPurchases"
          placeholder="Cari PO ID atau Supplier..."
          class="border rounded p-2 text-sm w-64"
        />
      </div>

      <div class="bg-white shadow-md rounded-lg overflow-hidden">
        <table class="min-w-full leading-normal">
          <thead>
            <tr
              class="bg-gray-100 text-gray-600 uppercase text-sm leading-normal"
            >
              <th class="py-3 px-6 text-left">ID</th>
              <th class="py-3 px-6 text-left">Tanggal</th>
              <th class="py-3 px-6 text-left">Supplier</th>
              <th class="py-3 px-6 text-right">Total</th>
              <th class="py-3 px-6 text-center">Status</th>
              <th class="py-3 px-6 text-center">Aksi</th>
            </tr>
          </thead>
          <tbody class="text-gray-600 text-sm font-light">
            <tr
              v-for="po in purchases"
              :key="po.ID"
              class="border-b border-gray-200 hover:bg-gray-50"
            >
              <td class="py-3 px-6 text-left whitespace-nowrap">{{ po.ID }}</td>
              <td class="py-3 px-6 text-left">
                {{ formatDate(po.Date) }}
              </td>
              <td class="py-3 px-6 text-left">
                {{ po.Supplier?.name || "N/A" }}
              </td>
              <td class="py-3 px-6 text-right">
                {{ formatCurrency(po.Total) }}
              </td>
              <td class="py-3 px-6 text-center">
                <span :class="getStatusClass(po.Status)">
                  {{ po.Status }}
                </span>
              </td>
              <td
                class="py-3 px-6 text-center flex justify-center items-center"
              >
                <button
                  v-if="po.Status === 'APPROVED'"
                  @click="openQcModal(po.ID)"
                  class="bg-green-500 hover:bg-green-600 text-white text-xs font-bold py-1 px-3 rounded disabled:opacity-50"
                >
                  QC & Terima
                </button>
                <span
                  v-else-if="po.Status === 'RECEIVED'"
                  class="text-green-500 font-bold text-xs"
                >
                  ✅ Diterima
                </span>
                <button
                  @click="showDetail(po.ID)"
                  class="ml-2 text-blue-500 hover:text-blue-700 text-xs"
                >
                  Detail
                </button>
              </td>
            </tr>
            <tr v-if="!purchases.length && !loading">
              <td colspan="6" class="py-4 text-center text-gray-500">
                Tidak ada Purchase Order yang ditemukan.
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="flex justify-between items-center mt-4">
        <p class="text-sm text-gray-600">Total: {{ totalItems }}</p>
        <div v-if="loading" class="text-blue-500">Memuat...</div>
      </div>
    </main>

    <div
      v-if="isModalOpen"
      class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50 flex justify-center items-center"
      @click="closeModal"
    >
      <div
        class="bg-white p-6 rounded-lg shadow-xl max-w-4xl w-full mx-4"
        @click.stop
      >
        <div class="flex justify-between items-center border-b pb-3 mb-4">
          <h3 class="text-xl font-bold">
            Detail Purchase Order #{{ selectedPurchase?.ID }}
          </h3>
          <button
            @click="closeModal"
            class="text-gray-500 hover:text-gray-800 text-2xl leading-none"
          >
            &times;
          </button>
        </div>

        <div v-if="selectedPurchase">
          <div class="grid grid-cols-2 gap-4 text-sm mb-4"></div>

          <h4 class="text-lg font-semibold mb-2">Item Pembelian</h4>
          <div class="overflow-x-auto border rounded">
            <table class="min-w-full text-sm">
              <thead>
                <tr class="bg-gray-50">
                  <th class="py-2 px-4 text-left">Produk</th>
                  <th class="py-2 px-4 text-right">Qty PO</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="item in selectedPurchase.Items"
                  :key="item.ID"
                  class="border-t"
                >
                  <td class="py-2 px-4">
                    {{ item.Product?.name || "Produk Hilang" }}
                  </td>
                  <td class="py-2 px-4 text-right">{{ item.Qty }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div class="mt-4 flex justify-end">
          <button
            @click="closeModal"
            class="bg-gray-300 hover:bg-gray-400 text-gray-800 font-bold py-2 px-4 rounded"
          >
            Tutup
          </button>
        </div>
      </div>
    </div>

    <div
      v-if="isQcModalOpen"
      class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50 flex justify-center items-center"
      @click="closeQcModal"
    >
      <div
        class="bg-white p-6 rounded-lg shadow-xl max-w-5xl w-full mx-4"
        @click.stop
      >
        <div class="flex justify-between items-center border-b pb-3 mb-4">
          <h3 class="text-xl font-bold text-blue-700">
            Laporan Quality Control (PO #{{ qcForm.purchase_id }})
          </h3>
          <button
            @click="closeQcModal"
            class="text-gray-500 hover:text-gray-800 text-2xl leading-none"
            :disabled="submitting"
          >
            &times;
          </button>
        </div>

        <form @submit.prevent="submitQcReport">
          <div class="mb-4">
            <p class="text-sm font-medium">
              Supplier:
              <span class="font-bold">{{ qcForm.supplierName }}</span>
            </p>
          </div>

          <h4 class="text-base font-semibold mb-3">Input Hasil QC per Item</h4>
          <div class="overflow-x-auto border rounded max-h-80">
            <table class="min-w-full text-sm">
              <thead>
                <tr class="bg-gray-50">
                  <th class="py-2 px-4 text-left w-2/5">Produk</th>
                  <th class="py-2 px-4 text-center w-1/12">Qty PO</th>
                  <th class="py-2 px-4 text-center w-1/12">Qty Baik</th>
                  <th class="py-2 px-4 text-center w-1/12">Qty Rusak</th>
                  <th class="py-2 px-4 text-center w-1/12">Total Terima</th>
                  <th class="py-2 px-4 text-left w-1/5">Catatan</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="(item, index) in qcForm.items"
                  :key="index"
                  class="border-t"
                >
                  <td class="py-2 px-4 font-medium">
                    {{ item.product_name }}
                  </td>
                  <td class="py-2 px-4 text-center font-semibold">
                    {{ item.qty_ordered }}
                  </td>
                  <td class="py-2 px-4 text-center">
                    <input
                      v-model.number="item.qty_good"
                      type="number"
                      min="0"
                      required
                      class="w-full border-gray-300 rounded-lg text-sm p-1.5 text-center"
                    />
                  </td>
                  <td class="py-2 px-4 text-center">
                    <input
                      v-model.number="item.qty_damaged"
                      type="number"
                      min="0"
                      class="w-full border-gray-300 rounded-lg text-sm p-1.5 text-center"
                    />
                  </td>
                  <td
                    class="py-2 px-4 text-center font-bold"
                    :class="{
                      'text-red-500': item.total_received > item.qty_ordered,
                    }"
                  >
                    {{ item.total_received }}
                    <p
                      v-if="item.total_received > item.qty_ordered"
                      class="text-red-500 text-xs mt-1"
                    >
                      Melebihi PO!
                    </p>
                  </td>
                  <td class="py-2 px-4">
                    <input
                      v-model="item.note"
                      type="text"
                      class="w-full border-gray-300 rounded-lg text-sm p-1.5"
                      placeholder="e.g. 5 Pcs Rusak karena goresan"
                    />
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <p
            v-if="qcError"
            class="text-red-500 text-sm mt-4 p-2 bg-red-50 border border-red-200 rounded"
          >
            **Error:** {{ qcError }}
          </p>

          <div class="mt-6 flex justify-end">
            <button
              type="button"
              @click="closeQcModal"
              class="px-6 py-2 border border-gray-300 text-gray-700 rounded-lg shadow-sm hover:bg-gray-100 mr-3"
              :disabled="submitting"
            >
              Batal
            </button>
            <button
              type="submit"
              :disabled="submitting || !isQcFormValid"
              class="px-6 py-2 bg-blue-600 text-white font-medium rounded-lg shadow-md hover:bg-blue-700 disabled:opacity-50 transition-colors"
            >
              {{ submitting ? "Memproses..." : "Selesaikan QC & Terima" }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from "vue";
import axios from "axios";
import GudangSidebar from "../components/GudangSidebar.vue"; // Sesuaikan path

const API_BASE = "http://localhost:8081/api/v1";
const token = localStorage.getItem("token");

const purchases = ref([]);
const loading = ref(false);
const totalItems = ref(0);
const searchQuery = ref("");
// FIX 1: Ubah default status menjadi "" (Semua) atau biarkan APPROVED.
// Pilihan: Kita biarkan APPROVED, tapi kita atasi perubahannya setelah submit.
const currentStatus = ref("APPROVED");

// State untuk Modal Detail (Original)
const selectedPurchase = ref(null);
const isModalOpen = ref(false);

// State untuk Modal QC (NEW)
const isQcModalOpen = ref(false);
const submitting = ref(false);
const qcError = ref("");
const qcForm = ref({
  purchase_id: null,
  supplierName: "",
  purchaseDate: "",
  items: [],
});

const isQcFormValid = computed(() => {
  return (
    qcForm.value.items.length > 0 &&
    qcForm.value.items.every(
      (item) =>
        item.qty_good !== null &&
        item.qty_good >= 0 &&
        item.qty_damaged !== null &&
        item.qty_damaged >= 0 && // Total diterima tidak boleh melebihi Qty PO
        item.total_received <= item.qty_ordered
    )
  );
});

// Tambahkan watch untuk memanggil loadPurchases saat searchQuery berubah
// Ini membantu memuat ulang data saat pengguna mengetik dan menekan Enter.
watch(searchQuery, (newVal) => {
  if (newVal === "") {
    loadPurchases();
  }
});

const page = ref(1);
const size = ref(10);

// --- Helper Functions ---
const formatDate = (dateString) => {
  if (!dateString) return "-";
  const date = new Date(dateString);
  const day = String(date.getDate()).padStart(2, "0");
  const month = String(date.getMonth() + 1).padStart(2, "0");
  const year = date.getFullYear();
  return `${day}/${month}/${year}`;
};

const formatCurrency = (amount) => {
  const numericAmount = Number(amount) || 0;

  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0,
  }).format(numericAmount);
};

const getStatusClass = (status) => {
  switch (status) {
    case "APPROVED":
      return "bg-yellow-200 text-yellow-800 py-1 px-3 rounded-full text-xs font-medium";
    case "RECEIVED":
      return "bg-green-200 text-green-800 py-1 px-3 rounded-full text-xs font-medium";
    case "REJECTED":
      return "bg-red-200 text-red-800 py-1 px-3 rounded-full text-xs font-medium";
    default:
      return "bg-gray-200 text-gray-800 py-1 px-3 rounded-full text-xs font-medium";
  }
};

const closeModal = () => {
  isModalOpen.value = false;
  selectedPurchase.value = null;
};

const closeQcModal = () => {
  isQcModalOpen.value = false;
  qcError.value = ""; // Reset form
  qcForm.value = {
    purchase_id: null,
    supplierName: "",
    purchaseDate: "",
    items: [],
  };
};

// Fungsi untuk memuat detail PO dan menampilkan Modal Detail
const showDetail = async (id) => {
  try {
    const res = await axios.get(`${API_BASE}/history/purchases/${id}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    selectedPurchase.value = res.data.data;
    isModalOpen.value = true;
  } catch (err) {
    console.error(
      "Error loading purchase detail:",
      err.response?.data || err.message
    );
    alert(`Gagal memuat detail Purchase Order #${id}.`);
  }
};

// --- API Calls ---
const loadPurchases = async () => {
  loading.value = true;
  purchases.value = [];
  try {
    const res = await axios.get(`${API_BASE}/history/purchases`, {
      headers: { Authorization: `Bearer ${token}` },
      params: {
        page: page.value,
        size: size.value,
        q: searchQuery.value,
        status: currentStatus.value,
      },
    });

    const data = res.data.data?.data || res.data.data || [];
    totalItems.value = res.data.data?.total || 0;
    purchases.value = data;
  } catch (err) {
    console.error("Error load purchases:", err.response?.data || err.message);
    alert("Gagal memuat daftar Purchase Order.");
  } finally {
    loading.value = false;
  }
};

// Fungsi untuk membuka Modal QC dan memuat data PO
const openQcModal = async (id) => {
  try {
    const res = await axios.get(`${API_BASE}/history/purchases/${id}`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    const po = res.data.data;

    if (!po || po.Status !== "APPROVED") {
      alert("PO ini tidak bisa diproses karena statusnya bukan APPROVED.");
      return;
    } // Siapkan data untuk form QC

    qcForm.value.purchase_id = po.ID;
    qcForm.value.supplierName = po.Supplier?.name || "N/A";
    qcForm.value.purchaseDate = po.Date;

    qcForm.value.items = po.Items.map((item) => {
      // Gunakan computed property untuk Total Received
      const itemProxy = {
        purchase_item_id: item.ID,
        product_id: item.Product_ID,
        product_name: item.Product?.name || "Produk Hilang",
        qty_ordered: item.Qty,
        qty_good: item.Qty, // Default: Qty baik sama dengan Qty dipesan
        qty_damaged: 0, // Default: Qty rusak 0
        note: "",
      };

      Object.defineProperty(itemProxy, "total_received", {
        enumerable: true,
        get: () => itemProxy.qty_good + itemProxy.qty_damaged,
      });
      return itemProxy;
    });

    isQcModalOpen.value = true;
  } catch (err) {
    console.error(
      "Error loading PO for QC:",
      err.response?.data || err.message
    );
    alert(`Gagal memuat detail PO #${id} untuk QC.`);
  }
};

// Fungsi untuk submit QC Report dan Good Receipt
const submitQcReport = async () => {
  // Periksa validasi sebelum submit
  if (!isQcFormValid.value) {
    qcError.value =
      "Kuantitas total yang diterima (Baik + Rusak) tidak boleh melebihi Qty dipesan.";
    return;
  }

  submitting.value = true;
  qcError.value = ""; // Format payload untuk API backend:

  const payloadItems = qcForm.value.items.map((item) => ({
    purchase_item_id: item.purchase_item_id,
    qty_good: item.qty_good, // Kuantitas yang LULUS QC (Masuk Stok)
    qty_damaged: item.qty_damaged, // Kuantitas Rusak (Diproses Retur)
    note: item.note,
  }));

  const payload = {
    items: payloadItems,
  };

  const id = qcForm.value.purchase_id;

  try {
    await axios.post(`${API_BASE}/gudang/purchases/${id}/qc-receive`, payload, {
      headers: { Authorization: `Bearer ${token}` },
    });

    alert(
      `QC Report untuk Purchase Order #${id} berhasil dicatat. Stok telah diperbarui dan status PO diubah menjadi RECEIVED.`
    );

    closeQcModal();
    // FIX 2: Perbarui filter status agar PO yang baru di-RECEIVE muncul.
    // Jika filter saat ini adalah APPROVED, kita ubah ke RECEIVED atau SEMUA.
    if (currentStatus.value === "APPROVED") {
      currentStatus.value = "RECEIVED";
    }

    loadPurchases(); // Muat ulang data dengan filter yang sudah diperbarui
  } catch (err) {
    console.error(
      "Error submitting QC Report:",
      err.response?.data || err.message
    );
    qcError.value = `Gagal menyelesaikan QC & Penerimaan: ${
      err.response?.data?.error ||
      err.response?.data?.message ||
      "Terjadi kesalahan pada server."
    }`;
  } finally {
    submitting.value = false;
  }
};

onMounted(() => {
  loadPurchases();
});
</script>
