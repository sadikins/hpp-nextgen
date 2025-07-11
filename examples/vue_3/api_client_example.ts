// examples/vue_3/api_client_example.ts
import axios, { AxiosInstance, AxiosResponse, AxiosError } from 'axios';
import { useAuthStore } from '@/stores/auth'; // Sesuaikan path Pinia store Anda
import router from '@/router'; // Pastikan Vue Router terimport jika digunakan

const apiClient: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:3000/api/v1',
  headers: {
    'Content-Type': 'application/json',
  },
});

// Interceptor untuk menambahkan token JWT ke setiap permintaan
apiClient.interceptors.request.use(
  (config) => {
    const authStore = useAuthStore();
    if (authStore.token) {
      config.headers.Authorization = `Bearer ${authStore.token}`;
    }
    return config;
  },
  (error: AxiosError) => {
    return Promise.reject(error);
  }
);

// Interceptor untuk menangani error respons global (misalnya, 401 Unauthorized)
apiClient.interceptors.response.use(
  (response: AxiosResponse) => response,
  (error: AxiosError) => {
    if (error.response && error.response.status === 401) {
      const authStore = useAuthStore();
      authStore.logout(); // Otomatis logout jika token tidak valid
      // Opsional: arahkan ke halaman login
      if (router.currentRoute.value.path !== '/login') {
         router.push('/login');
      }
    }
    return Promise.reject(error);
  }
);

export default apiClient;