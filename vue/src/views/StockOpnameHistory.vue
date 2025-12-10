<template>
  <div class="flex">
    <GudangSidebar />

    <main class="flex-1 ml-64 p-6">
      <h1 class="text-2xl font-bold mb-6">Riwayat Stock Opname</h1>

      <!-- Filter & Search -->
      <div class="flex justify-between items-center mb-4">
        <div class="flex gap-4">
          <label class="text-sm font-medium pt-2">Dari Tanggal:</label>
          <input
            type="date"
            v-model="startDate"
            @change="loadStockOpnames"
            class="border rounded p-2 text-sm"
          />
          <label class="text-sm font-medium pt-2">Sampai Tanggal:</label>
          <input
            type="date"
            v-model="endDate"
            @change="loadStockOpnames"
            class="border rounded p-2 text-sm"
          />
        </div>
        <input
          type="text"
          v-model="searchQuery"
          @keyup.enter="loadStockOpnames"
          placeholder="Cari ID atau User..."
          class="border rounded p-2 text-sm w-64"
        />
      </div>

      <!-- Table Riwayat -->
      <div class="bg-white shadow-md rounded-lg overflow-hidden mb-8">
        <table class="min-w-full leading-normal">
          <thead>
            <tr
              class="bg-gray-100 text-gray-600 uppercase text-sm leading-normal"
            >
              <th class="py-3 px-6 text-left">ID SO</th>
              <th class="py-3 px-6 text-left">Tanggal</th>
              <th class="py-3 px-6 text-left">Dibuat Oleh</th>
              <th class="py-3 px-6 text-right">Total Item Dicatat</th>
              <th class="py-3 px-6 text-center">Aksi</th>
            </tr>
          </thead>
          <tbody class="text-gray-600 text-sm font-light">
            <tr
              v-for="so in stockOpnames"
              :key="so.id"
              class="border-b border-gray-200 hover:bg-gray-50"
            >
              <td class="py-3 px-6 text-left whitespace-nowrap">{{ so.id }}</td>
              <td class="py-3 px-6 text-left">{{ formatDate(so.date) }}</td>
              <td class="py-3 px-6 text-left">{{ so.user?.Name || "N/A" }}</td>
              <td class="py-3 px-6 text-right">
                {{ so.items?.length || 0 }}
              </td>
              <td class="py-3 px-6 text-center">
                <button
                  @click="showDetail(so.id)"
                  class="text-blue-500 hover:text-blue-700 text-xs font-medium"
                >
                  Detail
                </button>
              </td>
            </tr>

            <tr v-if="!stockOpnames.length && !loading">
              <td colspan="5" class="py-4 text-center text-gray-500">
                Tidak ada riwayat Stock Opname ditemukan.
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Table Categories -->
      <div class="bg-white shadow-md rounded-lg overflow-hidden mb-8">
        <div class="p-6 border-b">
          <h2 class="text-xl font-bold mb-2">Daftar Kategori</h2>
          <p class="text-gray-600 text-sm">Master data kategori produk.</p>
        </div>
        <div class="overflow-x-auto">
          <table class="min-w-full leading-normal">
            <thead>
              <tr
                class="bg-gray-100 text-gray-600 uppercase text-sm leading-normal"
              >
                <th class="py-3 px-6 text-left">ID</th>
                <th class="py-3 px-6 text-left">Nama Kategori</th>
                <th class="py-3 px-6 text-left">Deskripsi</th>
                <th class="py-3 px-6 text-center">Aksi</th>
              </tr>
            </thead>
            <tbody class="text-gray-600 text-sm font-light">
              <tr
                v-for="cat in categories"
                :key="cat.id"
                class="border-b border-gray-200 hover:bg-gray-50"
              >
                <td class="py-3 px-6 text-left whitespace-nowrap">
                  {{ cat.id }}
                </td>
                <td class="py-3 px-6 text-left">{{ cat.name }}</td>
                <td class="py-3 px-6 text-left">
                  {{ cat.description || "-" }}
                </td>
                <td class="py-3 px-6 text-center">
                  <button
                    @click="editCategory(cat)"
                    class="text-blue-500 hover:text-blue-700 text-xs font-medium mr-2"
                  >
                    Edit
                  </button>
                  <button
                    @click="deleteCategory(cat.id)"
                    class="text-red-500 hover:text-red-700 text-xs font-medium"
                  >
                    Hapus
                  </button>
                </td>
              </tr>
              <tr v-if="!categories.length && !loadingCategories">
                <td colspan="4" class="py-4 text-center text-gray-500">
                  Tidak ada kategori ditemukan.
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Table Units -->
      <div class="bg-white shadow-md rounded-lg overflow-hidden">
        <div class="p-6 border-b">
          <h2 class="text-xl font-bold mb-2">Daftar Unit</h2>
          <p class="text-gray-600 text-sm">
            Master data unit pengukuran produk.
          </p>
        </div>
        <div class="overflow-x-auto">
          <table class="min-w-full leading-normal">
            <thead>
              <tr
                class="bg-gray-100 text-gray-600 uppercase text-sm leading-normal"
              >
                <th class="py-3 px-6 text-left">ID</th>
                <th class="py-3 px-6 text-left">Nama Unit</th>
                <th class="py-3 px-6 text-center">Aksi</th>
              </tr>
            </thead>
            <tbody class="text-gray-600 text-sm font-light">
              <tr
                v-for="unit in units"
                :key="unit.id"
                class="border-b border-gray-200 hover:bg-gray-50"
              >
                <td class="py-3 px-6 text-left whitespace-nowrap">
                  {{ unit.id }}
                </td>
                <td class="py-3 px-6 text-left">{{ unit.name }}</td>
                <td class="py-3 px-6 text-center">
                  <button
                    @click="editUnit(unit)"
                    class="text-blue-500 hover:text-blue-700 text-xs font-medium mr-2"
                  >
                    Edit
                  </button>
                  <button
                    @click="deleteUnit(unit.id)"
                    class="text-red-500 hover:text-red-700 text-xs font-medium"
                  >
                    Hapus
                  </button>
                </td>
              </tr>
              <tr v-if="!units.length && !loadingUnits">
                <td colspan="3" class="py-4 text-center text-gray-500">
                  Tidak ada unit ditemukan.
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- Info bawah tabel riwayat -->
      <div class="flex justify-between items-center mt-4 mb-8">
        <p class="text-sm text-gray-600">Total Riwayat: {{ totalItems }}</p>
        <div v-if="loading" class="text-blue-500">Memuat riwayat...</div>
      </div>

      <!-- Loading untuk tables master -->
      <div
        v-if="loadingCategories || loadingUnits"
        class="text-center text-blue-500 py-4"
      >
        Memuat data master...
      </div>
    </main>

    <!-- Modal Detail (tetap sama) -->
    <div
      v-if="isModalOpen"
      class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50 flex justify-center items-center"
      @click="closeModal"
    >
      <div
        class="bg-white p-6 rounded-lg shadow-xl max-w-6xl w-full mx-4"
        @click.stop
      >
        <div class="flex justify-between items-center border-b pb-3 mb-4">
          <h3 class="text-xl font-bold">
            Detail Stock Opname #{{ selectedSO?.id }}
          </h3>
          <button
            @click="closeModal"
            class="text-gray-500 hover:text-gray-800 text-2xl leading-none"
          >
            &times;
          </button>
        </div>

        <div v-if="selectedSO">
          <!-- Info Umum -->
          <div class="grid grid-cols-2 gap-4 text-sm mb-4">
            <div>
              <p><strong>ID SO:</strong> {{ selectedSO.id }}</p>
              <p><strong>Tanggal:</strong> {{ formatDate(selectedSO.date) }}</p>
            </div>
            <div>
              <p>
                <strong>Dibuat Oleh:</strong>
                {{ selectedSO.user?.Name || "N/A" }}
              </p>
              <p><strong>Keterangan:</strong> {{ selectedSO.notes || "-" }}</p>
            </div>
          </div>

          <!-- Detail Item -->
          <h4 class="text-lg font-semibold mb-2">Item Perubahan Stok</h4>
          <div class="overflow-x-auto border rounded">
            <table class="min-w-full text-sm">
              <thead>
                <tr class="bg-gray-50">
                  <th class="py-2 px-4 text-left">Produk</th>
                  <th class="py-2 px-4 text-left">Kategori</th>
                  <th class="py-2 px-4 text-left">Unit</th>
                  <th class="py-2 px-4 text-right">Stok Fisik</th>
                  <th class="py-2 px-4 text-right">Stok Sistem</th>
                  <th class="py-2 px-4 text-right">Selisih</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="item in selectedSO.items"
                  :key="item.id"
                  class="border-t"
                >
                  <td class="py-2 px-4">
                    {{ item.product?.name || "Produk Hilang" }}
                  </td>
                  <td class="py-2 px-4">
                    {{ item.product?.category?.name || "-" }}
                  </td>
                  <td class="py-2 px-4">
                    {{ item.product?.unit?.name || "-" }}
                  </td>
                  <td class="py-2 px-4 text-right">{{ item.physicalStock }}</td>
                  <td class="py-2 px-4 text-right">{{ item.systemStock }}</td>
                  <td
                    class="py-2 px-4 text-right"
                    :class="{
                      'text-red-600 font-semibold':
                        item.physicalStock - item.systemStock !== 0,
                    }"
                  >
                    {{ item.physicalStock - item.systemStock }}
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- Tombol Modal -->
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
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";
import GudangSidebar from "../components/GudangSidebar.vue"; // sesuaikan path

