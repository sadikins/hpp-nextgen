<script setup lang="ts">
import CustomerList from "@/components/CustomerList.vue";
import CustomerForm from "@/components/CustomerForm.vue";
</script>

<template>
  <div>
    <h1>Manajemen Pelanggan</h1>
    <CustomerForm />
    <CustomerList />
  </div>
</template>

<style scoped>
h1 {
  text-align: center;
}
</style>
