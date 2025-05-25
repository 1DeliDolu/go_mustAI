#!/bin/bash

# Local AI Project Setup Script
echo "🚀 Local AI Project kurulum başlıyor..."

# Ana proje dizini oluştur
mkdir -p local-ai-project
cd local-ai-project

echo "📁 Proje yapisi olusturuluyor..."

# Backend (Go) yapısı
mkdir -p backend/{cmd/server,internal/{handlers,services,models,storage,config},pkg/{utils,types},api,docs,scripts}

# Frontend (TypeScript/React) yapısı
mkdir -p frontend/{src/{components,pages,services,hooks,types,utils,store},public}

# Diğer dizinler
mkdir -p {data,models,uploads,logs}

echo "✅ Dizin yapisi olusturuldu"

# Backend dosyaları oluştur
echo "⚙️ Backend dosyalari olusturuluyor..."

# Go mod init
cd backend
go mod init local-ai-project/backend

# Backend dosyalarını oluştur (aşağıda detayları var)
cd ../

# Frontend setup
echo "🌐 Frontend setup..."
cd frontend
npx create-react-app . --template typescript --yes 2>/dev/null || echo "React app zaten mevcut"
cd ../

echo "📋 README dosyalari olusturuluyor..."

# Ana README
cat > README.md << 'EOF'
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
EOF

echo "✅ Temel proje yapisi olusturuldu!"
echo "📍 Dizin: $(pwd)"
echo ""
echo "Sonraki adimlar:"
echo "1. cd local-ai-project/backend && go mod tidy"
echo "2. cd ../frontend && npm install"
echo "3. Backend ve frontend kodlarini inceleyin"