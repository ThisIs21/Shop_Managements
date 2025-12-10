<template>
  <OwnerLayouts>
    <div class="flex flex-1 justify-center py-5">
      <div class="layout-content-container flex flex-col max-w-[960px] flex-1">
        <div class="flex flex-wrap justify-between gap-3 p-4">
          <p
            class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
          >
            Users
          </p>
          <router-link
            to="/add-user"
            class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-medium leading-normal"
          >
            <span class="truncate">Add User</span>
          </router-link>
        </div>
        <div class="px-4 py-3">
          <div
            class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
          >
            <table class="flex-1">
              <thead>
                <tr class="bg-white">
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal"
                  >
                    Name
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal"
                  >
                    Email
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal"
                  >
                    Role
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal"
                  >
                    Status
                  </th>
                  <th
                    class="px-4 py-3 text-left text-[#111418] text-sm font-medium leading-normal"
                  >
                    Actions
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr v-if="loading" class="border-t border-t-[#dbe0e6]">
                  <td
                    colspan="5"
                    class="h-[72px] px-4 py-2 text-center text-[#60758a] text-sm font-normal leading-normal"
                  >
                    Loading...
                  </td>
                </tr>
                <tr
                  v-else-if="users.length === 0"
                  class="border-t border-t-[#dbe0e6]"
                >
                  <td
                    colspan="5"
                    class="h-[72px] px-4 py-2 text-center text-[#60758a] text-sm font-normal leading-normal"
                  >
                    No users found.
                  </td>
                </tr>
                <tr
                  v-for="user in users"
                  :key="user.ID"
                  class="border-t border-t-[#dbe0e6]"
                >
                  <td
                    class="h-[72px] px-4 py-2 text-[#111418] text-sm font-normal leading-normal"
                  >
                    {{ user.Name }}
                  </td>
                  <td
                    class="h-[72px] px-4 py-2 text-[#111418] text-sm font-normal leading-normal"
                  >
                    {{ user.Email }}
                  </td>
                  <td
                    class="h-[72px] px-4 py-2 text-[#60758a] text-sm font-normal leading-normal"
                  >
                    {{ user.Role.Name }}
                  </td>
                  <td
                    class="h-[72px] px-4 py-2 text-sm font-normal leading-normal"
                  >
                    <button
                      :class="{
                        'bg-green-100 text-green-700': user.Status === 'active',
                        'bg-red-100 text-red-700': user.Status === 'inactive',
                      }"
                      class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 text-sm font-medium leading-normal w-full"
                    >
                      <span class="truncate">{{
                        user.Status === "active" ? "Active" : "Inactive"
                      }}</span>
                    </button>
                  </td>
                  <td
                    class="h-[72px] px-4 py-2 text-[#60758a] text-sm font-bold leading-normal tracking-[0.015em]"
                  >
                    <span
                      class="text-blue-500 hover:underline cursor-pointer"
                      @click="openEditModal(user)"
                      >Edit</span
                    >
                    /
                    <span
                      class="text-red-500 hover:underline cursor-pointer"
                      @click="confirmDelete(user.ID)"
                      >Delete</span
                    >
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>

        <!-- Modal untuk Edit -->
        <div
          v-if="showModal"
          class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
        >
          <div class="bg-white rounded-lg p-6 w-full max-w-md">
            <h2 class="text-xl font-bold mb-4">Edit User</h2>
            <form @submit.prevent="submitForm">
              <div class="mb-4">
                <label class="block text-sm font-medium text-[#111418] mb-1"
                  >Name</label
                >
                <input
                  v-model="formData.name"
                  type="text"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  placeholder="Enter name"
                  required
                />
              </div>
              <div class="mb-4">
                <label class="block text-sm font-medium text-[#111418] mb-1"
                  >Email</label
                >
                <input
                  v-model="formData.email"
                  type="email"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  placeholder="Enter email"
                  required
                />
              </div>
              <div class="mb-4">
                <label class="block text-sm font-medium text-[#111418] mb-1"
                  >Role</label
                >
                <select
                  v-model="formData.role"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  required
                >
                  <option value="" disabled>Select a role</option>
                  <option
                    v-for="role in roles"
                    :key="role.ID"
                    :value="role.Name"
                  >
                    {{ role.Name }}
                  </option>
                </select>
              </div>
              <div class="mb-4">
                <label class="block text-sm font-medium text-[#111418] mb-1"
                  >Status</label
                >
                <select
                  v-model="formData.status"
                  class="w-full rounded-lg border border-[#dbe0e6] p-2 text-sm"
                  required
                >
                  <option value="active">Active</option>
                  <option value="inactive">Inactive</option>
                </select>
              </div>
              <div class="flex justify-end gap-3">
                <button
                  type="button"
                  @click="closeModal"
                  class="px-4 py-2 bg-gray-200 text-[#111418] rounded-lg text-sm font-medium"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  class="px-4 py-2 bg-[#f0f2f5] text-[#111418] rounded-lg text-sm font-medium"
                >
                  Update
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
  </OwnerLayouts>
