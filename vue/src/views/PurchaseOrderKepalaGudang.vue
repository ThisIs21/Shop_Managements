<script setup>
import { ref, onMounted } from "vue";
import axios from "axios";
import KepalaGudangSidebar from "@/components/KepalaGudangSidebar.vue";

// URL Base API (Penting: Ganti dengan variabel lingkungan jika ada)
const API_BASE = "http://localhost:8081/api/v1";

// State
const purchases = ref([]);
const loading = ref(false);
const showApprovalModal = ref(false);
const showRejectModal = ref(false);
const showDetailsModal = ref(false);
const selectedPurchaseId = ref(null);
const selectedPurchase = ref(null);
const rejectReason = ref("");
const approving = ref(false);
const rejecting = ref(false);
const token = localStorage.getItem("token");
// Perhatikan: Sesuaikan inisialisasi role jika diperlukan
const role = localStorage.getItem("userRole") || "KEPALA_GUDANG";

// Fetch purchases (default to SUBMITTED status)
const fetchPurchases = async (status = "SUBMITTED") => {
  if (!token) {
    alert("No token found. Please log in again.");
    return;
  }
  if (role !== "KEPALA_GUDANG") {
    alert("Access denied: This page requires KEPALA_GUDANG role.");
    return;
  }

  loading.value = true;
  try {
    const res = await axios.get(`${API_BASE}/kepala-gudang/purchases`, {
      params: { page: 1, size: 20, q: "", status },
      headers: { Authorization: `Bearer ${token}` },
    });
    console.log("API Response (fetchPurchases):", res.data); // Debug full response
    // Sesuaikan dengan struktur API Anda. Asumsi res.data.data.data adalah array PO.
    purchases.value = res.data.data?.data || [];
    if (purchases.value.length === 0) {
      console.warn("No purchase data returned from API.");
    }
  } catch (err) {
    console.error("Failed to fetch purchases:", err);
    const status = err.response?.status || "Unknown";
    const message = err.response?.data?.error || err.message || "Network error";
    alert(`Failed to fetch purchases: ${status} - ${message}`);
    purchases.value = []; // Reset to empty array on error
  } finally {
    loading.value = false;
  }
};

