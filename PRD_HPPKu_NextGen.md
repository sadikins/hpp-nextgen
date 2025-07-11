Nama Produk: HPPKu NextGen (Nama Tentatif)Versi Dokumen: 1.1Tanggal: 10 Juli 2025Penulis: [Nama Anda/Tim]

1. Pendahuluan



1.1. Tujuan Proyek

Dokumen ini menguraikan persyaratan produk untuk pengembangan aplikasi HPPKu NextGen, sebuah sistem manajemen dan perhitungan Harga Pokok Produksi (HPP) untuk produk makanan individual dan pesanan event dalam periode tertentu. Tujuannya adalah untuk memigrasikan dan meningkatkan fungsionalitas inti dari sistem HPPKu yang ada ke dalam tech stack modern (Go Fiber & Vue 3) dengan fokus pada performa, skalabilitas, dan pengalaman pengguna yang lebih baik.

1.2. Visi Produk

Menjadi solusi HPP yang intuitif dan akurat bagi bisnis kuliner, memungkinkan mereka untuk dengan mudah menghitung biaya produksi per produk dan per event, serta memantau profitabilitas secara efisien.

1.3. Ruang Lingkup (MVP - Minimum Viable Product)

MVP ini akan mencakup fitur-fitur inti yang esensial untuk perhitungan dan manajemen HPP per produk dan per event, serta master data yang mendukung. Modul laporan dasar dan manajemen pengguna juga akan disertakan.

2. Target Audiens



Pengguna Utama: Pemilik bisnis kuliner, manajer produksi, akuntan, atau individu yang bertanggung jawab atas perhitungan biaya dan penetapan harga produk makanan atau layanan katering.

Peran Pengguna:

Admin: Memiliki akses penuh ke semua fitur, termasuk manajemen master data, pengguna, dan semua laporan.

User: Memiliki akses ke fitur perhitungan HPP makanan dan event, serta melihat laporan yang relevan. (Penyesuaian lebih lanjut pada hak akses dapat dilakukan di fase berikutnya).

3. Fitur Inti (MVP)

3.1. Modul Autentikasi & Otorisasi



PR-AUTH-001 - Login Pengguna: Pengguna dapat masuk ke sistem menggunakan kredensial yang valid.

Backend: Endpoint API untuk menerima kredensial, memvalidasi, dan mengembalikan JWT (JSON Web Token).

Frontend: Halaman login dengan form input username/email dan password.

PR-AUTH-002 - Logout Pengguna: Pengguna dapat keluar dari sistem, menghapus sesi/token.

Backend: Endpoint API untuk menginvalidasi token (jika diperlukan untuk logout di sisi server).

Frontend: Tombol logout yang menghapus token dari local storage atau session.

PR-AUTH-003 - Manajemen Pengguna (Admin Only): Admin dapat melakukan CRUD (Create, Read, Update, Delete) data pengguna dan menetapkan peran (role).

Backend: Endpoint API CRUD untuk entitas pengguna.

Frontend: Halaman daftar pengguna, form tambah/edit pengguna.

PR-AUTH-004 - Otorisasi Berbasis Peran (RBAC): Sistem akan membatasi akses fitur dan data berdasarkan peran pengguna.

Backend: Middleware Fiber untuk melindungi endpoint API berdasarkan peran.

Frontend: Navigasi dan elemen UI akan disembunyikan/dinonaktifkan jika pengguna tidak memiliki izin.

3.2. Modul Master Data



PR-MD-001 - Manajemen Bahan Baku: Pengguna (Admin) dapat melakukan CRUD untuk bahan baku (kode, nama, satuan, harga satuan).

PR-MD-002 - Manajemen Kategori Makanan: Pengguna (Admin) dapat melakukan CRUD untuk kategori makanan.

PR-MD-003 - Manajemen Makanan: Pengguna (Admin) dapat melakukan CRUD untuk data makanan (kode, nama, kategori, status HPP).

PR-MD-004 - Manajemen Parameter BOP: Pengguna (Admin) dapat melihat dan mengubah nilai persentase BOP.

PR-MD-005 - Manajemen Parameter BTKL: Pengguna (Admin) dapat melihat dan mengubah nilai BTKL.

PR-MD-006 - Manajemen Parameter Cost: Pengguna (Admin) dapat melihat dan mengubah nilai persentase cost (sales price).

PR-MD-007 - Manajemen Waktu Event: Pengguna (Admin) dapat melihat dan mengubah konfigurasi waktu event.

Catatan: Untuk semua fitur Master Data, diperlukan API endpoint CRUD di backend dan halaman daftar/form di frontend.

3.3. Modul Perhitungan HPP per Produk (Makanan)



PR-HPP-001 - Input Resep Makanan: Pengguna dapat memilih makanan dan menambahkan bahan baku beserta kuantitasnya untuk membentuk resep.

Backend: Endpoint API untuk menyimpan resep (reseps table).

Frontend: Form dinamis untuk menambah/menghapus baris bahan baku pada resep.

PR-HPP-002 - Perhitungan HPP Otomatis (per Makanan): Setelah resep disimpan, sistem secara otomatis menghitung total_bahan, total_btkl, total_bop, total_hpp, dan h_jual untuk makanan tersebut.

Backend: Logika perhitungan HPP sesuai rumus yang telah dianalisis sebelumnya, dan menyimpan hasilnya ke tabel hpps.

PR-HPP-003 - Daftar & Detail HPP Makanan: Menampilkan daftar makanan dengan status HPP dan detail perhitungan HPP untuk makanan tertentu.

PR-HPP-004 - Edit & Hapus HPP Makanan: Pengguna dapat mengubah resep makanan atau menghapus perhitungan HPP yang sudah ada.

