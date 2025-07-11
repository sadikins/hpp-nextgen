<script setup lang="ts">
import { onMounted, ref, watch } from "vue";
import { getOrders, deleteOrder, addOrder } from "@/api/ordersApi";
import type { Order } from "@/api/ordersApi";
import { eventBus } from "@/event";

const orders = ref<Order[]>([]);

const loadOrders = async () => {
  orders.value = await getOrders(); // Ambil data order dari API
};

const createOrder = async (newOrder: Order) => {
  const addedOrder = await addOrder(newOrder); // Tambahkan order ke database
  orders.value.push(addedOrder); // Tambahkan data baru ke state tanpa replace array
};

onMounted(() => {
  eventBus.on("order-added", (order: Order) => {
    orders.value.push(order);
  });
});
</script>

<template>
  <div>
    <h2>Daftar Pesanan</h2>
    <table class="w-100 border-collapse">
      <thead>
        <tr>
          <th class="border border-amber-200 p-8">Nama Pelanggan</th>
          <th class="border border-amber-200 p-8">Status</th>
          <th class="border border-amber-200 p-8">Harga</th>
          <th class="border border-amber-200 p-8">Aksi</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="order in orders" :key="order.id">
          <td class="border border-amber-200 p-8">{{ order.customerName }}</td>
          <td class="border border-amber-200 p-8">{{ order.status }}</td>
          <td class="border border-amber-200 p-8">{{ order.price }}</td>
          <td class="border border-amber-200 p-8">
            <button
              @click="removeOrder(order.id)"
              class="cursor-pointer border-none bg-red-500 px-10 py-5 text-white"
            >
              Hapus
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
