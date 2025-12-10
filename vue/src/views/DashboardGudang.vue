<template>
  <GudangLayouts>
    <div class="container mx-auto px-4 py-8 font-poppins">
      <!-- Header Section -->
      <div class="mb-8">
        <h1 class="text-2xl md:text-3xl font-semibold text-gray-900">
          Warehouse Staff Dashboard
        </h1>
        <p class="mt-2 text-base text-gray-600">
          Welcome back! Here's an overview of your warehouse activities.
        </p>
      </div>

      <!-- Summary Cards -->
      <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
        <div
          class="bg-white rounded-lg shadow-md p-6 border border-blue-100 hover:shadow-lg transition-shadow duration-200"
        >
          <p class="text-sm font-medium text-blue-600">Total Stock Opname</p>
          <div class="mt-2 flex items-center justify-between">
            <p class="text-2xl font-bold text-gray-900">
              {{ summary.totalOpnames || 0 }}
            </p>
            <svg
              class="w-7 h-7 text-blue-400"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"
              ></path>
            </svg>
          </div>
          <p class="text-xs text-gray-500 mt-3">Data tercatat sejak awal.</p>
        </div>
      </div>

      <!-- Add New Opname Button -->
      <div class="flex justify-end mb-8">
        <button
          @click="showModal = true"
          class="px-5 py-2.5 bg-blue-600 text-white font-medium rounded-lg shadow-md hover:bg-blue-700 transition-colors duration-200"
        >
          + Add New Opname
        </button>
      </div>

      <!-- Recent Stock Opname Table -->
      <div class="bg-white rounded-lg shadow-md border border-gray-100">
        <div
          class="flex justify-between items-center p-6 border-b border-gray-100"
        >
          <h2 class="text-xl font-semibold text-gray-800">
            Riwayat Stock Opname Terbaru
          </h2>
          <router-link
            to="/history/stock-opnames"
            class="px-4 py-1.5 text-sm bg-gray-200 text-gray-700 rounded-md hover:bg-gray-300 transition-colors duration-200"
          >
            Lihat Semua Riwayat
          </router-link>
        </div>

        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th
                  class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
                >
                  Date
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
                >
                  Staff
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
                >
                  Product Name
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
                >
                  Qty Fisik
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
                >
                  Qty System
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-semibold text-gray-600 uppercase tracking-wider"
                >
                  Selisih
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-100">
              <tr v-if="loading">
                <td
                  colspan="6"
                  class="px-6 py-4 text-center text-sm text-gray-500"
                >
                  Loading...
                </td>
              </tr>
              <tr v-else-if="opnameItems.length === 0">
                <td
                  colspan="6"
                  class="px-6 py-4 text-center text-sm text-gray-500"
                >
                  No data available.
                </td>
              </tr>
              <tr
                v-else
                v-for="item in opnameItems"
                :key="item.ID"
                class="hover:bg-gray-50 transition-colors duration-150"
              >
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">
                  {{ formatDate(item.Date) }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">
                  {{ item.StaffName || "Unknown" }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">
                  {{ item.ProductName || "N/A" }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">
                  {{ item.QtyPhysical }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">
                  {{ item.QtySystem }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-700">
                  <span
                    :class="{
                      'text-red-500 font-semibold':
                        item.QtyPhysical - item.QtySystem < 0,
                      'text-green-500 font-semibold':
                        item.QtyPhysical - item.QtySystem > 0,
                    }"
                  >
                    {{ item.QtyPhysical - item.QtySystem }}
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <p
          v-if="errorMessage && !loading"
          class="text-red-500 text-sm mt-4 p-4 bg-red-50 rounded-b-lg border-t border-red-200"
        >
          **Error:** {{ errorMessage }}
        </p>
      </div>

      <!-- Modal for Creating New Stock Opname -->
      <div
        v-if="showModal"
        class="fixed inset-0 z-50 overflow-y-auto bg-gray-900 bg-opacity-60 flex justify-center items-center p-4"
        @click.self="showModal = false"
      >
        <div
          class="bg-white rounded-lg shadow-xl w-full max-w-2xl p-8 relative transform transition-all duration-300"
          @click.stop
        >
          <h3 class="text-xl font-semibold text-gray-800 mb-6 border-b pb-3">
            Create New Stock Opname
          </h3>
          <button
            @click="showModal = false"
            class="absolute top-4 right-4 text-gray-400 hover:text-gray-600 text-2xl leading-none"
          >
            &times;
          </button>

          <form
            @submit.prevent="submitOpname"
            class="space-y-6 max-h-[70vh] overflow-y-auto pr-2"
          >
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1.5"
                  >Date</label
                >
                <input
                  v-model="opnameForm.date"
                  type="date"
                  required
                  class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500 py-2.5 px-3"
                  :min="currentDate"
                />
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-1.5"
                  >Note (Optional)</label
                >
                <textarea
                  v-model="opnameForm.note"
                  class="block w-full border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500 py-2.5 px-3"
                  rows="3"
                  placeholder="Enter any notes..."
                ></textarea>
              </div>
            </div>

            <div class="border p-4 rounded-lg bg-gray-50">
              <label class="block text-base font-semibold text-gray-800 mb-4"
                >Opname Items</label
              >
              <div
                v-for="(item, index) in opnameForm.items"
                :key="index"
                class="flex space-x-3 mb-4 items-center"
              >
                <select
                  v-model="item.product_id"
                  required
                  class="w-2/3 border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500 text-sm py-2.5 px-3"
                >
                  <option value="" disabled>Select Product</option>
                  <option
                    v-for="product in products"
                    :key="product.id"
                    :value="product.id"
                  >
                    {{ product.name }}
                  </option>
                </select>
                <input
                  v-model.number="item.qty_fisik"
                  type="number"
                  min="0"
                  required
                  class="w-1/3 border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500 text-sm py-2.5 px-3 text-center"
                  placeholder="Qty Fisik"
                />
                <button
                  @click.prevent="removeItem(index)"
                  type="button"
                  :disabled="opnameForm.items.length === 1"
                  class="text-red-500 hover:text-red-700 text-xl disabled:opacity-30 disabled:cursor-not-allowed"
                  title="Remove Item"
                >
                  &times;
                </button>
              </div>
              <button
                @click.prevent="addItem"
                type="button"
                class="mt-3 px-4 py-2 text-sm bg-green-500 text-white rounded-lg hover:bg-green-600 transition-colors duration-200"
              >
                + Add Item
              </button>
            </div>

            <div class="flex justify-end pt-4">
              <button
                type="button"
                @click="showModal = false"
                class="px-6 py-2.5 border border-gray-300 text-gray-700 rounded-lg shadow-sm hover:bg-gray-100 transition-colors duration-200 mr-3"
              >
                Cancel
              </button>
              <button
                type="submit"
                :disabled="submitting || !isFormValid"
                class="px-6 py-2.5 bg-blue-600 text-white font-medium rounded-lg shadow-md hover:bg-blue-700 disabled:opacity-50 transition-colors duration-200"
              >
                {{ submitting ? "Submitting..." : "Submit Opname" }}
              </button>
            </div>

            <p v-if="errorMessage" class="text-red-500 text-sm pt-3">
              {{ errorMessage }}
            </p>
          </form>
        </div>
      </div>
    </div>
  </GudangLayouts>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import GudangLayouts from "@/components/GudangLayouts.vue";
import axios from "axios";

// --- Variabel State ---
const showModal = ref(false);
const stockOpnames = ref([]); // Data mentah dari API
const opnameItems = ref([]); // Data flat untuk tabel
const loading = ref(false);
const products = ref([]);
const opnameForm = ref({
  date: new Date().toISOString().split("T")[0],
  note: "",
  items: [{ product_id: null, qty_fisik: 0 }],
});
const submitting = ref(false);
const errorMessage = ref("");
const summary = ref({ totalOpnames: 0 });

// --- Konfigurasi API ---
const api = axios.create({ baseURL: "http://localhost:8081/api/v1" });
const savedToken = localStorage.getItem("token");
if (savedToken) {
  api.defaults.headers.common["Authorization"] = `Bearer ${savedToken}`;
} else {
  errorMessage.value = "Authentication token not found. Please log in.";
}

const currentDate = computed(() => new Date().toISOString().split("T")[0]);
const isFormValid = computed(() => {
  return opnameForm.value.items.every(
    (item) => item.product_id && item.qty_fisik >= 0 && item.qty_fisik !== null
  );
});

// --- Fungsi Baru: Meratakan Struktur Data (Membaca tag JSON Go yang benar) ---
const flattenOpnames = (opnames) => {
  const flattened = [];

  opnames.forEach((opname) => {
    // Pastikan properti 'items' (huruf kecil) dari respons API
    const items =
      opname.items && opname.items.length > 0
        ? opname.items
        : [
            {
              product: { name: "N/A" },
              physicalStock: 0, // Menggunakan properti JSON yang benar
              systemStock: 0, // Menggunakan properti JSON yang benar
            },
          ];

    items.forEach((item) => {
      // Perhatikan: opname.date, opname.user?.name, opname.items, item.product?.name
      // sudah bekerja di perbaikan sebelumnya, kita fokus pada Qty.
      const uniqueId = item.id
        ? item.id
        : `${opname.id}-${item.product_id || Math.random()}`;

      const qtyPhysical = item.physicalStock || 0; // Mengambil dari "physicalStock"
      const qtySystem = item.systemStock || 0; // Mengambil dari "systemStock"

      flattened.push({
        ID: uniqueId,
        Date: opname.date,
        StaffName: opname.user?.name || opname.user?.Name || "Unknown",
        ProductName: item.product?.name || "N/A",

        // Memastikan variabel yang digunakan di template sudah terisi
        QtyPhysical: qtyPhysical,
        QtySystem: qtySystem,
      });
    });
  });
  return flattened;
};

// --- Fungsi Data Fetching Ringkasan ---
const fetchStockOpnameSummary = async () => {
  try {
    const response = await api.get("/history/stock-opnames", {
      params: { size: 1 },
    });
    const total = response.data.total || 0;
    summary.value.totalOpnames = total;
  } catch (error) {
    console.error("Gagal fetch ringkasan opnames:", error);
  }
};

// --- Fungsi Data Fetching Riwayat Terbaru (Hanya 5 data) ---
const fetchStockOpnames = async () => {
  loading.value = true;
  errorMessage.value = "";
  try {
    const response = await api.get("/history/stock-opnames", {
      params: { size: 5, sort: "date_desc" },
    });
    console.log("API Response:", response.data);

    const data = response.data.data || [];

    stockOpnames.value = data;
    opnameItems.value = flattenOpnames(data); // Panggil fungsi flatten
  } catch (error) {
    console.error("Gagal fetch stock opnames:", error);
    errorMessage.value = `Gagal memuat riwayat opname: ${
      error.response?.status
    } - ${error.response?.data?.error || "Terjadi kesalahan pada server."}`;
  } finally {
    loading.value = false;
  }
};

const fetchProducts = async () => {
  try {
    const response = await api.get("/products");
    products.value = response.data.data || [];
  } catch (error) {
    console.error("Gagal fetch products:", error);
  }
};

// --- Fungsi Item Form ---
const addItem = () => {
  opnameForm.value.items.push({ product_id: null, qty_fisik: 0 });
};

const removeItem = (index) => {
  if (opnameForm.value.items.length > 1) {
    opnameForm.value.items.splice(index, 1);
  }
};

// --- Fungsi Submit ---
const submitOpname = async () => {
  if (!isFormValid.value) {
    errorMessage.value =
      "Harap isi semua kolom produk dan kuantitas fisik yang valid.";
    return;
  }

  submitting.value = true;
  errorMessage.value = "";
  try {
    const payload = {
      date: opnameForm.value.date,
      note: opnameForm.value.note,
      items: opnameForm.value.items.filter(
        (item) => item.product_id && item.qty_fisik >= 0
      ),
    };
    const response = await api.post("/gudang/stock-opnames", payload);

    if (response.status === 201) {
      alert("Stock opname berhasil dibuat dan stok telah disesuaikan!");
      opnameForm.value = {
        date: new Date().toISOString().split("T")[0],
        note: "",
        items: [{ product_id: null, qty_fisik: 0 }],
      };
      showModal.value = false;
      fetchStockOpnameSummary();
      fetchStockOpnames();
    }
  } catch (error) {
    console.error("Gagal submit opname:", error);
    errorMessage.value = `Gagal membuat stock opname: ${
      error.response?.status
    } - ${error.response?.data?.error || "Terjadi kesalahan pada server."}`;
  } finally {
    submitting.value = false;
  }
};

const formatDate = (dateStr) => {
  if (!dateStr) return "N/A";
  try {
    const date = new Date(dateStr);
    return date.toLocaleDateString("id-ID", {
      day: "2-digit",
      month: "2-digit",
      year: "numeric",
    });
  } catch (e) {
    return "Invalid Date";
  }
};

// --- Lifecycle Hook ---
onMounted(() => {
  fetchStockOpnameSummary();
  fetchStockOpnames();
  fetchProducts();
});
</script>

<style scoped>
/* Import Poppins font */
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700&display=swap");

.font-poppins {
  font-family: "Poppins", sans-serif;
}

/* Custom scrollbar styles */
.max-h-\[70vh\]::-webkit-scrollbar {
  width: 8px;
}
.max-h-\[70vh\]::-webkit-scrollbar-thumb {
  background-color: #cbd5e1;
  border-radius: 4px;
}
.max-h-\[70vh\]::-webkit-scrollbar-track {
  background: #f1f1f1;
}
</style>
