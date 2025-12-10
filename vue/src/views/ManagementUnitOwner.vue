<template>
  <OwnerLayouts>
    <div class="flex flex-1 justify-center py-5">
      <div class="layout-content-container flex flex-col max-w-[960px] flex-1">
        <div class="flex flex-wrap justify-between gap-3 p-4">
          <p
            class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
          >
            Units
          </p>
          <button
            @click="openCreateModal"
            class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-medium leading-normal"
          >
            <span class="truncate">Add Unit</span>
          </button>
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
                    ID
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal"
                  >
                    Name
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
                    colspan="3"
                    class="h-[72px] px-4 py-2 text-center text-[#60758a] text-sm font-normal leading-normal"
                  >
                    Loading...
                  </td>
                </tr>
                <tr
                  v-else-if="units.length === 0"
                  class="border-t border-t-[#dbe0e6]"
                >
                  <td
                    colspan="3"
                    class="h-[72px] px-4 py-2 text-center text-[#60758a] text-sm font-normal leading-normal"
                  >
                    No units found.
                  </td>
                </tr>
                <tr
                  v-for="unit in units"
                  :key="unit.id"
                  class="border-t border-t-[#dbe0e6]"
                >
                  <td
                    class="h-[72px] px-4 py-2 text-[#111418] text-sm font-normal leading-normal"
                  >
                    {{ unit.id }}
                  </td>
                  <td
                    class="h-[72px] px-4 py-2 text-[#111418] text-sm font-normal leading-normal"
                  >
                    {{ unit.name }}
                  </td>
                  <td
                    class="h-[72px] px-4 py-2 text-[#60758a] text-sm font-bold leading-normal tracking-[0.015em]"
                  >
                    <span
                      class="text-blue-500 hover:underline cursor-pointer"
                      @click="openEditModal(unit)"
                      >Edit</span
                    >
                    /
                    <span
                      class="text-red-500 hover:underline cursor-pointer"
                      @click="confirmDelete(unit.id)"
                      >Delete</span
                    >
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <div
          v-if="showModal"
          class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
        >
          <div class="bg-white rounded-lg p-6 w-full max-w-md">
            <h2 class="text-xl font-bold mb-4">
              {{ isEditMode ? "Edit Unit" : "Create Unit" }}
            </h2>
            <form @submit.prevent="submitForm">
              <div class="mb-4">
                <label class="block text-sm font-medium text-[#111418] mb-1"
                  >Unit Name</label
                >
                <input
                  v-model="formData.name"
                  type="text"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  placeholder="Enter unit name"
                  required
                />
              </div>
              <div class="flex justify-end gap-3">
                <button
                  type="button"
                  @click="closeModal"
                  class="px-4 py-2 bg-gray-200 text-[#111418] rounded-lg text-sm font-medium"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  class="px-4 py-2 bg-[#f0f2f5] text-[#111418] rounded-lg text-sm font-medium"
                >
                  {{ isEditMode ? "Update" : "Create" }}
                </button>
              </div>
            </form>
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

const units = ref([]);
const loading = ref(true);
const showModal = ref(false);
const isEditMode = ref(false);
const formData = ref({ name: "", id: null });

const fetchUnits = async () => {
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("No JWT token found in localStorage.");
      loading.value = false;
      return;
    }

    const response = await axios.get(
      "http://localhost:8081/api/v1/owner/units",
      {
        headers: { Authorization: `Bearer ${token}` }, // Corrected syntax
      }
    );

    if (response.data && response.data.data) {
      units.value = response.data.data;
    } else {
      console.error('Invalid API response format. Expected "data" field.');
      units.value = [];
    }
  } catch (error) {
    console.error(
      "Error fetching units:",
      error.response ? error.response.data : error
    );
  } finally {
    loading.value = false;
  }
};

const openCreateModal = () => {
  isEditMode.value = false;
  formData.value = { name: "", id: null };
  showModal.value = true;
};

const openEditModal = (unit) => {
  isEditMode.value = true;
  formData.value = { name: unit.name, id: unit.id };
  showModal.value = true;
};

const closeModal = () => {
  showModal.value = false;
  formData.value = { name: "", id: null };
};

const submitForm = async () => {
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("No JWT token found in localStorage.");
      return;
    }

    const url = isEditMode.value
      ? `http://localhost:8081/api/v1/owner/units/${formData.value.id}` // Corrected syntax
      : "http://localhost:8081/api/v1/owner/units";
    const method = isEditMode.value ? "put" : "post";
    const payload = { name: formData.value.name };

    const response = await axios({
      method,
      url,
      data: payload,
      headers: { Authorization: `Bearer ${token}` }, // Corrected syntax
    });

    if (response.status === 200 || response.status === 201) {
      fetchUnits();
      closeModal();
    }
  } catch (error) {
    console.error(
      "Error submitting form:",
      error.response ? error.response.data : error
    );
  }
};

const confirmDelete = async (id) => {
  if (!confirm("Are you sure you want to delete this unit?")) return;

  try {
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("No JWT token found in localStorage.");
      return;
    }

    await axios.delete(`http://localhost:8081/api/v1/owner/units/${id}`, {
      // Corrected syntax
      headers: { Authorization: `Bearer ${token}` }, // Corrected syntax
    });

    fetchUnits();
  } catch (error) {
    console.error(
      "Error deleting unit:",
      error.response ? error.response.data : error
    );
  }
};

onMounted(() => {
  fetchUnits();
});
</script>

<style scoped>
/* Scoped styles consistent with the provided user management page */
</style>
