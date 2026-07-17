# CONVENTIONS

Version: 1.0

---

# Purpose

Dokumen ini berisi konvensi yang digunakan di seluruh project.

Tujuannya menjaga konsistensi dalam penamaan, struktur, dan gaya implementasi.

Jika tidak ada alasan yang kuat, seluruh AI dan developer harus mengikuti konvensi ini.

---

# General Principles

- Konsisten lebih penting daripada preferensi pribadi.
- Gunakan nama yang mudah dipahami.
- Hindari singkatan yang tidak umum.
- Gunakan bahasa Inggris untuk seluruh source code.
- Gunakan bahasa Indonesia atau Inggris untuk dokumentasi sesuai kebutuhan.

---

# Project Naming

Project Name

AI Roleplay Platform

Repository

ai-roleplay-platform

---

# Folder Naming

Gunakan

lowercase

Contoh

auth

conversation

character

memory

prompt

model

settings

shared

Jangan menggunakan

Temp

Misc

Other

New

Final

Latest

---

# File Naming

Gunakan snake_case.

Contoh

auth_handler.go

conversation_service.go

memory_repository.go

character_mapper.go

prompt_builder.go

Flutter mengikuti standar Dart.

Contoh

chat_screen.dart

character_card.dart

conversation_service.dart

---

# Go Package

Gunakan lowercase.

Contoh

package auth

package conversation

package memory

Jangan menggunakan underscore pada nama package.

---

# Struct Naming

Gunakan PascalCase.

Contoh

Conversation

ConversationService

PromptBuilder

CharacterProfile

---

# Interface Naming

Jangan selalu menambahkan suffix "Interface".

Lebih baik

ConversationRepository

ModelProvider

MemoryStore

dibanding

ConversationRepositoryInterface

---

# Variable Naming

Gunakan nama yang jelas.

Contoh

conversationID

characterID

currentUser

memoryEntries

Hindari

tmp

obj

data

value

item

---

# Constant Naming

Gunakan PascalCase untuk exported.

Gunakan camelCase untuk private.

Jika berkaitan dengan environment variable,

gunakan UPPER_SNAKE_CASE.

---

# Environment Variable

Contoh

DATABASE_URL

REDIS_URL

JWT_SECRET

OLLAMA_URL

MINIO_ENDPOINT

APP_ENV

APP_PORT

---

# REST API

Gunakan noun.

Contoh

/api/v1/conversations

/api/v1/characters

/api/v1/models

/api/v1/settings

Hindari

/getConversation

/createCharacter

/deleteMemory

Gunakan HTTP Method.

---

# HTTP Method

GET

Mengambil data

POST

Membuat data

PUT

Mengganti seluruh data

PATCH

Mengubah sebagian data

DELETE

Menghapus data

---

# JSON

Gunakan

camelCase

Contoh

conversationId

characterName

createdAt

updatedAt

---

# Database

Gunakan snake_case.

Contoh

conversation_messages

character_memories

user_profiles

Gunakan jamak (plural) untuk nama tabel.

---

# Primary Key

Semua tabel menggunakan UUID.

Kolom primary key selalu bernama

id

---

# Foreign Key

Gunakan format

character_id

conversation_id

user_id

---

# Timestamp

Semua entity memiliki

created_at

updated_at

deleted_at (jika menggunakan soft delete)

---

# Migration

Gunakan format

YYYYMMDDHHMMSS_description.sql

Contoh

20260718093000_create_users.sql

---

# Docker

Nama service menggunakan lowercase.

Contoh

backend

postgres

redis

minio

ollama

mailpit

---

# Git Branch

Gunakan format

feature/

fix/

refactor/

docs/

test/

chore/

Contoh

feature/character-memory

fix/login-refresh-token

docs/update-ai-context

---

# Commit Message

Gunakan Conventional Commits.

Contoh

feat(character): add memory support

fix(auth): refresh token validation

docs(ai): update architecture

refactor(memory): simplify service

---

# API Version

Gunakan prefix

/api/v1/

Jika terjadi breaking change,

buat versi baru.

Contoh

/api/v2/

---

# Error Code

Gunakan struktur yang konsisten.

Contoh

AUTH_INVALID_TOKEN

CHARACTER_NOT_FOUND

MEMORY_LIMIT_EXCEEDED

MODEL_UNAVAILABLE

---

# Logging

Gunakan structured logging.

Minimal berisi

timestamp

level

requestId

module

message

---

# Configuration

Seluruh konfigurasi berasal dari environment variable.

Tidak boleh ada credential di source code.

---

# Documentation

Dokumen menggunakan Markdown.

Judul menggunakan Title Case.

Gunakan heading secara konsisten.

---

# AI Reminder

Ketika AI membuat file baru:

- Ikuti konvensi ini.
- Jangan membuat style baru.
- Jangan mencampur beberapa gaya penamaan.
- Prioritaskan konsistensi.

---

# Final Principle

Developer seharusnya dapat menebak nama file, nama endpoint, dan nama package tanpa perlu membuka dokumentasi.

Jika penamaan terasa membingungkan, ubah menjadi lebih sederhana.