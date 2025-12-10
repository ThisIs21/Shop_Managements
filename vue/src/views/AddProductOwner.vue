<template>
  <OwnerLayouts>
    <div class="flex flex-1 flex-col py-5 p-8">
      <div class="layout-content-container flex flex-col flex-1">
        <form @submit.prevent="createProduct" class="flex flex-col flex-1">
          <div class="flex flex-wrap justify-between gap-3 p-4">
            <div class="flex flex-col gap-3">
              <p
                class="text-[#111418] tracking-light text-[32px] font-bold leading-tight"
              >
                Add Product
              </p>
            </div>
          </div>

          <div class="flex flex-col gap-4 p-4">
            <div>
              <label>Product Name</label>
              <input
                type="text"
                v-model="product.name"
                class="w-full border p-2 rounded"
                required
              />
            </div>

            <div>
              <label>Purchase Price</label>
              <input
                type="number"
                v-model.number="product.cost_price"
                class="w-full border p-2 rounded"
                required
              />
            </div>

            <div>
              <label>Selling Price</label>
              <input
                type="number"
                v-model.number="product.sell_price"
                class="w-full border p-2 rounded"
                required
              />
            </div>

            <div>
              <label>Stock</label>
              <input
                type="number"
                v-model.number="product.stock"
                class="w-full border p-2 rounded"
              />
            </div>

            <div>
              <label>Category</label>
              <select
                v-model.number="product.category_id"
                class="w-full border p-2 rounded"
                required
              >
                <option disabled value="">Select category</option>
                <option v-for="cat in categories" :key="cat.id" :value="cat.id">
                  {{ cat.name }}
                </option>
              </select>
            </div>

            <div>
              <label>Unit</label>
              <select
                v-model.number="product.unit_id"
                class="w-full border p-2 rounded"
                required
              >
                <option disabled value="">Select unit</option>
                <option v-for="unit in units" :key="unit.id" :value="unit.id">
                  {{ unit.name }}
                </option>
              </select>
            </div>

            <div>
              <label>Supplier</label>
              <select
                v-model.number="product.supplier_id"
                class="w-full border p-2 rounded"
                required
              >
                <option disabled value="">Select supplier</option>
                <option v-for="sup in suppliers" :key="sup.id" :value="sup.id">
                  {{ sup.name }}
                </option>
              </select>
            </div>
          </div>

          <div class="flex justify-end p-4 gap-3">
            <router-link to="/management-product">
              <button type="button" class="bg-gray-300 px-4 py-2 rounded">
                Cancel
              </button>
            </router-link>
            <button
              type="submit"
              class="bg-blue-500 text-white px-4 py-2 rounded"
            >
              Save
            </button>
          </div>
        </form>
      </div>
    </div>
  </OwnerLayouts>
</template>

<script setup>
import { ref, onMounted } from "vue";
import OwnerLayouts from "../components/OwnerLayouts.vue";
import { useRouter } from "vue-router";
import api from "/src/services/api.js";

const router = useRouter();

const product = ref({
  name: "",
  cost_price: null,
  sell_price: null,
  stock: 0,
  category_id: null,
  unit_id: null,
  supplier_id: null,
});

const categories = ref([]);
const units = ref([]);
const suppliers = ref([]);

const fetchData = async () => {
  try {
    const [catRes, unitRes, supRes] = await Promise.all([
      api.get("/owner/categories"),
      api.get("/owner/units"),
      api.get("/owner/suppliers"),
    ]);

    categories.value = catRes.data.data.map((c) => ({
      id: c.id,
      name: c.name,
    }));
    units.value = unitRes.data.data.map((u) => ({
      id: u.id,
      name: u.name,
    }));
    suppliers.value = supRes.data.data.map((s) => ({
      id: s.id,
      name: s.name,
    }));
  } catch (err) {
    console.error("Failed to fetch dropdown data:", err);
  }
};

const createProduct = async () => {
  try {
    await api.post("/owner/products", product.value);
    alert("Product added");
    router.push("/management-product");
  } catch (err) {
    console.error("Failed to add product:", err);
    alert("Failed to add product");
  }
};

onMounted(() => fetchData());
</script>

<style scoped>
/* Optional styling */
</style>