</template>

<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";
import OwnerLayouts from "../components/OwnerLayouts.vue";

const users = ref([]);
const roles = ref([]);
const loading = ref(true);
const showModal = ref(false);
const formData = ref({
  name: "",
  email: "",
  role: "",
  status: "active",
  id: null,
});

const fetchUsers = async () => {
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("No JWT token found in localStorage.");
      loading.value = false;
      return;
    }

    const apiUrl = "http://localhost:8081/api/v1/owner/users";

    console.log("Fetching users from:", apiUrl);
    console.log("Using token:", token);

    const response = await axios.get(apiUrl, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    console.log("API Response status:", response.status);
    console.log("API Response data:", response.data);

    if (response.data && response.data.data) {
      users.value = response.data.data;
    } else {
      console.error("Invalid API response format. Expected 'data' field.");
      users.value = [];
    }
  } catch (error) {
    if (error.response) {
      console.error(
        "Error response from server:",
        error.response.status,
        error.response.data
      );
      if (error.response.status === 404) {
        console.error("Endpoint Not Found. Check your API URL.");
      } else if (error.response.status === 401) {
        console.error("Unauthorized. The token is invalid or expired.");
      }
    } else {
      console.error("Network or other error:", error);
    }
  } finally {
    loading.value = false;
  }
};

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

    console.log("Roles API Response:", response.data);

    if (response.data && Array.isArray(response.data.data)) {
      roles.value = response.data.data;
      if (roles.value.length === 0) {
        console.warn("No roles found in the response.");
        alert("No roles available. Please add roles in the backend.");
      }
    } else {
      console.error(
        "Invalid API response format. Expected 'data' as an array."
      );
      roles.value = [];
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
      } else if (error.response.status === 401) {
        console.error("Unauthorized. Check the token.");
      }
    }
    roles.value = [];
  }
};

const openEditModal = (user) => {
  formData.value = {
    id: user.ID,
    name: user.Name,
    email: user.Email,
    role: user.Role.Name, // Gunakan nama role langsung
    status: user.Status,
  };
  showModal.value = true;
};

const closeModal = () => {
  showModal.value = false;
  formData.value = {
    name: "",
    email: "",
    role: "",
    status: "active",
    id: null,
  };
};

const submitForm = async () => {
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("No JWT token found in localStorage.");
      return;
    }

    const url = `http://localhost:8081/api/v1/owner/users/${formData.value.id}`;
    const payload = {
      name: formData.value.name,
      email: formData.value.email,
      role: formData.value.role, // Sesuaikan dengan backend
      status: formData.value.status,
    };

    console.log("Sending payload:", payload); // Debug payload

    const response = await axios.put(url, payload, {
      headers: { Authorization: `Bearer ${token}` },
    });

    if (response.status === 200) {
      fetchUsers(); // Refresh daftar users
      closeModal();
    }
  } catch (error) {
    console.error(
      "Error updating user:",
      error.response
        ? JSON.stringify(error.response.data, null, 2)
        : error.message
    );
    if (error.response && error.response.status === 400) {
      alert("Update failed: " + (error.response.data.error || "Invalid data"));
    } else {
      alert("Failed to update user. Please check the console for details.");
    }
  }
};

const confirmDelete = async (id) => {
  if (!confirm("Are you sure you want to delete this user?")) return;

  try {
    const token = localStorage.getItem("token");
    if (!token) {
      console.error("No JWT token found in localStorage.");
      return;
    }

    await axios.delete(`http://localhost:8081/api/v1/owner/users/${id}`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    fetchUsers(); // Refresh daftar users
  } catch (error) {
    console.error(
      "Error deleting user:",
      error.response ? error.response.data : error.message
    );
    alert("Failed to delete user. Please check the console for details.");
  }
};

onMounted(() => {
  fetchUsers();
  fetchRoles();
});
</script>

<style scoped>
/* Gaya CSS Anda tetap sama */
</style>
