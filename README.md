<h3>Field Service</h3>

<h3>Description</h3>

<p>...</p>

<h3>Directory Structure</h3>

```
skeleton
    L cmd                            → Contains the main entry point or initial configuration of the application
    L common                         → Stores common functions used throughout the application
    L config                         → Contains application configurations such as environment variables and other settings
    L constants                      → Stores global constant values used across the application
    L controllers                    → Manages control logic for handling HTTP requests
    L database                       → Contains files related to database management
        L seeders                    → Scripts for populating initial (seed) data into the database
    L domain                         → The application's domain module containing core domain elements
        L dto                        → Data Transfer Objects, used to define the structure of transferred data
        L models                     → Object models representing the application's or database's data structure
    L middlewares                    → Contains middleware for processing requests/responses before or after reaching the controller
    L repositories                   → Contains data access logic for interacting with the database
    L routes                         → Contains API route definitions
    L services                       → Stores the application's core business logic
```

```
skeleton/
├── cmd/                  → Titik awal (entry point) aplikasi. Biasanya berisi file `main.go`.
├── common/               → Berisi fungsi-fungsi umum (utility/helper) yang digunakan di berbagai bagian aplikasi.
├── config/               → Berisi konfigurasi dan inisialisasi layanan eksternal seperti database, Redis, MinIO, dsb.
├── constants/            → Menyimpan nilai konstanta global seperti enum, status code, atau pesan error.
├── controllers/          → Menangani logika untuk menerima dan merespons permintaan HTTP (jika REST digunakan).
├── database/             → Berisi logika terkait database seperti migrasi dan seeder (data awal).
│   └── seeders/          → Script untuk mengisi data awal (seeding).
├── domain/               → Berisi elemen inti domain aplikasi.
│   ├── dto/              → Data Transfer Object, mendefinisikan struktur data yang ditransfer.
│   └── models/           → Model representasi data dalam aplikasi atau database.
├── middlewares/          → Middleware untuk memproses permintaan sebelum atau sesudah controller/service.
├── repositories/         → Berisi logika untuk akses data dari database (query, insert, update, delete).
├── routes/               → Mendefinisikan rute-rute HTTP API (jika menggunakan REST).
├── services/             → Menyimpan logika bisnis utama aplikasi (core business logic).

```

## How to setup

```
- Clone this repository
- go mod tidy
```

## How to run

```bash
go run main.go
```

## How to build image docker
```bash
docker compose build
```

## How to run with docker

```bash
docker compose up -d
```



