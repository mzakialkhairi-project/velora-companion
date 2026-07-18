# DEVELOPMENT PRINCIPLES

Version: 1.0

---

# Purpose

Dokumen ini menjelaskan filosofi pengembangan project.

Tujuannya bukan membuat software yang paling kompleks.

Tujuannya adalah membuat software yang nyaman dikembangkan, mudah dipahami, dan menyenangkan untuk dipelihara.

Semua keputusan engineering harus mengacu pada prinsip-prinsip berikut.

---

# Core Philosophy

Software dibuat untuk manusia.

Code dibaca lebih sering daripada ditulis.

Karena itu, keterbacaan selalu lebih penting daripada kecerdikan.

---

# Keep It Simple

Pilih solusi paling sederhana yang dapat menyelesaikan masalah.

Jangan membuat sistem yang rumit hanya untuk mengantisipasi kemungkinan yang belum tentu terjadi.

Kompleksitas harus muncul karena kebutuhan nyata, bukan asumsi.

---

# Build Only What Is Needed

Project berkembang secara bertahap.

Jangan membuat fitur yang belum digunakan.

Jangan membuat abstraction yang belum memiliki minimal dua penggunaan nyata.

Hindari over engineering.

---

# Readability First

Code harus mudah dipahami.

Jika seseorang membutuhkan waktu lama untuk memahami sebuah fungsi, berarti fungsi tersebut terlalu rumit.

Prioritaskan:

- nama yang jelas
- struktur yang konsisten
- logika yang sederhana

---

# Consistency Over Preference

Gunakan pola yang sudah ada.

Jangan mengganti style hanya karena memiliki preferensi pribadi.

Consistency lebih penting daripada gaya masing-masing developer.

---

# Modular Thinking

Pisahkan fitur berdasarkan domain.

Setiap module memiliki tanggung jawab yang jelas.

Module tidak boleh mengetahui implementasi internal module lain.

---

# Small Changes

Lebih baik membuat perubahan kecil yang selesai daripada perubahan besar yang belum selesai.

Setiap commit sebaiknya memiliki tujuan yang jelas.

---

# Refactor Continuously

Refactor adalah bagian dari pengembangan.

Namun:

Jangan melakukan refactor tanpa alasan.

Refactor dilakukan ketika:

- code mulai sulit dipahami
- duplication mulai muncul
- tanggung jawab mulai bercampur

---

# Avoid Premature Optimization

Optimasi dilakukan setelah ditemukan bottleneck.

Jangan mengorbankan keterbacaan hanya demi optimasi yang belum dibutuhkan.

---

# Prefer Composition

Gunakan composition daripada inheritance.

Pisahkan tanggung jawab menjadi komponen kecil yang dapat digunakan kembali.

---

# One Responsibility

Setiap:

- package
- module
- struct
- service
- function

sebaiknya memiliki satu tanggung jawab utama.

---

# Explicit Is Better Than Implicit

Hindari code yang terlalu "magic".

Lebih baik sedikit verbose daripada sulit dipahami.

---

# AI Friendly Code

Project ini dikembangkan bersama AI.

Karena itu struktur project harus mudah dipahami oleh manusia maupun AI.

Gunakan nama yang jelas.

Hindari singkatan yang tidak umum.

Pisahkan domain dengan rapi.

---

# Documentation

Dokumentasi dibuat ketika memang membantu.

Tidak semua code membutuhkan dokumentasi panjang.

Yang perlu didokumentasikan adalah:

- keputusan penting
- arsitektur
- trade-off
- API publik

---

# Error Handling

Jangan mengabaikan error.

Selalu berikan pesan error yang membantu developer.

Error harus cukup informatif untuk proses debugging.

---

# Logging

Log digunakan untuk membantu observability.

Jangan menggunakan log sebagai pengganti error handling.

Gunakan level log secara konsisten.

---

# Testing

Testing dilakukan pada bagian yang memiliki logika penting.

Tidak semua fungsi sederhana membutuhkan unit test.

Prioritaskan kualitas test dibanding jumlah test.

---

# Dependencies

Sebelum menambahkan dependency baru, tanyakan:

- Apakah Standard Library sudah cukup?
- Apakah dependency masih aktif dikembangkan?
- Apakah dependency benar-benar dibutuhkan?

Jika ragu, jangan tambahkan.

---

# Git

Commit harus kecil.

Commit harus memiliki tujuan.

Hindari commit yang mencampur banyak perubahan berbeda.

---

# Performance

Performance penting.

Namun maintainability lebih penting.

Optimasi hanya dilakukan berdasarkan data.

Bukan asumsi.

---

# Security

Jangan hardcode credential.

Selalu gunakan environment variable.

Validasi seluruh input dari user.

Gunakan prinsip least privilege.

---

# User Experience

Teknologi bukan tujuan.

User experience adalah tujuan.

Keputusan engineering harus mempertimbangkan kenyamanan pengguna.

---

# Continuous Improvement

Project ini akan terus berkembang.

Jika ditemukan cara yang lebih baik:

- diskusikan
- evaluasi trade-off
- dokumentasikan keputusan

Jangan mengubah arah project secara mendadak.

---

# AI Reminder

Ketika membantu project ini:

- Jangan over engineering.
- Jangan membuat abstraction yang belum diperlukan.
- Ikuti pola yang sudah ada.
- Jelaskan trade-off ketika memberikan rekomendasi.
- Prioritaskan maintainability.
- Tulis code yang akan mudah dipahami enam bulan dari sekarang.

---

# Definition of Good Code

Good code adalah code yang:

- Mudah dibaca.
- Mudah diuji.
- Mudah diubah.
- Mudah di-debug.
- Mudah dijelaskan kepada orang lain.

Bukan code yang paling pendek.

Bukan code yang paling pintar.

Melainkan code yang paling mudah dipelihara.

---

# Final Principle

Setiap keputusan engineering harus menjawab satu pertanyaan:

"Apakah keputusan ini akan membuat project lebih mudah dikembangkan enam bulan dari sekarang?"

Jika jawabannya tidak yakin, pilih solusi yang lebih sederhana.