<script setup lang="ts">
import { addOrder } from "@/api/ordersApi.ts";
import type { Order } from "@/api/ordersApi.ts";
import { ref } from "vue";
import { eventBus } from "@/event";

const customerName = ref("");
const status = ref<Order["status"]>("Dalam Proses");
const price = ref(0);

const createOrder = async (newOrder: Order) => {
  const addedOrder = await addOrder(newOrder);
  eventBus.emit("order-added", addedOrder);
};

const submitOrder = () => {
  const newOrder: Order = {
    id: Date.now(),
    customerName: customerName.value,
    status: status.value,
    price: price.value,
  };
  addOrder(newOrder);
  // alert("Pesanan berhasil ditambahkan");
  customerName.value = "";
  price.value = 0;
};
</script>

<template>
  <form @submit.prevent="submitOrder" class="flex flex-col gap-8">
    <h2>Tambah Pesanan</h2>
    <label for="nama_pelanggan">Nama Pelanggan:</label>
    <input v-model="customerName" required />

    <label for="status">Status:</label>
    <select v-model="status">
      <option value="Dalam Proses">Dalam Proses</option>
      <option value="Selesai">Selesai</option>
      <option value="Siap Diambil">Siap diambil</option>
    </select>

    <label for="harga">Harga:</label>

    <input type="number" v-model="price" required />

    <button
      type="submit"
      class="cursor-pointer border-none bg-green-500 p-10 text-white"
    >
      Tambah
    </button>
  </form>
</template>
