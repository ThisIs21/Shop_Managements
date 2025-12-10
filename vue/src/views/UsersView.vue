<template>
  <div class="users-page">
    <h2>Daftar Pengguna</h2>
    <UsersList />
  </div>
</template>

<script>
import UsersList from "@/components/ProductList.vue";

export default {
  name: "UsersView",
  components: {
    UsersList,
  },
};
</script>

<style scoped>
.users-page {
  padding: 20px;
}
</style>
