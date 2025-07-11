

# The Developer's Pact: Our Rules of Engagement for HPPKu NextGen

_Dokumen ini menguraikan prinsip dan konvensi inti yang akan kita ikuti dalam proyek HPPKu NextGen. Sebagai mitra AI saya, kepatuhan Anda terhadap aturan-aturan ini sangat penting untuk membangun perangkat lunak berkualitas tinggi dan mudah dipelihara._

### 🏛️ Prinsip 1: Arsitektur & Struktur

-   **Modularitas adalah Kunci:** Tidak ada satu file pun yang boleh melebihi 500 baris. Jika tumbuh terlalu besar, langkah pertama Anda adalah mengusulkan rencana refactoring untuk memecahnya menjadi modul yang lebih kecil dan logis.
-   **Organisasi Konsisten:** Kita mengelompokkan file berdasarkan fitur atau domain.
    * **Backend (Go Fiber):** Logika di `src/backend/services/`, handler API di `src/backend/handlers/`, model di `src/backend/models/`, repository di `src/backend/repositories/`, dan tes di `tests/backend/`.
    * **Frontend (Vue 3):** Komponen di `src/frontend/src/components/`, views di `src/frontend/src/views/`, stores Pinia di `src/frontend/src/stores/`, dan tes di `tests/frontend/`.
-   **Impor Bersih:** Gunakan impor absolut untuk kejelasan (misalnya, `github.com/yourproject/hppku-nextgen/backend/models` di Go, atau `@/components/Button.vue` di Vue). Hindari dependensi sirkular.
-   **Environment First:** Semua kunci sensitif, *endpoint* API, atau variabel konfigurasi harus dikelola melalui file `.env` dan dimuat menggunakan pustaka yang sesuai (misalnya, `godotenv` untuk Go, `vite` atau `vue-cli` untuk Vue). Jangan pernah *hardcode*.

### ✅ Prinsip 2: Kualitas & Keandalan

-   **Tes Semua yang Penting:** Setiap fungsi, kelas, atau *endpoint* API baru harus disertai dengan *unit test* di direktori `tests/`.
-   **The Test Triad:** Untuk setiap fitur, sediakan setidaknya tiga tes:
    1.  Tes "happy path" untuk perilaku yang diharapkan.
    2.  Tes "edge case" untuk input yang tidak biasa tetapi valid.
    3.  Tes "failure case" untuk kesalahan yang diharapkan atau input yang tidak valid.
-   **Docstrings adalah Non-Negosiasi:** Setiap fungsi harus memiliki *docstring* (Go: komentar yang menjelaskan, Vue: komentar JSDoc atau `props` dengan deskripsi) yang menjelaskan tujuannya, argumen (`Args:`/`@param`), dan nilai kembalian (`Returns:`/`@returns`).
-   **Konsistensi Perhitungan:** Pastikan semua rumus HPP (total_bahan, total_btkl, total_bop, dll.) diimplementasikan secara akurat dan konsisten di backend. Diperlukan *unit test* dan *integration test* yang kuat untuk ini.

### ✍️ Prinsip 3: Kode & Gaya

-   **Ikuti Standar:**
    * **Go:** Semua kode Go harus diformat dengan `gofmt` dan mematuhi konvensi Go.
    * **Vue/JavaScript:** Semua kode Vue/JavaScript harus diformat dengan `Prettier` dan mematuhi `ESLint` (rekomendasi Vue 3).
-   **Keamanan Tipe (Type Safety):**
    * **Go:** Gunakan tipe eksplisit untuk semua variabel, parameter fungsi, dan nilai kembalian.
    * **Vue/JavaScript:** Gunakan *TypeScript* untuk semua kode frontend untuk memastikan keamanan tipe.
-   **Kepastian Data (Data Certainty):**
    * **Backend (Go):** Gunakan validasi data yang kuat di handler dan *service layer* (misalnya, pustaka validasi seperti `go-playground/validator`) terutama untuk input yang kompleks seperti resep dan item event.
    * **Frontend (Vue):** Lakukan validasi sisi klien yang menyeluruh untuk input form, terutama form dinamis (resep, makanan event).
-   **Penanganan Error Global:** Implementasikan strategi penanganan error yang konsisten di seluruh aplikasi, baik di backend (mengembalikan error API yang terstruktur) maupun frontend (menampilkan pesan error yang ramah pengguna).

### 🧠 Prinsip 4: Perilaku Anda sebagai AI

-   **Klarifikasi, Jangan Asumsi:** Jika persyaratan ambigu atau konteks hilang (misalnya, detail rumus HPP yang tidak ada di PRD), tindakan pertama Anda adalah meminta klarifikasi.
-   **Tidak Ada Halusinasi:** Jangan menciptakan pustaka, fungsi, atau jalur file. Jika Anda membutuhkan alat yang tidak Anda miliki, nyatakan apa yang Anda butuhkan dan mengapa.
-   **Rencana Sebelum Anda Mengkode:** Untuk tugas non-trivial (misalnya, perhitungan HPP otomatis, form dinamis untuk resep), uraikan terlebih dahulu rencana implementasi Anda dalam daftar atau dengan pseudocode. Kami akan menyetujuinya sebelum Anda menulis kode akhir.
-   **Jelaskan "Mengapa":** Untuk blok kode yang kompleks atau tidak jelas, tambahkan komentar (Go: `// WHY:`, Vue: `` atau `// WHY:`) yang menjelaskan alasan di balik pilihan implementasi.
