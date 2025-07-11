# Implementasi Fitur Login Pengguna

## 🎯 Tujuan
Mengimplementasikan fungsionalitas login pengguna untuk aplikasi HPPKu NextGen. Pengguna harus dapat masuk ke sistem menggunakan kredensial (username/email dan password) yang valid. Backend akan memvalidasi kredensial, menghasilkan JSON Web Token (JWT), dan mengembalikannya ke frontend. Frontend akan menyimpan JWT dan mengarahkannya ke dashboard.

## 🎨 Inspirasi & Contoh
* **Backend (Go Fiber):**
    * Untuk struktur handler dan rute, lihat `examples/go_fiber/handler_example.go` dan bagaimana `main.go` mengatur rute.
    * Untuk interaksi database (validasi pengguna), lihat `examples/go_fiber/model_example.go` (struktur model user) dan `examples/go_fiber/service_example.go` (logika bisnis otentikasi).
    * Jika ada contoh implementasi JWT sebelumnya, referensikan di sini (misalnya, `examples/go_fiber/auth/jwt_utils.go`).
* **Frontend (Vue 3):**
    * Untuk komponen form login dasar, lihat pola di `examples/vue_3/component_form_example.vue`.
    * Untuk interaksi API dan state management, lihat `examples/vue_3/api_client_example.vue` dan `examples/vue_3/pinia_store_example.js` (untuk menyimpan token dan status autentikasi).

## 📚 Pengetahuan yang Dibutuhkan
* **Backend:**
    * **Go Fiber:** Pembuatan rute POST (`/api/v1/auth/login`), parsing body JSON (`github.com/gofiber/fiber/v2/middleware/recover`).
    * **GORM:** Mengambil data pengguna dari database PostgreSQL (`users` table).
    * **JWT:** Pustaka `github.com/golang-jwt/jwt/v4` untuk pembuatan token. Konfigurasi secret key JWT dari `.env`.
    * **Hashing Password:** Pustaka `golang.org/x/crypto/bcrypt` untuk membandingkan password yang di-hash.
    * **Validasi Input:** Pustaka `github.com/go-playground/validator/v10` untuk memvalidasi input `email` dan `password`.
* **Frontend:**
    * **Vue 3 Composition API:** Mengelola state form login dan melakukan panggilan API.
    * **Pinia:** Mengelola state otentikasi global (token JWT, status login, data pengguna).
    * **Vue Router:** Mengarahkan pengguna ke halaman dashboard setelah login berhasil.
    * **Axios:** Klien HTTP untuk interaksi API.
    * **Tailwind CSS:** Untuk styling form login.
    * Pahami bagaimana menyimpan dan mengirim JWT dengan aman di setiap permintaan terautentikasi (misalnya, di header `Authorization`).

## ⚠️ Potensi Jebakan & Perangkap
* **Keamanan Password:** Pastikan password di-hash dengan kuat menggunakan `bcrypt` dan jangan pernah menyimpannya dalam teks biasa.
* **Penanganan Error API:** Sediakan pesan error yang jelas dan ramah pengguna di frontend jika login gagal (misalnya, kredensial tidak valid, masalah server). Hindari membocorkan detail error internal server.
* **Pengelolaan Sesi JWT:** Tentukan masa berlaku JWT (misalnya, 1 jam) dan strategi *refresh token* jika diperlukan di masa depan (tidak dalam MVP ini, tetapi perlu dipertimbangkan). Pastikan token dihapus dari *local storage* saat logout.
* **Validasi Input Frontend & Backend:** Lakukan validasi di kedua sisi untuk memastikan integritas data dan pengalaman pengguna yang baik.
* **Race Conditions/Loading States:** Tangani status *loading* di frontend saat proses login sedang berlangsung untuk mencegah *double submission*.