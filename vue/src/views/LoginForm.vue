<template>
  <div
    class="min-h-screen bg-gray-50 flex items-center justify-center p-4 font-poppins relative"
  >
    <div
      class="w-full max-w-4xl bg-white rounded-[32px] shadow-2xl p-0 overflow-hidden transform transition-all duration-300 relative"
      style="box-shadow: 0 10px 30px rgba(0, 0, 0, 0.08), 0 0 0 1px #f0f0f0"
    >
      <div class="flex flex-col md:flex-row min-h-[500px]">
        <div
          class="relative w-full md:w-5/12 min-h-full bg-gray-50 hidden md:block"
        >
          <div
            class="absolute top-0 bottom-0 left-0 w-8 bg-indigo-600 rounded-tl-[32px] rounded-bl-[32px] opacity-100"
          ></div>
          <div
            class="absolute inset-0 flex items-center justify-center p-4 overflow-hidden"
          >
            <img
              src="/images/loginweb.png"
              alt="Login Illustration"
              class="w-full h-auto object-contain max-h-[80%]"
            />
          </div>
        </div>

        <div class="w-full md:w-7/12 p-6 sm:p-10 flex flex-col justify-center">
          <div class="max-w-md mx-auto w-full">
            <div class="mb-10">
              <h1 class="text-3xl font-bold text-gray-900 tracking-tight">
                Welcome Back
              </h1>
              <p class="mt-2 text-sm text-gray-500">
                Hello, friend! task manager you can trust everything. Let's get
                in touch!
              </p>
            </div>

            <form @submit.prevent="handleLogin" class="space-y-6">
              <div>
                <input
                  id="email"
                  v-model="email"
                  type="email"
                  required
                  placeholder="Email"
                  class="w-full px-4 py-3 border border-gray-200 rounded-xl bg-gray-50 focus:ring-indigo-600 focus:border-indigo-600 transition duration-300 ease-in-out placeholder-gray-500"
                  style="border-radius: 8px"
                />
              </div>

              <div>
                <input
                  id="password"
                  v-model="password"
                  type="password"
                  required
                  placeholder="Password"
                  class="w-full px-4 py-3 border border-gray-200 rounded-xl bg-gray-50 focus:ring-indigo-600 focus:border-indigo-600 transition duration-300 ease-in-out placeholder-gray-500"
                  style="border-radius: 8px"
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
                  style="border-radius: 8px"
                >
                  Let's start!
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Bagian Logika (script setup) TIDAK ADA PERUBAHAN.
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
/* PERUBAHAN: Pastikan import Poppins ada */
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
  background: rgba(109, 40, 217, 0.3); /* Warna Indigo dipertahankan */
  border-radius: 10px;
}
</style>
