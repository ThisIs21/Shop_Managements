<template>
  <OwnerLayouts>
    <div class="flex flex-1 flex-col py-5 p-8">
      <div class="layout-content-container flex flex-col flex-1">
        <div class="flex flex-wrap justify-between items-center gap-3 p-4">
          <p
            class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
          >
            Purchase Returns
          </p>
          <button
            class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-medium leading-normal"
          >
            <span class="truncate">New Return</span>
          </button>
        </div>

        <div class="p-4 @container">
          <div
            class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
          >
            <table class="flex-1">
              <thead>
                <tr class="bg-white">
                  <th
                    class="table-column-120 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal"
                  >
                    Return No.
                  </th>
                  <th
                    class="table-column-240 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal"
                  >
                    Date
                  </th>
                  <th
                    class="table-column-360 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal"
                  >
                    Supplier
                  </th>
                  <th
                    class="table-column-480 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal"
                  >
                    Product
                  </th>
                  <th
                    class="table-column-600 px-4 py-3 text-left text-[#111418] w-60 text-sm font-medium leading-normal"
                  >
                    Status
                  </th>
                  <th
                    class="table-column-720 px-4 py-3 text-left text-[#111418] w-60 text-[#60758a] text-sm font-medium leading-normal"
                  >
                    Actions
                  </th>
                </tr>
              </thead>
              <tbody>
                <tr class="border-t border-t-[#dbe0e6]">
                  <td
                    class="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#111418] text-sm font-normal leading-normal"
                  >
                    RTN-20240715-001
                  </td>
                  <td
                    class="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    2024-07-15
                  </td>
                  <td
                    class="table-column-360 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    Supplier A
                  </td>
                  <td
                    class="table-column-480 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    Product X
                  </td>
                  <td
                    class="table-column-600 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal"
                  >
                    <button
                      class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-medium leading-normal w-full"
                    >
                      <span class="truncate">Pending</span>
                    </button>
                  </td>
                  <td
                    class="table-column-720 h-[72px] px-4 py-2 w-60 text-[#60758a] text-sm font-bold leading-normal tracking-[0.015em]"
                  >
                    View Details
                  </td>
                </tr>
                <tr class="border-t border-t-[#dbe0e6]">
                  <td
                    class="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#111418] text-sm font-normal leading-normal"
                  >
                    RTN-20240714-002
                  </td>
                  <td
                    class="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    2024-07-14
                  </td>
                  <td
                    class="table-column-360 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    Supplier B
                  </td>
                  <td
                    class="table-column-480 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    Product Y
                  </td>
                  <td
                    class="table-column-600 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal"
                  >
                    <button
                      class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-medium leading-normal w-full"
                    >
                      <span class="truncate">Approved</span>
                    </button>
                  </td>
                  <td
                    class="table-column-720 h-[72px] px-4 py-2 w-60 text-[#60758a] text-sm font-bold leading-normal tracking-[0.015em]"
                  >
                    View Details
                  </td>
                </tr>
                <tr class="border-t border-t-[#dbe0e6]">
                  <td
                    class="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#111418] text-sm font-normal leading-normal"
                  >
                    RTN-20240713-003
                  </td>
                  <td
                    class="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    2024-07-13
                  </td>
                  <td
                    class="table-column-360 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    Supplier C
                  </td>
                  <td
                    class="table-column-480 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    Product Z
                  </td>
                  <td
                    class="table-column-600 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal"
                  >
                    <button
                      class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-medium leading-normal w-full"
                    >
                      <span class="truncate">Rejected</span>
                    </button>
                  </td>
                  <td
                    class="table-column-720 h-[72px] px-4 py-2 w-60 text-[#60758a] text-sm font-bold leading-normal tracking-[0.015em]"
                  >
                    View Details
                  </td>
                </tr>
                <tr class="border-t border-t-[#dbe0e6]">
                  <td
                    class="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#111418] text-sm font-normal leading-normal"
                  >
                    RTN-20240712-004
                  </td>
                  <td
                    class="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    2024-07-12
                  </td>
                  <td
                    class="table-column-360 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    Supplier A
                  </td>
                  <td
                    class="table-column-480 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    Product X
                  </td>
                  <td
                    class="table-column-600 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal"
                  >
                    <button
                      class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-medium leading-normal w-full"
                    >
                      <span class="truncate">Pending</span>
                    </button>
                  </td>
                  <td
                    class="table-column-720 h-[72px] px-4 py-2 w-60 text-[#60758a] text-sm font-bold leading-normal tracking-[0.015em]"
                  >
                    View Details
                  </td>
                </tr>
                <tr class="border-t border-t-[#dbe0e6]">
                  <td
                    class="table-column-120 h-[72px] px-4 py-2 w-[400px] text-[#111418] text-sm font-normal leading-normal"
                  >
                    RTN-20240711-005
                  </td>
                  <td
                    class="table-column-240 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    2024-07-11
                  </td>
                  <td
                    class="table-column-360 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    Supplier B
                  </td>
                  <td
                    class="table-column-480 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                  >
                    Product Y
                  </td>
                  <td
                    class="table-column-600 h-[72px] px-4 py-2 w-60 text-sm font-normal leading-normal"
                  >
                    <button
                      class="flex min-w-[84px] max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-8 px-4 bg-[#f0f2f5] text-[#111418] text-sm font-medium leading-normal w-full"
                    >
                      <span class="truncate">Approved</span>
                    </button>
                  </td>
                  <td
                    class="table-column-720 h-[72px] px-4 py-2 w-60 text-[#60758a] text-sm font-bold leading-normal tracking-[0.015em]"
                  >
                    View Details
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </OwnerLayouts>
</template>

<script setup>
import OwnerLayouts from "../components/OwnerLayouts.vue";
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
</style>
