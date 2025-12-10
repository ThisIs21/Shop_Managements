<template>
  <div class="relative flex min-h-screen">
    <template v-if="showSidebar">
      <OwnerSidebar v-if="userRole === 'OWNER'" />

      <CashierSidebar v-else-if="userRole === 'KASIR'" />

      <GudangSidebar v-else-if="userRole === 'GUDANG'" />

      <PurchasingSidebar v-else-if="userRole === 'PEMBELIAN'" />

      <KepalaGudangSidebar v-else-if="userRole === 'KEPALA_GUDANG'" />

      <div v-else class="w-12 bg-gray-100 p-2 border-r"></div>
    </template>

    <main :class="mainClass">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { computed, ref, onMounted } from "vue";
import { useRoute } from "vue-router";

// Impor semua komponen sidebar spesifik Anda
import OwnerSidebar from "@/components/OwnerSidebar.vue";
import CashierSidebar from "@/components/CashierSidebar.vue";
import GudangSidebar from "@/components/GudangSidebar.vue";
import PurchasingSidebar from "@/components/PurchasingSidebar.vue";
import KepalaGudangSidebar from "@/components/KepalaGudangSidebar.vue";
// Ganti path impor di atas sesuai struktur proyek Anda

const route = useRoute();
const userRole = ref(null); // Gunakan ref untuk menyimpan role pengguna

onMounted(() => {
  // Ambil role dari localStorage saat komponen dimuat
  userRole.value = localStorage.getItem("userRole");
});

// Komputer untuk menentukan apakah sidebar harus ditampilkan
const showSidebar = computed(() => {
  // Logika: Tampilkan sidebar HANYA jika nama rute BUKAN 'Login' (rute Anda bernama 'Login')
  // dan role pengguna sudah dimuat.
  return route.name !== "Login" && userRole.value;
});

// Komputer untuk mengatur kelas pada <main> agar terpusat/ber-padding
const mainClass = computed(() => {
  // Jika showSidebar true (berarti ada sidebar yang dirender)
  if (showSidebar.value) {
    // Memberi ruang di kiri (pl-72) untuk sidebar
    return "flex-1 p-8 pl-72";
  } else {
    // Jika tidak ada sidebar (halaman login), gunakan w-full dan hapus padding yang mengganggu
    return "flex-1 w-full p-0";
  }
});
</script>

<style>
/* Gaya di bawah ini adalah gaya global.
  Pastikan mereka tidak bertabrakan dengan kelas Tailwind CSS Anda.
  Jika Anda menggunakan Tailwind, Anda bisa menghapus semua gaya di bawah ini.
*/
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
</style>

<!-- <template>
  <div class="relative flex min-h-screen">
    <AppSidebar />
    <main class="flex-1 p-8 pl-72">
      <router-view />
    </main>
  </div>
</template>

<script setup></script>

<style>
/* Gaya di bawah ini adalah gaya global.
  Pastikan mereka tidak bertabrakan dengan kelas Tailwind CSS Anda.
  Jika Anda menggunakan Tailwind, Anda bisa menghapus semua gaya di bawah ini.
*/
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
</style> -->
