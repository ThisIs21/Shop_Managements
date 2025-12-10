<template>
  <OwnerLayouts>
    <div class="p-8">
      <div class="flex flex-wrap items-center justify-between gap-3 p-4">
        <p
          class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
        >
          Add User
        </p>
        <router-link
          to="/management-user"
          class="flex items-center gap-2 text-[#60758a] text-base font-normal leading-normal"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="20px"
            height="20px"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path d="M19 12H5"></path>
            <path d="M12 19L5 12L12 5"></path>
          </svg>
          <span class="truncate">Kembali</span>
        </router-link>
      </div>

      <form
        @submit.prevent="createUser"
        class="flex flex-col w-[512px] max-w-[512px] py-5"
      >
        <div class="flex flex-col flex-wrap items-end gap-4 px-4 py-3">
          <label class="flex flex-col min-w-40 flex-1 w-full">
            <p class="text-[#111418] text-base font-medium leading-normal pb-2">
              Name
            </p>
            <input
              v-model="user.name"
              placeholder="Enter name"
              class="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dbe0e6] bg-white focus:border-[#dbe0e6] h-14 placeholder:text-[#60758a] p-[15px] text-base font-normal leading-normal"
            />
          </label>
        </div>
        <div class="flex flex-col flex-wrap items-end gap-4 px-4 py-3">
          <label class="flex flex-col min-w-40 flex-1 w-full">
            <p class="text-[#111418] text-base font-medium leading-normal pb-2">
              Email
            </p>
            <input
              v-model="user.email"
              type="email"
              placeholder="Enter email"
              class="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dbe0e6] bg-white focus:border-[#dbe0e6] h-14 placeholder:text-[#60758a] p-[15px] text-base font-normal leading-normal"
            />
          </label>
        </div>
        <div class="flex flex-col flex-wrap items-end gap-4 px-4 py-3">
          <label class="flex flex-col min-w-40 flex-1 w-full">
            <p class="text-[#111418] text-base font-medium leading-normal pb-2">
              Password
            </p>
            <input
              v-model="user.password"
              type="password"
              placeholder="Enter password"
              class="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dbe0e6] bg-white focus:border-[#dbe0e6] h-14 placeholder:text-[#60758a] p-[15px] text-base font-normal leading-normal"
            />
          </label>
        </div>
        <div class="flex flex-col flex-wrap items-end gap-4 px-4 py-3">
          <label class="flex flex-col min-w-40 flex-1 w-full">
            <p class="text-[#111418] text-base font-medium leading-normal pb-2">
              Role
            </p>
            <select
              v-model="user.role_id"
              class="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dbe0e6] bg-white focus:border-[#dbe0e6] h-14 placeholder:text-[#60758a] p-[15px] text-base font-normal leading-normal"
            >
              <option value="" disabled selected>Select a role</option>
              <option v-for="role in roles" :key="role.ID" :value="role.ID">
                {{ role.Name || "Unnamed Role" }}
              </option>
            </select>
          </label>
        </div>
        <div class="flex flex-col flex-wrap items-end gap-4 px-4 py-3">
          <label class="flex flex-col min-w-40 flex-1 w-full">
            <p class="text-[#111418] text-base font-medium leading-normal pb-2">
              Status
            </p>
            <select
              v-model="user.status"
              class="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dbe0e6] bg-white focus:border-[#dbe0e6] h-14 placeholder:text-[#60758a] p-[15px] text-base font-normal leading-normal"
            >
              <option value="active">Active</option>
              <option value="inactive">Inactive</option>
            </select>
          </label>
        </div>
        <div class="flex justify-stretch">
          <div class="flex flex-1 gap-3 flex-wrap px-4 py-3 justify-end">
            <router-link
              to="/management-user"
              class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-bold leading-normal tracking-[0.015em]"
            >
              <span class="truncate">Cancel</span>
            </router-link>
            <button
              type="submit"
              class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 px-4 bg-[#0d80f2] text-white text-sm font-bold leading-normal tracking-[0.015em]"
            >
              <span class="truncate">Save</span>
            </button>
          </div>
        </div>
      </form>
    </div>
  </OwnerLayouts>
</template>

<script setup>
import { ref, onMounted } from "vue";
import { useRouter } from "vue-router";
import axios from "axios";
import OwnerLayouts from "../components/OwnerLayouts.vue";

const user = ref({
  name: "",
  email: "",
  password: "",
  role_id: "",
  status: "active", // Default value
});

const roles = ref([]);
const router = useRouter();

// Fungsi untuk mengambil daftar role dari backend
const fetchRoles = async () => {
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("No JWT token found in localStorage.");
      return;
    }

    const apiUrl = "http://localhost:8081/api/v1/owner/roles";
    const response = await axios.get(apiUrl, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    console.log("Roles API Response:", response.data); // Debug response

    if (response.data && Array.isArray(response.data.data)) {
      roles.value = response.data.data;
      if (roles.value.length === 0) {
        console.warn("No roles found in the response.");
        alert("No roles available. Please add roles in the backend.");
      }
    } else if (response.data && Array.isArray(response.data)) {
      // Fallback jika data langsung dalam array
      roles.value = response.data;
      if (roles.value.length === 0) {
        console.warn("No roles found in the response.");
        alert("No roles available. Please add roles in the backend.");
      }
    } else {
      console.error(
        "Invalid API response format. Expected 'data' as an array."
      );
      roles.value = [];
      alert("Failed to load roles. Check the API response format.");
    }
  } catch (error) {
    console.error(
      "Error fetching roles:",
      error.response ? error.response.data : error.message
    );
    if (error.response) {
      console.error("Status:", error.response.status);
      if (error.response.status === 404) {
        console.error("Endpoint not found. Check the API URL.");
        alert("Role endpoint not found. Verify the URL.");
      } else if (error.response.status === 401) {
        console.error("Unauthorized. Check the token.");
        alert("Unauthorized. Please log in again.");
      } else if (error.response.status === 500) {
        console.error("Server error:", error.response.data);
        alert("Server error. Contact administrator.");
      }
    } else {
      console.error("Network error:", error.message);
      alert("Network error. Check your connection.");
    }
    roles.value = [];
  }
};

// Fungsi untuk mengirim data user baru ke backend
const createUser = async () => {
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("No JWT token found in localStorage.");
      return;
    }

    // Validasi sederhana
    if (
      !user.value.name ||
      !user.value.email ||
      !user.value.password ||
      !user.value.role_id
    ) {
      alert("Please fill in all required fields.");
      return;
    }

    const apiUrl = "http://localhost:8081/api/v1/owner/users";
    const response = await axios.post(apiUrl, user.value, {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
    });

    console.log("User created successfully:", response.data);

    // Navigasi kembali ke halaman manajemen pengguna setelah berhasil
    router.push("/management-user");
  } catch (error) {
    console.error(
      "Error creating user:",
      error.response ? error.response.data : error.message
    );
    alert("Failed to create user. Please check the console for details.");
  }
};

onMounted(() => {
  fetchRoles();
});
</script>
