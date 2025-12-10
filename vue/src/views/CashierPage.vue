<template>
  <div class="min-h-screen flex bg-white">
    <CashierSidebar v-if="isLoggedIn" />

    <main class="flex-1 p-6 max-w-7xl mx-auto w-full">
      <h2 class="text-2xl font-bold text-center mb-6">Cashier - Order</h2>

      <div
        v-if="!isLoggedIn"
        class="max-w-sm mx-auto border rounded p-6 shadow"
      >
        <h3 class="text-lg font-semibold mb-4">Login Kasir</h3>
        <input
          v-model="username"
          type="text"
          placeholder="Username"
          class="w-full border rounded p-2 mb-3 text-sm"
        />
        <input
          v-model="password"
          type="password"
          placeholder="Password"
          class="w-full border rounded p-2 mb-3 text-sm"
        />
        <button
          @click="login"
          class="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700"
        >
          Login
        </button>
      </div>

      <div v-else class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <section class="md:col-span-2">
          <h3 class="text-lg font-semibold mb-4">Available Products</h3>

          <div
            class="flex flex-col md:flex-row md:items-center md:justify-between gap-3 mb-4"
          >
            <input
              v-model="search"
              type="text"
              placeholder="Cari produk..."
              class="w-full md:w-1/2 border rounded p-2 text-sm"
            />
            <select
              v-model="selectedCategory"
              class="border rounded p-2 text-sm w-full md:w-48"
            >
              <option value="All">Semua Kategori</option>
              <option v-for="cat in categories" :key="cat" :value="cat">
                {{ cat }}
              </option>
            </select>
          </div>

          <div class="grid grid-cols-2 gap-4">
            <div
              v-for="p in filteredProducts"
              :key="p.id"
              class="border rounded-lg p-4 bg-white shadow flex flex-col justify-between"
            >
              <div>
                <h4 class="font-semibold">{{ p.name }}</h4>
                <p class="text-sm text-gray-600">
                  {{ formatCurrency(p.sell_price) }}
                </p>
                <p class="text-xs text-green-600">Stock: {{ p.stock }}</p>
              </div>
              <button
                @click="addToCart(p)"
                class="mt-3 bg-blue-500 text-white text-sm py-1 rounded hover:bg-blue-600"
                :disabled="p.stock === 0"
              >
                Add
              </button>
            </div>
          </div>
        </section>

        <section class="md:col-span-1 flex flex-col">
          <h3 class="text-lg font-semibold mb-4">Cart Order</h3>

          <div
            class="border rounded-lg p-3 bg-white shadow max-h-64 overflow-y-auto"
          >
            <table class="w-full text-sm">
              <thead class="sticky top-0 bg-gray-100">
                <tr>
                  <th class="p-2">Item</th>
                  <th class="p-2">Qty</th>
                  <th class="p-2">Subtotal</th>
                  <th class="p-2">X</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(item, i) in cart" :key="item.id" class="border-t">
                  <td class="p-2">{{ item.name }}</td>
                  <td class="p-2">
                    <input
                      type="number"
                      v-model.number="item.qty"
                      min="1"
                      :max="item.stock"
                      class="w-14 border rounded p-1"
                      @change="validateQuantity(item)"
                    />
                  </td>
                  <td class="p-2">
                    {{ formatCurrency(item.sell_price * item.qty) }}
                  </td>
                  <td class="p-2">
                    <button @click="removeFromCart(i)" class="text-red-500">
                      ✕
                    </button>
                  </td>
                </tr>
                <tr v-if="cart.length === 0">
                  <td colspan="4" class="text-center text-gray-500 py-4">
                    Cart kosong
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <div class="mt-6 border rounded-lg p-3 bg-white shadow">
            <h4 class="font-semibold mb-2">Order Summary</h4>
            <div class="flex justify-between text-sm">
              <span>Subtotal</span>
              <span>{{ formatCurrency(subtotal) }}</span>
            </div>
            <div class="flex justify-between text-sm">
              <span>Discount</span>
              <span>{{ formatCurrency(discount) }}</span>
            </div>
            <div class="flex justify-between text-sm font-semibold">
              <span>Total</span>
              <span>{{ formatCurrency(totalAmount) }}</span>
            </div>

            <div class="mt-4">
              <label class="block mb-2 text-sm">Customer Name (Optional)</label>
              <input
                v-model="customerName"
                type="text"
                placeholder="Nama pelanggan"
                class="w-full border rounded p-2 text-sm mb-3"
              />
              <label class="block mb-2 text-sm">Voucher (Optional)</label>
              <select
                v-model="voucherCode"
                class="w-full border rounded p-2 text-sm mb-3"
                @change="validateVoucher"
              >
                <option value="">Pilih Voucher</option>
                <option
                  v-for="voucher in vouchers"
                  :key="voucher.id"
                  :value="voucher.Code"
                >
                  {{ voucher.Code }} ({{ voucher.Type }}: {{ voucher.Value }})
                </option>
              </select>

              <label class="block mb-2 text-sm">Amount Received</label>
              <input
                v-model.number="amountReceived"
                type="number"
                placeholder="Masukkan nominal"
                class="w-full border rounded p-2 text-sm"
              />
              <p class="text-sm mt-1">
                Change:
                <span class="font-semibold">{{ formatCurrency(change) }}</span>
              </p>
            </div>

            <div class="flex flex-col gap-3 mt-6">
              <button
                @click="checkout"
                class="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700"
              >
                Pay Now
              </button>
            </div>
          </div>
        </section>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";
