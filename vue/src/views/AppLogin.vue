<template>
  <div
    class="min-h-screen bg-gray-50 flex items-center justify-center p-4 font-poppins relative overflow-hidden"
  >
    <div
      class="absolute inset-0 bg-gradient-to-br from-gray-50 via-gray-100 to-gray-200 opacity-50 -z-1"
    ></div>
    <div
      class="absolute top-0 left-0 w-64 h-64 bg-indigo-200 rounded-full mix-blend-multiply filter blur-xl opacity-30 animate-blob"
    ></div>
    <div
      class="absolute bottom-10 right-10 w-64 h-64 bg-purple-200 rounded-full mix-blend-multiply filter blur-xl opacity-30 animate-blob animation-delay-2000"
    ></div>
    <div
      class="w-full max-w-md bg-white rounded-xl shadow-2xl p-6 sm:p-10 transform transition-all duration-300 hover:shadow-3xl relative z-10"
    >
      <div class="mb-8 text-center">
        <h1
          class="text-3xl md:text-4xl font-extrabold text-gray-900 tracking-tight"
        >
          Welcome Back KasirKu
        </h1>
        <p class="mt-3 text-base text-gray-500">
          Please sign in to your account yaaa
        </p>
      </div>

      <form @submit.prevent="handleLogin" class="space-y-6">
        <div>
          <label
            for="email"
            class="block text-sm font-semibold text-gray-700 mb-2"
            >Email Address</label
          >
          <input
            id="email"
            v-model="email"
            type="email"
            required
            placeholder="Enter your email"
            class="w-full px-4 py-3 border border-gray-300 rounded-xl shadow-inner-sm focus:ring-indigo-600 focus:border-indigo-600 transition duration-300 ease-in-out placeholder-gray-400 focus:placeholder-gray-300"
          />
        </div>

        <div>
          <label
            for="password"
            class="block text-sm font-semibold text-gray-700 mb-2"
            >Password</label
          >
          <input
            id="password"
            v-model="password"
            type="password"
            required
            placeholder="Enter your password"
            class="w-full px-4 py-3 border border-gray-300 rounded-xl shadow-inner-sm focus:ring-indigo-600 focus:border-indigo-600 transition duration-300 ease-in-out placeholder-gray-400 focus:placeholder-gray-300"
          />
        </div>

        <div
          v-if="errorMessage"
          class="text-sm text-red-700 bg-red-100 p-3 rounded-xl border border-red-300 transition-all duration-300 ease-in-out transform scale-100"
        >
          {{ errorMessage }}
        </div>

        <div>
          <button
            type="submit"
            class="w-full px-4 py-3 bg-indigo-600 text-white text-lg font-bold rounded-xl shadow-lg shadow-indigo-300/50 hover:bg-indigo-700 focus:outline-none focus:ring-4 focus:ring-indigo-500 focus:ring-offset-2 transition-all duration-300 ease-in-out transform hover:scale-[1.01] hover:shadow-xl active:scale-[0.99]"
          >
            Sign In
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";
import api from "../services/api.js";

const router = useRouter();
const email = ref("");
const password = ref("");
const errorMessage = ref("");

const handleLogin = async () => {
  errorMessage.value = "";
  try {
    const response = await api.post("/auth/login", {
      email: email.value,
      password: password.value,
    });
    const { token, role } = response.data.data;
    localStorage.setItem("token", token);
    localStorage.setItem("userRole", role);

    switch (role) {
      case "OWNER":
        router.push("/dashboard-owner");
        break;
      case "KASIR":
        router.push("/dashboard-cashier");
        break;
      case "GUDANG":
        router.push("/dashboard-gudang");
        break;
      case "PEMBELIAN":
        router.push("/dashboard-pembelian");
        break;
      case "KEPALA_GUDANG":
        router.push("/dashboard-kepalagudang");
        break;
      default:
        router.push("/");
        break;
    }
  } catch (error) {
    if (error.response && error.response.status === 401) {
      errorMessage.value = "Email atau password salah!";
    } else {
      errorMessage.value = "Terjadi kesalahan saat login. Mohon coba lagi.";
    }
    console.error("Error saat login:", error);
  }
};
</script>

<style scoped>
/* Import Poppins font */
@import url("https://fonts.googleapis.com/css2?family=Poppins:wght@300;400;500;600;700;800&display=swap");

.font-poppins {
  font-family: "Poppins", sans-serif;
}

/* Custom scrollbar for inputs */
input::-webkit-scrollbar {
  width: 4px;
}
input::-webkit-scrollbar-track {
  background: transparent;
}
input::-webkit-scrollbar-thumb {
  background: rgba(
    109,
    40,
    217,
    0.3
  ); /* Sedikit warna indigo untuk scrollbar */
  border-radius: 10px;
}

/* Kunci Z-index */
.z-10 {
  z-index: 10;
}
.-z-1 {
  z-index: -1;
}

/* Animasi untuk latar belakang dekoratif */
@keyframes blob {
  0% {
    transform: translate(0, 0) scale(1);
  }
  33% {
    transform: translate(30px, -50px) scale(1.1);
  }
  66% {
    transform: translate(-20px, 20px) scale(0.9);
  }
  100% {
    transform: translate(0, 0) scale(1);
  }
}

.animate-blob {
  animation: blob 7s infinite cubic-bezier(0.68, -0.55, 0.265, 1.55);
}

.animation-delay-2000 {
  animation-delay: 2s;
}

/* Bayangan kustom untuk tampilan yang lebih modern */
.shadow-2xl {
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.1),
    0 10px 10px -5px rgba(0, 0, 0, 0.04);
}
.hover\:shadow-3xl:hover {
  box-shadow: 0 35px 60px -15px rgba(0, 0, 0, 0.2),
    0 25px 30px -15px rgba(0, 0, 0, 0.08);
}
.shadow-inner-sm {
  box-shadow: inset 0 1px 2px 0 rgba(0, 0, 0, 0.05);
}
</style>
