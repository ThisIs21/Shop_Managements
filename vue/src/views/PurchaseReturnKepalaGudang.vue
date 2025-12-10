<template>
  <div
    class="flex size-full min-h-screen flex-col bg-white font-manrope lg:flex-row"
  >
    <!-- Sidebar tetap ada -->
    <KepalaGudangSidebar />

    <div class="flex flex-1 flex-col">
      <main class="flex-1 p-8">
        <div class="layout-content-container flex flex-col flex-1">
          <div class="mb-6 flex flex-wrap justify-between items-center gap-3">
            <div>
              <p
                class="text-[#111418] tracking-light text-[32px] font-bold leading-tight"
              >
                Purchase Return Requests
              </p>
              <p class="text-[#60758a] text-sm font-normal leading-normal">
                Review and approve or reject purchase return requests from
                suppliers.
              </p>
            </div>
          </div>

          <!-- Tabel Data -->
          <div class="px-4 py-3 @container">
            <div
              class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
            >
              <table class="flex-1">
                <thead>
                  <tr class="bg-white">
                    <th
                      class="px-4 py-3 text-left w-[120px] text-sm font-medium"
                    >
                      Return No.
                    </th>
                    <th
                      class="px-4 py-3 text-left w-[120px] text-sm font-medium"
                    >
                      Purchase No.
                    </th>
                    <th
                      class="px-4 py-3 text-left w-[150px] text-sm font-medium"
                    >
                      Date
                    </th>
                    <th
                      class="px-4 py-3 text-left w-[200px] text-sm font-medium"
                    >
                      Supplier
                    </th>
                    <th
                      class="px-4 py-3 text-left w-[200px] text-sm font-medium"
                    >
                      Product
                    </th>
                    <th
                      class="px-4 py-3 text-left w-[120px] text-sm font-medium"
                    >
                      Total
                    </th>
                    <th
                      class="px-4 py-3 text-left w-[100px] text-sm font-medium"
                    >
                      Status
                    </th>
                    <th
                      class="px-4 py-3 text-left w-[150px] text-sm font-medium"
                    >
                      Actions
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-if="loading">
                    <td colspan="8" class="text-center py-8 text-gray-500">
                      Loading returns...
                    </td>
                  </tr>
                  <tr v-else-if="purchaseReturns.length === 0">
                    <td colspan="8" class="text-center py-8 text-gray-500">
                      Tidak ada retur pending.
                    </td>
                  </tr>
                  <tr
                    v-else
                    v-for="item in purchaseReturns"
                    :key="item.ID"
                    class="border-t border-t-[#dbe0e6]"
                  >
                    <td class="px-4 py-2 text-sm">
                      RTN-{{ formatReturnNumber(item.ID) }}
                    </td>
                    <td class="px-4 py-2 text-sm">
                      {{
                        item.PurchaseID
                          ? `PO-${String(item.PurchaseID).padStart(3, "0")}`
                          : "N/A"
                      }}
                    </td>
                    <td class="px-4 py-2 text-sm">
                      {{ formatDate(item.Date) }}
                    </td>
                    <td class="px-4 py-2 text-sm">
                      {{ getSupplierName(item.Items) || "N/A" }}
                    </td>
                    <td class="px-4 py-2 text-sm">
                      {{ getFirstProductName(item.Items) }}
                    </td>
                    <td class="px-4 py-2 text-sm">
                      {{ formatCurrency(item.Total) }}
                    </td>
                    <td class="px-4 py-2 text-sm">
                      <span
                        :class="getStatusClass(item.Status)"
                        class="px-3 py-1 rounded-lg text-xs font-medium"
                      >
                        {{ getStatusText(item.Status) }}
                      </span>
                    </td>
                    <td class="px-4 py-2 text-sm font-bold">
                      <div v-if="item.Status === 'PENDING'" class="flex gap-2">
                        <button
                          @click="openApprovalModal(item.ID)"
                          class="px-3 py-1 bg-green-500 text-white rounded text-xs hover:bg-green-600"
                        >
                          Approve
                        </button>
                        <button
                          @click="openRejectModal(item.ID)"
                          class="px-3 py-1 bg-red-500 text-white rounded text-xs hover:bg-red-600"
                        >
                          Reject
                        </button>
                        <button
                          @click="viewDetails(item.ID)"
                          class="px-3 py-1 bg-blue-500 text-white rounded text-xs hover:bg-blue-600"
                        >
                          View Details
                        </button>
                      </div>
                      <div v-else class="flex gap-2">
                        <button
                          @click="viewDetails(item.ID)"
                          class="px-3 py-1 bg-blue-500 text-white rounded text-xs hover:bg-blue-600"
                        >
                          View Details
                        </button>
                        <span class="text-gray-500">Completed</span>
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>

        <!-- Approval Modal -->
        <div
          v-if="showApprovalModal"
          class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
        >
          <div class="bg-white rounded-lg p-6 max-w-md w-full">
            <h2 class="text-xl font-bold mb-4">Approve Purchase Return</h2>
            <div class="mb-4">
              <p>
                <strong>Return No:</strong> RTN-{{
                  formatReturnNumber(selectedReturnId)
                }}
              </p>
              <p>
                <strong>Purchase No:</strong>
                {{
                  selectedReturn?.PurchaseID
                    ? `PO-${String(selectedReturn.PurchaseID).padStart(3, "0")}`
                    : "N/A"
                }}
              </p>
              <p>
                <strong>Supplier:</strong>
                {{ getSupplierName(selectedReturn?.Items) || "N/A" }}
              </p>
              <p>
                <strong>Total:</strong>
                {{ formatCurrency(selectedReturn?.Total || 0) }}
              </p>
              <p>
                <strong>Requested By:</strong> User ID
                {{ selectedReturn?.UserID || "N/A" }}
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

        <!-- Reject Modal -->
        <div
          v-if="showRejectModal"
          class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
        >
          <div class="bg-white rounded-lg p-6 max-w-md w-full">
            <h2 class="text-xl font-bold mb-4">Reject Purchase Return</h2>
            <div class="mb-4">
              <p>
                <strong>Return No:</strong> RTN-{{
                  formatReturnNumber(selectedReturnId)
                }}
              </p>
              <p>
                <strong>Purchase No:</strong>
                {{
                  selectedReturn?.PurchaseID
                    ? `PO-${String(selectedReturn.PurchaseID).padStart(3, "0")}`
                    : "N/A"
                }}
              </p>
              <p>
                <strong>Supplier:</strong>
                {{ getSupplierName(selectedReturn?.Items) || "N/A" }}
              </p>
              <p>
                <strong>Total:</strong>
                {{ formatCurrency(selectedReturn?.Total || 0) }}
              </p>
              <p>
                <strong>Requested By:</strong> User ID
                {{ selectedReturn?.UserID || "N/A" }}
              </p>
            </div>
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2"
                >Alasan Penolakan</label
              >
              <textarea
                v-model="rejectReason"
                placeholder="Masukkan alasan penolakan..."
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

        <!-- View Details Modal -->
        <div
          v-if="showDetailsModal"
          class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
        >
          <div
            class="bg-white rounded-lg p-6 max-w-2xl w-full max-h-[90vh] overflow-y-auto"
          >
            <div class="flex justify-between items-center mb-4">
              <h2 class="text-xl font-bold">Return Details</h2>
              <button
                @click="closeDetailsModal"
                class="text-gray-500 hover:text-gray-700 text-xl"
              >
                &times;
              </button>
            </div>
            <div v-if="selectedReturn" class="space-y-4">
              <p>
                <strong>Return No:</strong> RTN-{{
                  formatReturnNumber(selectedReturnId)
                }}
              </p>
              <p>
                <strong>Purchase No:</strong>
                {{
                  selectedReturn?.PurchaseID
                    ? `PO-${String(selectedReturn.PurchaseID).padStart(3, "0")}`
                    : "N/A"
                }}
              </p>
              <p>
                <strong>Supplier:</strong>
                {{ getSupplierName(selectedReturn?.Items) || "N/A" }}
              </p>
              <p>
                <strong>Date:</strong> {{ formatDate(selectedReturn.Date) }}
              </p>
              <p>
                <strong>Total:</strong>
                {{ formatCurrency(selectedReturn.Total) }}
              </p>
              <p>
                <strong>Requested By:</strong> User ID
                {{ selectedReturn.UserID || "N/A" }}
              </p>
              <p v-if="selectedReturn.Notes">
                <strong>Additional Notes:</strong> {{ selectedReturn.Notes }}
              </p>

              <h3 class="text-lg font-semibold mt-4">Return Items</h3>
              <div class="overflow-x-auto">
                <table class="min-w-full bg-white border border-gray-300">
                  <thead>
                    <tr class="bg-gray-100">
                      <th class="px-4 py-2 text-left text-gray-600">Product</th>
                      <th class="px-4 py-2 text-left text-gray-600">Qty</th>
                      <th class="px-4 py-2 text-left text-gray-600">Price</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr
                      v-for="(item, index) in selectedReturn.Items"
                      :key="index"
                      class="border-t border-gray-300"
                    >
                      <td class="px-4 py-2 text-gray-700">
                        {{ item.Product?.name || item.product_name || "N/A" }}
                      </td>
                      <td class="px-4 py-2 text-gray-700">{{ item.Qty }}</td>
                      <td class="px-4 py-2 text-gray-700">
                        {{ formatCurrency(item.Price) }}
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>

        <!-- History Modal -->
        <div
          v-if="showHistoryModal"
          class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4"
        >
          <div
            class="bg-white rounded-lg p-6 max-w-2xl w-full max-h-[90vh] overflow-y-auto"
          >
            <div class="flex justify-between items-center mb-4">
              <h2 class="text-xl font-bold">Purchase Return History</h2>
              <button
                @click="closeHistoryModal"
                class="text-gray-500 hover:text-gray-700 text-xl"
              >
                &times;
              </button>
            </div>
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2"
                >Filter by Status</label
              >
              <select
                v-model="historyFilter"
                class="w-full px-3 py-2 border border-gray-300 rounded-md"
              >
                <option value="">All</option>
                <option value="APPROVED">Approved</option>
                <option value="REJECTED">Rejected</option>
              </select>
            </div>
            <div class="overflow-x-auto">
              <table class="min-w-full bg-white border border-gray-300">
                <thead>
                  <tr class="bg-gray-100">
                    <th class="px-4 py-2 text-left text-gray-600">
                      Return No.
                    </th>
                    <th class="px-4 py-2 text-left text-gray-600">
                      Purchase No.
                    </th>
                    <th class="px-4 py-2 text-left text-gray-600">Date</th>
                    <th class="px-4 py-2 text-left text-gray-600">Supplier</th>
                    <th class="px-4 py-2 text-left text-gray-600">Total</th>
                    <th class="px-4 py-2 text-left text-gray-600">Status</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-if="loadingHistory">
                    <td colspan="6" class="text-center py-8 text-gray-500">
                      Loading history...
                    </td>
                  </tr>
                  <tr v-else-if="historyReturns.length === 0">
                    <td colspan="6" class="text-center py-8 text-gray-500">
                      No history available.
                    </td>
                  </tr>
                  <tr
                    v-else
                    v-for="item in historyReturns"
                    :key="item.id"
                    class="border-t border-gray-300"
                  >
                    <td class="px-4 py-2 text-sm">
                      RTN-{{ formatReturnNumber(item.id) }}
                    </td>
                    <td class="px-4 py-2 text-sm">
                      {{
                        item.purchase_id
                          ? `PO-${String(item.purchase_id).padStart(3, "0")}`
                          : "N/A"
                      }}
                    </td>
                    <td class="px-4 py-2 text-sm">
                      {{ formatDate(item.date) }}
                    </td>
                    <td class="px-4 py-2 text-sm">
                      {{ getSupplierNameFromHistory(item) || "N/A" }}
                    </td>
                    <td class="px-4 py-2 text-sm">
                      {{ formatCurrency(item.total) }}
                    </td>
                    <td class="px-4 py-2 text-sm">
                      <span
                        :class="getStatusClass(item.status)"
                        class="px-3 py-1 rounded-lg text-xs font-medium"
                      >
                        {{ getStatusText(item.status) }}
                      </span>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from "vue";
