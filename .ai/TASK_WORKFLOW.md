# TASK WORKFLOW

Version: 1.0

---

# Purpose

Dokumen ini menjelaskan bagaimana AI menyelesaikan sebuah task.

Tujuannya adalah menjaga kualitas implementasi sekaligus menghindari perubahan yang tidak perlu.

Workflow ini berlaku untuk seluruh jenis task.

- Feature
- Bug Fix
- Refactoring
- Documentation
- Optimization

---

# General Workflow

Semua task mengikuti urutan berikut.

Understand

↓

Analyze

↓

Plan

↓

Implement

↓

Review

↓

Complete

---

# Step 1 — Understand

Pahami permintaan user.

Pastikan mengetahui:

- tujuan task
- hasil yang diharapkan
- batasan yang diberikan

Jika ada informasi penting yang belum tersedia,

bertanya terlebih dahulu.

Jangan berasumsi.

---

# Step 2 — Analyze

Cari bagian project yang berkaitan.

Periksa:

- module
- service
- repository
- entity
- API
- database
- configuration

Pahami bagaimana fitur saat ini bekerja.

Jangan langsung membuat implementasi baru.

---

# Step 3 — Plan

Sebelum coding,

buat rencana singkat.

Contoh:

Task

↓

Module yang akan diubah

↓

File yang terdampak

↓

Pendekatan implementasi

↓

Risiko

Rencana tidak perlu panjang.

Yang penting jelas.

---

# Step 4 — Implement

Saat implementasi:

Ikuti architecture.

Ikuti coding rules.

Ikuti tech stack.

Gunakan kembali code yang sudah ada jika memungkinkan.

---

# Step 5 — Review

Sebelum task dianggap selesai,

periksa kembali.

Checklist:

Apakah code mudah dipahami?

Apakah ada duplication?

Apakah naming sudah konsisten?

Apakah architecture tetap terjaga?

Apakah ada bug yang terlihat?

---

# Step 6 — Complete

Jika implementasi selesai,

berikan ringkasan.

Ringkasan sebaiknya berisi:

Apa yang diubah.

Kenapa diubah.

Apakah ada dampak ke module lain.

Apakah ada pekerjaan lanjutan yang disarankan.

---

# Workflow for New Features

Ketika membuat feature baru:

1.

Pahami feature.

2.

Cari module yang sesuai.

3.

Tambahkan code pada module tersebut.

4.

Hindari membuat module baru jika belum diperlukan.

---

# Workflow for Bug Fix

1.

Identifikasi akar masalah.

2.

Jelaskan penyebab.

3.

Perbaiki akar masalah.

4.

Jangan hanya memperbaiki gejalanya.

---

# Workflow for Refactoring

Refactoring dilakukan jika:

- duplication muncul
- readability menurun
- responsibility bercampur

Refactoring bukan tujuan.

Refactoring mendukung maintainability.

---

# Workflow for Performance

Optimasi dilakukan berdasarkan data.

Bukan dugaan.

Jika tidak ada bottleneck,

hindari optimasi besar.

---

# Workflow for Documentation

Perbarui dokumentasi hanya jika:

- API berubah
- Architecture berubah
- Workflow berubah
- Decision penting berubah

---

# Decision Rules

Jika terdapat beberapa solusi,

urutkan berdasarkan:

1.

Sederhana

2.

Konsisten

3.

Maintainable

4.

Performa

5.

Scalable

---

# Before Creating New Files

Selalu tanyakan:

Apakah file baru benar-benar diperlukan?

Apakah file lama masih bisa digunakan?

Apakah penambahan file membuat project lebih mudah dipahami?

---

# Before Adding Dependencies

Selalu evaluasi:

Apakah Standard Library cukup?

Apakah dependency aktif dikembangkan?

Apakah dependency benar-benar dibutuhkan?

---

# Final Checklist

Sebelum menyatakan task selesai:

✓ Requirement terpenuhi.

✓ Code mudah dipahami.

✓ Tidak ada duplication yang jelas.

✓ Struktur project tetap konsisten.

✓ Tidak menambah kompleksitas yang tidak perlu.

✓ Penjelasan implementasi sudah diberikan.

Jika semua terpenuhi,

task dapat dianggap selesai.