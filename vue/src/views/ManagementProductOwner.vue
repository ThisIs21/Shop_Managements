<template>
  <OwnerLayouts>
    <div class="flex flex-1 flex-col py-5 p-8">
      <div class="layout-content-container flex flex-col max-w-[960px] flex-1">
        <div class="flex flex-wrap justify-between gap-3 p-4">
          <p
            class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
          >
            Product List
          </p>
          <router-link
            to="/add-product"
            class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-medium leading-normal"
          >
            <span class="truncate">Add Product</span>
          </router-link>
        </div>

        <div class="px-4 py-3 @container">
          <div
            class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
          >
            <table class="flex-1">
              <thead>
                <tr class="bg-white">
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal w-[200px]"
                  >
                    Name
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal w-[120px]"
                  >
                    Category
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal w-[120px]"
                  >
                    Unit
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal w-[120px]"
                  >
                    Supplier
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal w-[120px]"
                  >
                    Cost Price
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal w-[120px]"
                  >
                    Sell Price
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal w-[80px]"
                  >
                    Stock
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#60758a] text-sm font-medium leading-normal w-[160px]"
                  >
                    Actions
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr v-if="loading">
                  <td
                    colspan="8"
                    class="h-[72px] px-4 py-2 text-center text-[#60758a] text-sm font-normal leading-normal"
                  >
                    Loading...
                  </td>
                </tr>
                <tr v-else-if="products.length === 0">
                  <td
                    colspan="8"
                    class="h-[72px] px-4 py-2 text-center text-[#60758a] text-sm font-normal leading-normal"
                  >
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
                    class="px-4 py-2 text-[#111418] text-sm font-normal leading-normal"
                  >
                    {{ product.name }}
                  </td>
                  <td
                    class="px-4 py-2 text-[#60758a] text-sm font-normal leading-normal"
                  >
                    {{ product.category?.name || "-" }}
                  </td>
                  <td
                    class="px-4 py-2 text-[#60758a] text-sm font-normal leading-normal"
                  >
                    {{ product.unit?.name || "-" }}
                  </td>
                  <td
                    class="px-4 py-2 text-[#60758a] text-sm font-normal leading-normal"
                  >
                    {{ product.supplier?.name || "-" }}
                  </td>
                  <td
                    class="px-4 py-2 text-[#60758a] text-sm font-normal leading-normal"
                  >
                    {{ formatCurrency(product.cost_price) }}
                  </td>
                  <td
                    class="px-4 py-2 text-[#60758a] text-sm font-normal leading-normal"
                  >
                    {{ formatCurrency(product.sell_price) }}
                  </td>
                  <td
                    class="px-4 py-2 text-[#60758a] text-sm font-normal leading-normal"
                  >
                    {{ product.stock }}
                  </td>
                  <td
                    class="px-4 py-2 text-[#60758a] text-sm font-bold leading-normal tracking-[0.015em]"
                  >
                    <span
                      class="text-blue-500 hover:underline cursor-pointer"
                      @click="openEditModal(product)"
                      >Edit</span
                    >
                    /
                    <span
                      class="text-red-500 hover:underline cursor-pointer"
                      @click="confirmDelete(product.id)"
                      >Delete</span
                    >
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- Modal untuk Edit -->
        <div
          v-if="showModal"
          class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
        >
          <div class="bg-white rounded-lg p-6 w-full max-w-md">
            <h2 class="text-xl font-bold mb-4">Edit Product</h2>
            <form @submit.prevent="submitForm">
              <div class="mb-4">
                <label class="block text-sm font-medium text-[#111418] mb-1"
                  >Name</label
                >
                <input
                  v-model="formData.name"
                  type="text"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  placeholder="Enter product name"
                  required
                />
              </div>
              <div class="mb-4">
                <label class="block text-sm font-medium text-[#111418] mb-1"
                  >Category</label
                >
                <select
                  v-model="formData.category_id"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  required
                >
                  <option value="" disabled>Select a category</option>
                  <option
                    v-for="category in categories"
                    :key="category.id"
                    :value="category.id"
                  >
                    {{ category.name }}
                  </option>
                </select>
              </div>
              <div class="mb-4">
                <label class="block text-sm font-medium text-[#111418] mb-1"
                  >Unit</label
                >
                <select
                  v-model="formData.unit_id"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  required
                >
                  <option value="" disabled>Select a unit</option>
                  <option v-for="unit in units" :key="unit.id" :value="unit.id">
                    {{ unit.name }}
                  </option>
                </select>
              </div>
              <div class="mb-4">
                <label class="block text-sm font-medium text-[#111418] mb-1"
                  >Supplier</label
                >
                <select
                  v-model="formData.supplier_id"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  required
                >
                  <option value="" disabled>Select a supplier</option>
                  <option
                    v-for="supplier in suppliers"
                    :key="supplier.id"
                    :value="supplier.id"
                  >
                    {{ supplier.name }}
                  </option>
                </select>
              </div>
              <div class="mb-4">
                <label class="block text-sm font-medium text-[#111418] mb-1"
                  >Cost Price</label
                >
                <input
                  v-model.number="formData.cost_price"
                  type="number"
                  step="0.01"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  placeholder="Enter cost price"
                  required
                />
              </div>
              <div class="mb-4">
                <label class="block text-sm font-medium text-[#111418] mb-1"
                  >Sell Price</label
                >
                <input
                  v-model.number="formData.sell_price"
                  type="number"
                  step="0.01"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  placeholder="Enter sell price"
                  required
                />
              </div>
              <div class="mb-4">
                <label class="block text-sm font-medium text-[#111418] mb-1"
                  >Stock</label
                >
                <input
                  v-model.number="formData.stock"
                  type="number"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  placeholder="Enter stock"
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
                  Update
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

