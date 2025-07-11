# HPPKu NextGen (Gemini Edition)

Aplikasi HPPKu NextGen adalah sistem manajemen dan perhitungan Harga Pokok Produksi (HPP) untuk produk makanan individual dan pesanan event. Proyek ini bertujuan untuk memigrasikan dan meningkatkan fungsionalitas inti dari sistem HPPKu yang ada ke dalam *tech stack* modern (Go Fiber & Vue 3) dengan fokus pada performa, skalabilitas, dan pengalaman pengguna yang lebih baik.

Proyek ini dibangun dengan filosofi *Context Engineering*, memastikan bahwa AI asisten memiliki semua informasi yang diperlukan untuk membangun aplikasi ini secara end-to-end, mengikuti standar kualitas dan praktik terbaik.

---

💡 **Ingin tahu lebih banyak tentang filosofi di balik proyek ini?** Baca posting blog perkenalan tentang Context Engineering (jika ada, sesuaikan linknya).

---

## ⚙️ Konfigurasi

Sebelum Anda dapat menggunakan *framework* ini, Anda perlu mengkonfigurasinya dengan kunci API Google Gemini.

### 1. Dapatkan Kunci API Anda

Anda bisa mendapatkan kunci API gratis dari [Google AI Studio](https://aistudio.google.com/app/apikey).

### 2. Atur Variabel Lingkungan

*Framework* ini menggunakan variabel lingkungan untuk mengakses kunci API Anda dengan aman. Buka terminal Anda dan jalankan perintah yang sesuai untuk sistem Anda:

-   **macOS / Linux:**
    ```shell
    export GEMINI_API_KEY="YOUR_API_KEY_HERE"
    ```

-   **Windows (PowerShell):**
    ```powershell
    $env:GEMINI_API_KEY="YOUR_API_KEY_HERE"
    ```

**Catatan:** Variabel ini hanya diatur untuk sesi terminal Anda saat ini. Untuk solusi permanen, Anda perlu menambahkan baris ini ke file startup shell Anda (misalnya, `.zshrc`, `.bash_profile`, atau profil PowerShell Anda).

### 3. Instal Dependensi

Skrip menggunakan alat baris perintah standar agar berfungsi. Silakan instal menggunakan *package manager* sistem Anda.

-   `jq`: Untuk parsing respons JSON.
-   `awk`: Untuk parsing rencana AI. (Biasanya sudah terinstal di Linux/macOS).
-   `base64`: Untuk encoding/decoding konten. (Biasanya sudah terinstal di Linux/macOS).

## 📚 Daftar Isi

**[Apa itu Context Engineering?](#apa-itu-context-engineering)**

**[Struktur Template](#struktur-template)**

**[Panduan Langkah-demi-Langkah](#panduan-langkah-demi-langkah)**

**[Menulis File INITIAL.md yang Efektif](#menulis-file-initialmd-yang-efektif)**

**[Alur Kerja PRP](#alur-kerja-prp)**

**[Menggunakan Contoh Secara Efektif](#menggunakan-contoh-secara-efektif)**

**[Praktik Terbaik](#praktik-terbaik)**

## Apa itu Context Engineering?

*Context Engineering* mewakili pergeseran paradigma dari *prompt engineering* tradisional:

### Prompt Engineering vs Context Engineering

**Prompt Engineering:**
-   Fokus pada pemilihan kata yang cerdas dan frasa spesifik
-   Terbatas pada bagaimana Anda merumuskan tugas
-   Seperti memberi seseorang catatan tempel

**Context Engineering:**
-   Sistem lengkap untuk menyediakan konteks komprehensif
-   Mencakup dokumentasi, contoh, aturan, pola, dan validasi
-   Seperti menulis skenario lengkap dengan semua detailnya

### Mengapa Context Engineering Penting

1.  **Mengurangi Kegagalan AI:** Sebagian besar kegagalan agen bukanlah kegagalan model - itu adalah kegagalan konteks.
2.  **Memastikan Konsistensi:** AI mengikuti pola dan konvensi proyek Anda.
3.  **Memungkinkan Fitur Kompleks:** AI dapat menangani implementasi multi-langkah dengan konteks yang tepat.
4.  **Self-Correcting:** Loop validasi memungkinkan AI untuk memperbaiki kesalahannya sendiri.

## Struktur Template

hppku-nextgen/
│
├── .gemini/
│   ├── scripts/
│   │   ├── generate-prp.sh
│   │   └── execute-prp.sh
│   └── templates/
│       └── prp_template.md
│
├── PRPs/
│   ├── auth_login_prp.md
│   └── ...
│
├── examples/
│   ├── go_fiber/
│   │   ├── handler_example.go
│   │   ├── model_example.go
│   │   └── service_example.go
│   ├── vue_3/
│   │   ├── api_client_example.vue
│   │   ├── component_form_example.vue
│   │   └── pinia_store_example.js
│
├── src/
│   ├── backend/
│   │   ├── main.go
│   │   └── ... (struktur sesuai usulan di atas)
│   ├── frontend/
│   │   ├── src/
│   │   └── ... (struktur sesuai usulan di atas)
│
├── tests/
│   ├── backend/
│   │   ├── auth_test.go
│   │   └── ...
│   ├── frontend/
│   │   ├── unit/
│   │   └── e2e/
│
├── .env.example
├── .gitignore
├── GEMINI.md                 # Global rules and principles for the AI assistant.
├── INITIAL.md                # A blank template for writing new feature requests.
├── INITIAL_EXAMPLE.md        # An example of a completed feature request.
├── PRD_HPPKu_NextGen.md      # Dokumen PRD asli.
└── README.md                 # This file.


**Arah Masa Depan:**

Template ini menyediakan fondasi yang kuat untuk *Context Engineering*. Evolusi berikutnya dari alur kerja ini dapat melibatkan integrasi teknik AI yang lebih canggih, seperti *Retrieval-Augmented Generation (RAG)* untuk penelitian dokumentasi otomatis dan alat yang didorong oleh AI (*function calling*) untuk implementasi dan *testing loop* yang sepenuhnya otonom.

## Panduan Langkah-demi-Langkah

*Framework* ini menggunakan alur kerja dua langkah yang kuat untuk membangun perangkat lunak dengan Gemini.

### 1. Buat Permintaan Fitur

Untuk memulai fitur baru, Anda tidak mengedit template `INITIAL.md` secara langsung. Sebaliknya:

1.  **Salin** `INITIAL.md` ke file baru bernama `request.md` (misalnya, `requests/nama_fitur.md`).
2.  **Isi** `request.md` dengan detail fitur baru Anda, mengacu pada PRD HPPKu NextGen.

### 2. Hasilkan PRP (Rencana)

Sekarang, jalankan skrip `generate-prp.sh`. Skrip ini mengirimkan permintaan fitur Anda ke Gemini dan memintanya untuk bertindak sebagai insinyur senior, membuat cetak biru teknis terperinci yang disebut *Product Requirements Prompt* (PRP).

-   **Untuk macOS / Linux:**
    ```shell
    # Make the script executable first
    chmod +x .gemini/scripts/generate-prp.sh

    # Run the script
    ./.gemini/scripts/generate-prp.sh requests/nama_fitur.md
    ```

-   **Untuk Windows (menggunakan interpreter bash seperti Git Bash):**
    ```shell
    bash ./.gemini/scripts/generate-prp.sh requests/nama_fitur.md
    ```

Ini akan membuat file PRP baru dan permanen di dalam direktori `PRPs/`.

### 3. Eksekusi PRP (Agen Otomatis)

Di sinilah keajaiban terjadi. Jalankan skrip `execute-prp.sh` dengan jalur ke PRP yang baru dibuat.

-   **Untuk macOS / Linux:**
    ```shell
    # Make the script executable first
    chmod +x .gemini/scripts/execute-prp.sh

    # Run the script
    ./.gemini/scripts/execute-prp.sh PRPs/nama_fitur_prp.md
    ```

-   **Untuk Windows (menggunakan interpreter bash seperti Git Bash):**
    ```shell
    bash ./.gemini/scripts/execute-prp.sh PRPs/nama_fitur_prp.md
    ```

Skrip ini bertindak sebagai **agen AI**:

1.  Ia mengirimkan PRP terperinci ke Gemini untuk mendapatkan rencana implementasi langkah-demi-langkah.
2.  Ia memparsing respons AI, mengidentifikasi perintah shell untuk dijalankan dan file kode untuk dibuat.
3.  Ia kemudian mengeksekusi rencana ini langkah-demi-langkah, **menjeda untuk meminta konfirmasi Anda** sebelum menjalankan perintah apa pun atau menulis file apa pun.

Ini memungkinkan Anda untuk bersantai dan mengawasi saat AI membangun seluruh fitur untuk Anda, langsung di terminal lokal Anda.

## Menulis File INITIAL.md yang Efektif

File `INITIAL.md` adalah titik awal untuk fitur baru apa pun. Kualitas input Anda di sini secara langsung memengaruhi kualitas output AI. Berikut cara membuatnya efektif:

-   **🎯 Tujuan:** Jadilah spesifik dan jelas. Alih-alih "buat halaman login," tulis "buat halaman login dengan bidang email/password, tautan 'Lupa Password', dan integrasi Google OAuth." Semakin detail, semakin baik.
-   **🎨 Inspirasi & Contoh:** Ini adalah salah satu bagian yang paling kuat. Jika Anda memiliki file yang ada yang menunjukkan bagaimana Anda menangani klien API, koneksi database, atau penanganan kesalahan, rujuk di sini. Ini memberi AI pola konkret untuk diikuti, memastikan konsistensi.
-   **📚 Pengetahuan yang Dibutuhkan:** Jangan membuat AI menebak. Berikan tautan langsung ke dokumentasi API, pustaka, atau thread Stack Overflow yang tepat yang akan dibutuhkan. Ini menghemat waktu dan mencegah AI menggunakan informasi yang usang atau salah.
-   **⚠️ Potensi Jebakan & Perangkap:** Pikirkan apa yang mungkin salah. Apakah API memiliki batasan laju yang aneh? Apakah ada alur otentikasi yang rumit? Menyebutkan ini di awal membantu AI menghindari kesalahan umum dan membangun kode yang lebih kuat sejak awal.

## Alur Kerja PRP

Alur kerja PRP (*Product Requirements Prompt*) adalah proses dua langkah yang dirancang untuk memastikan AI memiliki rencana komprehensif sebelum menulis satu baris kode pun.

1.  **Generasi (`generate-prp.sh`):** Langkah pertama ini adalah tentang **perencanaan**. Skrip mengambil permintaan fitur `INITIAL.md` tingkat tinggi Anda dan meminta Gemini untuk mengembangkannya menjadi cetak biru teknis terperinci (PRP). Cetak biru ini mencakup struktur file yang diusulkan, rincian tugas, pseudocode, dan rencana validasi. Ini memaksa AI untuk memikirkan seluruh implementasi terlebih dahulu.

2.  **Eksekusi (`execute-prp.sh`):** Langkah kedua ini adalah tentang **implementasi**. Skrip mengambil PRP terperinci dan mengirimkannya kembali ke Gemini dengan instruksi yang jelas: "Bangun ini." Karena semua penelitian dan perencanaan sudah selesai, AI dapat fokus sepenuhnya pada penulisan kode yang bersih dan benar yang mengikuti cetak biru.

Pemisahan perencanaan dari implementasi ini adalah kunci keberhasilan alur kerja.

## Menggunakan Contoh Secara Efektif

Folder `examples/` adalah senjata rahasia Anda untuk memastikan konsistensi proyek. Asisten AI unggul dalam pengenalan pola.

### Apa yang membuat contoh yang baik?

-   **Pola Lengkap:** Tunjukkan contoh pola yang berfungsi penuh, bukan hanya cuplikan. Misalnya, berikan kelas klien API lengkap, bukan hanya satu fungsi.
-   **Struktur Kode:** Sertakan contoh bagaimana Anda menyusun kelas, mengatur impor, dan menamai variabel Anda.
-   **Penanganan Kesalahan:** Tunjukkan bagaimana Anda mengharapkan kesalahan ditangkap dan ditangani. Ini sering diabaikan tetapi sangat penting untuk kode kualitas produksi.
-   **Pengujian:** Berikan contoh file pengujian (`test_*.go`, `*.spec.js`) yang menunjukkan gaya pengujian pilihan Anda, termasuk cara menggunakan *mocks*.

Semakin banyak contoh berkualitas tinggi yang Anda berikan, semakin sedikit AI harus menebak, dan semakin banyak kode akhir akan terlihat seperti Anda yang menulisnya sendiri.

## Praktik Terbaik

-   **Jadilah Eksplisit:** Jangan pernah berasumsi AI mengetahui preferensi Anda. Semakin eksplisit Anda dalam file `INITIAL.md` dan `GEMINI.md` Anda, semakin baik hasilnya.
-   **Iterasi pada Proses:** Jika AI membuat kesalahan, jangan hanya memperbaiki kode. Pikirkan mengapa AI membuat kesalahan itu. Apakah aturan dalam `GEMINI.md` perlu lebih jelas? Apakah Anda memerlukan contoh yang lebih baik di folder `examples/`? Meningkatkan proses akan mencegah kesalahan yang sama terjadi lagi.
-   **Percayai Alur Kerja:** Mungkin terlihat seperti pekerjaan ekstra untuk menulis `INITIAL.md` yang detail dan meninjau PRP, tetapi investasi awal ini menghemat banyak waktu untuk debugging dan refactoring nanti.

---