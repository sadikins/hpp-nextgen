<script setup lang="ts">
  import { useAuthStore } from '@/stores/auth'; // Sesuaikan path store
  import { onMounted } from 'vue';
  // import { useRouter } from 'vue-router'; // Opsional, jika Anda ingin navigasi programatik

  const authStore = useAuthStore();
  // const router = useRouter(); // Opsional

  // Memastikan user terautentikasi atau mencoba fetch user dari token yang ada
  onMounted(() => {
    if (!authStore.isAuthenticated && authStore.token) {
      authStore.fetchUser(); // Coba fetch user jika ada token tapi state user kosong
    }
    // Jika tidak terautentikasi dan tidak ada token, router guard akan menangani redirect ke /login
  });
  </script>

<template>
    <div class="min-h-screen bg-gray-100 flex flex-col">
      <nav class="bg-indigo-600 text-white p-4 shadow-md">
        <div class="container mx-auto flex justify-between items-center">
          <router-link to="/dashboard" class="text-2xl font-bold">HPPKu NextGen</router-link>
          <div class="space-x-4">
            <router-link to="/raw-materials" class="hover:text-indigo-200">Bahan Baku</router-link>
            <router-link to="/foods" class="hover:text-indigo-200">Makanan</router-link>
            <router-link to="/events" class="hover:text-indigo-200">Event Order</router-link>
            <router-link v-if="authStore.isAdmin" to="/users" class="hover:text-indigo-200">Manajemen User</router-link>
            <button v-if="authStore.isAuthenticated" @click="authStore.logout()" class="hover:text-indigo-200">Logout</button>
            <router-link v-else to="/login" class="hover:text-indigo-200">Login</router-link>
          </div>
        </div>
      </nav>

      <main class="flex-grow container mx-auto p-6">
        <router-view />
      </main>

      <footer class="bg-gray-800 text-white p-4 text-center">
        &copy; 2025 HPPKu NextGen. All rights reserved.
      </footer>
    </div>
  </template>

