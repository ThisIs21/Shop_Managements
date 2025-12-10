<template>
  <KepalaGudangLayouts>
    <div class="flex flex-wrap justify-between gap-3 p-4">
      <p
        class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
      >
        Warehouse Manager Dashboard
      </p>
    </div>

    <div class="p-4">
      <div
        class="flex flex-col gap-2 rounded-lg p-6 bg-[#f0f2f5] min-h-[150px] justify-center items-center"
      >
        <p
          class="text-[#111418] tracking-light text-3xl font-bold leading-tight"
        >
          Selamat datang, Kepala Gudang!
        </p>
      </div>
    </div>
  </KepalaGudangLayouts>
</template>

<script setup>
import KepalaGudangLayouts from "@/components/KepalaGudangLayouts.vue";

// Semua logika (ref, computed, axios, onMounted, fetchDashboardData) telah dihapus
// agar dashboard menjadi STATIS dan tidak ada lagi panggilan yang menghasilkan error.
</script>
