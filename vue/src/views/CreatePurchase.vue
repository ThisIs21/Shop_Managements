<template>
  <div class="flex">
    <!-- Sidebar -->
    <PurchasingSidebar />

    <!-- Konten utama -->
    <main class="flex-1 ml-64 p-8 bg-gray-50 min-h-screen">
      <!-- Header -->
      <div class="mb-8">
        <h1 class="text-[32px] font-bold text-[#111418] leading-tight">
          Create Purchase Order
        </h1>
        <p class="text-[#60758a] text-sm font-normal leading-normal mt-1">
          Create a new purchase order and manage your inventory efficiently
        </p>
      </div>

      <!-- Quick Stats -->
      <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-8">
        <div class="rounded-lg p-6 border border-[#dbe0e6] bg-white">
          <p class="text-[#60758a] text-sm font-medium">Total Items in Cart</p>
          <p class="text-[#111418] text-3xl font-bold leading-tight">
            {{ form.items.length }}
          </p>
          <p class="text-blue-600 text-xs font-semibold mt-2">
            {{ getTotalQty() }} units
          </p>
        </div>
        <div class="rounded-lg p-6 border border-[#dbe0e6] bg-white">
          <p class="text-[#60758a] text-sm font-medium">Total Cost</p>
          <p class="text-[#111418] text-3xl font-bold leading-tight">
            {{ formatCurrency(getTotalCost()) }}
          </p>
          <p
            class="text-green-600 text-xs font-semibold mt-2"
            v-if="getTotalCost() > 0"
          >
            Ready to submit
          </p>
          <p class="text-gray-400 text-xs font-semibold mt-2" v-else>
            Add items to calculate
          </p>
        </div>
        <div class="rounded-lg p-6 border border-[#dbe0e6] bg-white">
          <p class="text-[#60758a] text-sm font-medium">Suppliers Involved</p>
          <p class="text-[#111418] text-3xl font-bold leading-tight">
            {{ getUniqueSuppliersCount() }}
          </p>
          <p class="text-amber-600 text-xs font-semibold mt-2">
            Multiple orders
          </p>
        </div>
        <div class="rounded-lg p-6 border border-[#dbe0e6] bg-white">
          <p class="text-[#60758a] text-sm font-medium">Products Available</p>
          <p class="text-[#111418] text-3xl font-bold leading-tight">
            {{ products.length }}
          </p>
          <p class="text-purple-600 text-xs font-semibold mt-2">In inventory</p>
        </div>
      </div>

      <!-- Main Content Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <!-- Form Section - Left (2 cols) -->
        <div class="lg:col-span-2">
          <form @submit.prevent="submitPurchase">
            <!-- Tanggal -->
            <div class="mb-6 bg-white p-6 rounded-lg border border-[#dbe0e6]">
              <label class="block mb-2 text-sm font-semibold text-[#111418]"
                >Purchase Date</label
              >
              <input
                type="date"
                v-model="form.date"
                class="w-full border border-[#dbe0e6] rounded p-3 text-sm focus:outline-none focus:border-blue-500"
                required
              />
              <p class="text-[#60758a] text-xs mt-2">
                Select the date for this purchase order
              </p>
            </div>

            <!-- Produk -->
            <div class="mb-6 bg-white p-6 rounded-lg border border-[#dbe0e6]">
              <label class="block mb-3 text-sm font-semibold text-[#111418]"
                >Add Products</label
              >
              <div class="space-y-3">
                <div>
                  <select
                    v-model="selectedProduct"
                    class="w-full border border-[#dbe0e6] rounded p-3 text-sm focus:outline-none focus:border-blue-500"
                  >
                    <option value="">-- Select Product --</option>
                    <option v-for="p in products" :key="p.id" :value="p">
                      {{ p.name }} ({{ p.supplier?.name || "No Supplier" }}) -
                      Cost: {{ formatCurrency(p.cost_price) }}
                    </option>
                  </select>
                  <p class="text-[#60758a] text-xs mt-1">
                    Choose a product to add to your order
                  </p>
                </div>
                <div class="flex gap-2">
                  <input
                    type="number"
                    v-model.number="qty"
                    min="1"
                    placeholder="Quantity"
                    class="flex-1 border border-[#dbe0e6] rounded p-3 text-sm focus:outline-none focus:border-blue-500"
                  />
                  <button
                    type="button"
                    @click="addItem"
                    class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-3 rounded font-medium text-sm transition"
                    :disabled="!selectedProduct"
                  >
                    + Add
                  </button>
                </div>
              </div>
            </div>

            <!-- List Item -->
            <div
              v-if="form.items.length"
              class="bg-white p-6 rounded-lg border border-[#dbe0e6]"
            >
              <h2 class="font-semibold text-[#111418] mb-4 text-lg">
                Order Items ({{ form.items.length }})
              </h2>
              <div class="overflow-x-auto">
                <table class="w-full text-sm">
                  <thead>
                    <tr class="bg-gray-100 border-b border-[#dbe0e6]">
                      <th class="p-3 text-left text-[#111418] font-semibold">
                        Product
                      </th>
                      <th class="p-3 text-left text-[#111418] font-semibold">
                        Supplier
                      </th>
                      <th class="p-3 text-right text-[#111418] font-semibold">
                        Qty
                      </th>
                      <th class="p-3 text-right text-[#111418] font-semibold">
                        Price
                      </th>
                      <th class="p-3 text-right text-[#111418] font-semibold">
                        Subtotal
                      </th>
                      <th class="p-3 text-center text-[#111418] font-semibold">
                        Action
                      </th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr
                      v-for="(item, i) in form.items"
                      :key="i"
                      class="border-b border-[#dbe0e6] hover:bg-gray-50"
                    >
                      <td class="p-3 text-[#111418]">{{ item.name }}</td>
                      <td class="p-3 text-[#60758a]">
                        {{ item.supplier_name || "-" }}
                      </td>
                      <td class="p-3 text-right text-[#111418] font-medium">
                        {{ item.qty }}
                      </td>
                      <td class="p-3 text-right text-[#111418] font-medium">
                        {{ formatCurrency(item.price) }}
                      </td>
                      <td class="p-3 text-right text-[#111418] font-semibold">
                        {{ formatCurrency(item.qty * item.price) }}
                      </td>
                      <td class="p-3 text-center">
                        <button
                          type="button"
                          class="text-red-600 hover:text-red-800 hover:underline transition text-sm font-medium"
                          @click="removeItem(i)"
                        >
                          Remove
                        </button>
                      </td>
                    </tr>
                  </tbody>
                  <tfoot>
                    <tr class="bg-blue-50 border-t-2 border-[#dbe0e6]">
                      <td class="p-3 font-semibold text-[#111418]" colspan="4">
                        TOTAL
                      </td>
                      <td
                        class="p-3 text-right font-bold text-[#111418] text-lg"
                      >
                        {{ formatCurrency(getTotalCost()) }}
                      </td>
                      <td class="p-3"></td>
                    </tr>
                  </tfoot>
                </table>
              </div>
            </div>

            <!-- Submit Section -->
            <div class="bg-white p-6 rounded-lg border border-[#dbe0e6]">
              <h3 class="font-semibold text-[#111418] mb-4">
                Ready to Submit?
              </h3>
              <p
                v-if="!form.date"
                class="text-amber-600 text-sm mb-4 flex items-center gap-2"
              >
                <span>⚠️</span> Please select a purchase date
              </p>
              <p
                v-if="!form.items.length"
                class="text-amber-600 text-sm mb-4 flex items-center gap-2"
              >
                <span>⚠️</span> Add at least one product to continue
              </p>
              <p
                v-if="hasMissingSuppliers()"
                class="text-red-600 text-sm mb-4 flex items-center gap-2"
              >
                <span>❌</span> Some products don't have suppliers assigned
              </p>
              <p
                v-else-if="form.items.length"
                class="text-green-600 text-sm mb-4 flex items-center gap-2"
              >
                <span>✓</span> All items ready for submission
              </p>
              <button
                type="submit"
                class="w-full bg-green-600 hover:bg-green-700 text-white px-4 py-3 rounded font-semibold transition disabled:opacity-50 disabled:cursor-not-allowed"
                :disabled="
                  submitting ||
                  !form.items.length ||
                  !form.date ||
                  hasMissingSuppliers()
                "
              >
                <span v-if="submitting">{{ getSubmittingText() }}</span>
                <span v-else>{{
                  form.items.length ? "✓ Submit Order" : "Submit Order"
                }}</span>
              </button>
              <p class="text-[#60758a] text-xs mt-3 text-center">
                {{ form.items.length }}
                {{ form.items.length === 1 ? "item" : "items" }} will be sent to
                {{ getUniqueSuppliersCount() }}
                {{ getUniqueSuppliersCount() === 1 ? "supplier" : "suppliers" }}
              </p>
            </div>
          </form>
        </div>

        <!-- Sidebar - Right (1 col) -->
        <div class="space-y-6">
          <!-- Top Suppliers Shortcut -->
          <div class="bg-white p-6 rounded-lg border border-[#dbe0e6]">
            <h3 class="font-semibold text-[#111418] mb-4">Top Suppliers</h3>
            <div class="space-y-2">
              <div
                v-if="topSuppliers.length === 0"
                class="text-[#60758a] text-sm text-center py-4"
              >
                No supplier data
              </div>
              <button
                v-for="supplier in topSuppliers"
                :key="supplier.id"
                type="button"
                @click="filterBySupplier(supplier.id)"
                class="w-full text-left p-3 rounded border border-gray-200 hover:border-blue-500 hover:bg-blue-50 transition text-sm"
              >
                <p class="font-medium text-[#111418]">{{ supplier.name }}</p>
                <p class="text-[#60758a] text-xs">
                  {{ getSupplierProductCount(supplier.id) }} products
                </p>
              </button>
            </div>
          </div>

          <!-- Recent Categories -->
          <div class="bg-white p-6 rounded-lg border border-[#dbe0e6]">
            <h3 class="font-semibold text-[#111418] mb-4">Quick Filter</h3>
            <div class="space-y-2">
              <button
                type="button"
                @click="clearFilter()"
                class="w-full text-left p-3 rounded border border-gray-200 hover:border-purple-500 hover:bg-purple-50 transition text-sm font-medium text-[#111418]"
              >
                📦 All Products ({{ products.length }})
              </button>
              <button
                type="button"
                @click="filterLowStock()"
                class="w-full text-left p-3 rounded border border-gray-200 hover:border-red-500 hover:bg-red-50 transition text-sm"
              >
                <p class="font-medium text-[#111418]">⚠️ Low Stock Items</p>
                <p class="text-[#60758a] text-xs">Quick reorder</p>
              </button>
            </div>
          </div>

          <!-- Tips Section -->
          <div class="bg-blue-50 p-6 rounded-lg border border-blue-200">
            <h4
              class="font-semibold text-blue-900 mb-3 flex items-center gap-2"
            >
              <span>💡</span> Tips
            </h4>
            <ul class="text-blue-800 text-xs space-y-2">
              <li>✓ Group items by supplier for faster processing</li>
              <li>✓ Ensure all products have suppliers assigned</li>
              <li>✓ Review prices before submitting</li>
              <li>✓ Check stock levels regularly</li>
            </ul>
          </div>

          <!-- Summary Card -->
          <div
            class="bg-gradient-to-br from-green-50 to-green-100 p-6 rounded-lg border border-green-200"
          >
            <h4 class="font-semibold text-green-900 mb-3">Order Summary</h4>
            <div class="space-y-2 text-sm text-green-800">
              <div class="flex justify-between">
                <span>Total Items:</span>
                <span class="font-semibold">{{ form.items.length }}</span>
              </div>
              <div class="flex justify-between">
                <span>Total Units:</span>
                <span class="font-semibold">{{ getTotalQty() }}</span>
              </div>
              <div class="flex justify-between">
                <span>Total Cost:</span>
                <span class="font-semibold">{{
                  formatCurrency(getTotalCost())
                }}</span>
              </div>
              <div
                class="border-t border-green-200 pt-2 mt-2 flex justify-between font-bold"
              >
                <span>Suppliers:</span>
                <span>{{ getUniqueSuppliersCount() }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import axios from "axios";
import PurchasingSidebar from "../components/PurchasingSidebar.vue";

const API_BASE = "http://localhost:8081/api/v1";
const token = localStorage.getItem("token");

const products = ref([]);
const filteredProducts = ref([]);
const selectedProduct = ref("");
const qty = ref(1);
const submitting = ref(false);
const currentFilter = ref("all");

const form = ref({
  date: "",
  items: [], // { product_id, name, qty, price, supplier_id, supplier_name }
  submit: true,
});

// Format currency
const formatCurrency = (value) => {
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
  }).format(value || 0);
};

