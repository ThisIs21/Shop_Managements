<template>
  <div class="flex">
    <!-- Sidebar -->
    <PurchasingSidebar />

    <!-- Konten utama -->
    <main class="flex-1 ml-64 p-6">
      <h1 class="text-2xl font-bold mb-6">Create Purchase Order</h1>

      <form @submit.prevent="submitPurchase">
        <!-- Tanggal -->
        <div class="mb-4">
          <label class="block mb-2 text-sm font-medium">Tanggal</label>
          <input
            type="date"
            v-model="form.date"
            class="w-full border rounded p-2"
            required
          />
        </div>

        <!-- Produk -->
        <div class="mb-4">
          <label class="block mb-2 text-sm font-medium">Tambah Produk</label>
          <div class="flex gap-2">
            <select
              v-model="selectedProduct"
              class="border rounded p-2 min-w-[260px]"
            >
              <option value="">-- Pilih Produk --</option>
              <option v-for="p in products" :key="p.id" :value="p">
                {{ p.name }} ({{ p.supplier?.name || "Tanpa Supplier" }})
              </option>
            </select>
            <input
              type="number"
              v-model.number="qty"
              min="1"
              placeholder="Qty"
              class="w-24 border rounded p-2"
            />
            <button
              type="button"
              @click="addItem"
              class="bg-blue-500 text-white px-3 py-1 rounded"
            >
              Tambah
            </button>
          </div>
        </div>

        <!-- List Item -->
        <div v-if="form.items.length" class="mb-4">
          <h2 class="font-semibold mb-2">Daftar Item</h2>
          <table class="w-full border">
            <thead>
              <tr class="bg-gray-100">
                <th class="p-2 border">Produk</th>
                <th class="p-2 border">Supplier</th>
                <th class="p-2 border">Qty</th>
                <th class="p-2 border">Harga</th>
                <th class="p-2 border">Subtotal</th>
                <th class="p-2 border w-24">Aksi</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(item, i) in form.items" :key="i">
                <td class="border p-2">{{ item.name }}</td>
                <td class="border p-2">{{ item.supplier_name || "-" }}</td>
                <td class="border p-2 text-right">{{ item.qty }}</td>
                <td class="border p-2 text-right">{{ item.price }}</td>
                <td class="border p-2 text-right">
                  {{ item.qty * item.price }}
                </td>
                <td class="border p-2 text-center">
                  <button
                    type="button"
                    class="text-red-600"
                    @click="removeItem(i)"
                  >
                    Hapus
                  </button>
                </td>
              </tr>
            </tbody>
            <tfoot>
              <tr>
                <td class="border p-2 font-semibold" colspan="4">Total</td>
                <td class="border p-2 text-right font-semibold">
                  {{ form.items.reduce((a, b) => a + b.qty * b.price, 0) }}
                </td>
                <td class="border p-2"></td>
              </tr>
            </tfoot>
          </table>
        </div>

        <!-- Tombol -->
        <div class="mt-6 flex items-center gap-3">
          <button
            type="submit"
            class="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded disabled:opacity-60"
            :disabled="submitting || !form.items.length || !form.date"
          >
            {{ submitting ? "Menyimpan..." : "Simpan" }}
          </button>
          <span v-if="submitting" class="text-sm text-gray-500"
            >Mengirim per supplier...</span
          >
        </div>
      </form>
    </main>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";
import PurchasingSidebar from "../components/PurchasingSidebar.vue";

const API_BASE = "http://localhost:8081/api/v1";
const token = localStorage.getItem("token");

const products = ref([]);
const selectedProduct = ref("");
const qty = ref(1);
const submitting = ref(false);

const form = ref({
  date: "",
  items: [], // { product_id, name, qty, price, supplier_id, supplier_name }
  submit: true,
});

// Load produk
const loadProducts = async () => {
  try {
    const res = await axios.get(`${API_BASE}/products`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    products.value = res.data.data || [];
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
