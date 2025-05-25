# Local AI Project

Yerel yapay zeka sistemi - Dokuman yukleme, wiki arama ve AI sorgulama

## Proje Yapisi

```
local-ai-project/
├── backend/          # Go backend API
├── frontend/         # TypeScript/React frontend
├── data/            # Veritabani dosyalari
├── models/          # AI modelleri
├── uploads/         # Yuklenen dokumanlar
└── logs/           # Log dosyalari
```

## Kurulum

### Backend
```bash
cd backend
go mod tidy
go run cmd/server/main.go
```

### Frontend
```bash
cd frontend
npm install
npm start
```

## API Endpoints

- `GET /api/v1/health` - Sistem durumu
- `GET /api/v1/models` - Model listesi
- `POST /api/v1/models/download` - Model indirme
- `POST /api/v1/documents/upload` - Dokuman yukleme
- `POST /api/v1/query` - AI sorgulama
- `GET /api/v1/wiki/search` - Wiki arama

## Teknolojiler

**Backend:** Go, Gin, SQLite, Ollama
**Frontend:** React, TypeScript, Tailwind CSS
