# TECH STACK

Version: 1.0

---

# Purpose

Dokumen ini menjelaskan teknologi yang digunakan dalam project beserta alasan pemilihannya.

Tujuan utama dokumen ini adalah menjaga konsistensi teknis selama pengembangan.

AI tidak boleh mengganti teknologi utama tanpa alasan yang kuat.

---

# Technology Philosophy

Pemilihan teknologi didasarkan pada prinsip berikut:

- Sederhana dipahami.
- Stabil dan mature.
- Memiliki komunitas yang besar.
- Mudah dipelajari.
- Mudah dipelihara.
- Mendukung pengembangan jangka panjang.
- Tidak memilih teknologi hanya karena sedang populer.

---

# Architecture Style

Current Architecture

Modular Monolith

Reason

- Cocok untuk personal project.
- Mudah dikembangkan.
- Mudah dipahami.
- Mudah dipecah menjadi microservices jika suatu saat diperlukan.

Current Status

Recommended

---

# Backend

Language

Go

Framework

Gin

Reason

- Performa tinggi.
- Concurrency sangat baik.
- Binary tunggal.
- Deployment sederhana.
- Sangat cocok sebagai AI Gateway.

Alternatives

Echo

Fiber

Chi

Current Choice

Gin

---

# ORM

Selected

GORM

Reason

- Dokumentasi baik.
- Komunitas besar.
- Mudah digunakan.
- Cocok untuk project pribadi.

Alternatives

SQLC

Ent

Bun

Future

Jika performa query menjadi masalah, SQLC dapat digunakan secara bertahap tanpa mengganti keseluruhan project.

---

# Database

Selected

PostgreSQL

Reason

- Open Source.
- Sangat stabil.
- Mendukung JSON.
- Mendukung pgvector.
- Cocok untuk AI.

Current Version

PostgreSQL 17

---

# Vector Database

Selected

pgvector

Reason

Tidak menggunakan database vector terpisah.

Semua data aplikasi dan embedding berada di PostgreSQL.

Keuntungan:

- Deployment lebih sederhana.
- Backup lebih mudah.
- Tidak perlu sinkronisasi dua database.

---

# Cache

Selected

Redis

Purpose

- Session
- Cache
- Queue ringan
- Streaming

---

# Object Storage

Selected

MinIO

Purpose

- Avatar
- Character Image
- User Upload
- Voice
- Attachment

Reason

S3 Compatible.

Mudah dipindahkan ke AWS S3 di masa depan.

---

# Mobile

Framework

Flutter

Reason

Single codebase.

Android dan iOS.

Performa tinggi.

UI konsisten.

---

# Admin Dashboard

Framework

Next.js

Reason

Cepat dibuat.

Ekosistem React sangat besar.

Mudah membuat dashboard.

Status

Optional.

Tidak menjadi prioritas awal.

---

# AI Engine

Current

Ollama

Reason

- Instalasi mudah.
- Mendukung banyak model.
- API sederhana.
- Cocok untuk development.

Future

llama.cpp

vLLM

SGLang

OpenAI

Claude

Gemini

Semua diakses melalui abstraction layer.

---

# Default Models

Roleplay

Qwen2.5-Kunou-14B

General Chat

Qwen3

Coding

Claude

GPT

Embedding

BGE-M3

---

# AI Provider Strategy

Project tidak bergantung pada satu provider.

Semua provider harus dapat ditukar melalui Model Router.

Contoh

Local

- Ollama

Cloud

- OpenAI
- Anthropic
- Gemini
- OpenRouter

---

# Authentication

JWT

Reason

Sederhana.

Stateless.

Cocok untuk mobile.

---

# API Style

REST API

Streaming menggunakan

Server Sent Events (SSE)

atau

WebSocket

Gunakan REST sebagai default.

Gunakan streaming hanya untuk AI Response.

---

# Communication

Client

↓

REST

↓

Backend

↓

Model Router

↓

LLM

---

# Container

Selected

Docker

Docker Compose

Reason

Semua dependency dijalankan menggunakan container.

Tidak ada instalasi manual selain Docker.

---

# Development Environment

Required

Docker Desktop

Go

Flutter

Git

VS Code

Recommended Extensions

Cline

GitHub Copilot (optional)

Docker

Go

Flutter

---

# Project Structure

Backend

/backend

Flutter

/mobile

AI Context

/.ai

Deployment

/docker

---

# Configuration

Semua konfigurasi menggunakan environment variable.

Contoh

DATABASE_URL

REDIS_URL

OLLAMA_URL

JWT_SECRET

MINIO_ENDPOINT

Tidak boleh ada credential di source code.

---

# Logging

Gunakan structured logging.

Level

Debug

Info

Warn

Error

Panic

---

# Testing

Manual Testing

Required

Unit Test

Recommended

Integration Test

Recommended

End-to-End Test

Future

---

# Dependency Policy

Gunakan dependency seminimal mungkin.

Sebelum menambahkan library baru:

- Apakah bisa menggunakan Standard Library?
- Apakah dependency masih aktif?
- Apakah komunitasnya besar?
- Apakah benar-benar diperlukan?

Jika jawabannya tidak jelas, jangan gunakan dependency tersebut.

---

# Upgrade Policy

Gunakan versi stabil.

Hindari menggunakan release candidate atau experimental pada project utama.

Upgrade dilakukan secara bertahap.

---

# Future Technology

Kemungkinan akan ditambahkan:

- Voice Recognition
- Text to Speech
- Image Generation
- OCR
- MCP Support
- Tool Calling
- Plugin System

Namun tidak menjadi bagian dari MVP.

---

# AI Notes

Ketika AI membuat kode:

- Ikuti tech stack ini.
- Jangan mengganti framework tanpa alasan.
- Jangan menambahkan dependency baru jika belum diperlukan.
- Selalu jelaskan trade-off ketika mengusulkan perubahan teknologi.
- Prioritaskan konsistensi dibanding mencoba teknologi terbaru.