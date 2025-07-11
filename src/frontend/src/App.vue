<script setup lang="ts"></script>

<template>
  <nav>
    <router-link to="/">Home</router-link> |
    <router-link to="/orders">Pesanan</router-link> |
    <router-link to="/customers">Pelanggan</router-link>
  </nav>
  <router-view />
</template>

<style scoped>
nav {
  background-color: #333;
  padding: 10px;
}
nav a {
  color: white;
  text-decoration: none;
  margin: 0 10px;
}
</style>
