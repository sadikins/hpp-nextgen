<script setup lang="ts">
import { ref, onMounted } from "vue";
import { getCustomers } from "@/api/customersApi";
import type { Customer } from "@/api/customersApi";

const customers = ref<Customer[]>([]);

const loadCustomers = () => {
  customers.value = getCustomers();
};

onMounted(loadCustomers);
</script>

<template>
  <div>
    <h2>Daftar Pelanggan</h2>
    <table>
      <thead>
        <tr>
          <th>Nama</th>
          <th>No. Telepon</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="customer in customers" :key="customer.id">
          <td>{{ customer.name }}</td>
          <td>{{ customer.phone }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
table {
  width: 100%;
  border-collapse: collapse;
  margin-top: 10px;
}
th,
td {
  border: 1px solid #ddd;
  padding: 8px;
}
</style>
