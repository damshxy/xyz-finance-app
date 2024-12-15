# XYZ Finance App

XYZ Finance App adalah manajemen transaksi dan konsumen yang menggunakan **Go**, **PostgreSQL**, dan **Docker**. Aplikasi ini memungkinkan admin untuk mengelola konsumen dan transaksi, serta memberikan fitur refund untuk transaksi yang sudah dilakukan.

## Fitur Utama

- **Manajemen Konsumen**:
     - Membuat konsumen baru.
     - Melihat profil konsumen berdasarkan ID.
     - Mengupdate batas kredit konsumen

- **Manajem Transaksi**:
     - Membuat transaksi untuk konsumen dengan pengecekan batas kredit.
     - Melihat transaksi berdasarkan konsumen.
     - Refund transaksi dan mengembalikan batas kredit.

## Teknologi yang digunakan

- **Go (Golang)**: Backend Aplikasi
- **PostgreSQL**: Database untuk menyimpan data konsumen dan transaksi.
- **Docker**: Untuk containerisasi aplikasi dan database.
- **Docker Compose**: Untuk mengelola container aplikasi dan database dengan mudah.

## Prasyarat

- Pastikan **Docker** dan **Docker Compose** sudah terinstall di sistem anda. Jika belu,, ikuti petunjuk instalasi di sini:
     - [Docker Install](https://docs.docker.com/get-docker)
     - [Docker Compose install](https://docs.docker.com/compose/install)

## Instalasi

1. **Clone Repository**:

     Clone repositori ini ke mesin lokal anda:
     ```bash
     git clone https://github.com/damshxy/xyz-finance-app.git
     cd xyz-finance-app
     ```

2. **Build Dan Jalankan dengan Docker Compose**:

     Jalankan perintah berikut untuk membangun dan menjalakan aplikasi menggukan Docker Compose:
     ```bash
     sudo docker compose up --build
     ```
     Docker Compose akan membangun aplikasi Go, menjalankan container untuk aplikasi dan database postgreSQL, dan menghubungkan aplikasi ke database.

3. **Akses Aplikasi**:

     Setalah Docker Compose selesai, aplikasi akan berjalan di port **4000**. Anda bisa mengaksesnya melalui browser atau alat API sepert Postman di:
     ```bash
     http://localhost:4000
     ```

## Penggunaan API

### **Consumer API**

1. **Create Consumer**

Endpoint: **POST /consumers**

Body Request:
```json
     {
          "nik": "1204329037983724",
          "full_name": "John Doe",
          "legal_name": "Steve Doe",
          "birth_place": "America",
          "birth_date": "1990-01-01",
          "salary": 5000000.00,
          "ktp_photo": "https://example.com/ktp.jpg",
          "selfie_photo": "https://example.com/selfie.jpg"
     }

```

Response:
```json
     {
          "message": "Consumer created successfully"
     }
```

2. **Get Consumer By ID**

Endpoint: **GET /consumers/:id**

Response:
```json
     {
          "consumer": {
               "id": 2,
               "nik": "1204329037983724",
               "full_name": "John Doe",
               "legal_name": "Steve Doe",
               "birth_place": "America",
               "birth_date": "1990-01-01",
               "salary": 5000000,
               "ktp_photo": "https://example.com/ktp.jpg",
               "selfie_photo": "https://example.com/selfie.jpg",
               "credit_limit": 1000000
          },
          "message": "Consumer retrieved successfully"
     }
```

3. **Update Consumer Credit Limit**

Endpoint: **PATCH /consumers/:id/credit**

Body Request:
```json
     {
          "new_credit_limit": 3000000.00
     }
```

Response:
```json
     {
          "message": "Consumer updated successfully"
     }
```

### **Transaction API**

1. **Create Transaction**

Endpoint: **POST /transactions**

Body Request:
```json
     {
          "consumer_id": 2,
          "contract_number": "CN987654321",
          "otr": 2000000.00,
          "admin_fee": 500000.00,
          "installment": 250000.00,
          "interest": 5.00,
          "asset_name": "Honda Brio"
     }
```

Response:
```json
     {
          "message": "Transaction created successfully"
     }
```

2. **Get Transaction By Consumer ID**

Endpoint: **GET /transactions/consumer/:consumer_id**

Response:
```json
     {
         "message": "Transactions retrieved successfully",
          "transactions": [
               {
                    "id": 1,
                    "consumer_id": 2,
                    "consumer": {
                         "id": 2,
                         "nik": "1204329037983724",
                         "full_name": "John Doe",
                         "legal_name": "Steve Doe",
                         "birth_place": "America",
                         "birth_date": "1990-01-01",
                         "salary": 5000000,
                         "ktp_photo": "https://example.com/ktp.jpg",
                         "selfie_photo": "https://example.com/selfie.jpg",
                         "credit_limit": 1000000
                    },
                    "contract_number": "CN987654321",
                    "otr": 2000000,
                    "admin_fee": 500000,
                    "installment": 250000,
                    "interest": 5,
                    "asset_name": "Honda Brio",
                    "refunded": false
               }
          ]
     }
```

3. **Refund Transaction**

Endpoint: **POST /transaction/:transaction_id/refund**

Response:
```json
     {
          "message": "Transaction marked as refunded successfully"
     }
```