import axios from "axios";
import CashierSidebar from "@/components/CashierSidebar.vue";

// State definitions
const username = ref("");
const password = ref("");
const isLoggedIn = ref(!!localStorage.getItem("token"));
const customerName = ref("");
const voucherCode = ref("");
const products = ref([]);
const cart = ref([]);
const search = ref("");
const selectedCategory = ref("All");
const categories = ref([]);
const vouchers = ref([]);
const amountReceived = ref(0);
const discount = ref(0);

// Axios instance
const api = axios.create({ baseURL: "http://localhost:8081/api/v1" });
const savedToken = localStorage.getItem("token");
if (savedToken) {
  api.defaults.headers.common["Authorization"] = `Bearer ${savedToken}`;
}

// Login function
const login = async () => {
  try {
    const res = await axios.post("http://localhost:8081/api/v1/auth/login", {
      username: username.value,
      password: password.value,
    });
    const token = res.data.token;
    localStorage.setItem("token", token);
    isLoggedIn.value = true;
    api.defaults.headers.common["Authorization"] = `Bearer ${token}`;
    await loadProducts();
    await loadVouchers();
  } catch (err) {
    alert("Login gagal: " + (err.response?.data?.error || err.message));
  }
};

// Load products
const loadProducts = async () => {
  try {
    const res = await api.get("/products");
    products.value = res.data.data || [];
    categories.value = [
      ...new Set(products.value.map((p) => p.category?.name || "Unknown")),
    ];
  } catch (err) {
    if (err.response?.status === 401) {
      isLoggedIn.value = false;
      localStorage.removeItem("token");
    }
  }
};

// Load vouchers
const loadVouchers = async () => {
  try {
    const res = await api.get("/vouchers");
    if (res.data && res.data.data) {
      vouchers.value = res.data.data.filter((v) => v.Active);
    }
  } catch (err) {
    console.error("Error loading vouchers:", err);
  }
};

// Filter products
const filteredProducts = computed(() => {
  return products.value.filter((p) => {
    const matchSearch = p.name
      .toLowerCase()
      .includes(search.value.toLowerCase());
    const matchCategory =
      selectedCategory.value === "All" ||
      p.category?.name === selectedCategory.value;
    return matchSearch && matchCategory;
  });
});

// Add to cart with stock validation
const addToCart = (product) => {
  if (product.stock === 0) {
    alert(`Stock produk ${product.name} habis!`);
    return;
  }
  const existing = cart.value.find((i) => i.id === product.id);
  if (existing) {
    if (existing.qty < product.stock) existing.qty++;
    else alert(`Stock ${product.name} hanya tersedia ${product.stock} unit!`);
  } else {
    cart.value.push({ ...product, qty: 1 });
  }
};

// Remove from cart
const removeFromCart = (index) => cart.value.splice(index, 1);

