// examples/vue_3/pinia_store_example.ts
import { defineStore } from 'pinia';
import apiClient from '@/services/apiClient'; // Sesuaikan path API client Anda
import router from '@/router'; // Pastikan Vue Router terimport jika digunakan

// Interface untuk data pengguna
interface User {
  id: number;
  email: string;
  username: string;
  role: string; // 'admin' | 'user'
  // tambahkan properti user lainnya sesuai API
}

// Interface untuk state store autentikasi
interface AuthState {
  user: User | null;
  token: string | null;
  loading: boolean;
  error: string | null;
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    user: null,
    token: localStorage.getItem('authToken'), // Ambil token dari local storage
    loading: false,
    error: null,
  }),
  getters: {
    isAuthenticated: (state): boolean => !!state.token,
    isAdmin: (state): boolean => state.user?.role === 'admin',
  },
  actions: {
    async login(credentials: { email?: string; username?: string; password: string }) {
      this.loading = true;
      this.error = null;
      try {
        const response = await apiClient.post<{ token: string; user: User }>('/auth/login', credentials);
        this.token = response.data.token;
        this.user = response.data.user;
        localStorage.setItem('authToken', this.token); // Simpan token
        router.push('/dashboard'); // Arahkan ke dashboard setelah login
      } catch (err: any) {
        this.error = err.response?.data?.error || 'Login failed';
        this.user = null;
        this.token = null;
        localStorage.removeItem('authToken');
      } finally {
        this.loading = false;
      }
    },
    logout() {
      this.user = null;
      this.token = null;
      localStorage.removeItem('authToken');
      router.push('/login'); // Arahkan ke halaman login
    },
    // Aksi untuk memuat ulang detail pengguna jika diperlukan (misalnya dari token)
    async fetchUser() {
        if (!this.token) {
            this.user = null;
            return;
        }
        try {
            const response = await apiClient.get<{ user: User }>('/auth/me'); // Endpoint untuk mendapatkan detail user
            this.user = response.data.user;
        } catch (err) {
            console.error('Failed to fetch user:', err);
            this.logout(); // Logout jika gagal fetch user (token mungkin invalid)
        }
    }
  },
});