3.4. Modul Perhitungan HPP per Periode Produksi (Event Order)



PR-EVENT-001 - Pembuatan Event Order: Pengguna dapat membuat event baru dengan mengisi detail event (nama, rentang tanggal) dan memilih makanan yang terlibat beserta kuantitas porsinya.

Backend: Endpoint API untuk menyimpan event (eveents table) dan makanan yang terkait (eveent_makanans table).

PR-EVENT-002 - Perhitungan HPP Event Otomatis: Sistem secara otomatis menghitung total_bahan, total_btkl, total_bop, total_produksi (HPP event), biaya_perporsi, harga_jual_perporsi, dan revenue untuk seluruh event.

Backend: Logika agregasi biaya dan perhitungan total untuk event, menyimpan hasilnya ke tabel eveents.

PR-EVENT-003 - Daftar & Detail Event Order: Menampilkan daftar event yang sudah dibuat dan detail perhitungannya.

PR-EVENT-004 - Edit & Hapus Event Order: Pengguna dapat mengubah detail event atau menghapus event order.

3.5. Modul Laporan (Reporting)



PR-REP-001 - Laporan HPP Makanan: Menampilkan laporan HPP untuk setiap makanan dengan opsi cetak ke PDF.

PR-REP-002 - Laporan Event Order: Menampilkan laporan rincian biaya dan pendapatan per event dengan opsi cetak ke PDF.

PR-REP-003 - Filter Laporan: Menyediakan opsi filter pada laporan (misalnya berdasarkan kategori makanan untuk HPP, rentang tanggal untuk Event).

Backend: Endpoint API untuk menghasilkan data laporan yang difilter.

Frontend: UI untuk pemilihan filter dan pemicu cetak.

3.6. Dashboard



PR-DASH-001 - Tampilan Ringkasan: Menampilkan ringkasan statistik penting seperti jumlah makanan, jumlah event, ringkasan pendapatan, dan mungkin grafik sederhana.

4. Non-Fungsional Requirements



Performa: Respon API tidak lebih dari 500ms untuk operasi CRUD dan perhitungan standar. Laporan yang kompleks tidak lebih dari 5 detik.

Keamanan:

Penggunaan HTTPS untuk semua komunikasi.

Sanitasi dan validasi input untuk mencegah serangan injeksi.

Hashing password yang kuat.

Autentikasi JWT yang aman.

Skalabilitas: Arsitektur mendukung penambahan pengguna dan data tanpa penurunan performa yang signifikan.

Usability: Antarmuka pengguna yang intuitif, bersih, dan mudah dipelajari.

Maintainability: Kode bersih, terdokumentasi, dan mudah untuk dimodifikasi atau ditambahkan fitur baru.

Reliabilitas: Sistem harus stabil dengan penanganan kesalahan yang memadai.

5. Technical Stack



Backend Framework: Go (Fiber)

Frontend Framework: Vue 3 (Composition API dengan <script setup>)

Database: PostgreSQL

ORM/Database Library (Go): GORM

State Management (Vue): Pinia

Routing (Vue): Vue Router

UI/CSS Framework (Vue): Tailwind CSS

PDF Generation: JsPDF (frontend-based)

6. Tantangan Teknis Utama yang Harus Diantisipasi



Backend (Go dengan Fiber):

Pemilihan dan Implementasi ORM/Query Builder: Go tidak memiliki ORM bawaan sekuat Eloquent. Implementasi GORM untuk relasi kompleks (resep, event-makanan) akan krusial.

Validasi Input & Penanganan Error API: Mengimplementasikan validasi input yang robust, terutama untuk data array (resep, item event), dan penanganan error yang konsisten di seluruh API.

Manajemen Konfigurasi Global (BOP, BTKL, Cost): Strategi pengelolaan nilai-nilai global ini agar konsisten dan up-to-date, serta potensi implikasinya pada perhitungan historis.

Frontend (Vue 3 dengan Script Setup):

Manajemen State (Pinia): Mengelola state yang kompleks (data master, data HPP, data event) agar sinkron dan performant di seluruh aplikasi menggunakan Pinia.

Form Kompleks (Resep & Event Makanan): Mendesain dan mengimplementasikan form dengan input dinamis (tambah/hapus baris bahan baku/makanan) dan validasi sisi klien yang menyeluruh.

Interaksi API Asynchronous: Efisiensi fetching data, penanganan loading/error states, dan cascading selects saat memilih data terkait.

Desain Komponen Reusable: Membangun komponen UI yang reusable dan fleksibel dengan Tailwind CSS untuk menjaga konsistensi dan kecepatan pengembangan.

Generasi PDF (Laporan) dengan JsPDF: Mengkonversi konten HTML/Vue menjadi PDF di sisi frontend menggunakan JsPDF, yang mungkin memerlukan penyesuaian tata letak dan styling yang cermat.

Tantangan Umum (Backend & Frontend):

Konsistensi Perhitungan: Memastikan semua rumus HPP diimplementasikan secara akurat dan konsisten di backend, terutama agregasi untuk event. Diperlukan unit/integration tests yang kuat.

Skalabilitas: Mengoptimalkan query database dan performa perhitungan untuk menangani peningkatan volume data di masa mendatang.

Deployment: Mengelola deployment aplikasi Go (binary) dan Vue (static assets), serta mengonfigurasi reverse proxy (misalnya Nginx) untuk melayani keduanya.

7. Pertimbangan Masa Depan (Diluar MVP)



Fitur inventaris bahan baku.

Manajemen pembelian bahan baku.

Modul laporan keuangan lebih lanjut (laba rugi, neraca).

Integrasi dengan sistem POS.

Notifikasi dan peringatan (misalnya, stok bahan baku rendah).

Fitur impor/ekspor data.