// Get total quantity
const getTotalQty = () => {
  return form.value.items.reduce((sum, item) => sum + item.qty, 0);
};

// Get total cost
const getTotalCost = () => {
  return form.value.items.reduce((sum, item) => sum + item.qty * item.price, 0);
};

// Get unique suppliers count
const getUniqueSuppliersCount = () => {
  const suppliers = new Set(form.value.items.map((item) => item.supplier_id));
  return suppliers.size;
};

// Check if there are missing suppliers
const hasMissingSuppliers = () => {
  return form.value.items.some((item) => !item.supplier_id);
};

// Get submitting text
const getSubmittingText = () => {
  const count = getUniqueSuppliersCount();
  return `Sending to ${count} ${count === 1 ? "supplier" : "suppliers"}...`;
};

// Get supplier product count
const getSupplierProductCount = (supplierId) => {
  return products.value.filter((p) => p.supplier_id === supplierId).length;
};

// Top suppliers computed
const topSuppliers = computed(() => {
  const supplierMap = {};
  products.value.forEach((p) => {
    if (p.supplier_id) {
      if (!supplierMap[p.supplier_id]) {
        supplierMap[p.supplier_id] = {
          id: p.supplier_id,
          name: p.supplier?.name || "Unknown",
          count: 0,
        };
      }
      supplierMap[p.supplier_id].count++;
    }
  });
  return Object.values(supplierMap)
    .sort((a, b) => b.count - a.count)
    .slice(0, 5);
});