const products = ref([]);
const categories = ref([]);
const units = ref([]);
const suppliers = ref([]);
const loading = ref(true);
const showModal = ref(false);
const formData = ref({
  id: null,
  name: "",
  category_id: null,
  unit_id: null,
  supplier_id: null,
  cost_price: 0,
  sell_price: 0,
  stock: 0,
});

const fetchProducts = async () => {
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      alert("Harap login dulu!");
      loading.value = false;
      return;
    }
    const apiUrl = "http://localhost:8081/api/v1/owner/products";
    const response = await axios.get(apiUrl, {
      headers: { Authorization: `Bearer ${token}` },
    });
    console.log("Products API Response:", response.data);
    products.value = response.data.data || [];
  } catch (error) {
    console.error(
      "Error fetching products:",
      error.response?.data || error.message
    );
    alert(
      `Gagal ambil data produk: ${
        error.response?.data?.message || error.message
      }`
    );
  } finally {
    loading.value = false;
  }
};

const fetchCategories = async () => {
  try {
    const token = localStorage.getItem("token");
    if (!token) return;
    const apiUrl = "http://localhost:8081/api/v1/owner/categories";
    const response = await axios.get(apiUrl, {
      headers: { Authorization: `Bearer ${token}` },
    });
    console.log("Categories API Response:", response.data);
    categories.value = response.data.data || [];
  } catch (error) {
    console.error(
      "Error fetching categories:",
      error.response?.data || error.message
    );
  }
};

const fetchUnits = async () => {
  try {
    const token = localStorage.getItem("token");
    if (!token) return;
    const apiUrl = "http://localhost:8081/api/v1/owner/units";
    const response = await axios.get(apiUrl, {
      headers: { Authorization: `Bearer ${token}` },
    });
    console.log("Units API Response:", response.data);
    units.value = response.data.data || [];
  } catch (error) {
    console.error(
      "Error fetching units:",
      error.response?.data || error.message
    );
  }
};

const fetchSuppliers = async () => {
  try {
    const token = localStorage.getItem("token");
    if (!token) return;
    const apiUrl = "http://localhost:8081/api/v1/owner/suppliers";
    const response = await axios.get(apiUrl, {
      headers: { Authorization: `Bearer ${token}` },
    });
    console.log("Suppliers API Response:", response.data);
    suppliers.value = response.data.data || [];
  } catch (error) {
    console.error(
      "Error fetching suppliers:",
      error.response?.data || error.message
    );
  }
};

const formatCurrency = (value) => {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
  }).format(value);
};

const openEditModal = (product) => {
  formData.value = {
    id: product.id,
    name: product.name,
    category_id: Number(product.category_id) || null,
    unit_id: Number(product.unit_id) || null,
    supplier_id: Number(product.supplier_id) || null,
    cost_price: product.cost_price,
    sell_price: product.sell_price,
    stock: product.stock,
  };
  console.log("Form data saat buka modal:", formData.value);
  showModal.value = true;
};

const closeModal = () => {
  showModal.value = false;
  formData.value = {
    id: null,
    name: "",
    category_id: null,
    unit_id: null,
    supplier_id: null,
    cost_price: 0,
    sell_price: 0,
    stock: 0,
  };
};

const submitForm = async () => {
  console.log("Payload sebelum kirim:", formData.value);
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      alert("Harap login dulu!");
      return;
    }
    if (
      !formData.value.name ||
      !formData.value.category_id ||
      !formData.value.unit_id ||
      !formData.value.supplier_id
    ) {
      alert("Isi semua kolom wajib!");
      return;
    }
    if (formData.value.cost_price <= 0 || formData.value.sell_price <= 0) {
      alert("Harga harus lebih dari 0!");
      return;
    }
    const url = `http://localhost:8081/api/v1/owner/products/${formData.value.id}`;
    const payload = {
      name: formData.value.name,
      category_id: Number(formData.value.category_id),
      unit_id: Number(formData.value.unit_id),
      supplier_id: Number(formData.value.supplier_id),
      cost_price: Number(formData.value.cost_price),
      sell_price: Number(formData.value.sell_price),
      stock: Number(formData.value.stock),
    };
    console.log("Mengirim payload:", payload);
    const response = await axios.put(url, payload, {
      headers: { Authorization: `Bearer ${token}` },
    });
    console.log("Response dari server:", response.data);
    if (response.data.success) {
      await fetchProducts();
      closeModal();
      alert("Produk berhasil diupdate!");
    }
  } catch (error) {
    console.error(
      "Error updating product:",
      error.response?.data || error.message
    );
    alert(
      `Gagal update produk: ${error.response?.data?.message || error.message}`
    );
  }
};

const confirmDelete = async (id) => {
  if (!confirm("Yakin hapus produk ini?")) return;
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      alert("Harap login dulu!");
      return;
    }
    const response = await axios.delete(
      `http://localhost:8081/api/v1/owner/products/${id}`,
      { headers: { Authorization: `Bearer ${token}` } }
    );
    console.log("Delete response:", response.data);
    if (response.data.success) {
      await fetchProducts();
      alert("Produk berhasil dihapus!");
    }
  } catch (error) {
    console.error(
      "Error deleting product:",
      error.response?.data || error.message
    );
    alert(
      `Gagal hapus produk: ${error.response?.data?.message || error.message}`
    );
  }
};

onMounted(() => {
  fetchProducts();
  fetchCategories();
  fetchUnits();
  fetchSuppliers();
});
</script>

<style scoped>
/* Gaya CSS tetap sama */
</style>
