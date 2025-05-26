# Llama2 Local Go Projesi

Bu proje, Go diliyle yazılmış ve Llama 2 modelini yerel olarak çalıştırabilen bir yapay zeka uygulamasıdır.

## Özellikler

- Llama 2 modelini yerel olarak çalıştırma
- Türkçe dil desteği
- Basit komut satırı arayüzü
- Yapılandırılabilir model parametreleri

## Kurulum

1. **Gereksinimler:**

   - Go 1.21 veya üzeri
   - GCC derleyicisi (go-llama.cpp için)

2. **Model dosyasını edinin:**

   - Llama 2 modelini indirip `models/` klasörüne yerleştirin
   - Örnek: `llama-2-7b.Q4_0.bin`

3. **Bağımlılıkları yükleyin:**

   ```bash
   go mod tidy
   ```

## Kullanım

```bash
go run main.go
```

## Proje Yapısı

```text
demo/
├── main.go          # Ana uygulama
├── config.go        # Yapılandırma ayarları
├── go.mod           # Go modül dosyası
├── README.md        # Bu dosya
└── models/          # Model dosyaları
    └── llama-2-7b.Q4_0.bin
```

## Yapılandırma

Model yolu ve diğer parametreler `config.go` dosyasından değiştirilebilir.

## Notlar

- İlk çalıştırma model yükleme nedeniyle zaman alabilir
- Model dosyaları büyük olduğu için yeterli disk alanı olduğundan emin olun
- CPU'da çalıştırma GPU'ya göre daha yavaş olabilir
