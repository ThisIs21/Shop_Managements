<template>
  <OwnerLayouts>
    <div class="flex flex-1 justify-center py-5">
      <div class="layout-content-container flex flex-col max-w-[960px] flex-1">
        <div class="flex flex-wrap justify-between gap-3 p-4">
          <p
            class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
          >
            Vouchers
          </p>
          <router-link
            to="/add-voucher"
            class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-medium leading-normal"
          >
            <span class="truncate">Add Voucher</span>
          </router-link>
        </div>
        <div class="px-4 py-3">
          <div
            class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
          >
            <table class="flex-1">
              <thead>
                <tr class="bg-white">
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal"
                  >
                    Code
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal"
                  >
                    Type
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal"
                  >
                    Value
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal"
                  >
                    Active
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal"
                  >
                    Actions
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr v-if="loading" class="border-t border-t-[#dbe0e6]">
                  <td
                    colspan="5"
                    class="h-[72px] px-4 py-2 text-center text-[#60758a] text-sm font-normal leading-normal"
                  >
                    Loading...
                  </td>
                </tr>
                <tr
                  v-else-if="vouchers.length === 0"
                  class="border-t border-t-[#dbe0e6]"
                >
                  <td
                    colspan="5"
                    class="h-[72px] px-4 py-2 text-center text-[#60758a] text-sm font-normal leading-normal"
                  >
                    No vouchers found.
                  </td>
                </tr>
                <tr
                  v-for="voucher in vouchers"
                  :key="voucher.ID"
                  class="border-t border-t-[#dbe0e6]"
                >
                  <td
                    class="h-[72px] px-4 py-2 text-[#111418] text-sm font-normal leading-normal"
                  >
                    {{ voucher.Code }}
                  </td>
                  <td
                    class="h-[72px] px-4 py-2 text-[#111418] text-sm font-normal leading-normal"
                  >
                    {{ voucher.Type }}
                  </td>
                  <td
                    class="h-[72px] px-4 py-2 text-[#111418] text-sm font-normal leading-normal"
                  >
                    {{ voucher.Value }}
                  </td>
                  <td
                    class="h-[72px] px-4 py-2 text-[#111418] text-sm font-normal leading-normal"
                  >
                    {{ voucher.Active }}
                  </td>
                  <td
                    class="h-[72px] px-4 py-2 text-[#60758a] text-sm font-bold leading-normal tracking-[0.015em]"
                  >
                    <router-link
                      :to="`/edit-voucher/${voucher.ID}`"
                      class="text-blue-500 hover:underline cursor-pointer"
                      >Edit</router-link
                    >
                    /
                    <span
                      @click="deleteVoucher(voucher.ID)"
                      class="text-red-500 hover:underline cursor-pointer"
                      >Delete</span
                    >
                  </td>
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
import { ref, onMounted } from "vue";
import axios from "axios";
import OwnerLayouts from "../components/OwnerLayouts.vue";

const vouchers = ref([]);
const loading = ref(true);

const API_URL = "http://192.168.100.111:8081/api/v1/vouchers";

const fetchVouchers = async () => {
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("No JWT token found in localStorage.");
      loading.value = false;
      return;
    }
    const response = await axios.get(API_URL, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (response.data && response.data.data) {
      vouchers.value = response.data.data;
    } else {
      vouchers.value = [];
    }
  } catch (error) {
    console.error("Error fetching vouchers:", error);
  } finally {
    loading.value = false;
  }
};

const deleteVoucher = async (id) => {
  if (confirm("Are you sure you want to delete this voucher?")) {
    try {
      const token = localStorage.getItem("token");
      if (!token) {
        console.error("No JWT token found in localStorage.");
        return;
      }
      await axios.delete(`${API_URL}/${id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      fetchVouchers();
    } catch (error) {
      console.error("Error deleting voucher:", error);
      alert("Failed to delete voucher.");
    }
  }
};

onMounted(() => {
  fetchVouchers();
});
</script>
