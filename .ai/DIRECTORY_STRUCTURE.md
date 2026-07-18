# DIRECTORY STRUCTURE

Version: 1.0

---

# Purpose

Dokumen ini menjelaskan struktur repository project.

Tujuannya adalah menjaga repository tetap rapi, mudah dipahami, dan mudah dikembangkan.

Semua AI assistant harus mengikuti struktur ini ketika membuat file atau folder baru.

---

# Repository Philosophy

Repository harus:

- sederhana
- konsisten
- mudah dinavigasi
- mudah dipahami developer baru
- mudah dipahami AI

Folder dibuat berdasarkan tanggung jawab.

Bukan berdasarkan kebiasaan.

---

# Root Structure

```
ai-roleplay-platform/

│
├── .ai/
├── backend/
├── mobile/
├── docker/
├── docs/
├── scripts/
├── tools/
├── assets/
│
├── README.md
├── LICENSE
├── .gitignore
├── .dockerignore
└── docker-compose.yml
```

---

# Root Folder Responsibilities

## .ai/

Seluruh konteks AI.

Tidak berisi source code.

---

## backend/

Seluruh source code backend Go.

---

## mobile/

Seluruh source code Flutter.

---

## docker/

Dockerfile dan konfigurasi container.

---

## docs/

Dokumentasi project.

Tidak termasuk AI Context.

---

## scripts/

Automation script.

Misalnya:

- build
- migrate
- seed
- backup

---

## tools/

Utility kecil yang membantu development.

Contoh:

generator

converter

migration helper

---

## assets/

Asset project.

Logo.

Placeholder.

Demo image.

---

# Backend Structure

```
backend/

├── cmd/
├── configs/
├── internal/
├── migrations/
├── seeders/
├── storage/
├── tests/
│
├── Dockerfile
├── go.mod
└── go.sum
```

---

# cmd/

Entry point aplikasi.

Contoh

api

worker

cli

Setiap executable memiliki folder sendiri.

---

# configs/

Configuration default.

Template config.

Bukan credential.

---

# internal/

Seluruh business logic.

Merupakan folder terbesar.

---

# migrations/

SQL migration.

---

# seeders/

Initial data.

Development data.

---

# storage/

Temporary storage.

Upload lokal.

Development only.

---

# tests/

Integration test.

E2E test.

Testing helper.

---

# Internal Structure

```
internal/

├── auth/
├── user/
├── character/
├── conversation/
├── memory/
├── prompt/
├── model/
├── embedding/
├── rag/
├── settings/
├── shared/
└── bootstrap/
```

---

# Module Philosophy

Setiap module merepresentasikan satu domain.

Contoh

conversation

bukan

conversation_handler

conversation_service

---

# Module Structure

```
conversation/

handler.go

service.go

repository.go

entity.go

dto.go

mapper.go

validator.go

routes.go
```

Jika module berkembang,

boleh dipecah menjadi beberapa file.

Namun tetap berada di folder yang sama.

---

# Shared

Folder shared hanya berisi hal yang benar-benar digunakan lintas module.

Contoh

logger

config

errors

response

utils yang benar-benar generic

Jangan menjadikan shared sebagai tempat "membuang" code.

---

# Bootstrap

Berisi proses startup aplikasi.

Contoh

database

redis

router

middleware

dependency injection

---

# Mobile Structure

```
mobile/

lib/

assets/

android/

ios/

web/

linux/

macos/

windows/

test/

pubspec.yaml
```

---

# Flutter lib/

```
lib/

core/

features/

shared/

main.dart
```

---

# Features

Contoh

chat

character

settings

profile

login

home

Setiap feature berdiri sendiri.

---

# Feature Structure

```
chat/

presentation/

application/

domain/

data/
```

Gunakan hanya jika feature mulai besar.

Jika masih kecil,

buat sederhana.

---

# Docs

```
docs/

api/

database/

deployment/

architecture/

images/
```

---

# Docker

```
docker/

backend/

postgres/

redis/

minio/

ollama/

mailpit/
```

Setiap service memiliki Dockerfile sendiri jika diperlukan.

---

# Scripts

Contoh

```
scripts/

dev.sh

build.sh

seed.sh

backup.sh

reset.sh
```

---

# Naming Rules

Gunakan:

lowercase

snake_case

untuk folder.

Gunakan nama yang singkat.

Jangan menggunakan:

misc

other

temp

new

final

latest

---

# Creating New Modules

Buat module baru jika:

- memiliki domain sendiri
- memiliki business logic sendiri
- akan berkembang

Jangan membuat module baru hanya untuk satu file.

---

# Creating New Folders

Sebelum membuat folder baru,

tanyakan:

Apakah folder yang ada sudah cukup?

Apakah folder baru meningkatkan keterbacaan?

Apakah folder baru benar-benar memiliki tanggung jawab sendiri?

---

# Repository Growth

Project boleh berkembang.

Namun struktur repository harus tetap sederhana.

Jika struktur mulai membingungkan,

lakukan refactoring.

---

# AI Reminder

Saat membuat file baru:

- ikuti struktur repository
- jangan membuat folder baru tanpa alasan
- jangan membuat folder generic
- tempatkan code sedekat mungkin dengan domainnya
- prioritaskan keterbacaan dibanding jumlah folder

---

# Final Principle

Repository yang baik adalah repository yang dapat dipahami dalam beberapa menit oleh developer baru.

Jika sebuah folder membutuhkan penjelasan panjang, evaluasi kembali apakah strukturnya sudah tepat.