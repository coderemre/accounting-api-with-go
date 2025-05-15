# 🧾 Accounting API with Go

A fully containerized, production-ready financial backend system written in Go.  
Supports user registration, authentication, balance tracking, transaction processing, distributed tracing, logging, caching, and real-time observability.


## What’s Inside 🚀
### Endpoints:
	•	User Authentication: Register and login endpoints with JWT integration.
	•	Transaction Management: Manage user balances and transfer funds.


### Database:
	•	Tables:
	•	users: Stores user information and hashed passwords.
	•	transactions: Tracks balance transfers and timestamps.
	•	balances: Maintains the current balance of each user.
	•	Environment Configuration:
	•	Fully configurable with .env for flexible setups.

---

## 🔧 Features

- **User Management**: Registration, login, and JWT-secured sessions
- **Transaction Engine**: Transfer and balance operations with validation
- **Observability**:
  - 📈 Prometheus metrics (`/metrics` endpoint)
  - 📊 Grafana dashboards (Go runtime stats)
  - 🔍 Jaeger-based distributed tracing (OpenTelemetry)
  - 📄 Zerolog structured JSON logging
- **DevOps Ready**:
  - 🐳 Docker Compose (multi-service stack)
  - 🛠️ Makefile commands (`make dev`, `make prod`, etc.)
- **Caching**: Redis integration for performance and queuing

---

## 🚀 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/coderemre/accounting-api-with-go.git
cd accounting-api-with-go
```

### 2. Create `.env` file

```env
	DB_HOST=127.0.0.1
	LOG_LEVEL=debug
	DB_HOST=mysql
	DB_USER=root
	DB_PASSWORD=password
	DB_NAME=bank_app
	REDIS_HOST=redis
	PORT=8080
	METRICS_PORT=2112
	DB_PORT=3306
	PROMETHEUS_PORT=9090
	JAEGER_UI_PORT=16686
	JAEGER_OTLP_PORT=4318
	GRAFANA_PORT=3000
	REDIS_PORT=6379
	DATABASE_DSN=root:password@tcp(mysql:3306)/bank_app?parseTime=true
```

---

## 🧪 Development & Deployment

### 🔁 Available Makefile Commands

| Command         | Description                                                  |
|----------------|--------------------------------------------------------------|
| `make dev`     | Build and start all services with live logs                  |
| `make prod`    | Start all services in detached mode (production ready)       |
| `make stop`    | Stop and remove all services and volumes                     |
| `make restart` | Restart all services with a fresh build                      |
| `make logs`    | Tail logs for all running services                           |

> All services are containerized. No dependencies needed on host machine.

---

## 📊 Monitoring Setup

### Prometheus + Grafana

- Prometheus scrapes from `/metrics`
- Grafana visualizes runtime and memory metrics

Access Grafana at: [http://localhost:3000](http://localhost:3000)  
Login: `admin / admin`  
Import dashboards manually or use JSON exports.

📸 **Example Dashboard:**  
![Metrics Dashboard](./screenshots/Metrics_Dashboard.png)

---

## 🔍 Distributed Tracing

Every handler initializes OpenTelemetry spans. Traces are sent to Jaeger and grouped per route.

Jaeger UI: [http://localhost:16686](http://localhost:16686)

---

## 📂 Logging

All log output is in structured JSON using [Zerolog](https://github.com/rs/zerolog).  
Each log includes timestamp, level, and message for production monitoring tools.

Example:
```json
{
  "level": "info",
  "time": "2025-05-12T10:00:00Z",
  "message": "User logged in successfully",
  "user_id": 23
}
```

---

## 🧠 Architecture Overview

- **API Server**: Handles all HTTP logic and services
- **Database**: MySQL (`users`, `transactions`, `balances` tables)
- **Metrics**: `/metrics` exposed for Prometheus
- **Tracing**: OpenTelemetry with Jaeger
- **Cache**: Redis integration
- **Docker Stack**: App, MySQL, Redis, Prometheus, Grafana, Jaeger

---

## 📌 Some Endpoints

- `POST /register` – Create a new user
- `POST /login` – Authenticate and receive JWT
- `GET /profile` – Get current user info (JWT required)
- `POST /transfer` – Transfer balance between users
- `GET /balance` – Check current balance

---

## ✅ TODO (Future Enhancements)

- [ ] Event sourcing for all transactions
- [ ] Redis caching for frequently accessed records
- [ ] Scheduled payments (job queue)
- [ ] Multi-currency support
- [ ] Circuit breaker + fallback logic
- [ ] Read-replica support and load balancing

---

## 📄 License

MIT — free to use and modify.