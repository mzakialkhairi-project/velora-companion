# DECISIONS

## ADR-0001

**Decision**

Flutter dipilih sebagai framework mobile.

**Reason**

- Satu codebase.
- Performa baik.
- Dukungan desktop.

**Trade-off**

- Ukuran aplikasi lebih besar dibanding native.

---

## ADR-0002

**Decision**

Go dipilih untuk backend.

**Reason**

- Mudah di-deploy.
- Cepat.
- Cocok sebagai AI Gateway.

---

## ADR-0003

**Decision**

Menggunakan PostgreSQL + pgvector.

**Reason**

- Tidak perlu vector database terpisah.
- Deployment lebih sederhana.

---

## ADR-0004

**Decision**

Gunakan environment variable (.env) sebagai satu-satunya source konfigurasi.

**Library**

`github.com/caarlos0/env/v11`

**Must Not**

- Jangan menggunakan Viper.
- Jangan menggunakan config.yaml.
- Jangan membuat configuration file selain .env.

---

## ADR-0005

**Decision**

Gunakan standard library Go `log/slog` untuk logging.

**Must Not**

- Jangan menggunakan zerolog.
- Jangan menggunakan logrus.
- Jangan menggunakan zap.

---

## ADR-0006

**Decision**

Gunakan `pgx/v5` sebagai PostgreSQL driver dengan `GORM` sebagai ORM.

**Notes**

- MVP tetap menggunakan GORM agar implementasi lebih cepat.
- Jangan menambahkan abstraction tambahan di atas GORM.

---

## ADR-0007

**Decision**

Gunakan `golang-migrate` untuk database migration.

**Status**

Ditunda hingga milestone Database Migration.

---

## ADR-0008

**Decision**

Manual Dependency Injection.

**Must Not**

- Jangan menggunakan Wire.
- Jangan menggunakan Fx.
- Jangan menggunakan Dig.
- Jangan menggunakan Service Locator.

---

## ADR-0009

**Decision**

Tambahkan health check endpoints:

- `GET /` - Informasi dasar aplikasi
- `GET /health` - HTTP 200 jika aplikasi berjalan
- `GET /ready` - HTTP 200 jika PostgreSQL dan Redis siap, HTTP 503 jika tidak

---

## ADR-0010

**Decision**

WAJIB implementasi Graceful Shutdown.

Server harus menutup secara graceful:

- HTTP Server
- PostgreSQL connection
- Redis connection

Ketika menerima SIGINT atau SIGTERM.

---

## ADR-0011

**Deferred**

Dependency berikut belum diperlukan pada Backend Foundation:

- JWT Library (ditunda hingga Authentication)
- UUID Library (ditunda hingga entity pertama dibuat)
