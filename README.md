# Go Log Analyzer Microservice

A production-ready log analyzer built using Golang, Gin, Docker, and structured logging.

---

## Features

- REST API using Gin
- Structured logging with Zap
- Middleware for request logging
- Health check endpoint
- Environment-based configuration
- Graceful shutdown
- Docker Compose support

---

## 🛠 Tech Stack

- Golang
- Gin
- Docker
- Zap Logger
---
## ▶️ Run Locally

go run ./cmd/loganalyzer

---


## 🐳 Run with Docker


docker-compose up --build

## 📡 API Endpoints

### Health Check

GET /health

### Response:

```json
{
  "status": "ok"
}
```
Get Logs

GET /logs?file=/app/logs.txt

### Optional filter:

GET /logs?file=/app/logs.txt&level=error

## 📊 Example Response
```json
{
  "status": "success",
  "data": {
    "error": 2,
    "info": 2,
    "warning": 1
  }
}
```

## 🧱 Architecture
      Client
        ↓
      Gin Router
        ↓
      Middleware
        ↓
      Analyzer Service
        ↓
      Logs File

## 📌 Future Improvements
1. Add database support (PostgreSQL)
2. Add authentication layer
3. Add metrics and monitoring

👨‍💻 Author
Built by Rajitha


