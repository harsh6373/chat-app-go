# 💬 chat-app-go

A high-performance, real-time chat backend built with **Go Fiber**, **WebSockets**, **PostgreSQL**, and **Redis**. Designed for scalability, security, and extensibility.

---

## 🚀 Features

- Real-time messaging with WebSockets
- 1-to-1 and broadcast messaging
- User authentication with JWT
- Chat history persisted in PostgreSQL
- Redis integration for caching / pub-sub
- WebSocket connection handling
- Modular and scalable project structure

---

## 🧱 Tech Stack

- **Go (Fiber Framework)**
- **PostgreSQL**
- **Redis**
- **JWT for Authentication**
- **WebSockets for Real-Time**

---

## 🏁 Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/yourusername/chat-app-go.git
cd chat-app-go
```

### 2. Configure environment

Create a `.env` file based on:

```env
PORT=3000
DATABASE_URL=postgres://user:password@localhost:5432/chatapp
REDIS_ADDR=localhost:6379
JWT_SECRET=your_jwt_secret
```

### 3. Install dependencies

```bash
go mod tidy
```

### 4. Run the application

```bash
go run cmd/main.go
```

---

## 📁 Project Structure

```
chat-app-go/
├── cmd/               # Entry point (main.go)
├── config/            # Configuration (env, db, redis)
├── controllers/       # Route handlers
├── middleware/        # Authentication middleware
├── models/            # DB models
├── routes/            # Route setup
├── services/          # Business logic
├── utils/             # Utility functions (JWT etc.)
├── ws/                # WebSocket logic
├── .env               # Environment variables
├── .gitignore         # Git ignore rules
└── README.md          # Documentation
```

---

## 🛠️ Future Improvements

- Read receipts
- Typing indicators
- Group chat support
- Admin dashboard
- Dockerization