// Open approval modal
const openApprovalModal = async (id) => {
  if (!id) {
    alert("Invalid purchase ID.");
    return;
  }
  try {
    const res = await axios.get(`${API_BASE}/kepala-gudang/purchases/${id}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    console.log("Purchase Details Response (Approval):", res.data); // Debug response
    selectedPurchase.value = res.data.data; // Adjust for nested data
    selectedPurchaseId.value = id;
    showApprovalModal.value = true;
  } catch (err) {
    console.error("Failed to fetch purchase details:", err);
    const status = err.response?.status || "Unknown";
    const message = err.response?.data?.error || err.message || "Network error";
    alert(`Failed to fetch purchase details: ${status} - ${message}`);
  }
};

// Confirm approval
const confirmApprove = async () => {
  if (!selectedPurchaseId.value) return;

  approving.value = true;
  try {
    const res = await axios.post(
      `${API_BASE}/kepala-gudang/purchases/${selectedPurchaseId.value}/approve`,
      {},
      { headers: { Authorization: `Bearer ${token}` } }
    );
    if (res.status === 200) {
      alert("Purchase order approved successfully!");
      closeApprovalModal();
      await fetchPurchases(); // Ensure refresh after approval
    }
  } catch (err) {
    console.error("Failed to approve purchase:", err);
    const status = err.response?.status || "Unknown";
    const message = err.response?.data?.error || err.message || "Network error";
    alert(`Failed to approve purchase: ${status} - ${message}`);
  } finally {
    approving.value = false;
  }
};

// Open reject modal
const openRejectModal = async (id) => {
  if (!id) {
    alert("Invalid purchase ID.");
    return;
  }
  try {
    const res = await axios.get(`${API_BASE}/kepala-gudang/purchases/${id}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    console.log("Purchase Details Response (Reject):", res.data); // Debug response
    selectedPurchase.value = res.data.data; // Adjust for nested data
    selectedPurchaseId.value = id;
    rejectReason.value = "";
    showRejectModal.value = true;
  } catch (err) {
    console.error("Failed to fetch purchase details:", err);
    const status = err.response?.status || "Unknown";
    const message = err.response?.data?.error || err.message || "Network error";
    alert(`Failed to fetch purchase details: ${status} - ${message}`);
  }
};

// Confirm rejection
const confirmReject = async () => {
  if (!selectedPurchaseId.value || !rejectReason.value.trim()) {
    alert("Reason for rejection is required!");
    return;
  }

  rejecting.value = true;
  try {
    const res = await axios.post(
      `${API_BASE}/kepala-gudang/purchases/${selectedPurchaseId.value}/reject`,
      { reason: rejectReason.value.trim() },
      { headers: { Authorization: `Bearer ${token}` } }
    );
    if (res.status === 200) {
      alert("Purchase order rejected successfully!");
      closeRejectModal();
      await fetchPurchases(); // Ensure refresh after rejection
    }
  } catch (err) {
    console.error("Failed to reject purchase:", err);
    const status = err.response?.status || "Unknown";
    const message = err.response?.data?.error || err.message || "Network error";
    alert(`Failed to reject purchase: ${status} - ${message}`);
  } finally {
    rejecting.value = false;
  }
};

// View details
const viewDetails = async (id) => {
  if (!id) {
    alert("Invalid purchase ID.");
    return;
  }
  try {
    const res = await axios.get(`${API_BASE}/kepala-gudang/purchases/${id}`, {
      headers: { Authorization: `Bearer ${token}` },
    });
    console.log("Purchase Details Response (View):", res.data); // Debug response
    selectedPurchase.value = res.data.data; // Adjust for nested data
    selectedPurchaseId.value = id;
    showDetailsModal.value = true;
  } catch (err) {
    console.error("Failed to fetch purchase details:", err);
    const status = err.response?.status || "Unknown";
    const message = err.response?.data?.error || err.message || "Network error";
    alert(`Failed to fetch purchase details: ${status} - ${message}`);
  }
};

// Close modals
const closeApprovalModal = () => {
  showApprovalModal.value = false;
  selectedPurchase.value = null;
  selectedPurchaseId.value = null;
};
const closeRejectModal = () => {
  showRejectModal.value = false;
  selectedPurchase.value = null;
  selectedPurchaseId.value = null;
};
const closeDetailsModal = () => {
  showDetailsModal.value = false;
  selectedPurchase.value = null;
  selectedPurchaseId.value = null;
};

// Format PO number
const formatPONumber = (id) => {
  return `PO-${String(id).padStart(3, "0")}`;
};

// Format tanggal
const formatDate = (dateStr) => {
  if (!dateStr) return "-";
  // Hapus Z jika ada untuk menghindari masalah zona waktu lokal
  const date = new Date(dateStr.endsWith("Z") ? dateStr : dateStr + "Z");
  return date.toLocaleDateString("id-ID", {
    day: "2-digit",
    month: "short",
    year: "numeric",
  });
};

// Format mata uang
const formatCurrency = (val) => {
  if (val === null || val === undefined) return "-";
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    minimumFractionDigits: 0,
  }).format(val);
};

// Format status
const getStatusText = (status) => {
  switch (status) {
    case "SUBMITTED":
      return "Submitted";
    case "APPROVED":
      return "Approved";
    case "REJECTED":
      return "Rejected";
    case "DRAFT":
      return "Draft";
    default:
      return status || "-";
  }
};