// Validate quantity to prevent exceeding stock
const validateQuantity = (item) => {
  if (item.qty > item.stock) {
    item.qty = item.stock;
    alert(`Stock ${item.name} hanya tersedia ${item.stock} unit!`);
  } else if (item.qty < 1) {
    item.qty = 1;
    alert("Jumlah minimal 1!");
  }
};

// Compute subtotal
const subtotal = computed(() =>
  cart.value.reduce((sum, item) => sum + item.sell_price * item.qty, 0)
);

// Validate voucher and compute discount
const validateVoucher = async () => {
  if (!voucherCode.value) {
    discount.value = 0;
    return;
  }

  const selectedVoucher = vouchers.value.find(
    (v) => v.Code === voucherCode.value
  );

  if (!selectedVoucher) {
    alert("Voucher tidak ditemukan!");
    discount.value = 0;
    voucherCode.value = "";
    return;
  }

  switch (selectedVoucher.Type) {
    case "PERCENT":
      discount.value = (selectedVoucher.Value / 100) * subtotal.value;
      break;
    case "AMOUNT":
      discount.value = Math.min(selectedVoucher.Value, subtotal.value);
      break;
    default:
      discount.value = 0;
  }
};

// Compute total amount
const totalAmount = computed(() =>
  Math.max(0, subtotal.value - discount.value)
);

// Compute change
const change = computed(() =>
  Math.max(0, amountReceived.value - totalAmount.value)
);

// Checkout function
const checkout = async () => {
  if (cart.value.length === 0) {
    alert("Cart masih kosong!");
    return;
  }
  if (amountReceived.value < totalAmount.value) {
    alert("Uang kurang!");
    return;
  }

  try {
    const payload = {
      customer_name: customerName.value || null,
      voucher_code: voucherCode.value || null,
      paid_amount: amountReceived.value,
      items: cart.value.map((item) => ({
        product_id: item.id,
        qty: item.qty,
      })),
    };

    const res = await api.post("/kasir/sales", payload);
    alert(`Pembayaran berhasil! ID Transaksi: ${res.data.id}`);

    // Cetak struk otomatis
    const receipt = `
      <div style="font-family: Arial, sans-serif; max-width: 300px; margin: 0 auto; padding: 20px; border: 1px solid #ccc;">
        <h2 style="text-align: center;">Struk Pembayaran</h2>
        <p>Tanggal: ${new Date().toLocaleString("id-ID")}</p>
        <p>ID Transaksi: ${res.data.id}</p>
        <p>Pelanggan: ${customerName.value || "Umum"}</p>
        <h3>Detail Pembelian</h3>
        ${cart.value
          .map(
            (item) => `
          <p>${item.name} x${item.qty} = ${formatCurrency(
              item.sell_price * item.qty
            )}</p>
        `
          )
          .join("")}
        <p>Subtotal: ${formatCurrency(subtotal.value)}</p>
        <p>Diskon: ${formatCurrency(discount.value)}</p>
        <p>Total: ${formatCurrency(totalAmount.value)}</p>
        <p>Dibayar: ${formatCurrency(amountReceived.value)}</p>
        <p>Kembali: ${formatCurrency(change.value)}</p>
        <p style="text-align: center; margin-top: 20px;">Terima kasih telah berbelanja!</p>
      </div>
    `;
    const printWindow = window.open("", "", "width=300,height=400");
    printWindow.document.write(receipt);
    printWindow.document.close();
    printWindow.print();
    printWindow.close();

    // Reset dan auto reload
    cart.value = [];
    amountReceived.value = 0;
    customerName.value = "";
    voucherCode.value = "";
    discount.value = 0;
    await loadProducts(); // Refresh stok
    await loadVouchers(); // Refresh daftar voucher
  } catch (err) {
    console.error(err);
    const errorMsg =
      err.response?.data?.message || err.message || "Transaksi gagal!";
    alert(errorMsg);
  }
};

// Format currency with clearer thousand separator
const formatCurrency = (num) =>
  new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0, // Pastikan tidak ada desimal yang tidak perlu
    maximumFractionDigits: 0, // Pastikan hanya pemisah ribuan
  }).format(num);

// Load products on mount
onMounted(() => {
  if (isLoggedIn.value) {
    loadProducts();
    loadVouchers();
  }
});
</script>

<style scoped>
/* Opsional: Sesuaikan gaya sesuai kebutuhan */
</style>
