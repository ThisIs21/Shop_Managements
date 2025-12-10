<template>
  <div
    class="relative flex size-full min-h-screen flex-col bg-white group/design-root overflow-x-hidden"
    style="font-family: Manrope, 'Noto Sans', sans-serif"
  >
    <div class="layout-container flex h-full grow flex-col">
      <div class="flex flex-1">
        <KepalaGudangSidebar />

        <main class="flex-1 overflow-y-auto">
          <slot></slot>
        </main>
      </div>
    </div>
  </div>
</template>

<script setup>
import KepalaGudangSidebar from "@/components/KepalaGudangSidebar.vue";
</script>

<style scoped>
/* Anda bisa menambahkan gaya khusus di sini jika diperlukan */
</style>
