<template>
  <OwnerLayouts>
    <div class="flex flex-1 justify-center py-5">
      <div class="layout-content-container flex flex-col max-w-[960px] flex-1">
        <!-- Header -->
        <div class="flex flex-wrap justify-between gap-3 p-4">
          <p
            class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
          >
            Suppliers
          </p>
          <button
            @click="openCreateModal"
            class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-medium leading-normal"
          >
            <span class="truncate">Add Supplier</span>
          </button>
        </div>

        <!-- Table -->
        <div class="px-4 py-3">
          <div
            class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
          >
            <table class="flex-1">
              <thead>
                <tr class="bg-white">
                  <th class="px-4 py-3 text-left text-sm font-medium">ID</th>
                  <th class="px-4 py-3 text-left text-sm font-medium">Name</th>
                  <th class="px-4 py-3 text-left text-sm font-medium">
                    Contact
                  </th>
                  <th class="px-4 py-3 text-left text-sm font-medium">
                    Address
                  </th>
                  <th class="px-4 py-3 text-left text-sm font-medium">
                    Actions
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr v-if="loading" class="border-t border-t-[#dbe0e6]">
                  <td colspan="5" class="text-center py-4">Loading...</td>
                </tr>
                <tr
                  v-else-if="suppliers.length === 0"
                  class="border-t border-t-[#dbe0e6]"
                >
                  <td colspan="5" class="text-center py-4">
                    No suppliers found.
                  </td>
                </tr>
                <tr
                  v-for="supplier in suppliers"
                  :key="supplier.id"
                  class="border-t border-t-[#dbe0e6]"
                >
                  <td class="px-4 py-2">{{ supplier.id }}</td>
                  <td class="px-4 py-2">{{ supplier.name }}</td>
                  <td class="px-4 py-2">{{ supplier.contact }}</td>
                  <td class="px-4 py-2">{{ supplier.address }}</td>
                  <td class="px-4 py-2">
                    <span
                      class="text-blue-500 hover:underline cursor-pointer"
                      @click="openEditModal(supplier)"
                      >Edit</span
                    >
                    /
                    <span
                      class="text-red-500 hover:underline cursor-pointer"
                      @click="confirmDelete(supplier.id)"
                      >Delete</span
                    >
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
          <!-- Pagination Controls -->
          <div class="flex items-center justify-between mt-3">
            <div class="text-sm text-slate-600">
              Showing page {{ page }} / {{ totalPages }}
            </div>
            <div class="flex items-center gap-2">
              <button
                @click="prevPage"
                class="px-3 py-1 border rounded bg-white"
              >
                Prev
              </button>
              <select v-model="perPage" class="border rounded p-1 text-sm">
                <option :value="5">5</option>
                <option :value="10">10</option>
                <option :value="25">25</option>
              </select>
              <button
                @click="nextPage"
                class="px-3 py-1 border rounded bg-white"
              >
                Next
              </button>
            </div>
          </div>
        </div>

        <!-- Modal -->
        <div
          v-if="showModal"
          class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
        >
          <div class="bg-white rounded-lg p-6 w-full max-w-md">
            <h2 class="text-xl font-bold mb-4">
              {{ isEditMode ? "Edit Supplier" : "Create Supplier" }}
            </h2>
            <form @submit.prevent="submitForm">
              <div class="mb-3">
                <label class="block text-sm font-medium mb-1">Name</label>
                <input
                  v-model="formData.name"
                  type="text"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  placeholder="Enter supplier name"
                  required
                />
              </div>
              <div class="mb-3">
                <label class="block text-sm font-medium mb-1">Phone</label>
                <input
                  v-model="formData.contact"
                  type="text"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  placeholder="Enter contact number"
                  required
                />
              </div>
              <div class="mb-3">
                <label class="block text-sm font-medium mb-1">Address</label>
                <textarea
                  v-model="formData.address"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  placeholder="Enter supplier address"
                  required
                ></textarea>
              </div>
              <div class="flex justify-end gap-3">
                <button
                  type="button"
                  @click="closeModal"
                  class="px-4 py-2 bg-gray-200 rounded-lg text-sm font-medium"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  class="px-4 py-2 bg-[#f0f2f5] rounded-lg text-sm font-medium"
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
import { ref, onMounted, computed, watch } from "vue";
import axios from "axios";
import OwnerLayouts from "../components/OwnerLayouts.vue";

const suppliers = ref([]);
const loading = ref(true);
const showModal = ref(false);
const isEditMode = ref(false);
const formData = ref({ id: null, name: "", contact: "", address: "" });
// Pagination
const suppliersAll = ref([]);
const page = ref(1);
const perPage = ref(10);
const totalPages = computed(() =>
  Math.max(1, Math.ceil((suppliersAll.value || []).length / perPage.value))
);

const updatePagination = () => {
  const start = (page.value - 1) * perPage.value;
  suppliers.value = suppliersAll.value.slice(start, start + perPage.value);
};

const fetchSuppliers = async () => {
  try {
    const token = localStorage.getItem("token");
    const res = await axios.get(
      "http://localhost:8081/api/v1/owner/suppliers",
      {
        headers: { Authorization: `Bearer ${token}` },
      }
    );
    suppliersAll.value = res.data?.data || [];
    page.value = 1;
    updatePagination();
  } catch (err) {
    console.error("Error fetching suppliers:", err.response?.data || err);
  } finally {
    loading.value = false;
  }
};

const openCreateModal = () => {
  isEditMode.value = false;
  formData.value = { id: null, name: "", contact: "", address: "" };
  showModal.value = true;
};

const openEditModal = (supplier) => {
  isEditMode.value = true;
  formData.value = { ...supplier };
  showModal.value = true;
};

const closeModal = () => {
  showModal.value = false;
};

const submitForm = async () => {
  try {
    const token = localStorage.getItem("token");
    const url = isEditMode.value
      ? `http://localhost:8081/api/v1/owner/suppliers/${formData.value.id}`
      : "http://localhost:8081/api/v1/owner/suppliers";
    const method = isEditMode.value ? "put" : "post";
    const payload = {
      name: formData.value.name,
      phone: formData.value.contact,
      address: formData.value.address,
    };
    await axios({
      method,
      url,
      data: payload,
      headers: { Authorization: `Bearer ${token}` },
    });
    fetchSuppliers();
    closeModal();
  } catch (err) {
    console.error("Error submitting supplier:", err.response?.data || err);
  }
};

const confirmDelete = async (id) => {
  if (!confirm("Are you sure you want to delete this supplier?")) return;
  try {
    const token = localStorage.getItem("token");
    await axios.delete(`http://localhost:8081/api/v1/owner/suppliers/${id}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    fetchSuppliers();
  } catch (err) {
    console.error("Error deleting supplier:", err.response?.data || err);
  }
};

onMounted(() => {
  fetchSuppliers();
});

watch([page, perPage], () => updatePagination());

const prevPage = () => {
  if (page.value > 1) page.value--;
};
const nextPage = () => {
  if (page.value < totalPages.value) page.value++;
};
</script>
