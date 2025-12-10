<template>
  <OwnerLayouts>
    <div class="flex flex-1 flex-col py-5 p-8">
      <div class="layout-content-container flex flex-col flex-1">
        <h1 class="text-2xl font-bold mb-4 text-white">Stock Snapshot</h1>

        <div class="px-4 py-3">
          <div
            class="flex overflow-hidden rounded-lg border border-[#333] bg-[#1f1f1f]"
          >
            <table class="flex-1 text-white">
              <thead>
                <tr>
                  <th class="px-4 py-3 text-left">Product Name</th>
                  <th class="px-4 py-3 text-left">Product Code</th>
                  <th class="px-4 py-3 text-left">Category</th>
                  <th class="px-4 py-3 text-left">Purchase Price</th>
                  <th class="px-4 py-3 text-left">Selling Price</th>
                  <th class="px-4 py-3 text-left">Stock Quantity</th>
                  <th class="px-4 py-3 text-left">Location</th>
                </tr>
              </thead>
              <tbody>
                <tr v-if="loading">
                  <td colspan="7" class="text-center py-4">Loading...</td>
                </tr>
                <tr
                  v-for="(item, index) in tableData"
                  :key="index"
                  class="border-t border-[#333]"
                >
                  <td class="px-4 py-2">{{ item.name }}</td>
                  <td class="px-4 py-2">-</td>
                  <td class="px-4 py-2">{{ item.category?.name || "-" }}</td>
                  <td class="px-4 py-2">
                    Rp {{ formatRupiah(item.cost_price) }}
                  </td>
                  <td class="px-4 py-2">
                    Rp {{ formatRupiah(item.sell_price) }}
                  </td>
                  <td class="px-4 py-2">{{ item.stock }}</td>
                  <td class="px-4 py-2">{{ item.supplier?.address || "-" }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </OwnerLayouts>
</template>

<script setup>
import OwnerLayouts from "../components/OwnerLayouts.vue";
import { ref, onMounted } from "vue";
import axios from "axios";

const api = axios.create({ baseURL: "http://localhost:8081/api/v1" });
const savedToken = localStorage.getItem("token");

if (savedToken) {
  api.defaults.headers.common["Authorization"] = `Bearer ${savedToken}`;
}

const tableData = ref([]);
const loading = ref(false);

const fetchStock = async () => {
  loading.value = true;
  try {
    const res = await api.get("/owner/products"); // Pastikan endpoint ini benar
    tableData.value = res.data.data || [];
  } catch (err) {
    console.error("Failed to fetch stock", err);
  }
  loading.value = false;
};

const formatRupiah = (value) => new Intl.NumberFormat("id-ID").format(value);

onMounted(fetchStock);
</script>
