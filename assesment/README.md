# E-Commerce Aqua Developer

E-Commerce adalah sistem komersial berbasis internet yang bisa mengelola pengguna untuk melakukan transaksi pada produk eFishery yang tersedia, serta memonitoring hasil transaksi pada keranjang pengguna secara daring.

## Problem and Motivation

Para mitra eFishery menemukan kesulitan dalam memenuhi kebutuhan untuk pembudidayaan dan pemeliharaan udang dan ikan yang dikelolanya, serta tidak bisa memonitor hasil pembelian produk yang telah dipesan. Maka dari itu dibutuhkannya sebuah sistem untuk mengotomatisasikan permasalahan ini.

## Cara menjalankan aplikasi

Untuk menjalankan aplikasi, pertama-tama clone repository ini dengan command:

```bash
git clone https://github.com/uranghishna/Koje-Aqua-Developer.git
```

Aplikasi E-Commerce ini tersedia pada folder "assesment", abaikan folder lain.
Kemudian buat file dengan ekstensi .env pada folder "assesment" dan isi dengan:

```bash
BASE_URL="http://localhost:8000"
```

Setelah itu buat database bernama "assesment", setelah membuat database lalu jalankan command berikut:

```bash
go run .
```

Jika muncul error, maka jalankan command "go get" untuk mendapatkan semua package yang dibutuhkan. contoh:

```bash
 go get -u github.com/swaggo/swag/cmd/swag
```

Setelah itu jalankan kembali programnya dengan command "go run ." .

## Cara menggunakan

Untuk melihat endpoint API yang tersedia, bisa menggunakan link swagger pada local. contoh:

```bash
 http://localhost:8000/swagger/index.html
```

atau ekstrak collection.json yang sudah tersedia pada aplikasi di postman.
