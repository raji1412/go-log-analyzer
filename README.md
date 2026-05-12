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

## 📦 Project Structure
cmd/
|- logAnalyser
internal/
├── api/
├── analyzer/
├── config/
└── logger/
Dockerfile
docker-compose.yml
.env
