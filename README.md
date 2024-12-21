# SISTEM MANAGEMENT RESTO
Sistem Manajemen Resto adalah backend API yang dirancang untuk mengelola data restoran, termasuk menu, pesanan, detail pesanan, dan transaksi. Aplikasi ini mendukung operasi CRUD dan menyediakan dokumentasi API menggunakan Swagger.

![GitHub Logo](https://cdn.prod.website-files.com/6100d0111a4ed76bc1b9fd54/62217e885f52b860da9f00cc_Apa%20Itu%20Golang%3F%20Apa%20Saja%20Fungsi%20Dan%20Keunggulannya%20-%20Binar%20Academy.jpeg)

## Fitur Utama
- **Auth Basic:** Implementasi Auth Basic ketika signin dan memberikan otorisasi pada setiap API   
- **Manajemen Menu Item:** Terdapat proses implementasi CRUD pada menu item dengan otorisasi dari basic auth.
- **Manajemen Pesanan:** Terdapat proses implementasi CRUD pada pesanan dengan otorisasi dari basic auth.
- **Manajemen Detail Pesanan:** Terdapat proses implementasi CRUD pada detail pesanan dengan otorisasi dari basic auth.
- **Manajemen Transaksi:** Terdapat proses implementasi CRUD pada transaksi dengan otorisasi dari basic auth.
- **Dokumentasi API:** Swagger UI untuk eksplorasi API.

## Teknologi
- **Bahasa:** Golang
- **Golang Migrate:**
    ```bach
    go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
- **Golang Httprouter:**
    ```bach
    github.com/julienschmidt/httprouter
- **Golang Validate:**
    ```bach
    github.com/go-playground/validator/v10
- **Database:** PostgreSQL
    ```bach
    github.com/lib/pq
- **Dokumentasi API:** Swagger UI
    ```bach
    github.com/swaggo/swag/cmd/swag
    github.com/swaggo/http-swagger
## Instalasi
1. Clone repository:
   ```bash
   git clone https://github.com/AgustiBayu/sistem-management-resto.git
   
2. cd sistem-manajemen-restoran
3. go mod tidy
4. Atur konfigurasi database di file app.
5. Jalankan aplikasi:
    ```bash
    go run main.go

## Dokumentasi API
Dokumentasi API tersedia di Swagger UI.

**URL Lokal:** [Swagger UI - URL](http://localhost:3000/swagger/index.html)

Untuk melihat spesifikasi API secara langsung, Anda dapat menggunakan [Swagger Editor](https://editor.swagger.io/).