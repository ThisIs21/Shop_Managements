<template>
  <OwnerLayouts>
    <div class="flex flex-1 flex-col py-5 p-8">
      <div class="layout-content-container flex flex-col flex-1">
        <div class="flex flex-wrap justify-between gap-3 p-4">
          <p
            class="text-[#111418] tracking-light text-[32px] font-bold leading-tight min-w-72"
          >
            Reports Dashboard
          </p>
        </div>

        <h2
          class="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5"
        >
          Sales Reports
        </h2>
        <div
          class="grid grid-cols-[repeat(auto-fit,minmax(158px,1fr))] gap-3 p-4"
        >
          <RouterLink
            to="/reports-sales-summary"
            class="flex flex-1 gap-3 rounded-lg border border-[#dbe0e6] bg-white p-4 flex-col hover:bg-gray-50 transition-colors cursor-pointer"
          >
            <div
              class="text-[#111418]"
              data-icon="PresentationChart"
              data-size="24px"
              data-weight="regular"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24px"
                height="24px"
                fill="currentColor"
                viewBox="0 0 256 256"
              >
                <path
                  d="M216,40H136V24a8,8,0,0,0-16,0V40H40A16,16,0,0,0,24,56V176a16,16,0,0,0,16,16H79.36L57.75,219a8,8,0,0,0,12.5,10l29.59-37h56.32l29.59,37a8,8,0,1,0,12.5-10l-21.61-27H216a16,16,0,0,0,16-16V56A16,16,0,0,0,216,40Zm0,136H40V56H216V176Zm-112-56v24a8,8,0,0,1-16,0V120a8,8,0,0,1,16,0Zm32-16v40a8,8,0,0,1-16,0V104a8,8,0,0,1,16,0Zm32-16v56a8,8,0,0,1-16,0V88a8,8,0,0,1,16,0Z"
                ></path>
              </svg>
            </div>
            <div class="flex flex-col gap-1">
              <h2 class="text-[#111418] text-base font-bold leading-tight">
                Sales Summary
              </h2>
              <p class="text-[#60758a] text-sm font-normal leading-normal">
                Overview of total sales, revenue, and profit within a specified
                period.
              </p>
            </div>
          </RouterLink>
        </div>

        <h2
          class="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5"
        >
          Purchase Reports
        </h2>
        <div
          class="grid grid-cols-[repeat(auto-fit,minmax(158px,1fr))] gap-3 p-4"
        >
          <RouterLink
            to="/reports-purchase-summary"
            class="flex flex-1 gap-3 rounded-lg border border-[#dbe0e6] bg-white p-4 flex-col hover:bg-gray-50 transition-colors cursor-pointer"
          >
            <div
              class="text-[#111418]"
              data-icon="ShoppingCart"
              data-size="24px"
              data-weight="regular"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24px"
                height="24px"
                fill="currentColor"
                viewBox="0 0 256 256"
              >
                <path
                  d="M222.14,58.87A8,8,0,0,0,216,56H54.68L49.79,29.14A16,16,0,0,0,34.05,16H16a8,8,0,0,0,0,16h18L59.56,172.29a24,24,0,0,0,5.33,11.27,28,28,0,1,0,44.4,8.44h45.42A27.75,27.75,0,0,0,152,204a28,28,0,1,0,28-28H83.17a8,8,0,0,1-7.87-6.57L72.13,152h116a24,24,0,0,0,23.61-19.71l12.16-66.86A8,8,0,0,0,222.14,58.87ZM96,204a12,12,0,1,1-12-12A12,12,0,0,1,96,204Zm96,0a12,12,0,1,1-12-12A12,12,0,0,1,192,204Zm4-74.57A8,8,0,0,1,188.1,136H69.22L57.59,72H206.41Z"
                ></path>
              </svg>
            </div>
            <div class="flex flex-col gap-1">
              <h2 class="text-[#111418] text-base font-bold leading-tight">
                Purchase Summary
              </h2>
              <p class="text-[#60758a] text-sm font-normal leading-normal">
                Summary of all purchases made, including total cost and supplier
                details.
              </p>
            </div>
          </RouterLink>
        </div>

        <h2
          class="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5"
        >
          Return Reports
        </h2>
        <div
          class="grid grid-cols-[repeat(auto-fit,minmax(158px,1fr))] gap-3 p-4"
        >
          <RouterLink
            to="/reports-returns-sale"
            class="flex flex-1 gap-3 rounded-lg border border-[#dbe0e6] bg-white p-4 flex-col hover:bg-gray-50 transition-colors cursor-pointer"
          >
            <div
              class="text-[#111418]"
              data-icon="ArrowULeftUp"
              data-size="24px"
              data-weight="regular"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24px"
                height="24px"
                fill="currentColor"
                viewBox="0 0 256 256"
              >
                <path
                  d="M208,80v88a64,64,0,0,1-128,0V51.31L45.66,85.66A8,8,0,0,1,34.34,74.34l48-48a8,8,0,0,1,11.32,0l48,48a8,8,0,0,1-11.32,11.32L96,51.31V168a48,48,0,0,0,96,0V80a8,8,0,0,1,16,0Z"
                ></path>
              </svg>
            </div>
            <div class="flex flex-col gap-1">
              <h2 class="text-[#111418] text-base font-bold leading-tight">
                Sale Return Summary
              </h2>
              <p class="text-[#60758a] text-sm font-normal leading-normal">
                Details of products returned by customers.
              </p>
            </div>
          </RouterLink>
          <RouterLink
            to="/reports-returns-purchase"
            class="flex flex-1 gap-3 rounded-lg border border-[#dbe0e6] bg-white p-4 flex-col hover:bg-gray-50 transition-colors cursor-pointer"
          >
            <div class="flex flex-col gap-1">
              <h2 class="text-[#111418] text-base font-bold leading-tight">
                Purchase Return Summary
              </h2>
              <p class="text-[#60758a] text-sm font-normal leading-normal">
                Details of products returned to suppliers.
              </p>
            </div>
          </RouterLink>
        </div>

        <h2
          class="text-[#111418] text-[22px] font-bold leading-tight tracking-[-0.015em] px-4 pb-3 pt-5"
        >
          Inventory Reports
        </h2>
        <div
          class="grid grid-cols-[repeat(auto-fit,minmax(158px,1fr))] gap-3 p-4"
        >
          <RouterLink
            to="/report-stock-opname"
            class="flex flex-1 gap-3 rounded-lg border border-[#dbe0e6] bg-white p-4 flex-col hover:bg-gray-50 transition-colors cursor-pointer"
          >
            <div
              class="text-[#111418]"
              data-icon="Package"
              data-size="24px"
              data-weight="regular"
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="24px"
                height="24px"
                fill="currentColor"
                viewBox="0 0 256 256"
              >
                <path
                  d="M223.68,66.15,135.68,18a15.88,15.88,0,0,0-15.36,0l-88,48.17a16,16,0,0,0-8.32,14v95.64a16,16,0,0,0,8.32,14l88,48.17a15.88,15.88,0,0,0,15.36,0l88-48.17a16,16,0,0,0,8.32-14V80.18A16,16,0,0,0,223.68,66.15ZM128,32l80.34,44-29.77,16.3-80.35-44ZM128,120,47.66,76l33.9-18.56,80.34,44ZM40,90l80,43.78v85.79L40,175.82Zm176,85.78h0l-80,43.79V133.82l32-17.51V152a8,8,0,0,0,16,0V107.55L216,90v85.77Z"
                ></path>
              </svg>
            </div>
            <div class="flex flex-col gap-1">
              <h2 class="text-[#111418] text-base font-bold leading-tight">
                Stock Opname Summary
              </h2>
              <p class="text-[#60758a] text-sm font-normal leading-normal">
                Summary of stock adjustments made during stocktaking.
              </p>
            </div>
          </RouterLink>
          <RouterLink
            to="/stocksnapshot"
            class="flex flex-1 gap-3 rounded-lg border border-[#dbe0e6] bg-white p-4 flex-col hover:bg-gray-50 transition-colors cursor-pointer"
          >
            <div class="flex flex-col gap-1">
              <h2 class="text-[#111418] text-base font-bold leading-tight">
                Stock Snapshot
              </h2>
              <p class="text-[#60758a] text-sm font-normal leading-normal">
                Current stock levels of all products.
              </p>
            </div>
          </RouterLink>
          <RouterLink
            to="/inventory-value"
            class="flex flex-1 gap-3 rounded-lg border border-[#dbe0e6] bg-white p-4 flex-col hover:bg-gray-50 transition-colors cursor-pointer"
          >
            <div class="flex flex-col gap-1">
              <h2 class="text-[#111418] text-base font-bold leading-tight">
                Inventory Value
              </h2>
              <p class="text-[#60758a] text-sm font-normal leading-normal">
                Total value of current inventory based on cost price.
              </p>
            </div>
          </RouterLink>
        </div>
      </div>
    </div>
  </OwnerLayouts>
</template>

<script setup>
import OwnerLayouts from "../components/OwnerLayouts.vue";
import { RouterLink } from "vue-router";
</script>
