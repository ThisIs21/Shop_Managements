<template>
  <OwnerLayouts>
    <div class="p-8">
      <div class="flex flex-wrap items-center justify-between gap-3 p-4">
        <p
          class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
        >
          {{ isEdit ? "Edit Category" : "Add Category" }}
        </p>
        <router-link
          to="/list-category"
          class="flex items-center gap-2 text-[#60758a] text-base font-normal leading-normal"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="20px"
            height="20px"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path d="M19 12H5"></path>
            <path d="M12 19L5 12L12 5"></path>
          </svg>
          <span class="truncate">Kembali</span>
        </router-link>
      </div>

      <form
        @submit.prevent="handleSubmit"
        class="flex flex-col w-[512px] max-w-[512px] py-5"
      >
        <div class="flex flex-col flex-wrap items-end gap-4 px-4 py-3">
          <label class="flex flex-col min-w-40 flex-1 w-full">
            <p class="text-[#111418] text-base font-medium leading-normal pb-2">
              Name
            </p>
            <input
              v-model="category.name"
              placeholder="Enter category name"
              class="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dbe0e6] bg-white focus:border-[#dbe0e6] h-14 placeholder:text-[#60758a] p-[15px] text-base font-normal leading-normal"
            />
          </label>
        </div>
        <div class="flex justify-stretch">
          <div class="flex flex-1 gap-3 flex-wrap px-4 py-3 justify-end">
            <router-link
              to="/management-category"
              class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-bold leading-normal tracking-[0.015em]"
            >
              <span class="truncate">Cancel</span>
            </router-link>
            <button
              type="submit"
              class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 px-4 bg-[#0d80f2] text-white text-sm font-bold leading-normal tracking-[0.015em]"
            >
              <span class="truncate">Save</span>
            </button>
          </div>
        </div>
      </form>
    </div>
  </OwnerLayouts>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import { useRoute, useRouter } from "vue-router";
import axios from "axios";
import OwnerLayouts from "../components/OwnerLayouts.vue";

const category = ref({
  name: "",
});

const route = useRoute();
const router = useRouter();
const isEdit = computed(() => route.params.id !== undefined);

const API_URL = "http://192.168.100.111:8081/api/v1/categories";

const fetchCategory = async (id) => {
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("No JWT token found in localStorage.");
      return;
    }
    const response = await axios.get(`${API_URL}/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (response.data && response.data.data) {
      // API balikin field "Name"
      category.value.name = response.data.data.Name;
    }
  } catch (error) {
    console.error("Error fetching category:", error);
    alert("Failed to fetch category data.");
  }
};

const handleSubmit = async () => {
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("No JWT token found in localStorage.");
      return;
    }

    if (!category.value.name) {
      alert("Please fill in the category name.");
      return;
    }

    if (isEdit.value) {
      await axios.put(`${API_URL}/${route.params.id}`, category.value, {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });
      alert("Category updated successfully.");
    } else {
      await axios.post(API_URL, category.value, {
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });
      alert("Category created successfully.");
    }
    router.push("/management-category");
  } catch (error) {
    console.error("Error saving category:", error);
    alert("Failed to save category. Please check the console.");
  }
};

onMounted(() => {
  if (isEdit.value) {
    fetchCategory(route.params.id);
  }
});
</script>
