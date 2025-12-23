<template>
  <OwnerLayouts>
    <div class="p-8 bg-gradient-to-br from-slate-50 to-slate-100 min-h-screen">
      <div class="mb-8">
        <h1
          class="text-slate-800 text-4xl font-bold leading-tight tracking-tight"
        >
          Master Data Management
        </h1>
        <p class="text-slate-600 text-sm font-normal leading-normal mt-2">
          Kelola semua data master sistem Anda dengan mudah dan cepat
        </p>
      </div>

      <!-- Quick Stats Row -->
      <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4 mb-8">
        <div
          class="bg-white rounded-lg p-4 shadow-sm border border-slate-200 text-center"
        >
          <p class="text-3xl font-bold text-blue-600">{{ userCount }}</p>
          <p class="text-slate-600 text-xs mt-1">User</p>
        </div>
        <div
          class="bg-white rounded-lg p-4 shadow-sm border border-slate-200 text-center"
        >
          <p class="text-3xl font-bold text-emerald-600">{{ productCount }}</p>
          <p class="text-slate-600 text-xs mt-1">Produk</p>
        </div>
        <div
          class="bg-white rounded-lg p-4 shadow-sm border border-slate-200 text-center"
        >
          <p class="text-3xl font-bold text-amber-600">{{ categoryCount }}</p>
          <p class="text-slate-600 text-xs mt-1">Kategori</p>
        </div>
        <div
          class="bg-white rounded-lg p-4 shadow-sm border border-slate-200 text-center"
        >
          <p class="text-3xl font-bold text-purple-600">{{ supplierCount }}</p>
          <p class="text-slate-600 text-xs mt-1">Supplier</p>
        </div>
        <div
          class="bg-white rounded-lg p-4 shadow-sm border border-slate-200 text-center"
        >
          <p class="text-3xl font-bold text-pink-600">{{ unitCount }}</p>
          <p class="text-slate-600 text-xs mt-1">Unit</p>
        </div>
        <div
          class="bg-white rounded-lg p-4 shadow-sm border border-slate-200 text-center"
        >
          <p class="text-3xl font-bold text-red-600">{{ voucherCount }}</p>
          <p class="text-slate-600 text-xs mt-1">Voucher</p>
        </div>
      </div>

      <!-- Master Data Cards Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <!-- User Management Card -->
        <router-link
          to="/management-user"
          class="bg-gradient-to-br from-blue-50 to-blue-100 rounded-xl p-6 border border-blue-200 shadow-sm hover:shadow-lg transition-all hover:scale-105 cursor-pointer group"
        >
          <div class="flex items-start justify-between mb-4">
            <div class="text-4xl group-hover:scale-110 transition-transform">
              👤
            </div>
            <span
              class="px-3 py-1 bg-blue-200 text-blue-700 rounded-full text-xs font-medium"
            >
              {{ userCount }} items
            </span>
          </div>
          <h3 class="text-slate-800 text-xl font-bold mb-2">Manage User</h3>
          <p class="text-slate-600 text-sm mb-4">
            Kelola data pengguna sistem, role, dan status akses
          </p>
          <div
            class="flex items-center text-blue-600 font-medium text-sm group-hover:gap-2 transition-all"
          >
            Buka <span class="ml-1">→</span>
          </div>
        </router-link>

        <!-- Product Management Card -->
        <router-link
          to="/management-product"
          class="bg-gradient-to-br from-emerald-50 to-emerald-100 rounded-xl p-6 border border-emerald-200 shadow-sm hover:shadow-lg transition-all hover:scale-105 cursor-pointer group"
        >
          <div class="flex items-start justify-between mb-4">
            <div class="text-4xl group-hover:scale-110 transition-transform">
              📦
            </div>
            <span
              class="px-3 py-1 bg-emerald-200 text-emerald-700 rounded-full text-xs font-medium"
            >
              {{ productCount }} items
            </span>
          </div>
          <h3 class="text-slate-800 text-xl font-bold mb-2">Products</h3>
          <p class="text-slate-600 text-sm mb-4">
            Kelola daftar produk, harga jual, harga beli, dan stok
          </p>
          <div
            class="flex items-center text-emerald-600 font-medium text-sm group-hover:gap-2 transition-all"
          >
            Buka <span class="ml-1">→</span>
          </div>
        </router-link>

        <!-- Category Management Card -->
        <router-link
          to="/management-category"
          class="bg-gradient-to-br from-amber-50 to-amber-100 rounded-xl p-6 border border-amber-200 shadow-sm hover:shadow-lg transition-all hover:scale-105 cursor-pointer group"
        >
          <div class="flex items-start justify-between mb-4">
            <div class="text-4xl group-hover:scale-110 transition-transform">
              🏷️
            </div>
            <span
              class="px-3 py-1 bg-amber-200 text-amber-700 rounded-full text-xs font-medium"
            >
              {{ categoryCount }} items
            </span>
          </div>
          <h3 class="text-slate-800 text-xl font-bold mb-2">Kategori</h3>
          <p class="text-slate-600 text-sm mb-4">
            Kelola kategori produk untuk organisasi yang lebih baik
          </p>
          <div
            class="flex items-center text-amber-600 font-medium text-sm group-hover:gap-2 transition-all"
          >
            Buka <span class="ml-1">→</span>
          </div>
        </router-link>

        <!-- Supplier Management Card -->
        <router-link
          to="/management-supplier"
          class="bg-gradient-to-br from-purple-50 to-purple-100 rounded-xl p-6 border border-purple-200 shadow-sm hover:shadow-lg transition-all hover:scale-105 cursor-pointer group"
        >
          <div class="flex items-start justify-between mb-4">
            <div class="text-4xl group-hover:scale-110 transition-transform">
              🚚
            </div>
            <span
              class="px-3 py-1 bg-purple-200 text-purple-700 rounded-full text-xs font-medium"
            >
              {{ supplierCount }} items
            </span>
          </div>
          <h3 class="text-slate-800 text-xl font-bold mb-2">Supplier</h3>
          <p class="text-slate-600 text-sm mb-4">
            Kelola data supplier, kontak, dan alamat
          </p>
          <div
            class="flex items-center text-purple-600 font-medium text-sm group-hover:gap-2 transition-all"
          >
            Buka <span class="ml-1">→</span>
          </div>
        </router-link>

        <!-- Unit Management Card -->
        <router-link
          to="/management-unit"
          class="bg-gradient-to-br from-pink-50 to-pink-100 rounded-xl p-6 border border-pink-200 shadow-sm hover:shadow-lg transition-all hover:scale-105 cursor-pointer group"
        >
          <div class="flex items-start justify-between mb-4">
            <div class="text-4xl group-hover:scale-110 transition-transform">
              ⚖️
            </div>
            <span
              class="px-3 py-1 bg-pink-200 text-pink-700 rounded-full text-xs font-medium"
            >
              {{ unitCount }} items
            </span>
          </div>
          <h3 class="text-slate-800 text-xl font-bold mb-2">Unit</h3>
          <p class="text-slate-600 text-sm mb-4">
            Kelola satuan ukuran (Pcs, Kg, Liter, dll)
          </p>
          <div
            class="flex items-center text-pink-600 font-medium text-sm group-hover:gap-2 transition-all"
          >
            Buka <span class="ml-1">→</span>
          </div>
        </router-link>

        <!-- Voucher Management Card -->
        <router-link
          to="/management-voucher"
          class="bg-gradient-to-br from-red-50 to-red-100 rounded-xl p-6 border border-red-200 shadow-sm hover:shadow-lg transition-all hover:scale-105 cursor-pointer group"
        >
          <div class="flex items-start justify-between mb-4">
            <div class="text-4xl group-hover:scale-110 transition-transform">
              🎟️
            </div>
            <span
              class="px-3 py-1 bg-red-200 text-red-700 rounded-full text-xs font-medium"
            >
              {{ voucherCount }} items
            </span>
          </div>
          <h3 class="text-slate-800 text-xl font-bold mb-2">Voucher</h3>
          <p class="text-slate-600 text-sm mb-4">
            Kelola voucher diskon dan promosi penjualan
          </p>
          <div
            class="flex items-center text-red-600 font-medium text-sm group-hover:gap-2 transition-all"
          >
            Buka <span class="ml-1">→</span>
          </div>
        </router-link>
      </div>

      <!-- Info Box -->
      <div
        class="mt-8 bg-white rounded-xl p-6 shadow-sm border border-slate-200"
      >
        <h3 class="text-slate-800 text-lg font-bold mb-3">💡 Tips</h3>
        <ul class="text-slate-600 text-sm space-y-2">
          <li>✓ Klik card di atas untuk mengelola setiap data master</li>
          <li>
            ✓ Gunakan tombol "Add" di setiap halaman untuk membuat data baru
          </li>
          <li>
            ✓ Data yang ada ditampilkan dalam tabel dengan opsi edit/delete
          </li>
          <li>
            ✓ Pastikan semua data master lengkap sebelum membuat transaksi
          </li>
        </ul>
      </div>
    </div>
  </OwnerLayouts>
