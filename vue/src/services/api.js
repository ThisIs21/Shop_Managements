import axios from "axios";

// Buat instance Axios dengan konfigurasi dasar
const api = axios.create({
  // Atur URL dasar untuk semua permintaan API.
  // Ganti dengan alamat IP lokal dan port backend Go Anda.
  // Contoh: "http://192.168.1.10:8081/api/v1"
  baseURL: "http://localhost:8081/api/v1",

  // Atur header default untuk semua permintaan
  headers: {
    "Content-Type": "application/json",
  },
});

// Tambahkan interceptor untuk menyisipkan token ke setiap header permintaan
api.interceptors.request.use(
  (config) => {
    // Ambil token dari local storage
    const token = localStorage.getItem("token");

    // Jika token ada, tambahkan ke header Authorization
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// Ekspor instance api yang sudah dikonfigurasi
export default api;
