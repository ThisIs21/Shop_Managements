<template>
  <div class="product-list-container">
    <h2>Daftar Produk</h2>

    <div v-if="loading" class="status-message">Memuat data... ⏳</div>
    <div v-else-if="error" class="status-message error">
      Terjadi kesalahan saat mengambil data: {{ error.message }} 💔
    </div>

    <div v-else-if="products.length > 0">
      <ul class="product-items">
        <li v-for="product in products" :key="product.id" class="product-item">
          <strong>ID:</strong> {{ product.id }} <br />
          <strong>Nama:</strong> {{ product.name }} <br />
          <strong>Harga:</strong> Rp{{ product.price }}
        </li>
      </ul>
    </div>

    <div v-else class="status-message">
      Tidak ada data produk yang ditemukan. 🤷‍♂
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  // Ganti nama komponen menjadi ProductList
  name: "ProductList",
  data() {
    return {
      // Ganti variabel users menjadi products
      products: [],
      loading: true,
      error: null,
    };
  },
  methods: {
    async fetchProducts() {
      try {
        const token = localStorage.getItem("authToken");
        if (!token) {
          this.error = new Error(
            "Token otorisasi tidak ditemukan. Silakan login."
          );
          this.loading = false;
          return;
        }

        const response = await axios.get(
          "http://192.168.100.111:8081/api/v1/products",
          {
            headers: {
              // PERBAIKAN: Gunakan template literal dengan backtick (` `)
              Authorization: "Bearer ${token}",
            },
          }
        );
        // Sesuaikan dengan struktur respons dari API kamu
        this.products = response.data.data;
      } catch (err) {
        this.error = err;
      } finally {
        this.loading = false;
      }
    },
  },
  mounted() {
    this.fetchProducts();
  },
};
</script>

<style scoped>
.product-list-container {
  max-width: 600px;
  margin: 40px auto;
  padding: 20px;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  font-family: sans-serif;
  text-align: left;
}

h2 {
  text-align: center;
  color: #333;
}

.status-message {
  text-align: center;
  padding: 20px;
  font-style: italic;
  color: #666;
}

.status-message.error {
  color: #d9534f;
  font-weight: bold;
}

.product-items {
  list-style-type: none;
  padding: 0;
}

.product-item {
  background: #f9f9f9;
  border: 1px solid #eee;
  border-radius: 4px;
  padding: 15px;
  margin-bottom: 10px;
}
</style>
