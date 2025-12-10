<template>
  <aside
    class="w-64 bg-white border-r border-[#dbe0e6] p-4 fixed h-full top-0 left-0 overflow-y-auto flex flex-col"
  >
    <!-- Header -->
    <div class="flex items-center gap-4 pb-4 mb-4 border-b border-[#dbe0e6]">
      <div class="size-6 text-[#0d80f2]">
        <svg viewBox="0 0 48 48" fill="none" xmlns="http://www.w3.org/2000/svg">
          <path
            fill-rule="evenodd"
            clip-rule="evenodd"
            d="M12.0799 24L4 19.2479L9.95537 8.75216L18.04 13.4961L18.0446 4H29.9554L29.96 13.4961L38.0446 8.75216L44 19.2479L35.92 24L44 28.7521L38.0446 39.2479L29.96 34.5039L29.9554 44H18.0446L18.04 34.5039L9.95537 39.2479L4 28.7521L12.0799 24Z"
            fill="currentColor"
          ></path>
        </svg>
      </div>
      <h2 class="text-xl font-bold leading-tight">Warehouse Manager</h2>
    </div>

    <!-- Navigasi -->
    <nav class="flex flex-col gap-2">
      <router-link
        to="/dashboard-kepalagudang"
        class="flex items-center gap-3 px-4 py-2 rounded-lg text-sm font-medium text-gray-700 hover:bg-gray-100 hover:text-gray-900 transition-colors"
        :class="{
          'bg-blue-100 text-blue-600 font-semibold hover:bg-blue-200':
            $route.path === '/dashboard-kepalagudang',
        }"
      >
        Dashboard
      </router-link>

      <router-link
        to="/purchaseorder-kepalagudang"
        class="flex items-center gap-3 px-4 py-2 rounded-lg text-sm font-medium text-gray-700 hover:bg-gray-100 hover:text-gray-900 transition-colors"
        :class="{
          'bg-blue-100 text-blue-600 font-semibold hover:bg-blue-200':
            $route.path === '/purchaseorder-kepalagudang',
        }"
      >
        Purchase Order Pending Approval
      </router-link>

      <router-link
        to="/sales-return"
        class="flex items-center gap-3 px-4 py-2 rounded-lg text-sm font-medium text-gray-700 hover:bg-gray-100 hover:text-gray-900 transition-colors"
        :class="{
          'bg-blue-100 text-blue-600 font-semibold hover:bg-blue-200':
            $route.path === '/sales-return',
        }"
      >
        Sales Return Requests
      </router-link>

      <router-link
        to="/purchasereturn-kepalagudang"
        class="flex items-center gap-3 px-4 py-2 rounded-lg text-sm font-medium text-gray-700 hover:bg-gray-100 hover:text-gray-900 transition-colors"
        :class="{
          'bg-blue-100 text-blue-600 font-semibold hover:bg-blue-200':
            $route.path === '/purchasereturn-kepalagudang',
        }"
      >
        Purchase Return Requests
      </router-link>

      <router-link
        to="/product-catalog"
        class="flex items-center gap-3 px-4 py-2 rounded-lg text-sm font-medium text-gray-700 hover:bg-gray-100 hover:text-gray-900 transition-colors"
        :class="{
          'bg-blue-100 text-blue-600 font-semibold hover:bg-blue-200':
            $route.path === '/product-catalog',
        }"
      >
        Product Catalog
      </router-link>
    </nav>

    <div class="flex-grow flex items-center justify-center p-4">
      <img
        src="/images/kepalagudang.png"
        alt="Kepala Gudang Illustration"
        class="w-full h-auto max-w-[80%] opacity-80"
      />
    </div>

    <!-- Tombol Logout -->
    <div class="mt-auto pt-4 border-t border-[#dbe0e6]">
      <button
        @click="handleLogout"
        class="flex items-center w-full gap-3 p-2 rounded-lg text-[#d9534f] hover:bg-[#f0f2f5] transition-colors"
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
import { useRouter } from "vue-router";

const router = useRouter();

const handleLogout = () => {
  localStorage.removeItem("token");
  localStorage.removeItem("userRole");
  router.push("/");
};
</script>
