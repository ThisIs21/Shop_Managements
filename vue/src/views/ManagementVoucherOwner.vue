<template>
  <OwnerLayouts>
    <div class="flex flex-1 justify-center py-5">
      <div class="layout-content-container flex flex-col max-w-[960px] flex-1">
        <!-- Title -->
        <div class="flex flex-wrap justify-between gap-3 p-4">
          <p
            class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
          >
            Management Voucher
          </p>
          <button
            @click="openCreateModal"
            class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-medium leading-normal"
          >
            <span class="truncate">Add Voucher</span>
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
                  <th class="px-4 py-3 text-left text-sm font-medium">Code</th>
                  <th class="px-4 py-3 text-left text-sm font-medium">Type</th>
                  <th class="px-4 py-3 text-left text-sm font-medium">Value</th>
                  <th class="px-4 py-3 text-left text-sm font-medium">
                    Active
                  </th>
                  <th class="px-4 py-3 text-left text-sm font-medium">
                    Actions
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr v-if="loading" class="border-t">
                  <td
                    colspan="6"
                    class="h-[72px] text-center text-sm text-[#60758a]"
                  >
                    Loading...
                  </td>
                </tr>
                <tr v-else-if="vouchers.length === 0" class="border-t">
                  <td
                    colspan="6"
                    class="h-[72px] text-center text-sm text-[#60758a]"
                  >
                    No vouchers found.
                  </td>
                </tr>
                <tr
                  v-for="voucher in vouchers"
                  :key="voucher.ID"
                  class="border-t"
                >
                  <td class="px-4 py-2 text-sm">{{ voucher.ID }}</td>
                  <td class="px-4 py-2 text-sm">{{ voucher.Code }}</td>
                  <td class="px-4 py-2 text-sm">{{ voucher.Type }}</td>
                  <td class="px-4 py-2 text-sm">{{ voucher.Value }}</td>
                  <td class="px-4 py-2 text-sm">
                    <span
                      :class="
                        voucher.Active ? 'text-green-600' : 'text-red-600'
                      "
                    >
                      {{ voucher.Active ? "Active" : "Inactive" }}
                    </span>
                  </td>
                  <td class="px-4 py-2 text-sm font-bold text-[#60758a]">
                    <span
                      class="text-blue-500 hover:underline cursor-pointer"
                      @click="openEditModal(voucher)"
                      >Edit</span
                    >
                    /
                    <span
                      class="text-red-500 hover:underline cursor-pointer"
                      @click="confirmDelete(voucher.ID)"
                      >Delete</span
                    >
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- Modal -->
        <div
          v-if="showModal"
          class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
        >
          <div class="bg-white rounded-lg p-6 w-full max-w-md">
            <h2 class="text-xl font-bold mb-4">
              {{ isEditMode ? "Edit Voucher" : "Create Voucher" }}
            </h2>
            <form @submit.prevent="submitForm">
              <!-- Code -->
              <div class="mb-4">
                <label class="block text-sm font-medium mb-1">Code</label>
                <input
                  v-model="formData.code"
                  type="text"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  placeholder="Enter voucher code"
                  required
                />
              </div>

              <!-- Type -->
              <div class="mb-4">
                <label class="block text-sm font-medium mb-1">Type</label>
                <select
                  v-model="formData.type"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  required
                >
                  <option value="PERCENT">PERCENT</option>
                  <option value="AMOUNT">AMOUNT</option>
                </select>
              </div>

              <!-- Value -->
              <div class="mb-4">
                <label class="block text-sm font-medium mb-1">Value</label>
                <input
                  v-model.number="formData.value"
                  type="number"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  placeholder="Enter voucher value"
                  required
                />
              </div>

              <!-- Active -->
              <div class="mb-4">
                <label class="block text-sm font-medium mb-1">Active</label>
                <select
                  v-model="formData.active"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                >
                  <option :value="true">Active</option>
                  <option :value="false">Inactive</option>
                </select>
              </div>

              <!-- Buttons -->
              <div class="flex justify-end gap-3">
                <button
                  type="button"
                  @click="closeModal"
                  class="px-4 py-2 bg-gray-200 text-sm rounded-lg"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  class="px-4 py-2 bg-blue-500 text-white text-sm rounded-lg"
                >
                  {{ isEditMode ? "Update" : "Create" }}
                </button>
              </div>
            </form>
          </div>
        </div>
        <!-- Pagination Controls -->
        <div class="flex items-center justify-between mt-3 px-4">
          <div class="text-sm text-slate-600">
            Showing page {{ page }} / {{ totalPages }}
          </div>
          <div class="flex items-center gap-2">
            <button
              @click="
                () => {
                  if (page > 1) page--;
                }
              "
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
              @click="
                () => {
                  if (page < totalPages) page++;
                }
              "
              class="px-3 py-1 border rounded bg-white"
            >
              Next
            </button>
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

const vouchers = ref([]);
const loading = ref(true);
const showModal = ref(false);
const isEditMode = ref(false);
const formData = ref({
  id: null,
  code: "",
  type: "PERCENT",
  value: 0,
  active: true,
});
// Pagination
const vouchersAll = ref([]);
const page = ref(1);
const perPage = ref(10);
const totalPages = computed(() =>
  Math.max(1, Math.ceil((vouchersAll.value || []).length / perPage.value))
);

const updatePagination = () => {
  const start = (page.value - 1) * perPage.value;
  vouchers.value = vouchersAll.value.slice(start, start + perPage.value);
};

const fetchVouchers = async () => {
  try {
    const token = localStorage.getItem("token");
    const res = await axios.get("http://localhost:8081/api/v1/vouchers", {
      headers: { Authorization: `Bearer ${token}` },
    });
    vouchersAll.value = res.data.data || [];
    page.value = 1;
    updatePagination();
  } catch (err) {
    console.error("Error fetching vouchers:", err);
  } finally {
    loading.value = false;
  }
};

const openCreateModal = () => {
  isEditMode.value = false;
  formData.value = {
    id: null,
    code: "",
    type: "PERCENT",
    value: 0,
    active: true,
  };
  showModal.value = true;
};

const openEditModal = (voucher) => {
  isEditMode.value = true;
  formData.value = {
    id: voucher.ID,
    code: voucher.Code,
    type: voucher.Type,
    value: voucher.Value,
    active: voucher.Active,
  };
  showModal.value = true;
};

const closeModal = () => {
  showModal.value = false;
};

const submitForm = async () => {
  try {
    const token = localStorage.getItem("token");
    const url = isEditMode.value
      ? `http://localhost:8081/api/v1/owner/vouchers/${formData.value.id}`
      : "http://localhost:8081/api/v1/owner/vouchers";
    const method = isEditMode.value ? "put" : "post";
    const payload = {
      code: formData.value.code,
      type: formData.value.type,
      value: formData.value.value,
      active: formData.value.active,
    };
    await axios({
      method,
      url,
      data: payload,
      headers: { Authorization: `Bearer ${token}` },
    });
    fetchVouchers();
    closeModal();
  } catch (err) {
    console.error("Error saving voucher:", err);
  }
};

const confirmDelete = async (id) => {
  if (!confirm("Are you sure want to delete this voucher?")) return;
  try {
    const token = localStorage.getItem("token");
    await axios.delete(`http://localhost:8081/api/v1/owner/vouchers/${id}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    fetchVouchers();
  } catch (err) {
    console.error("Error deleting voucher:", err);
  }
};

onMounted(() => {
  fetchVouchers();
});

watch([page, perPage], () => updatePagination());
</script>
