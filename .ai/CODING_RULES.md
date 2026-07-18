# CODING RULES

Version: 1.0

---

# Purpose

Dokumen ini mendefinisikan standar penulisan kode untuk seluruh project.

Tujuannya adalah menghasilkan code yang:

- Konsisten
- Mudah dibaca
- Mudah dipelihara
- Mudah dipahami oleh manusia maupun AI

Rule ini berlaku untuk seluruh repository.

---

# General Philosophy

Code adalah komunikasi.

Tulis code seolah-olah orang yang akan membacanya adalah dirimu sendiri enam bulan dari sekarang.

Prioritaskan kejelasan dibanding kepintaran.

---

# Readability

Gunakan nama yang deskriptif.

Hindari nama seperti:

data

temp

obj

test

item

foo

bar

Gunakan nama yang menjelaskan isi.

Contoh:

conversationRepository

memoryService

characterProfile

---

# Function

Setiap function hanya memiliki satu tanggung jawab.

Jika function mulai sulit dijelaskan dalam satu kalimat, pecah menjadi beberapa function.

Function sebaiknya:

- pendek
- fokus
- mudah diuji

---

# Naming

Gunakan bahasa Inggris.

Gunakan PascalCase untuk:

Struct

Interface

Public Function

Gunakan camelCase untuk:

Variable

Private Function

Gunakan snake_case hanya jika mengikuti kebutuhan tertentu seperti database.

---

# File Naming

Gunakan nama file yang sederhana.

Contoh:

service.go

repository.go

handler.go

entity.go

mapper.go

validator.go

routes.go

Jangan membuat nama file yang terlalu panjang.

---

# Comments

Comment digunakan untuk menjelaskan alasan.

Bukan menjelaskan apa yang sudah terlihat jelas.

Hindari:

// increment i

i++

Lebih baik:

// Retry only for transient errors.

---

# Magic Value

Jangan hardcode value penting.

Gunakan:

constant

configuration

environment variable

sesuai kebutuhan.

---

# Duplication

Jika code mulai di-copy lebih dari satu kali,

evaluasi apakah memang perlu dibuat reusable.

Namun jangan membuat abstraction terlalu cepat.

---

# Error Handling

Selalu tangani error.

Jangan mengabaikan error.

Pesan error harus membantu debugging.

Gunakan wrapping jika diperlukan.

---

# Logging

Gunakan logging seperlunya.

Log harus membantu debugging.

Jangan memenuhi log dengan informasi yang tidak berguna.

---

# Validation

Semua input dari user harus divalidasi.

Validation dilakukan sedekat mungkin dengan titik masuk data.

---

# Configuration

Semua konfigurasi berasal dari environment variable.

Tidak boleh ada credential di source code.

---

# API Response

Gunakan format response yang konsisten.

Success

Error

Validation Error

Unauthorized

Not Found

Internal Error

harus memiliki struktur yang seragam.

---

# Dependency

Gunakan dependency sesedikit mungkin.

Sebelum menambahkan library baru,

jawab pertanyaan berikut:

Apakah Standard Library sudah cukup?

Apakah library benar-benar diperlukan?

Apakah library masih aktif dikembangkan?

---

# Module

Setiap module memiliki tanggung jawab yang jelas.

Jangan mencampurkan logic conversation ke module character.

Jangan mencampurkan memory ke authentication.

---

# Repository

Repository hanya berinteraksi dengan database.

Tidak boleh berisi business logic.

---

# Service

Service adalah tempat seluruh business logic.

Semua keputusan aplikasi berada di sini.

---

# Handler

Handler hanya bertugas:

- menerima request
- validasi awal
- memanggil service
- mengembalikan response

---

# DTO

Gunakan DTO untuk komunikasi antar layer.

Jangan mengembalikan entity database langsung ke client.

---

# Entity

Entity merepresentasikan domain.

Entity bukan response API.

---

# Mapper

Gunakan mapper jika transformasi mulai kompleks.

Hindari transformasi besar di handler.

---

# Testing

Prioritaskan test pada:

Business Logic

Memory

Prompt Builder

Model Router

Authentication

---

# AI Generated Code

AI boleh menghasilkan code.

Namun AI harus:

- mengikuti struktur project
- tidak membuat duplicate logic
- tidak mengubah architecture tanpa alasan
- menjelaskan trade-off jika mengusulkan perubahan besar

---

# Refactoring

Refactor dilakukan ketika:

- duplication muncul
- code sulit dibaca
- tanggung jawab bercampur

Jangan melakukan refactor hanya karena preferensi pribadi.

---

# Code Review Checklist

Sebelum menganggap task selesai:

- Apakah code mudah dipahami?
- Apakah ada duplication?
- Apakah naming sudah jelas?
- Apakah error ditangani?
- Apakah architecture tetap konsisten?
- Apakah module lain tidak ikut rusak?

Jika ada jawaban "tidak",

perbaiki sebelum melanjutkan.

---

# Final Rule

Tulis code yang sederhana.

Jika ada dua solusi,

pilih solusi yang lebih mudah dipahami.