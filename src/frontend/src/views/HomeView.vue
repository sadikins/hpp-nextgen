<script setup lang="ts">
import { computed } from "vue";
import { getOrders } from "@/api/ordersApi";
import { getCustomers } from "@/api/customersApi";

const totalOrders = computed(() => getOrders().length);
const totalCustomers = computed(() => getCustomers().length);
</script>

<template>
  <div class="home">
    <h1>Selamat Datang di Laundry Management</h1>
    <div class="stats">
      <div class="card">
        <h2>Total Pesanan</h2>
        <p>{{ totalOrders }}</p>
      </div>
      <div class="card">
        <h2>Total Pelanggan</h2>
        <p>{{ totalCustomers }}</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.home {
  text-align: center;
}
.stats {
  display: flex;
  justify-content: center;
  gap: 20px;
  margin-top: 20px;
}
.card {
  background-color: #6e62e5;
  color: white;
  padding: 20px;
  border-radius: 10px;
  min-width: 150px;
}
</style>