// Get status class
const getStatusClass = (status) => {
  switch (status) {
    case "SUBMITTED":
      return "bg-yellow-200 text-yellow-800";
    case "APPROVED":
      return "bg-green-200 text-green-800";
    case "REJECTED":
      return "bg-red-200 text-red-800";
    case "DRAFT":
      return "bg-gray-200 text-gray-800";
    default:
      return "bg-gray-200 text-gray-800";
  }
};

onMounted(() => {
  fetchPurchases();
});
</script>

<template>
  <div
    class="flex size-full min-h-screen flex-col bg-white font-manrope lg:flex-row"
  >
    <KepalaGudangSidebar />

    <div class="flex flex-1 flex-col">
      <main class="flex-1 p-8">
        <div class="layout-content-container flex flex-col flex-1">
          <div class="mb-6 flex flex-wrap justify-between items-center gap-3">
            <div>
              <p
                class="text-[#111418] tracking-light text-[32px] font-bold leading-tight"
              >
                Purchase Orders Approval
              </p>
              <p class="text-[#60758a] text-sm font-normal leading-normal">
                Review and approve or reject pending purchase orders.
              </p>
            </div>
          </div>

          <div class="px-4 py-3 @container">
            <div
              class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
            >
              <table class="flex-1 w-full">
                <thead>
                  <tr class="bg-white">
                    <th
                      class="px-4 py-3 text-left text-[#111418] w-[120px] text-sm font-medium"
                    >
                      PO Number
                    </th>
                    <th
                      class="px-4 py-3 text-left text-[#111418] w-[200px] text-sm font-medium"
                    >
                      Supplier
                    </th>
                    <th
                      class="px-4 py-3 text-left text-[#111418] w-[150px] text-sm font-medium"
                    >
                      Date
                    </th>
                    <th
                      class="px-4 py-3 text-left text-[#111418] w-[150px] text-sm font-medium"
                    >
                      Total
                    </th>
                    <th
                      class="px-4 py-3 text-left text-[#111418] w-[120px] text-sm font-medium"
                    >
                      Status
                    </th>
                    <th
                      class="px-4 py-3 text-left text-[#111418] w-[150px] text-sm font-medium"
                    >
                      Actions
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr
                    v-for="po in purchases"
                    :key="po.ID"
                    class="border-t border-t-[#dbe0e6]"
                  >
                    <td class="h-[72px] px-4 py-2 text-[#111418] text-sm">
                      {{ formatPONumber(po.ID) }}
                    </td>
                    <td class="h-[72px] px-4 py-2 text-[#60758a] text-sm">
                      {{ po.Supplier?.name || "-" }}
                    </td>
                    <td class="h-[72px] px-4 py-2 text-[#60758a] text-sm">
                      {{ formatDate(po.Date) }}
                    </td>
                    <td class="h-[72px] px-4 py-2 text-[#60758a] text-sm">
                      {{ formatCurrency(po.Total) }}
                    </td>
                    <td class="h-[72px] px-4 py-2 text-sm">
                      <span
                        :class="getStatusClass(po.Status)"
                        class="px-3 py-1 rounded-lg text-xs font-medium"
                      >
                        {{ getStatusText(po.Status) }}
                      </span>
                    </td>
                    <td class="h-[72px] px-4 py-2 text-sm font-bold">
                      <div v-if="po.Status === 'SUBMITTED'" class="flex gap-2">
                        <button
                          @click="openApprovalModal(po.ID)"
                          :disabled="approving"
                          class="px-3 py-1 bg-green-500 text-white rounded text-xs hover:bg-green-600 disabled:opacity-50"
                        >
                          {{ approving ? "Approving..." : "Approve" }}
                        </button>
                        <button
                          @click="openRejectModal(po.ID)"
                          :disabled="rejecting"
                          class="px-3 py-1 bg-red-500 text-white rounded text-xs hover:bg-red-600 disabled:opacity-50"
                        >
                          {{ rejecting ? "Rejecting..." : "Reject" }}
                        </button>
                        <button
                          @click="viewDetails(po.ID)"
                          class="px-3 py-1 bg-blue-500 text-white rounded text-xs hover:bg-blue-600"
                        >
                          View Details
                        </button>
                      </div>
                      <div v-else class="flex gap-2">
                        <button
                          @click="viewDetails(po.ID)"
                          class="px-3 py-1 bg-blue-500 text-white rounded text-xs hover:bg-blue-600"
                        >
                          View Details
                        </button>
                        <span class="text-gray-500">{{
                          getStatusText(po.Status)
                        }}</span>
                      </div>
                    </td>
                  </tr>

                  <tr v-if="!loading && purchases.length === 0">
                    <td
                      colspan="6"
                      class="text-center py-4 text-[#60758a] text-sm"
                    >
                      No purchase orders found
                    </td>
                  </tr>

                  <tr v-if="loading">
                    <td
                      colspan="6"
                      class="text-center py-4 text-[#60758a] text-sm"
                    >
                      Loading...
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>

          <div
            v-if="showApprovalModal"
            class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
          >
            <div class="bg-white rounded-lg p-6 max-w-md w-full">
              <h2 class="text-xl font-bold mb-4">Approve Purchase Order</h2>
              <div class="mb-4">
                <p>
                  <strong>PO Number:</strong>
                  {{ formatPONumber(selectedPurchaseId) }}
                </p>
                <p>
                  <strong>Supplier:</strong>
                  {{ selectedPurchase?.Supplier?.name || "-" }}
                </p>
                <p>
                  <strong>Date:</strong>
                  {{ formatDate(selectedPurchase?.Date) }}
                </p>
                <p>
                  <strong>Total:</strong>
                  {{ formatCurrency(selectedPurchase?.Total) }}
                </p>
                <p>
                  <strong>Requested By:</strong> User ID
                  {{ selectedPurchase?.UserID || "-" }}
                </p>
              </div>
              <div class="flex justify-end gap-2">
                <button
                  @click="closeApprovalModal"
                  class="px-4 py-2 bg-gray-500 text-white rounded-md hover:bg-gray-600"
                >
                  Cancel
                </button>
                <button
                  @click="confirmApprove"
                  :disabled="approving"
                  class="px-4 py-2 bg-green-500 text-white rounded-md hover:bg-green-600 disabled:opacity-50"
                >
                  {{ approving ? "Approving..." : "Approve" }}
                </button>
              </div>
            </div>
          </div>

          <div
            v-if="showRejectModal"
            class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
          >
            <div class="bg-white rounded-lg p-6 max-w-md w-full">
              <h2 class="text-xl font-bold mb-4">Reject Purchase Order</h2>
              <div class="mb-4">
                <p>
                  <strong>PO Number:</strong>
                  {{ formatPONumber(selectedPurchaseId) }}
                </p>
                <p>
                  <strong>Supplier:</strong>
                  {{ selectedPurchase?.Supplier?.name || "-" }}
                </p>
                <p>
                  <strong>Date:</strong>
                  {{ formatDate(selectedPurchase?.Date) }}
                </p>
                <p>
                  <strong>Total:</strong>
                  {{ formatCurrency(selectedPurchase?.Total) }}
                </p>
                <p>
                  <strong>Requested By:</strong> User ID
                  {{ selectedPurchase?.UserID || "-" }}
                </p>
              </div>
              <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700 mb-2"
                  >Reason for Rejection</label
                >
                <textarea
                  v-model="rejectReason"
                  placeholder="Enter reason for rejection..."
                  rows="3"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-red-500"
                  required
                ></textarea>
              </div>
              <div class="flex justify-end gap-2">
                <button
                  @click="closeRejectModal"
                  class="px-4 py-2 bg-gray-500 text-white rounded-md hover:bg-gray-600"
                >
                  Cancel
                </button>
                <button
                  @click="confirmReject"
                  :disabled="rejecting || !rejectReason.trim()"
                  class="px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 disabled:opacity-50"
                >
                  {{ rejecting ? "Rejecting..." : "Reject" }}
                </button>
              </div>
            </div>
          </div>

          <div
            v-if="showDetailsModal"
            class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
          >
            <div
              class="bg-white rounded-lg p-6 max-w-2xl w-full max-h-[90vh] overflow-y-auto"
            >
              <div class="flex justify-between items-center mb-4">
                <h2 class="text-xl font-bold">Purchase Order Details</h2>
                <button
                  @click="closeDetailsModal"
                  class="text-gray-500 hover:text-gray-700 text-xl"
                >
                  &times;
                </button>
              </div>
              <div v-if="selectedPurchase" class="space-y-4">
                <p>
                  <strong>PO Number:</strong>
                  {{ formatPONumber(selectedPurchase.ID) }}
                </p>
                <p>
                  <strong>Supplier:</strong>
                  {{ selectedPurchase.Supplier?.name || "-" }}
                </p>
                <p>
                  <strong>Date:</strong> {{ formatDate(selectedPurchase.Date) }}
                </p>
                <p>
                  <strong>Total:</strong>
                  {{ formatCurrency(selectedPurchase.Total) }}
                </p>
                <p>
                  <strong>Status:</strong>
                  {{ getStatusText(selectedPurchase.Status) }}
                </p>
                <p>
                  <strong>Requested By:</strong> User ID
                  {{ selectedPurchase.UserID || "-" }}
                </p>
                <p v-if="selectedPurchase.ApprovedBy">
                  <strong>Approved By:</strong> User ID
                  {{ selectedPurchase.ApprovedBy }}
                </p>
                <p v-if="selectedPurchase.ApprovedAt">
                  <strong>Approved At:</strong>
                  {{ formatDate(selectedPurchase.ApprovedAt) }}
                </p>

                <h3 class="text-lg font-semibold mt-4">Purchase Items</h3>
                <div class="overflow-x-auto">
                  <table class="min-w-full bg-white border border-gray-300">
                    <thead>
                      <tr class="bg-gray-100">
                        <th class="px-4 py-2 text-left text-gray-600">
                          Product
                        </th>
                        <th class="px-4 py-2 text-left text-gray-600">
                          Quantity
                        </th>
                        <th class="px-4 py-2 text-left text-gray-600">Price</th>
                        <th class="px-4 py-2 text-left text-gray-600">
                          Subtotal
                        </th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr
                        v-for="(item, index) in selectedPurchase.Items"
                        :key="index"
                        class="border-t border-gray-300"
                      >
                        <td class="px-4 py-2 text-gray-700">
                          {{ item.Product?.name || "-" }}
                        </td>
                        <td class="px-4 py-2 text-gray-700">
                          {{ item.Qty || "-" }}
                        </td>
                        <td class="px-4 py-2 text-gray-700">
                          {{ formatCurrency(item.Price) }}
                        </td>
                        <td class="px-4 py-2 text-gray-700">
                          {{ formatCurrency(item.Subtotal) }}
                        </td>
                      </tr>
                      <tr
                        v-if="
                          !selectedPurchase.Items ||
                          selectedPurchase.Items.length === 0
                        "
                      >
                        <td
                          colspan="4"
                          class="px-4 py-2 text-center text-gray-700"
                        >
                          No items found
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
              <div v-else class="text-red-500">
                No data available for this purchase.
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<style scoped>
@container (max-width: 120px) {
  .table-column-120 {
    display: none;
  }
}
@container (max-width: 240px) {
  .table-column-240 {
    display: none;
  }
}
@container (max-width: 360px) {
  .table-column-360 {
    display: none;
  }
}
@container (max-width: 480px) {
  .table-column-480 {
    display: none;
  }
}
@container (max-width: 600px) {
  .table-column-600 {
    display: none;
  }
}
</style>
