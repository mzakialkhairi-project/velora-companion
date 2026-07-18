# AI Playbook

Version: 1.0

---

# Purpose

Dokumen ini berisi panduan operasional untuk seluruh AI assistant yang bekerja pada project Velora.

Tujuannya adalah menjaga konsistensi implementasi, menghindari over-engineering, dan memastikan setiap perubahan tetap sesuai dengan visi project.

Dokumen ini berlaku untuk seluruh AI, termasuk Claude, GPT, Cline, Devin, maupun Local LLM.

---

# Project Mindset

Velora adalah **personal portfolio project**, bukan enterprise software.

Prioritas utama:

- kode yang bersih
- struktur yang sederhana
- pengalaman pengguna yang baik
- maintainability

Bukan:

- kompleksitas
- banyaknya fitur
- over engineering

---

# Before Starting

Sebelum mengerjakan task:

1. Pahami requirement.
2. Identifikasi tujuan perubahan.
3. Baca dokumentasi yang relevan.
4. Pastikan perubahan masih berada dalam MVP Scope.

---

# Required Reading

Minimal baca:

- PROJECT.md
- ARCHITECTURE.md
- CODING_RULES.md
- CONVENTIONS.md

Jika perubahan memengaruhi ruang lingkup project, baca juga:

- ROADMAP.md
- MVP_SCOPE.md
- DECISIONS.md

---

# Core Principles

Selalu:

- ikuti Clean Architecture
- tulis kode yang mudah dibaca
- gunakan struktur yang konsisten
- minimalkan dependency
- prioritaskan maintainability

Lebih baik implementasi sederhana yang stabil daripada solusi kompleks yang sulit dipahami.

---

# Never Do

Jangan:

- mengubah arsitektur tanpa persetujuan
- mengganti technology stack
- menambah dependency besar tanpa alasan
- membuat folder generic (misc, temp, helper, dll.)
- melakukan refactor besar di luar scope task
- mengimplementasikan fitur di luar MVP tanpa diminta

---

# Preferred Workflow

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

# Decision Guidelines

Jika terdapat beberapa solusi:

1. Pilih yang paling sederhana.
2. Pilih yang paling mudah dipelihara.
3. Pilih yang paling konsisten dengan project.
4. Hindari optimisasi dini.

---

# Code Quality Checklist

Sebelum menyelesaikan task, pastikan:

- requirement terpenuhi
- tidak ada code duplication yang tidak perlu
- nama variabel jelas
- struktur folder tetap rapi
- error handling memadai
- tidak ada dead code
- dokumentasi diperbarui jika diperlukan

---

# Escalation Rules

Hentikan implementasi dan minta keputusan developer apabila:

- terjadi perubahan arsitektur
- membutuhkan dependency baru yang signifikan
- memerlukan breaking change
- requirement tidak jelas atau saling bertentangan

---

# Communication Style

Ketika memberikan rekomendasi:

- jelaskan alasan teknis
- sampaikan trade-off
- hindari jawaban yang terlalu kompleks
- fokus pada solusi yang realistis

---

# Success Criteria

Task dianggap selesai apabila:

- sesuai requirement
- mengikuti dokumentasi project
- mudah dipahami developer lain
- tidak menambah kompleksitas yang tidak perlu
- siap dikembangkan pada iterasi berikutnya

---

# Final Principle

Selalu ingat:

> **Velora adalah project yang mengutamakan kualitas, bukan kuantitas.**

Setiap keputusan harus membuat project menjadi lebih sederhana, lebih konsisten, dan lebih mudah dipelihara.

---

Created for **Velora**

Maintained by **mzakiaklhairi**