import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import OrdersView from "@/views/OrdersView.vue";
import CustomersView from "@/views/CustomersView.vue";
const routes = [
  { path: "/", component: HomeView },
  { path: "/orders", component: OrdersView },
  { path: "/customers", component: CustomersView },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
