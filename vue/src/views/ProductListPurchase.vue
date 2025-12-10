<script setup>
import { ref, onMounted, computed, watch } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";
import PurchasingSidebar from "../components/PurchasingSidebar.vue";

const router = useRouter();
const products = ref([]);
const searchQuery = ref("");
const selectedCategory = ref("");
const selectedStockFilter = ref("");
const categories = ref([]);
const loading = ref(true);
const productError = ref(null);

const api = axios.create({
  baseURL: "http://localhost:8081",
  headers: {
    "Content-Type": "application/json",
  },
});

api.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Ambil produk
const fetchProducts = async () => {
  try {
    const params = {};

    if (selectedCategory.value) {
      params.category_id = selectedCategory.value; // kirim ID kategori
    }
    if (searchQuery.value) {
      params.search = searchQuery.value;
    }

    const response = await api.get("/api/v1/products", { params });
    products.value = response.data.data.map((product) => ({
      id: product.id || 0,
      name: product.name || "Produk Tidak Dikenal",
      code:
        product.code || `P-${(product.id || 0).toString().padStart(3, "0")}`,
      purchasePrice: `Rp ${(product.cost_price || 0).toLocaleString("id-ID")}`,
      sellingPrice: `Rp ${(product.sell_price || 0).toLocaleString("id-ID")}`,
      stock: product.stock || 0,
      categoryId: product.category?.id || null,
      category: product.category?.name || "Tanpa Kategori",
    }));
  } catch (err) {
    console.error("Gagal mengambil data produk:", err);
    productError.value =
      err.response?.status === 403
        ? "Akses ditolak: Anda tidak memiliki izin untuk mengambil data produk."
        : "Gagal memuat produk. Silakan coba lagi.";
  }
};

// Ambil kategori
const fetchCategories = async () => {
  try {
    const response = await api.get("/api/v1/categories");
    if (Array.isArray(response.data.data)) {
      categories.value = response.data.data;
    } else {
      categories.value = [];
    }
  } catch (err) {
    console.error("Gagal mengambil data kategori:", err);
  }
};

// Filter produk (frontend)
const filteredProducts = computed(() => {
  let filtered = products.value;

  if (searchQuery.value) {
    filtered = filtered.filter((product) =>
      product.name.toLowerCase().includes(searchQuery.value.toLowerCase())
    );
  }

  if (selectedCategory.value) {
    filtered = filtered.filter(
      (product) => product.categoryId == selectedCategory.value
    );
  }

  if (selectedStockFilter.value === "stok-ada") {
    filtered = filtered.filter((product) => product.stock > 0);
  } else if (selectedStockFilter.value === "stok-kosong") {
    filtered = filtered.filter((product) => product.stock === 0);
  }

  return filtered;
});

// Re-fetch kalau filter kategori berubah
watch(selectedCategory, fetchProducts);

const goToAddProduct = () => {
  router.push("/management-product");
};

onMounted(async () => {
  await Promise.all([fetchProducts(), fetchCategories()]);
  loading.value = false;
});
</script>

<template>
  <div
    class="relative flex size-full min-h-screen flex-col bg-white group/design-root overflow-x-hidden"
    style="font-family: Manrope, 'Noto Sans', sans-serif"
  >
    <div class="layout-container flex h-full grow flex-col">
      <div class="flex flex-1 justify-center">
        <PurchasingSidebar />
        <div
          class="layout-content-container flex flex-col max-w-[960px] flex-1"
        >
          <!-- Header -->
          <div class="flex flex-wrap justify-between gap-3 p-4">
            <p
              class="text-[#111418] text-[32px] font-bold leading-tight min-w-72"
            >
              Daftar Produk
            </p>
            <button
              @click="goToAddProduct"
              class="flex min-w-[84px] cursor-pointer items-center justify-center rounded-lg h-8 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-medium"
            >
              Tambah Produk
            </button>
          </div>

          <!-- Search -->
          <div class="px-4 py-3">
            <label class="flex flex-col min-w-40 h-12 w-full">
              <div class="flex w-full flex-1 items-stretch rounded-lg h-full">
                <div
                  class="text-[#60758a] flex bg-[#f0f2f5] items-center justify-center pl-4 rounded-l-lg"
                >
                  🔍
                </div>
                <input
                  v-model="searchQuery"
                  placeholder="Cari produk"
                  class="flex w-full flex-1 rounded-r-lg bg-[#f0f2f5] px-4 text-base focus:outline-0"
                />
              </div>
            </label>
          </div>

          <!-- Filters -->
          <div class="flex gap-3 p-3 flex-wrap pr-4">
            <!-- Dropdown Kategori -->
            <select
              v-model="selectedCategory"
              class="flex h-8 items-center rounded-lg bg-[#f0f2f5] pl-4 pr-2 text-sm font-medium"
              :disabled="!categories.length"
            >
              <option value="">Semua Kategori</option>
              <option v-for="cat in categories" :key="cat.id" :value="cat.id">
                {{ cat.name }}
              </option>
            </select>

            <!-- Dropdown Stok -->
            <select
              v-model="selectedStockFilter"
              class="flex h-8 items-center rounded-lg bg-[#f0f2f5] pl-4 pr-2 text-sm font-medium"
            >
              <option value="">Semua Stok</option>
              <option value="stok-ada">Stok Ada</option>
              <option value="stok-kosong">Stok Kosong</option>
            </select>
          </div>

          <!-- Table -->
          <div class="px-4 py-3">
            <div v-if="loading" class="text-center py-4 text-gray-500">
              Memuat data produk...
            </div>
            <div v-else-if="productError" class="text-center py-4 text-red-500">
              {{ productError }}
            </div>
            <div
              v-else-if="filteredProducts.length === 0"
              class="text-center py-4 text-gray-500"
            >
              Tidak ada produk ditemukan.
            </div>
            <div
              v-else
              class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
            >
              <table class="flex-1">
                <thead>
                  <tr class="bg-white">
                    <th class="px-4 py-3 text-left text-sm font-medium">
                      Nama
                    </th>
                    <th class="px-4 py-3 text-left text-sm font-medium">
                      Kode
                    </th>
                    <th class="px-4 py-3 text-left text-sm font-medium">
                      Harga Beli
                    </th>
                    <th class="px-4 py-3 text-left text-sm font-medium">
                      Harga Jual
                    </th>
                    <th class="px-4 py-3 text-left text-sm font-medium">
                      Stok
                    </th>
                    <th class="px-4 py-3 text-left text-sm font-medium">
                      Kategori
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr
                    v-for="product in filteredProducts"
                    :key="product.id"
                    class="border-t"
                  >
                    <td class="px-4 py-2 text-sm">{{ product.name }}</td>
                    <td class="px-4 py-2 text-sm text-[#60758a]">
                      {{ product.code }}
                    </td>
                    <td class="px-4 py-2 text-sm text-[#60758a]">
                      {{ product.purchasePrice }}
                    </td>
                    <td class="px-4 py-2 text-sm text-[#60758a]">
                      {{ product.sellingPrice }}
                    </td>
                    <td class="px-4 py-2 text-sm text-[#60758a]">
                      {{ product.stock }}
                    </td>
                    <td class="px-4 py-2 text-sm text-[#60758a]">
                      {{ product.category }}
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