// Filter by supplier
const filterBySupplier = (supplierId) => {
  currentFilter.value = `supplier-${supplierId}`;
  filteredProducts.value = products.value.filter(
    (p) => p.supplier_id === supplierId
  );
};

// Filter low stock
const filterLowStock = () => {
  currentFilter.value = "low-stock";
  filteredProducts.value = products.value.filter((p) => (p.stock || 0) < 10);
};

// Clear filter
const clearFilter = () => {
  currentFilter.value = "all";
  filteredProducts.value = products.value;
};

// Load produk
const loadProducts = async () => {
  try {
    const res = await axios.get(`${API_BASE}/products`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    products.value = res.data.data || [];
    filteredProducts.value = products.value;
  } catch (err) {
    console.error("Error load products:", err.response?.data || err.message);
    alert("Gagal memuat produk.");
  }
};

const addItem = () => {
  if (!selectedProduct.value) return;
  if (!qty.value || qty.value <= 0) return;

  // Ambil supplier dari produk (boleh null, tapi tidak bisa disubmit tanpa supplier)
  const sId = selectedProduct.value.supplier_id ?? null;
  const sName = selectedProduct.value.supplier?.name ?? null;

  form.value.items.push({
    product_id: selectedProduct.value.id,
    name: selectedProduct.value.name,
    qty: qty.value,
    price: selectedProduct.value.cost_price || selectedProduct.value.price || 0,
    supplier_id: sId,
    supplier_name: sName,
  });

  selectedProduct.value = "";
  qty.value = 1;
};

const removeItem = (idx) => {
  form.value.items.splice(idx, 1);
};

// Helper: group items by supplier_id
const groupItemsBySupplier = (items) => {
  const map = new Map();
  for (const it of items) {
    if (!it.supplier_id) {
      // Jika ada item tanpa supplier_id, hentikan dan beri tahu user
      throw new Error(
        `Produk "${it.name}" belum memiliki supplier. Lengkapi data supplier di master produk terlebih dahulu.`
      );
    }
    if (!map.has(it.supplier_id)) map.set(it.supplier_id, []);
    map.get(it.supplier_id).push(it);
  }
  return map;
};

// Submit: kirim per supplier ke /api/v1/pembelian/purchases
const submitPurchase = async () => {
  if (!form.value.date) {
    alert("Tanggal wajib diisi.");
    return;
  }
  if (!form.value.items.length) {
    alert("Tambahkan minimal satu produk.");
    return;
  }

  let groups;
  try {
    groups = groupItemsBySupplier(form.value.items);
  } catch (e) {
    alert(e.message);
    return;
  }

  submitting.value = true;

  // Build requests per supplier
  const dateIso = new Date(form.value.date).toISOString();
  const requests = [];
  const supplierNames = [];

  for (const [supplier_id, items] of groups.entries()) {
    requests.push(
      axios.post(
        `${API_BASE}/pembelian/purchases`,
        {
          supplier_id: Number(supplier_id),
          date: dateIso,
          submit: true, // status SUBMITTED (approve belakangan)
          items: items.map((i) => ({
            product_id: i.product_id,
            qty: i.qty,
            price: i.price,
          })),
        },
        { headers: { Authorization: `Bearer ${token}` } }
      )
    );
    // simpan nama supplier untuk ringkasan
    supplierNames.push(items[0].supplier_name || `#${supplier_id}`);
  }

  try {
    const results = await Promise.allSettled(requests);

    const ok = [];
    const fail = [];
    results.forEach((r, idx) => {
      const supName = supplierNames[idx];
      if (r.status === "fulfilled") ok.push(supName);
      else {
        const msg =
          r.reason?.response?.data?.error ||
          r.reason?.response?.data?.message ||
          r.reason?.message ||
          "Unknown error";
        fail.push(`${supName}: ${msg}`);
      }
    });

    if (ok.length && !fail.length) {
      alert(`Berhasil menyimpan purchase untuk supplier: ${ok.join(", ")}.`);
      // reset form
      form.value = { date: "", items: [], submit: true };
    } else if (ok.length && fail.length) {
      alert(
        `Sebagian berhasil:\nSukses: ${ok.join(", ")}\nGagal: ${fail.join(
          " | "
        )}`
      );
      // Hapus item yang sudah sukses (berdasarkan supplier yang sukses)
      form.value.items = form.value.items.filter(
        (i) => !ok.includes(i.supplier_name || `#${i.supplier_id}`)
      );
    } else {
      alert(`Gagal menyimpan semua purchase:\n${fail.join(" | ")}`);
    }
  } catch (err) {
    console.error("Error submit batch:", err);
    alert("Gagal menyimpan purchase.");
  } finally {
    submitting.value = false;
  }
};

onMounted(() => {
  loadProducts();
});
</script>