import axios from "axios";
import KepalaGudangSidebar from "@/components/KepalaGudangSidebar.vue";

// State
const purchaseReturns = ref([]);
const loading = ref(false);
const showApprovalModal = ref(false);
const showRejectModal = ref(false);
const showDetailsModal = ref(false);
const showHistoryModal = ref(false);
const selectedReturnId = ref(null);
const selectedReturn = ref(null);
const rejectReason = ref("");
const approving = ref(false);
const rejecting = ref(false);
const loadingHistory = ref(false);
const historyReturns = ref([]);
const historyFilter = ref("");

// Ambil role dari localStorage
const role = localStorage.getItem("role") || "kepalagudang";

// Mapping role ke path backend
const endpointMap = {
  kepalagudang: "kepala-gudang",
  owner: "owner",
  admin: "admin",
};

// Helper buat dapetin base URL sesuai role
const getBaseUrl = () => {
  const path = endpointMap[role] || "kepala-gudang";
  return `http://localhost:8081/api/v1/${path}`;
};

// Inisialisasi axios dengan baseURL
const api = axios.create({ baseURL: "http://localhost:8081/api/v1" });
const savedToken = localStorage.getItem("token");
if (savedToken) {
  api.defaults.headers.common["Authorization"] = `Bearer ${savedToken}`;
}

