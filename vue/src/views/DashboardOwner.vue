<template>
  <OwnerLayouts>
    <div class="p-8">
      <div class="mb-6">
        <h1
          class="text-[#111418] text-[32px] font-bold leading-tight tracking-[-0.015em]"
        >
          Dashboard Owner
        </h1>
        <p class="text-[#60758a] text-sm font-normal leading-normal mt-1">
          Selamat datang kembali! Ini adalah ringkasan aktivitas toko Anda.
        </p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-8">
        <div class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6]">
          <p class="text-[#111418] text-base font-medium leading-normal">
            Total Penjualan
          </p>
          <p class="text-[#111418] text-2xl font-bold leading-tight">
            Rp 12.500.000
          </p>
        </div>
        <div class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6]">
          <p class="text-[#111418] text-base font-medium leading-normal">
            Total Pembelian
          </p>
          <p class="text-[#111418] text-2xl font-bold leading-tight">
            Rp 8.750.000
          </p>
        </div>
        <div class="flex flex-col gap-2 rounded-lg p-6 border border-[#dbe0e6]">
          <p class="text-[#111418] text-base font-medium leading-normal">
            Total Retur
          </p>
          <p class="text-[#111418] text-2xl font-bold leading-tight">
            Rp 500.000
          </p>
        </div>
      </div>
    </div>
  </OwnerLayouts>
</template>

<script setup>
import OwnerLayouts from "../components/OwnerLayouts.vue";
</script>

<style scoped>
/* Aturan container query dihapus */
</style>
