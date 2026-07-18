# ARCHITECTURE

Version: 1.0

---

# Purpose

Dokumen ini menjelaskan struktur arsitektur project secara keseluruhan.

Tujuannya adalah agar semua AI assistant memahami bagaimana setiap bagian saling berhubungan sebelum mulai membuat kode.

Architecture harus mudah dipahami.

Architecture boleh berkembang.

Namun perubahan besar harus tetap menjaga konsistensi project.

---

# Architecture Style

Project menggunakan pendekatan:

Modular Monolith

Setiap fitur dipisahkan menjadi module yang memiliki tanggung jawab sendiri.

Semua module berjalan dalam satu backend.

Jika suatu saat dibutuhkan, module dapat dipisahkan menjadi service terpisah tanpa mengubah API publik.

---

# High Level Architecture

                    Flutter Mobile
                           │
                           │
                    REST / Streaming
                           │
                           ▼
                  Go Backend API
                           │
     ┌─────────────────────┼─────────────────────┐
     │                     │                     │
 Conversation         Character            Authentication
     │                     │                     │
     ├──────────────┐       │                     │
     ▼              ▼       ▼                     ▼
 Memory         Prompt Engine          User Management
     │
     ▼
 Model Router
     │
     ├──────── Ollama
     ├──────── OpenAI
     ├──────── Claude
     ├──────── Gemini
     └──────── OpenRouter

---

# Project Layers

Presentation

↓

Application

↓

Domain

↓

Infrastructure

Namun project tidak menggunakan Clean Architecture yang terlalu ketat.

Tujuan utama adalah code mudah dipahami.

---

# Backend Modules

Backend dibagi berdasarkan domain.

Contoh:

auth

conversation

character

memory

prompt

model

embedding

rag

settings

user

shared

Setiap module memiliki folder sendiri.

---

# Standard Module Structure

module/

handler/

service/

repository/

entity/

dto/

routes/

errors/

validator/

mapper/

Module boleh berkembang sesuai kebutuhan.

Namun struktur dasar harus tetap konsisten.

---

# Backend Flow

Mobile

↓

HTTP Request

↓

Router

↓

Handler

↓

Service

↓

Repository

↓

Database

↓

Response

Business logic hanya berada pada Service.

Repository hanya bertanggung jawab terhadap database.

---

# AI Flow

User Message

↓

Conversation

↓

Memory

↓

Prompt Builder

↓

Character Context

↓

Model Router

↓

LLM

↓

Streaming Response

↓

Conversation Saved

---

# Model Router

Model Router bertugas memilih model AI.

Contoh:

Roleplay

↓

Qwen2.5 Kunou

General Chat

↓

Qwen3

Cloud Chat

↓

Claude

Coding

↓

Claude / GPT

Seluruh backend tidak boleh mengetahui provider secara langsung.

Backend hanya berbicara dengan Model Router.

---

# Character System

Setiap Character memiliki:

Identity

Personality

Greeting

Speaking Style

System Prompt

Knowledge

Memory

Relationship

Avatar

World Information

Character tidak mengetahui implementasi model AI.

Character hanya menghasilkan context.

---

# Prompt Builder

Prompt Builder menggabungkan:

System Prompt

+

Character

+

World

+

Memory

+

Conversation

+

User Message

↓

Final Prompt

Prompt Builder adalah satu-satunya tempat yang bertanggung jawab membangun prompt.

---

# Memory System

Memory dibagi menjadi beberapa bagian.

Conversation Memory

Riwayat chat saat ini.

Long-term Memory

Informasi penting yang ingin diingat.

Character Memory

Informasi khusus karakter.

Summary Memory

Ringkasan percakapan panjang.

Memory dapat berkembang tanpa mengubah module lain.

---

# Database

PostgreSQL digunakan sebagai database utama.

Semua entity disimpan di PostgreSQL.

Embedding menggunakan pgvector.

Tidak menggunakan database vector terpisah.

---

# Storage

MinIO digunakan untuk:

Avatar

Attachment

Image

Voice

Future Assets

---

# Cache

Redis digunakan untuk:

Session

Cache

Queue ringan

Temporary data

---

# API Design

REST digunakan sebagai default.

Streaming digunakan hanya untuk AI response.

Semua endpoint menggunakan format response yang konsisten.

---

# Mobile Architecture

Flutter

↓

Presentation

↓

Application

↓

Data

↓

API

↓

Backend

State management dipilih berdasarkan kebutuhan.

Belum diputuskan pada tahap awal.

---

# Authentication

JWT

Refresh Token

Session dikelola backend.

---

# Configuration

Semua konfigurasi menggunakan environment variable.

Tidak boleh ada credential di source code.

---

# Dependency Rule

Module hanya boleh bergantung pada:

Shared

Interface

Common Utility

Module tidak boleh mengakses implementasi internal module lain.

---

# Error Handling

Semua error harus:

Jelas

Konsisten

Mudah dipahami

Tidak membocorkan informasi sensitif.

---

# Logging

Gunakan structured logging.

Setiap request penting memiliki log.

Jangan membuat log yang berlebihan.

---

# Scalability

Architecture dirancang agar mudah berkembang.

Namun:

Sederhana lebih penting daripada scalable sejak awal.

Optimasi dilakukan ketika benar-benar diperlukan.

---

# Future Expansion

Voice Chat

Text to Speech

Speech to Text

Image Understanding

Image Generation

Plugin

Tool Calling

Desktop Client

Web Client

Cloud Sync

Semua fitur baru harus mengikuti architecture yang sudah ada.

---

# Architecture Principles

- Module memiliki satu tanggung jawab utama.
- Business logic tidak berada di Handler.
- Database tidak diakses langsung dari Handler.
- Prompt hanya dibuat oleh Prompt Builder.
- Model hanya dipanggil melalui Model Router.
- Memory tidak boleh tersebar di banyak tempat.
- Setiap layer memiliki tanggung jawab yang jelas.
- Struktur project harus mudah dipahami oleh manusia maupun AI.

---

# Final Rule

Jika terdapat lebih dari satu solusi arsitektur, pilih solusi yang:

- paling sederhana,
- paling konsisten,
- paling mudah dipelihara,
- dan paling sesuai dengan tujuan project.