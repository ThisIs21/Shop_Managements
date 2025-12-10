<template>
  <aside
    class="w-64 bg-white border-r border-[#dbe0e6] p-4 fixed h-full top-0 left-0 overflow-y-auto flex flex-col"
  >
    <div>
      <div class="flex items-center gap-4 pb-4 mb-4 border-b border-[#dbe0e6]">
        <div class="size-6 text-[#0d80f2]">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="currentColor"
            viewBox="0 0 256 256"
          >
            <path
              d="M116,216H48a8,8,0,0,1-8-8V88H216v24a8,8,0,0,1,16,0V88a16,16,0,0,0-16-16H208V48a16,16,0,0,0-16-16H64A16,16,0,0,0,48,48V72H40a16,16,0,0,0-16,16V208a16,16,0,0,0,16,16h76a8,8,0,0,0,0-16ZM64,48H192V72H64ZM224,144v64a8,8,0,0,1-16,0V144a8,8,0,0,1,16,0Zm-48,0v64a8,8,0,0,1-16,0V144a8,8,0,0,1,16,0Zm-48,0v64a8,8,0,0,1-16,0V144a8,8,0,0,1,16,0Z"
            ></path>
          </svg>
        </div>
        <h2 class="text-xl font-bold leading-tight">Inventory</h2>
      </div>

      <nav class="flex w-full flex-col gap-1">
        <router-link
          to="/dashboard-gudang"
          class="flex items-center gap-3 rounded-lg px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-100 hover:text-gray-900 transition-colors"
          :class="{
            'bg-blue-100 text-blue-600 font-semibold hover:bg-blue-200':
              $route.path === '/dashboard-gudang',
          }"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            fill="currentColor"
            viewBox="0 0 256 256"
          >
            <path
              d="M224,208H176a16,16,0,0,1-16-16V128a16,16,0,0,1,16-16h48a16,16,0,0,1,16,16v64A16,16,0,0,1,224,208ZM176,128v64h48V128ZM112,208H64a16,16,0,0,1-16-16V48A16,16,0,0,1,64,32h48a16,16,0,0,1,16,16V192A16,16,0,0,1,112,208ZM64,48V192h48V48ZM40,208H16a8,8,0,0,0,0,16H240a8,8,0,0,0,0-16H216V192a8,8,0,0,0-16,0v16H152V128a8,8,0,0,0-16,0v80H96V48a8,8,0,0,0-16,0v160H40a8,8,0,0,0,0,16Z"
            ></path>
          </svg>
          Dashboard
        </router-link>

        <router-link
          to="/gudang/goods-receipt"
          class="flex items-center gap-3 rounded-lg px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-100 hover:text-gray-900 transition-colors"
          :class="{
            'bg-blue-100 text-blue-600 font-semibold hover:bg-blue-200':
              $route.path === '/gudang/goods-receipt',
          }"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            fill="currentColor"
            viewBox="0 0 256 256"
          >
            <path
              d="M224,128V48a8,8,0,0,0-8-8H40a8,8,0,0,0-8,8v80a8,8,0,0,0,16,0V56H208v72a8,8,0,0,0,16,0ZM96,224a8,8,0,0,1-8-8V176a8,8,0,0,1,16,0v40A8,8,0,0,1,96,224Zm64,0a8,8,0,0,1-8-8V176a8,8,0,0,1,16,0v40A8,8,0,0,1,160,224Zm40-56H56V144a8,8,0,0,1-16,0V208a16,16,0,0,0,16,16H200a16,16,0,0,0,16-16V144a8,8,0,0,1-16,0Z"
            ></path>
          </svg>
          Penerimaan Barang
        </router-link>

        <router-link
          to="/productsgudang"
          class="flex items-center gap-3 rounded-lg px-4 py-2 text-sm font-medium text-gray-700 hover:bg-gray-100 hover:text-gray-900 transition-colors"
          :class="{
            'bg-blue-100 text-blue-600 font-semibold hover:bg-blue-200':
              $route.path === '/productsgudang',
          }"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            fill="currentColor"
            viewBox="0 0 256 256"
          >
            <path
              d="M232,80.79V192a16,16,0,0,1-16,16H40a16,16,0,0,1-16-16V80.79A16.09,16.09,0,0,1,32.22,68.7L119.5,23.32a16,16,0,0,1,17,0l87.28,45.38A16.09,16.09,0,0,1,232,80.79ZM216,88.42,136,46.12V120H216ZM120,46.12,40,88.42V120h80V46.12ZM40,136H216v56.79l-87.28-45.38a16,16,0,0,0-17,0L40,192.79Z"
            ></path>
          </svg>
          Products
        </router-link>
        <router-link
          to="/purchase-returns"
          class="flex items-center gap-3 p-2 rounded-lg text-[#60758a] hover:bg-[#f0f2f5] transition-colors"
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
            <path
              d="M12 2v20M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"
            ></path>
          </svg>
          <span class="text-sm font-medium">Purchase Returns</span>
        </router-link>
      </nav>
    </div>

    <div class="flex-grow flex items-center justify-center p-4">
      <img
        src="/images/gudang.png"
        alt="Warehouse Management Illustration"
        class="w-full h-auto max-w-[80%] opacity-80"
      />
    </div>

    <div class="pt-4 border-t border-[#dbe0e6]">
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
  // Hapus token dan userRole dari localStorage
  localStorage.removeItem("token");
  localStorage.removeItem("userRole");

  // Arahkan pengguna kembali ke halaman login
  router.push("/");
};
</script>

<style scoped>
/* Opsional: Sesuaikan gaya sesuai kebutuhan */
</style>
