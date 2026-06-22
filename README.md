# 🎬 CineReview

Aplikasi manajemen ulasan film berbasis terminal yang ditulis dalam bahasa **Go (Golang)**. CineReview memungkinkan pengguna untuk menambahkan, menghapus, mengubah, mencari, mengurutkan, dan melihat statistik koleksi film secara interaktif.

---

## 📋 Daftar Isi

- [Fitur](#-fitur)
- [Prasyarat](#-prasyarat)
- [Instalasi](#-instalasi)
- [Cara Menjalankan](#-cara-menjalankan)
- [Panduan Penggunaan](#-panduan-penggunaan)
- [Algoritma yang Digunakan](#-algoritma-yang-digunakan)
- [Struktur Data](#-struktur-data)

---

## ✨ Fitur

| No | Fitur | Keterangan |
|----|-------|------------|
| 1 | **Tambah Film** | Menambahkan satu atau lebih data film sekaligus |
| 2 | **Hapus Film** | Menghapus data film berdasarkan nomor urut |
| 3 | **Ubah Film** | Memperbarui seluruh informasi film yang dipilih |
| 4 | **Tampilkan Film** | Menampilkan semua data film yang tersimpan |
| 5 | **Cari Film** | Mencari film berdasarkan judul atau genre |
| 6 | **Urutkan Film** | Mengurutkan film berdasarkan tahun rilis atau rating |
| 7 | **Statistik** | Menampilkan jumlah film per genre atau rata-rata rating |

---

## 🔧 Prasyarat

Pastikan perangkat Anda telah terinstall:

- **Go** versi `1.18` atau lebih baru
- Sistem operasi: Windows, macOS, atau Linux

### Cek versi Go

```bash
go version
```

Jika Go belum terinstall, unduh dari situs resmi: [https://go.dev/dl/](https://go.dev/dl/)

---

## 📦 Instalasi

### 1. Clone atau Unduh File

Jika menggunakan Git:

```bash
git clone https://github.com/username/cine-review.git
cd cine-review
```

Atau letakkan file `main.go` ke dalam sebuah folder, misalnya:

```
cine-review/
└── main.go
```

### 2. Inisialisasi Module Go (jika belum ada `go.mod`)

```bash
go mod init cine-review
```

---

## ▶️ Cara Menjalankan

### Jalankan langsung (tanpa kompilasi)

```bash
go run main.go
```

### Kompilasi terlebih dahulu, lalu jalankan

```bash
# Kompilasi
go build -o cine-review main.go

# Jalankan (Linux / macOS)
./cine-review

# Jalankan (Windows)
cine-review.exe
```

---

## 📖 Panduan Penggunaan

Setelah program berjalan, Anda akan melihat tampilan menu utama:

```
Selamat Datang di CineReview🙌
=======================
 1. Menambahkan
 2. Menghapus
 3. Mengubah
 4. Tampilkan Film
 5. Cari Film
 6. Urutkan Film
 7. Statistik
 0. Keluar
=======================
Input Pilihan Anda:
```

Ketik angka pilihan dan tekan **Enter**.

---

### Menu 1 — Menambahkan Film

1. Pilih `1` lalu tekan Enter
2. Masukkan jumlah film yang ingin ditambahkan
3. Untuk setiap film, isi:
   - **Judul Film** — nama film *(otomatis dikonversi ke huruf kecil)*
   - **Genre** — genre film (contoh: `action`, `horror`, `drama`)
   - **Tahun Rilis** — tahun berupa angka (contoh: `2023`)
   - **Deskripsi** — sinopsis singkat film
   - **Skor Film** — nilai antara `0.0` hingga `10.0`

**Contoh input:**
```
Masukkan Jumlah Film Yang Mau Ditambahkan: 1
--- Film ke-1 ---
Masukkan Judul Film: Interstellar
Masukkan Genre Film: scifi
Masukkan Tahun Rilis: 2014
Masukkan Deskripsi Film: Perjalanan antariksa menembus lubang cacing
Masukkan Skor Film (0-10): 9.5
```

---

### Menu 2 — Menghapus Film

1. Pilih `2` lalu tekan Enter
2. Daftar film akan ditampilkan beserta nomor urutnya
3. Masukkan nomor film yang ingin dihapus

**Contoh:**
```
Data yang mau dihapus: 2
Data berhasil dihapus
```

---

### Menu 3 — Mengubah Film

1. Pilih `3` lalu tekan Enter
2. Daftar film akan ditampilkan
3. Masukkan nomor (ID) film yang ingin diubah
4. Isi ulang seluruh informasi film (judul, genre, tahun, deskripsi, skor)

---

### Menu 4 — Tampilkan Film

Pilih `4` untuk menampilkan seluruh data film yang tersimpan dalam format:

```
1. Judul: interstellar | Genre: scifi | Tahun: 2014 | Deskripsi: Perjalanan antariksa... | Skor: 9.5
```

---

### Menu 5 — Cari Film

1. Pilih `5` lalu tekan Enter
2. Pilih kriteria pencarian:
   - `1` — Berdasarkan **Judul**
   - `2` — Berdasarkan **Genre**
3. Masukkan kata kunci pencarian
4. Pilih metode pencarian:
   - `1` — **Binary Search** *(lebih cepat, data diurutkan otomatis sementara)*
   - `2` — **Sequential Search** *(menelusuri satu per satu, mendukung banyak hasil)*

> **Catatan:** Pencarian judul bersifat case-insensitive (huruf besar/kecil diabaikan). Pencarian genre bersifat case-sensitive.

---

### Menu 6 — Urutkan Film

1. Pilih `6` lalu tekan Enter
2. Pilih dasar pengurutan:
   - `1` — Berdasarkan **Tahun Rilis** *(dari terbaru ke terlama)*
   - `2` — Berdasarkan **Rating** *(dari tertinggi ke terendah)*
3. Pilih algoritma pengurutan:
   - `1` — **Selection Sort**
   - `2` — **Insertion Sort**

> **Catatan:** Hasil pengurutan hanya ditampilkan sementara dan tidak mengubah urutan data asli.

---

### Menu 7 — Statistik

Pilih `7` lalu tekan Enter, kemudian pilih:

- `1` — **Jumlah film per genre**: Masukkan nama genre untuk melihat daftar dan total film dalam genre tersebut
- `2` — **Rata-rata rating**: Menampilkan total jumlah film dan rata-rata skor seluruh film

**Contoh output:**
```
Total Film : 5
Rata-rata Rating: 8.2
```

---

### Menu 0 — Keluar

Pilih `0` untuk mengakhiri program.

```
Terima kasih telah menggunakan CineReview! Sampai jumpa di review film selanjutnya 🎬👋
```

---

## ⚙️ Algoritma yang Digunakan

| Algoritma | Digunakan Untuk |
|-----------|-----------------|
| **Sequential Search** | Mencari film berdasarkan judul atau genre (semua hasil) |
| **Binary Search** | Mencari film berdasarkan judul atau genre (data diurutkan dulu) |
| **Bubble Sort** | Mengurutkan sementara berdasarkan judul / genre (untuk Binary Search) |
| **Selection Sort** | Mengurutkan film berdasarkan tahun atau rating |
| **Insertion Sort** | Mengurutkan film berdasarkan tahun atau rating |

---

## 🗂️ Struktur Data

Program menggunakan `struct` bernama `CineReview` untuk menyimpan informasi setiap film:

```go
type CineReview struct {
    judulFilm  string   // Judul film
    genre      string   // Genre film
    tahunRilis int      // Tahun rilis
    deskripsi  string   // Deskripsi / sinopsis
    skorFilm   float64  // Skor film (0.0 - 10.0)
}
```

Data disimpan dalam **slice** (`[]CineReview`) selama program berjalan. Data tidak disimpan secara permanen ke file — semua data akan hilang saat program ditutup.

---

## ⚠️ Catatan Penting

- Data film **tidak tersimpan secara permanen**. Setiap kali program dijalankan ulang, data akan kosong kembali.
- Pencarian **judul** menggunakan huruf kecil secara otomatis. Pastikan input kata kunci juga menggunakan huruf kecil agar cocok.
- Pencarian dan pengurutan berdasarkan **genre** bersifat *case-sensitive*.
- Skor film yang valid adalah antara `0.0` hingga `10.0`.
