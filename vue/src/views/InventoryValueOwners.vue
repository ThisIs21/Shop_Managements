<template>
  <OwnerLayouts>
    <div class="layout-content-container flex flex-col max-w-[960px] flex-1">
      <!-- HEADER -->
      <div class="flex flex-wrap justify-between gap-3 p-4">
        <div class="flex min-w-72 flex-col gap-3">
          <p
            class="text-[#111418] tracking-light text-[32px] font-bold leading-tight"
          >
            Inventory Value Report
          </p>
          <p class="text-[#60758a] text-sm font-normal leading-normal">
            View the total value of your inventory based on purchase price
          </p>
        </div>
        <div class="flex items-center">
          <router-link to="/report-owners">
            <button
              class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-bold leading-normal tracking-[0.015em]"
            >
              <span class="truncate">Kembali ke Report Owners</span>
            </button>
          </router-link>
        </div>
      </div>

      <!-- ALERT ERROR -->
      <div
        v-if="errorMessage"
        class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative mx-4"
        role="alert"
      >
        <span class="block sm:inline">{{ errorMessage }}</span>
      </div>

      <!-- FILTER -->
      <h3
        class="text-[#111418] text-lg font-bold leading-tight tracking-[-0.015em] px-4 pb-2 pt-4"
      >
        Filters
      </h3>
      <div class="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
        <label class="flex flex-col min-w-40 flex-1">
          <p class="text-[#111418] text-base font-medium leading-normal pb-2">
            Category
          </p>
          <select
            v-model="selectedCategory"
            @change="filterByCategory"
            class="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dbe0e6] bg-white focus:border-[#dbe0e6] h-14 placeholder:text-[#60758a] p-[15px] text-base font-normal leading-normal"
          >
            <option value="">All Categories</option>
            <option v-for="cat in categories" :key="cat" :value="cat">
              {{ cat }}
            </option>
          </select>
        </label>
      </div>

      <!-- EXPORT BUTTONS -->
      <div class="flex justify-stretch">
        <div class="flex flex-1 gap-3 flex-wrap px-4 py-3 justify-end">
          <button
            @click="exportData('csv')"
            class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-bold leading-normal tracking-[0.015em]"
            :disabled="loading"
          >
            <span class="truncate">Export to CSV</span>
          </button>
          <button
            @click="exportData('pdf')"
            class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-bold leading-normal tracking-[0.015em]"
            :disabled="loading"
          >
            <span class="truncate">Export to PDF</span>
          </button>
        </div>
      </div>

      <!-- TABLE -->
      <div class="px-4 py-3 @container">
        <div
          class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
        >
          <table class="flex-1">
            <thead>
              <tr class="bg-white">
                <th
                  class="px-4 py-3 text-left text-[#111418] text-sm font-medium"
                >
                  Product Name
                </th>
                <th
                  class="px-4 py-3 text-left text-[#111418] text-sm font-medium"
                >
                  Product Code
                </th>
                <th
                  class="px-4 py-3 text-left text-[#111418] text-sm font-medium"
                >
                  Category
                </th>
                <th
                  class="px-4 py-3 text-left text-[#111418] text-sm font-medium"
                >
                  Unit
                </th>
                <th
                  class="px-4 py-3 text-left text-[#111418] text-sm font-medium"
                >
                  Stock Quantity
                </th>
                <th
                  class="px-4 py-3 text-left text-[#111418] text-sm font-medium"
                >
                  Purchase Price
                </th>
                <th
                  class="px-4 py-3 text-left text-[#111418] text-sm font-medium"
                >
                  Total Stock Value
                </th>
              </tr>
            </thead>
            <tbody>
              <!-- LOADING SPINNER -->
              <tr v-if="loading">
                <td colspan="7" class="text-center py-6">
                  <div class="flex justify-center items-center space-x-2">
                    <svg
                      class="animate-spin h-6 w-6 text-gray-600"
                      xmlns="http://www.w3.org/2000/svg"
                      fill="none"
                      viewBox="0 0 24 24"
                    >
                      <circle
                        class="opacity-25"
                        cx="12"
                        cy="12"
                        r="10"
                        stroke="currentColor"
                        stroke-width="4"
                      ></circle>
                      <path
                        class="opacity-75"
                        fill="currentColor"
                        d="M4 12a8 8 0 018-8v8H4z"
                      ></path>
                    </svg>
                    <span>Loading data...</span>
                  </div>
                </td>
              </tr>

              <!-- DATA -->
              <tr
                v-for="item in filteredInventory"
                :key="item.product_code"
                class="border-t border-t-[#dbe0e6]"
                v-else
              >
                <td class="px-4 py-2 text-[#111418] text-sm">
                  {{ item.product_name }}
                </td>
                <td class="px-4 py-2 text-[#60758a] text-sm">
                  {{ item.product_code }}
                </td>
                <td class="px-4 py-2 text-[#60758a] text-sm">
                  {{ item.category }}
                </td>
                <td class="px-4 py-2 text-[#60758a] text-sm">
                  {{ item.unit }}
                </td>
                <td class="px-4 py-2 text-[#60758a] text-sm">
                  {{ item.stock_quantity }}
                </td>
                <td class="px-4 py-2 text-[#60758a] text-sm">
                  ${{ item.purchase_price.toFixed(2) }}
                </td>
                <td class="px-4 py-2 text-[#60758a] text-sm">
                  ${{ item.total_value.toFixed(2) }}
                </td>
              </tr>

              <!-- EMPTY -->
              <tr v-if="!loading && filteredInventory.length === 0">
                <td colspan="7" class="text-center py-4 text-[#60758a]">
                  No data available
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </OwnerLayouts>
</template>

