<script setup lang="ts">
import { ref } from "vue";
import { addCustomer, Customer } from "@/api/customersApi";

const name = ref("");
const phone = ref("");

const submitCustomer = () => {
  const newCustomer: Customer = {
    id: Date.now(),
    name: name.value,
    phone: phone.value,
  };
  addCustomer(newCustomer);
  alert("Pelanggan berhasil ditambahkan!");
  name.value = "";
  phone.value = "";
};
</script>

<template>
  <form @submit.prevent="submitCustomer">
    <h2>Tambah Pelanggan</h2>
    <label>Nama:</label>
    <input v-model="name" required />

    <label>No. Telepon:</label>
    <input type="tel" v-model="phone" required />

    <button type="submit">Tambah</button>
  </form>
</template>

<style scoped>
form {
  display: flex;
  flex-direction: column;
  gap: 8px;
}
button {
  background-color: blue;
  color: white;
  padding: 10px;
  cursor: pointer;
}
</style>