// Fetch data dari backend sesuai role
const fetchPurchaseReturns = async (status = "PENDING") => {
  loading.value = true;
  try {
    const url = `${getBaseUrl()}/purchase-returns`;
    const res = await api.get(url, {
      params: { status, page: 1, size: 10 },
    });

    console.log("Respon backend Purchase Returns:", res.data);

    if (res.data.data && Array.isArray(res.data.data)) {
      purchaseReturns.value = res.data.data;
    } else {
      purchaseReturns.value = [];
    }

    console.log("Purchase Returns loaded:", purchaseReturns.value);
  } catch (err) {
    console.error(
      "Gagal ambil data Purchase Returns:",
      err.response?.data || err
    );
    alert(
      `Gagal mengambil data retur: ${err.response?.status || ""} - ${
        err.response?.data?.error || err.message
      }`
    );
    purchaseReturns.value = [];
  } finally {
    loading.value = false;
  }
};

const fetchPurchaseReturnHistory = async () => {
  loadingHistory.value = true;
  try {
    const url = `${getBaseUrl()}/purchase-returns/history`;
    const res = await api.get(url, {
      params: { status: historyFilter.value, page: 1, size: 10 },
    });
    if (res.data.data && Array.isArray(res.data.data)) {
      historyReturns.value = res.data.data;
    } else {
      historyReturns.value = [];
    }
    console.log("Purchase Return History loaded:", historyReturns.value);
  } catch (err) {
    console.error(
      "Gagal ambil data Purchase Return History:",
      err.response?.data || err
    );
    alert(
      `Gagal mengambil data riwayat retur: ${err.response?.status || ""} - ${
        err.response?.data?.error || err.message
      }`
    );
    historyReturns.value = [];
  } finally {
    loadingHistory.value = false;
  }
};