<script>
import OwnerLayouts from "@/components/OwnerLayouts.vue";
import axios from "axios";

export default {
  name: "InventoryValueOwners",
  components: { OwnerLayouts },
  data() {
    return {
      inventory: [],
      filteredInventory: [],
      categories: [],
      selectedCategory: "",
      loading: false,
      errorMessage: "",
    };
  },
  methods: {
    async fetchInventory() {
      this.loading = true;
      this.errorMessage = "";
      try {
        const token = localStorage.getItem("token");
        if (!token) {
          this.errorMessage =
            "Token autentikasi tidak ditemukan. Silakan login kembali.";
          return;
        }

        const response = await axios.get(
          "http://localhost:8081/api/v1/owner/reports/inventory-detail",
          {
            headers: { Authorization: `Bearer ${token}` },
          }
        );

        this.inventory = response.data.data || [];
        this.filteredInventory = [...this.inventory];
        this.categories = [
          ...new Set(this.inventory.map((item) => item.category)),
        ];
      } catch (err) {
        console.error("Failed to load inventory:", err);
        this.errorMessage =
          "Gagal mengambil data inventaris. Periksa koneksi atau coba lagi.";
      } finally {
        this.loading = false;
      }
    },
    filterByCategory() {
      if (!this.selectedCategory) {
        this.filteredInventory = [...this.inventory];
      } else {
        this.filteredInventory = this.inventory.filter(
          (item) => item.category === this.selectedCategory
        );
      }
    },
    async exportData(type) {
      if (this.loading) return;

      this.loading = true;
      this.errorMessage = "";
      try {
        const token = localStorage.getItem("token");
        if (!token) {
          this.errorMessage =
            "Token autentikasi tidak ditemukan. Silakan login kembali.";
          return;
        }

        const response = await axios.get(
          `http://localhost:8081/api/v1/owner/reports/inventory-detail?export=${type}`,
          {
            headers: { Authorization: `Bearer ${token}` },
            responseType: type === "pdf" ? "blob" : "text",
          }
        );

        const url = window.URL.createObjectURL(
          new Blob([response.data], {
            type: type === "pdf" ? "application/pdf" : "text/csv",
          })
        );
        const link = document.createElement("a");
        link.href = url;
        link.setAttribute("download", `inventory_report.${type}`);
        document.body.appendChild(link);
        link.click();
        document.body.removeChild(link);
        window.URL.revokeObjectURL(url);
      } catch (err) {
        console.error("Export failed:", err);
        this.errorMessage =
          "Gagal mengekspor data. Periksa koneksi atau coba lagi.";
      } finally {
        this.loading = false;
      }
    },
  },
  mounted() {
    this.fetchInventory();
  },
};
</script>
