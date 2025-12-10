<template>
  <aside
    class="w-64 bg-white border-r border-gray-200 p-4 fixed h-full top-0 left-0 overflow-y-auto"
  >
    <div class="flex items-center gap-4 pb-4 mb-4 border-b border-gray-200">
      <div class="size-6">
        <svg
          viewBox="0 0 48 48"
          fill="currentColor"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path d="M44 4H30.6666V17.3334H17.3334V30.6666H4V44H44V4Z" />
        </svg>
      </div>
      <h2 class="text-xl font-bold leading-tight">CashierApp</h2>
    </div>
    <nav>
      <ul>
        <li>
          <router-link
            to="/dashboard-cashier"
            class="flex items-center gap-3 p-2 rounded-lg text-gray-700 hover:bg-gray-100"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24px"
              height="24px"
              fill="currentColor"
              viewBox="0 0 256 256"
            >
              <path
                d="M228,110.16,140.1,38.74a16,16,0,0,0-20.2,0L28,110.16A16,16,0,0,0,20,124.08V208a16,16,0,0,0,16,16H220a16,16,0,0,0,16-16V124.08A16,16,0,0,0,228,110.16ZM128,48.06,204,112,128,176,52,112Zm80,168H48V129.58L128,65.58l80,64ZM128,160a32,32,0,1,0-32-32A32,32,0,0,0,128,160Zm0-48a16,16,0,1,1-16,16A16,16,0,0,1,128,112Z"
              />
            </svg>
            Dashboard
          </router-link>
        </li>
        <li class="mt-2">
          <router-link
            to="/cashier"
            class="flex items-center gap-3 p-2 rounded-lg text-gray-700 hover:bg-gray-100"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24px"
              height="24px"
              fill="currentColor"
              viewBox="0 0 256 256"
            >
              <path
                d="M216,40H40A16,16,0,0,0,24,56V200a16,16,0,0,0,16,16H216a16,16,0,0,0,16-16V56A16,16,0,0,0,216,40ZM40,56H216v24H40Zm176,144H40V96H216Z"
              />
            </svg>
            Payments
          </router-link>
        </li>
        <li class="mt-2">
          <router-link
            to="/return"
            class="flex items-center gap-3 p-2 rounded-lg text-gray-700 hover:bg-gray-100"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24px"
              height="24px"
              fill="currentColor"
              viewBox="0 0 256 256"
            >
              <path
                d="M224,128a96,96,0,0,1-94.71,96H128A95.38,95.38,0,0,1,62.1,197.8a8,8,0,0,1,11-11.63A80,80,0,1,0,71.43,71.39a3.07,3.07,0,0,1-.26.25L44.59,96H72a8,8,0,0,1,0,16H24a8,8,0,0,1-8-8V56a8,8,0,0,1,16,0V85.8L60.25,60A96,96,0,0,1,224,128Z"
              />
            </svg>
            Pengembalian
          </router-link>
        </li>
      </ul>
    </nav>

    <div class="flex-grow flex items-center justify-center p-4">
      <img
        src="/images/cashier.png"
        alt="Kepala Gudang Illustration"
        class="w-full h-auto max-w-[80%] opacity-80"
      />
    </div>

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