// Open Approval Modal
const openApprovalModal = async (id) => {
  try {
    const url = `${getBaseUrl()}/purchase-returns/${id}`;
    const response = await api.get(url, {
      params: { include_approver: true },
    });
    selectedReturn.value = response.data;
    selectedReturnId.value = id;
    showApprovalModal.value = true;
  } catch (error) {
    console.error(
      "Error fetching purchase return:",
      error.response ? error.response.data : error.message
    );
    alert(
      "Gagal memuat detail retur: " +
        (error.response?.data?.error || error.message)
    );
  }
};

// Close Approval Modal
const closeApprovalModal = () => {
  showApprovalModal.value = false;
  selectedReturn.value = null;
  selectedReturnId.value = null;
};

// Confirm Approve
const confirmApprove = async () => {
  if (!selectedReturnId.value) return;

  approving.value = true;
  try {
    const url = `${getBaseUrl()}/purchase-returns/${
      selectedReturnId.value
    }/approve`;
    const response = await api.post(url, {});
    if (response.status === 200) {
      alert("Retur berhasil disetujui!");
      closeApprovalModal();
      fetchPurchaseReturns();
    }
  } catch (error) {
    console.error(
      "Error approving purchase return:",
      error.response ? error.response.data : error.message
    );
    alert(
      "Gagal menyetujui retur: " +
        (error.response?.data?.error || error.message)
    );
  } finally {
    approving.value = false;
  }
};

// Open Reject Modal
const openRejectModal = async (id) => {
  try {
    const url = `${getBaseUrl()}/purchase-returns/${id}`;
    const response = await api.get(url, {
      params: { include_approver: true },
    });
    selectedReturn.value = response.data;
    selectedReturnId.value = id;
    rejectReason.value = "";
    showRejectModal.value = true;
  } catch (error) {
    console.error(
      "Error fetching purchase return:",
      error.response ? error.response.data : error.message
    );
    alert(
      "Gagal memuat detail retur: " +
        (error.response?.data?.error || error.message)
    );
  }
};

