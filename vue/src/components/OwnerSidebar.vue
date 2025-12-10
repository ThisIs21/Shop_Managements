<template>
  <aside
    class="w-64 bg-white border-r border-[#dbe0e6] p-4 fixed h-full top-0 left-0 overflow-y-auto flex flex-col"
  >
    <div class="flex items-center gap-4 pb-4 mb-4 border-b border-[#dbe0e6]">
      <div class="size-6 text-[#0d80f2]">
        <svg
          viewBox="0 0 24 24"
          fill="currentColor"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M22 6c0-1.1-.9-2-2-2H4c-1.1 0-2 .9-2 2v10c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V6zm-2 0l-8 5-8-5h16z"
          ></path>
          <circle cx="12" cy="14" r="2"></circle>
        </svg>
      </div>
      <h2 class="text-xl font-bold leading-tight">Dashboard Owner</h2>
    </div>

    <nav class="flex-grow">
      <ul>
        <li>
          <router-link
            to="/dashboard-owner"
            class="flex items-center gap-3 p-2 rounded-lg text-[#111418] hover:bg-[#f0f2f5] transition-colors"
            :class="{ 'bg-[#f0f2f5]': route.path === '/dashboard-owner' }"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <rect x="3" y="3" width="7" height="9"></rect>
              <rect x="14" y="3" width="7" height="5"></rect>
              <rect x="14" y="12" width="7" height="9"></rect>
              <rect x="3" y="16" width="7" height="5"></rect>
            </svg>
            <span class="text-sm font-medium">Dashboard</span>
          </router-link>
        </li>

        <li class="mt-2">
          <button
            @click="toggleManagementDropdown"
            class="flex items-center justify-between w-full gap-3 p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5] transition-colors"
            :class="{ 'text-[#111418] bg-[#f0f2f5]': isManagementDropdownOpen }"
          >
            <div class="flex items-center gap-3">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
              >
                <path d="M3 12h18"></path>
                <path d="M3 6h18"></path>
                <path d="M3 18h18"></path>
              </svg>
              <span class="text-sm font-medium">Management</span>
            </div>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-4 w-4 transition-transform"
              :class="{ 'rotate-90': isManagementDropdownOpen }"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 5l7 7-7 7"
              />
            </svg>
          </button>
          <ul v-if="isManagementDropdownOpen" class="pl-8 mt-1 space-y-1">
            <li>
              <router-link
                to="/management-user"
                class="block p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5]"
                :class="{ 'bg-[#f0f2f5]': route.path === '/management-user' }"
                >Management User</router-link
              >
            </li>
            <li>
              <router-link
                to="/management-product"
                class="block p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5]"
                :class="{
                  'bg-[#f0f2f5]': route.path === '/management-product',
                }"
                >Management Product</router-link
              >
            </li>
            <li>
              <router-link
                to="/management-category"
                class="block p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5]"
                :class="{
                  'bg-[#f0f2f5]': route.path === '/management-category',
                }"
                >Management Category</router-link
              >
            </li>
            <li>
              <router-link
                to="/management-supplier"
                class="block p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5]"
                :class="{
                  'bg-[#f0f2f5]': route.path === '/management-supplier',
                }"
                >Management Supplier</router-link
              >
            </li>
            <li>
              <router-link
                to="/management-unit"
                class="block p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5]"
                :class="{ 'bg-[#f0f2f5]': route.path === '/management-unit' }"
                >Management Unit</router-link
              >
            </li>
            <li>
              <router-link
                to="/management-voucher"
                class="block p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5]"
                :class="{
                  'bg-[#f0f2f5]': route.path === '/management-voucher',
                }"
                >Management Voucher</router-link
              >
            </li>
          </ul>
        </li>

        <li class="mt-2">
          <button
            @click="toggleReportDropdown"
            class="flex items-center justify-between w-full gap-3 p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5] transition-colors"
            :class="{ 'text-[#111418] bg-[#f0f2f5]': isReportDropdownOpen }"
          >
            <div class="flex items-center gap-3">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-5 w-5"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-width="2"
                stroke-linecap="round"
                stroke-linejoin="round"
              >
                <rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect>
                <line x1="16" y1="2" x2="16" y2="6"></line>
                <line x1="8" y1="2" x2="8" y2="6"></line>
                <line x1="3" y1="10" x2="21" y2="10"></line>
              </svg>
              <span class="text-sm font-medium">Laporan</span>
            </div>
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-4 w-4 transition-transform"
              :class="{ 'rotate-90': isReportDropdownOpen }"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 5l7 7-7 7"
              />
            </svg>
          </button>

          <ul v-if="isReportDropdownOpen" class="pl-8 mt-1 space-y-1">
            <li>
              <router-link
                to="/report-sales-summary"
                class="block p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5]"
                :class="{
                  'bg-[#f0f2f5]': route.path === '/report-sales-summary',
                }"
                >Sales Summary</router-link
              >
            </li>
            <li>
              <router-link
                to="/report-purchase-summary"
                class="block p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5]"
                :class="{
                  'bg-[#f0f2f5]': route.path === '/report-purchase-summary',
                }"
                >Purchase Summary</router-link
              >
            </li>
            <li>
              <router-link
                to="/report-sale-return-summary"
                class="block p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5]"
                :class="{
                  'bg-[#f0f2f5]': route.path === '/report-sale-return-summary',
                }"
                >Sale Return Summary</router-link
              >
            </li>
            <li>
              <router-link
                to="/report-purchase-return-summary"
                class="block p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5]"
                :class="{
                  'bg-[#f0f2f5]':
                    route.path === '/report-purchase-return-summary',
                }"
                >Purchase Return Summary</router-link
              >
            </li>
            <li>
              <router-link
                to="/report-stock-opname-summary"
                class="block p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5]"
                :class="{
                  'bg-[#f0f2f5]': route.path === '/report-stock-opname-summary',
                }"
                >Stock Opname Summary</router-link
              >
            </li>
            <li>
              <router-link
                to="/report-stock-snapshot"
                class="block p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5]"
                :class="{
                  'bg-[#f0f2f5]': route.path === '/report-stock-snapshot',
                }"
                >Stock Snapshot</router-link
              >
            </li>
            <li>
              <router-link
                to="/report-inventory-value"
                class="block p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5]"
                :class="{
                  'bg-[#f0f2f5]': route.path === '/report-inventory-value',
                }"
                >Inventory Value</router-link
              >
            </li>
          </ul>
        </li>
      </ul>
    </nav>

    <div class="flex-grow flex items-center justify-center p-4">
      <img
        src="/images/owner.png"
        alt="Warehouse Management Illustration"
        class="w-full h-auto max-w-[80%] opacity-80"
      />
    </div>

    <div class="mt-auto pt-4 border-t border-[#dbe0e6]">
      <button
        @click="handleLogout"
        class="flex items-center w-full gap-3 p-2 rounded-lg text-[#d9534f] hover:bg-[#f0f2f5]"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-5 w-5"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"></path>
          <polyline points="16 17 21 12 16 7"></polyline>
          <line x1="21" y1="12" x2="9" y2="12"></line>
        </svg>
        <span class="text-sm font-medium">Logout</span>
      </button>
    </div>
  </aside>
