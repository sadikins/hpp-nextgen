<script setup lang="ts">
import OrderList from "@/components/OrderList.vue";
import OrderForm from "@/components/OrderForm.vue";
</script>

<template>
  <div>
    <h1>Manajemen Pesanan</h1>
    <OrderForm />
    <OrderList />
  </div>
</template>

<style scoped>
h1 {
  text-align: center;
}
</style>