</template>

<script setup>
import OwnerLayouts from "@/components/OwnerLayouts.vue";
import { ref, onMounted } from "vue";
import api from "@/services/api";

const userCount = ref(0);
const productCount = ref(0);
const categoryCount = ref(0);
const supplierCount = ref(0);
const unitCount = ref(0);
const voucherCount = ref(0);

const fetchCounts = async () => {
  try {
    // Fetch all data untuk count dengan fallback error
    const [users, products, categories, suppliers, units, vouchers] =
      await Promise.all([
        api.get("/owner/users").catch(() => ({ data: [] })),
        api.get("/owner/products").catch(() => ({ data: [] })),
        api.get("/owner/categories").catch(() => ({ data: [] })),
        api.get("/owner/suppliers").catch(() => ({ data: [] })),
        api.get("/owner/units").catch(() => ({ data: [] })),
        api.get("/vouchers").catch(() => ({ data: [] })),
      ]);

    userCount.value = (users.data?.data || users.data || []).length;
    productCount.value = (products.data?.data || products.data || []).length;
    categoryCount.value = (
      categories.data?.data ||
      categories.data ||
      []
    ).length;
    supplierCount.value = (suppliers.data?.data || suppliers.data || []).length;
    unitCount.value = (units.data?.data || units.data || []).length;
    voucherCount.value = (vouchers.data?.data || vouchers.data || []).length;
  } catch (err) {
    console.error("Failed to fetch counts", err);
    // Set dummy values untuk demo
    userCount.value = 5;
    productCount.value = 24;
    categoryCount.value = 8;
    supplierCount.value = 12;
    unitCount.value = 6;
    voucherCount.value = 3;
  }
};

onMounted(() => {
  fetchCounts();
});
</script>

<style scoped>
/* Minimal scoped styles */
</style>
