# Инструкции по развертыванию

## 🚀 Быстрый старт

### 1. Клонирование репозитория
```bash
git clone https://github.com/YOUR_USERNAME/price-tags-generator.git
cd price-tags-generator
```

### 2. Запуск Backend (Go)
```bash
cd backend
go mod tidy
go run main.go
```
Сервер запустится на `http://localhost:8080`

### 3. Запуск Frontend (Vue)
```bash
cd front
npm install
npm run dev
```
Приложение будет доступно на `http://localhost:5173`

## 🐳 Docker (опционально)

### Backend
```bash
cd backend
docker build -t price-tags-backend .
docker run -p 8080:8080 price-tags-backend
```

### Frontend
```bash
cd front
docker build -t price-tags-frontend .
docker run -p 3000:3000 price-tags-frontend
```

## 🔧 Настройка окружения

### Backend
Создайте файл `.env` в папке `backend/`:
```env
PORT=8080
UPLOAD_DIR=./uploads
TEMPLATES_DIR=./templates
```

### Frontend
Создайте файл `.env` в папке `front/`:
```env
VITE_API_URL=http://localhost:8080
```

## 📦 Продакшн сборка

### Frontend
```bash
cd front
npm run build
```

### Backend
```bash
cd backend
go build -o price-tags-server main.go
./price-tags-server
```

## 🗂 Структура данных

### Шаблоны ценников
Хранятся в `backend/templates/pricetags/` в формате SVG

### Загруженные файлы
Сохраняются в `backend/uploads/`

### Товары
Хранятся в памяти (можно расширить до базы данных)

## 🔍 Отладка

### Backend логи
```bash
cd backend
go run main.go --debug
```

### Frontend логи
Откройте DevTools в браузере (F12)

## 📋 Требования

- **Go** 1.19+
- **Node.js** 16+
- **npm** 8+

## 🆘 Поддержка

При возникновении проблем:
1. Проверьте логи в консоли
2. Убедитесь, что порты 8080 и 5173 свободны
3. Проверьте права доступа к папкам uploads и templates
