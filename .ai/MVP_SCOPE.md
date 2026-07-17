# MVP SCOPE

Version: 1.0

---

# Purpose

Dokumen ini mendefinisikan batasan MVP.

Semua AI harus menghormati scope ini.

Jika sebuah fitur tidak berada di dalam MVP,

jangan mengimplementasikannya kecuali diminta.

---

# MVP Goal

Membuat aplikasi AI Chat yang nyaman digunakan.

Sudah mendukung roleplay sederhana.

Mudah dikembangkan.

---

# Included

## Backend

- Authentication
- User
- Conversation
- Character
- AI Gateway
- Model Router
- Ollama

---

## Mobile

- Login
- Home
- Chat
- Character
- Settings

---

## AI

- Streaming Chat
- Character Prompt
- Model Switching
- Conversation History

---

## Character

- Name
- Avatar
- Description
- Personality
- Greeting
- Prompt

---

## Database

- PostgreSQL

- pgvector

---

## AI Provider

- Ollama

Cloud provider akan ditambahkan nanti.

---

# Not Included

- Voice Chat
- RAG
- Tool Calling
- Image Generation
- Image Understanding
- Video
- Group Chat
- Multi Character
- Plugin
- MCP
- Cloud Sync
- Desktop
- Web
- Analytics

---

# UI

Fokus pada:

- sederhana
- cepat
- nyaman

Tidak mengejar animasi kompleks.

---

# Success Criteria

MVP dianggap selesai apabila:

User dapat login.

↓

Memilih Character.

↓

Memilih Model.

↓

Mengirim pesan.

↓

Mendapat response streaming.

↓

Percakapan tersimpan.

↓

Dapat membuka percakapan kembali.

Itu sudah cukup.

---

# Out of Scope

Semua fitur di luar dokumen ini dianggap bukan bagian dari MVP.

Tambahkan hanya jika memang diperlukan.