</template>

<script setup>
import { ref, watch } from "vue";
import { useRoute, useRouter } from "vue-router";

const route = useRoute();
const router = useRouter();

// 1. STATE LOKAL UNTUK MANAGEMENT (Disesuaikan)
const isManagementDropdownOpen = ref(route.path.startsWith("/management-"));

// 2. STATE LOKAL BARU UNTUK LAPORAN
const isReportDropdownOpen = ref(route.path.startsWith("/report-"));

// 3. WATCHER: Mengawasi perubahan rute
watch(
  () => route.path,
  (newPath) => {
    // Logic untuk Management: buka jika rute Management
    if (newPath.startsWith("/management-")) {
      isManagementDropdownOpen.value = true;
    } else if (!newPath.startsWith("/report-")) {
      // Tutup Management jika pindah ke Dashboard atau Logout (opsional)
      // isManagementDropdownOpen.value = false;
    }

    // Logic BARU untuk Laporan: buka jika rute Laporan
    if (newPath.startsWith("/report-")) {
      isReportDropdownOpen.value = true;
    } else if (!newPath.startsWith("/management-")) {
      // Tutup Laporan jika pindah ke Dashboard atau Logout (opsional)
      // isReportDropdownOpen.value = false;
    }
  }
);

// FUNGSI TOGGLE MANAGEMENT (Disesuaikan untuk menutup Laporan)
const toggleManagementDropdown = () => {
  isManagementDropdownOpen.value = !isManagementDropdownOpen.value;
  // Pastikan menu Laporan tertutup saat Management dibuka/ditutup
  if (isManagementDropdownOpen.value) {
    isReportDropdownOpen.value = false;
  }
};

// FUNGSI TOGGLE BARU UNTUK LAPORAN
const toggleReportDropdown = () => {
  isReportDropdownOpen.value = !isReportDropdownOpen.value;
  // Pastikan menu Management tertutup saat Laporan dibuka/ditutup
  if (isReportDropdownOpen.value) {
    isManagementDropdownOpen.value = false;
  }
};

const handleLogout = () => {
  localStorage.removeItem("token");
  localStorage.removeItem("userRole");
  router.push("/");
};
</script>
