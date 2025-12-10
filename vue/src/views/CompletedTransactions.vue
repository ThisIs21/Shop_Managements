<template>
  <div
    class="relative flex size-full min-h-screen flex-col bg-white group/design-root overflow-x-hidden"
    style="
      --select-button-svg: url('data:image/svg+xml,%3csvg xmlns=%27http://www.w3.org/2000/svg%27 width=%2724px%27 height=%2724px%27 fill=%27rgb(96,117,138)%27 viewBox=%270 0 256 256%27%3e%3cpath d=%27M181.66,170.34a8,8,0,0,1,0,11.32l-48,48a8,8,0,0,1-11.32,0l-48-48a8,8,0,0,1,11.32-11.32L128,212.69l42.34-42.35A8,8,0,0,1,181.66,170.34Zm-96-84.68L128,43.31l42.34,42.35a8,8,0,0,0,11.32-11.32l-48-48a8,8,0,0,0-11.32,0l-48,48A8,8,0,0,0,85.66,85.66Z%27%3e%3c/path%3e%3c/svg%3e');
      font-family: Manrope, 'Noto Sans', sans-serif;
    "
  >
    <div class="layout-container flex h-full grow flex-col">
      <header
        class="fixed top-0 left-0 right-0 z-50 flex items-center justify-between whitespace-nowrap border-b border-solid border-b-[#f0f2f5] bg-white px-10 py-3"
      >
        <div class="flex items-center gap-4 text-[#111418]">
          <div class="size-4">
            <svg
              viewBox="0 0 48 48"
              fill="currentColor"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path d="M4 4H17.3334V17.3334H30.6666V30.6666H44V44H4V4Z" />
            </svg>
          </div>
          <h2
            class="text-[#111418] text-lg font-bold leading-tight tracking-[-0.015em]"
          >
            SalesApp
          </h2>
        </div>
        <div class="flex flex-1 justify-end gap-8">
          <div class="flex items-center gap-9">
            <a
              class="text-[#111418] text-sm font-medium leading-normal"
              href="#"
              >Dashboard</a
            >
            <a
              class="text-[#111418] text-sm font-medium leading-normal"
              href="#"
              >Products</a
            >
            <a
              class="text-[#111418] text-sm font-medium leading-normal"
              href="#"
              >Orders</a
            >
            <a
              class="text-[#111418] text-sm font-medium leading-normal"
              href="#"
              >Customers</a
            >
          </div>
          <button
            class="flex max-w-[480px] cursor-pointer items-center justify-center overflow-hidden rounded-lg h-10 bg-[#f0f2f5] text-[#111418] gap-2 text-sm font-bold leading-normal tracking-[0.015em] min-w-0 px-2.5"
          >
            <div
              class="text-[#111418]"
              data-icon="Bell"
              data-size="20px"
              data-weight="regular"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="20px"
                height="20px"
                fill="currentColor"
                viewBox="0 0 256 256"
              >
                <path
                  d="M221.8,175.94C216.25,166.38,208,139.33,208,104a80,80,0,1,0-160,0c0,35.34-8.26,62.38-13.81,71.94A16,16,0,0,0,48,200H88.81a40,40,0,0,0,78.38,0H208a16,16,0,0,0,13.8-24.06ZM128,216a24,24,0,0,1-22.62-16h45.24A24,24,0,0,1,128,216ZM48,184c7.7-13.24,16-43.92,16-80a64,64,0,1,1,128,0c0,36.05,8.28,66.73,16,80Z"
                />
              </svg>
            </div>
          </button>
          <div
            class="bg-center bg-no-repeat aspect-square bg-cover rounded-full size-10"
            style="
              background-image: url('https://lh3.googleusercontent.com/aida-public/AB6AXuCl99u8jlyb9iXtsIlBdNWkVYg3NndO9Jg4kDDDPQbP6YlZI71dx4zC2HEuFly2UHn8uvRLeHEwTaR_QkS6hhvUBhk8TCMtdcmjEBnUh4YOwE8KEg2mm24F9RkrPQazy2lSijOzX7KLuPofspli5Gtvi3X7XVpX3LSMoslqdd74OyPr3PlbeUFm86ejc47TMcpXckZkjz1Ayo_ATOEYokIszeSVAuXKu2l_Tebctt0X7lQIV1i3Qh3_s_Y_OoQBrflzer_mPIw-XZk');
            "
          ></div>
        </div>
      </header>

      <div class="px-40 flex flex-1 justify-center py-5 mt-16">
        <div
          class="layout-content-container flex flex-col max-w-[960px] flex-1"
        >
          <div class="flex flex-wrap justify-between gap-3 p-4">
            <div class="flex min-w-72 flex-col gap-3">
              <p
                class="text-[#111418] tracking-light text-[32px] font-bold leading-tight"
              >
                Completed Transactions
              </p>
              <p class="text-[#60758a] text-sm font-normal leading-normal">
                View all completed transactions with details such as time, total
                price, and payment method.
              </p>
            </div>
          </div>
          <div class="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
            <label class="flex flex-col min-w-40 flex-1">
              <p
                class="text-[#111418] text-base font-medium leading-normal pb-2"
              >
                Date
              </p>
              <div class="flex w-full flex-1 items-stretch rounded-lg">
                <input
                  type="date"
                  class="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dbe0e6] bg-white focus:border-[#dbe0e6] h-14 placeholder:text-[#60758a] p-[15px] text-base font-normal leading-normal"
                />
              </div>
            </label>
          </div>
          <div class="flex max-w-[480px] flex-wrap items-end gap-4 px-4 py-3">
            <label class="flex flex-col min-w-40 flex-1">
              <p
                class="text-[#111418] text-base font-medium leading-normal pb-2"
              >
                Shift
              </p>
              <select
                class="form-input flex w-full min-w-0 flex-1 resize-none overflow-hidden rounded-lg text-[#111418] focus:outline-0 focus:ring-0 border border-[#dbe0e6] bg-white focus:border-[#dbe0e6] h-14 p-[15px] text-base font-normal leading-normal"
              >
                <option value="one">Shift 1</option>
                <option value="two">Shift 2</option>
                <option value="three">Shift 3</option>
              </select>
            </label>
          </div>
          <div class="px-4 py-3 @container">
            <div
              class="flex overflow-hidden rounded-lg border border-[#dbe0e6] bg-white"
            >
              <table class="flex-1">
                <thead>
                  <tr class="bg-white">
                    <th
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-120 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal"
                    >
                      Time
                    </th>
                    <th
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-240 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal"
                    >
                      Total Price
                    </th>
                    <th
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-360 px-4 py-3 text-left text-[#111418] w-[400px] text-sm font-medium leading-normal"
                    >
                      Payment Method
                    </th>
                  </tr>
                </thead>
                <tbody>
                  <tr class="border-t border-t-[#dbe0e6]">
                    <td
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-120 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                    >
                      10:30 AM
                    </td>
                    <td
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-240 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                    >
                      $50.00
                    </td>
                    <td
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-360 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                    >
                      Cash
                    </td>
                  </tr>
                  <tr class="border-t border-t-[#dbe0e6]">
                    <td
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-120 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                    >
                      11:15 AM
                    </td>
                    <td
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-240 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                    >
                      $75.00
                    </td>
                    <td
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-360 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                    >
                      Credit Card
                    </td>
                  </tr>
                  <tr class="border-t border-t-[#dbe0e6]">
                    <td
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-120 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                    >
                      12:00 PM
                    </td>
                    <td
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-240 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                    >
                      $30.00
                    </td>
                    <td
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-360 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                    >
                      Mobile Payment
                    </td>
                  </tr>
                  <tr class="border-t border-t-[#dbe0e6]">
                    <td
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-120 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                    >
                      1:45 PM
                    </td>
                    <td
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-240 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                    >
                      $120.00
                    </td>
                    <td
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-360 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                    >
                      Cash
                    </td>
                  </tr>
                  <tr class="border-t border-t-[#dbe0e6]">
                    <td
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-120 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                    >
                      2:30 PM
                    </td>
                    <td
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-240 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                    >
                      $60.00
                    </td>
                    <td
                      class="table-e9436867-983a-4023-9c5b-73bb44eff294-column-360 h-[72px] px-4 py-2 w-[400px] text-[#60758a] text-sm font-normal leading-normal"
                    >
                      Credit Card
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Bagian script di sini bisa kosong atau diisi dengan logika data jika diperlukan
</script>

<style>
/* Blok style untuk responsivitas tabel */
@container (max-width: 120px) {
  .table-e9436867-983a-4023-9c5b-73bb44eff294-column-120 {
    display: none;
  }
}
@container (max-width: 240px) {
  .table-e9436867-983a-4023-9c5b-73bb44eff294-column-240 {
    display: none;
  }
}
@container (max-width: 360px) {
  .table-e9436867-983a-4023-9c5b-73bb44eff294-column-360 {
    display: none;
  }
}
</style>
