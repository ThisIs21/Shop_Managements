<template>
  <div class="flex size-full min-h-screen">
    <AppSidebar />
    <div
      class="flex h-full grow flex-col bg-white p-6"
      style="font-family: Manrope, 'Noto Sans', sans-serif"
    >
      <div
        class="layout-content-container flex flex-col max-w-[960px] flex-1 px-4 py-5"
      >
        <div class="flex flex-wrap justify-between gap-3 p-4">
          <p
            class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
          >
            Products
          </p>
          <button
            class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-medium leading-normal"
          >
            <span class="truncate">Add Product</span>
          </button>
        </div>
        <div class="px-4 py-3">
          <label class="flex flex-col min-w-40 h-12 w-full">
            <div class="flex w-full flex-1 items-stretch rounded-lg h-full">
              <div
                class="text-[#60758a] flex border-none bg-[#f0f2f5] items-center justify-center pl-4 rounded-l-lg border-r-0"
                data-icon="MagnifyingGlass"
                data-size="24px"
                data-weight="regular"
              >
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="24px"
                  height="24px"
                  fill="currentColor"
                  viewBox="0 0 256 256"
                >
                  <path
                    d="M229.66,218.34l-50.07-50.06a88.11,88.11,0,1,0-11.31,11.31l50.06,50.07a8,8,0,0,0,11.32-11.32ZM40,112a72,72,0,1,1,72,72A72.08,72.08,0,0,1,40,112Z"
                  ></path>
                </svg>
              </div>
              <input
                placeholder="Search products by name or code"
                class="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border-none bg-[#f0f2f5] focus:border-none h-full placeholder:text-[#60758a] px-4 rounded-l-none border-l-0 pl-2 text-base font-normal leading-normal"
                value=""
              />
            </div>
          </label>
        </div>
        <div class="px-4 py-3 @container">
          <div
            class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
          >
            <table class="flex-1">
              <thead>
                <tr class="bg-white">
                  <th
                    class="table-75f3fd5d-fb3d-4782-ab8e-d54b0e1a052d-column-120 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal"
                  >
                    Product Name
                  </th>
                  <th
                    class="table-75f3fd5d-fb3d-4782-ab8e-d54b0e1a052d-column-240 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal"
                  >
                    Selling Price
                  </th>
                  <th
                    class="table-75f3fd5d-fb3d-4782-ab8e-d54b0e1a052d-column-360 px-4 py-3 text-left text-[#111418] w-60 text-sm font-medium leading-normal"
                  >
                    Stock Availability
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr v-if="loading">
                  <td colspan="3" class="text-center py-4">
                    Loading products...
                  </td>
                </tr>
                <tr v-else-if="products.length === 0">
                  <td colspan="3" class="text-center py-4">
                    No products found.
                  </td>
                </tr>
                <tr
                  v-else
                  v-for="product in products"
                  :key="product.id"
                  class="border-t border-t-[#dbe0e6]"
                >
                  <td
                    class="table-75f3fd5d-fb3d-4782-ab8e-d54b0e1a052d-column-120 h-[72px] px-4 py-2 w-[400px] text-[#111418] text-sm font-normal leading-normal"
                  >
                    {{ product.name }}
                  </td>
                  <td
                    class="table-75f3fd5d-fb3d-4782-ab8e-d54b0e1a052d-column-240 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    ${{ product.selling_price }}
                  </td>
                  <td
                    class="table-75f3fd5d-fb3d-4782-ab8e-d54b0e1a052d-column-360 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal"
                  >
                    <button
                      class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 text-sm font-medium leading-normal w-full"
                      :class="{
                        'bg-[#4CAF50] text-white': product.stock > 0,
                        'bg-[#F44336] text-white': product.stock <= 0,
                      }"
                    >
                      <span class="truncate">{{
                        product.stock > 0 ? "In Stock" : "Out of Stock"
                      }}</span>
                    </button>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import AppSidebar from "../components/CashierSidebar.vue";
import axios from "axios";

// Definisikan state reaktif
const products = ref([]);
const loading = ref(true);

// Fungsi untuk mengambil data produk
const getProducts = async () => {
  try {
    const response = await axios.get("http://localhost:8081/api/v1/products", {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
    });
    products.value = response.data;
  } catch (error) {
    console.error("Failed to fetch products:", error);
    // Tambahkan logika penanganan error, misalnya menampilkan pesan ke pengguna
  } finally {
    loading.value = false;
  }
};

// Panggil fungsi saat komponen dimuat
onMounted(() => {
  getProducts();
});
</script>

<style>
@container (max-width: 120px) {
  .table-75f3fd5d-fb3d-4782-ab8e-d54b0e1a052d-column-120 {
    display: none;
  }
}
@container (max-width: 240px) {
  .table-75f3fd5d-fb3d-4782-ab8e-d54b0e1a052d-column-240 {
    display: none;
  }
}
@container (max-width: 360px) {
  .table-75f3fd5d-fb3d-4782-ab8e-d54b0e1a052d-column-360 {
    display: none;
  }
}
</style>
