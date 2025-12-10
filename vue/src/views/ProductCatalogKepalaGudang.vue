<template>
  <div class="flex h-screen bg-gray-100">
    <!-- Sidebar -->
    <KepalaGudangSidebar />

    <!-- Main Content -->
    <div class="flex-1 flex flex-col">
      <!-- Header -->
      <header class="bg-white shadow p-4 flex justify-between items-center">
        <h1 class="text-2xl font-bold">Product Catalog - Kepala Gudang</h1>
      </header>

      <!-- Content -->
      <main class="p-6 flex-1 overflow-y-auto">
        <!-- Daftar Produk -->
        <section>
          <h2 class="text-xl font-semibold mb-4">Daftar Produk</h2>

          <div class="overflow-x-auto bg-white p-4 rounded shadow">
            <table class="w-full border-collapse text-sm">
              <thead class="bg-gray-100">
                <tr>
                  <th class="border p-2 text-left">Nama</th>
                  <th class="border p-2 text-left">Kategori</th>
                  <th class="border p-2 text-right">Harga Beli</th>
                  <th class="border p-2 text-right">Harga Jual</th>
                  <th class="border p-2 text-right">Stok</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="prod in products" :key="prod.id">
                  <td class="border p-2">{{ prod.name }}</td>
                  <td class="border p-2">
                    {{ prod.category?.name ?? prod.category ?? "-" }}
                  </td>
                  <td class="border p-2 text-right">
                    {{
                      prod.cost_price != null
                        ? "Rp " + formatPrice(prod.cost_price)
                        : "-"
                    }}
                  </td>
                  <td class="border p-2 text-right">
                    {{
                      prod.sell_price != null
                        ? "Rp " + formatPrice(prod.sell_price)
                        : "-"
                    }}
                  </td>
                  <td class="border p-2 text-right">{{ prod.stock ?? "-" }}</td>
                </tr>

                <tr v-if="products.length === 0">
                  <td class="text-center p-4" colspan="5">
                    No products available
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </section>
      </main>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import KepalaGudangSidebar from "@/components/KepalaGudangSidebar.vue";

export default {
  name: "ProductCatalogKepalaGudang",
  components: { KepalaGudangSidebar },
  data() {
    return {
      products: [],
      api: axios.create({
        baseURL: "http://localhost:8081/api/v1",
        timeout: 10000,
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token") || ""}`,
        },
      }),
    };
  },
  methods: {
    async fetchProducts() {
      try {
        const res = await this.api.get("/products");
        this.products = Array.isArray(res.data)
          ? res.data
          : res.data.data ?? [];
      } catch (err) {
        console.error("Gagal ambil produk:", err.response?.data ?? err.message);
        this.products = [];
      }
    },
    formatPrice(value) {
      if (value == null || isNaN(value)) return "0";
      return new Intl.NumberFormat("id-ID").format(value);
    },
  },
  mounted() {
    this.fetchProducts();
  },
};
</script>