const API_BASE = "http://localhost:8081/api/v1"; // ganti sesuai backend
const token = localStorage.getItem("token");

const stockOpnames = ref([]);
const loading = ref(false);
const totalItems = ref(0);
const searchQuery = ref("");
const startDate = ref(null);
const endDate = ref(null);

// Tambahan untuk Categories dan Units
const categories = ref([]);
const units = ref([]);
const loadingCategories = ref(false);
const loadingUnits = ref(false);

// Modal
const selectedSO = ref(null);
const isModalOpen = ref(false);

// Pagination
const page = ref(1);
const size = ref(10);

// Helper date
const formatDate = (dateString) => {
  if (!dateString) return "-";
  return new Date(dateString).toLocaleDateString("id-ID", {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
  });
};

const closeModal = () => {
  isModalOpen.value = false;
  selectedSO.value = null;
};

// Fungsi untuk Categories
const loadCategories = async () => {
  loadingCategories.value = true;
  try {
    const res = await axios.get(`${API_BASE}/categories`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    categories.value = res.data.data || [];
  } catch (err) {
    console.error("Error load categories:", err.response?.data || err);
    alert("Gagal memuat daftar kategori.");
  } finally {
    loadingCategories.value = false;
  }
};

// Fungsi untuk Units
const loadUnits = async () => {
  loadingUnits.value = true;
  try {
    const res = await axios.get(`${API_BASE}/units`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    units.value = res.data.data || [];
  } catch (err) {
    console.error("Error load units:", err.response?.data || err);
    alert("Gagal memuat daftar unit.");
  } finally {
    loadingUnits.value = false;
  }
};

// Fungsi aksi (placeholder, implementasikan sesuai kebutuhan)
const editCategory = (cat) => {
  console.log("Edit category:", cat);
  // Implementasi edit modal atau redirect
};
const deleteCategory = async (id) => {
  if (confirm("Yakin hapus kategori ini?")) {
    try {
      await axios.delete(`${API_BASE}/categories/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      loadCategories();
    } catch (err) {
      alert("Gagal menghapus kategori.");
    }
  }
};
const editUnit = (unit) => {
  console.log("Edit unit:", unit);
  // Implementasi edit modal atau redirect
};
const deleteUnit = async (id) => {
  if (confirm("Yakin hapus unit ini?")) {
    try {
      await axios.delete(`${API_BASE}/units/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      loadUnits();
    } catch (err) {
      alert("Gagal menghapus unit.");
    }
  }
};

// === API Riwayat ===
const loadStockOpnames = async () => {
  loading.value = true;
  stockOpnames.value = [];
  try {
    const res = await axios.get(`${API_BASE}/history/stock-opnames`, {
      headers: { Authorization: `Bearer ${token}` },
      params: {
        page: page.value,
        size: size.value,
        q: searchQuery.value,
        start_date: startDate.value,
        end_date: endDate.value,
      },
    });

    stockOpnames.value = res.data.data || [];
    totalItems.value = res.data.total || stockOpnames.value.length;
  } catch (err) {
    console.error("Error load stock opname:", err.response?.data || err);
    alert("Gagal memuat daftar Riwayat Stock Opname.");
  } finally {
    loading.value = false;
  }
};

const showDetail = async (id) => {
  try {
    const res = await axios.get(`${API_BASE}/history/stock-opnames/${id}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    selectedSO.value = res.data.data;
    isModalOpen.value = true;
  } catch (err) {
    console.error("Error load detail:", err.response?.data || err);
    alert(`Gagal memuat detail Stock Opname #${id}`);
  }
};

onMounted(() => {
  loadStockOpnames();
  loadCategories();
  loadUnits();
});
</script>