// Close Reject Modal
const closeRejectModal = () => {
  showRejectModal.value = false;
  selectedReturn.value = null;
  selectedReturnId.value = null;
  rejectReason.value = "";
};

// Confirm Reject
const confirmReject = async () => {
  if (!selectedReturnId.value || !rejectReason.value.trim()) {
    alert("Alasan penolakan wajib diisi!");
    return;
  }

  rejecting.value = true;
  try {
    const url = `${getBaseUrl()}/purchase-returns/${
      selectedReturnId.value
    }/reject`;
    const response = await api.post(url, {
      reason: rejectReason.value.trim(),
    });
    if (response.status === 200) {
      alert("Retur berhasil ditolak!");
      closeRejectModal();
      fetchPurchaseReturns();
    }
  } catch (error) {
    console.error(
      "Error rejecting purchase return:",
      error.response ? error.response.data : error.message
    );
    alert(
      "Gagal menolak retur: " + (error.response?.data?.error || error.message)
    );
  } finally {
    rejecting.value = false;
  }
};

// View Details
const viewDetails = async (id) => {
  try {
    const url = `${getBaseUrl()}/purchase-returns/${id}`;
    const response = await api.get(url, {
      params: { include_approver: true },
    });
    selectedReturn.value = response.data;
    selectedReturnId.value = id;
    showDetailsModal.value = true;
  } catch (error) {
    console.error(
      "Error fetching purchase return details:",
      error.response ? error.response.data : error.message
    );
    alert(
      "Gagal memuat detail retur: " +
        (error.response?.data?.error || error.message)
    );
  }
};

// Close Details Modal
const closeDetailsModal = () => {
  showDetailsModal.value = false;
  selectedReturn.value = null;
  selectedReturnId.value = null;
};

// Close History Modal
const closeHistoryModal = () => {
  showHistoryModal.value = false;
  historyReturns.value = [];
  historyFilter.value = "";
};

// Watch filter changes
watch(historyFilter, () => {
  fetchPurchaseReturnHistory();
});

// Format Return Number
const formatReturnNumber = (id) => {
  return `2025-${String(id).padStart(3, "0")}`;
};

// Get Status Class
const getStatusClass = (status) => {
  switch (status) {
    case "PENDING":
      return "bg-yellow-200 text-yellow-800";
    case "APPROVED":
      return "bg-green-200 text-green-800";
    case "REJECTED":
      return "bg-red-200 text-red-800";
    default:
      return "bg-gray-200 text-gray-800";
  }
};

// Get Status Text
const getStatusText = (status) => {
  switch (status) {
    case "PENDING":
      return "Pending";
    case "APPROVED":
      return "Approved";
    case "REJECTED":
      return "Rejected";
    default:
      return status;
  }
};

// Get First Product Name
const getFirstProductName = (items) => {
  if (items && items.length > 0) {
    return items[0].Product?.name || items[0].product_name || "N/A";
  }
  return "N/A";
};

// Get Supplier Name from Items
const getSupplierName = (items) => {
  if (items && items.length > 0) {
    return items[0].Product?.supplier?.name || "N/A";
  }
  return "N/A";
};

// Get Supplier Name from History
const getSupplierNameFromHistory = (item) => {
  return item.supplier_id ? "Supplier ID: " + item.supplier_id : "N/A";
};

// Format Date
const formatDate = (dateStr) => {
  if (!dateStr) return "N/A";
  const date = new Date(dateStr);
  return date.toLocaleDateString("id-ID", {
    day: "2-digit",
    month: "2-digit",
    year: "numeric",
  });
};

// Format Currency
const formatCurrency = (amount) => {
  return amount !== undefined && amount !== null
    ? new Intl.NumberFormat("id-ID", {
        style: "currency",
        currency: "IDR",
        minimumFractionDigits: 0,
      }).format(amount)
    : "Rp 0";
};

onMounted(() => {
  fetchPurchaseReturns();
});
</script>

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
@container (max-width: 720px) {
  .table-column-720 {
    display: none;
  }
}
@container (max-width: 840px) {
  .table-column-840 {
    display: none;
  }
}
</style>
