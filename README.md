# 🚀Limitify - Scalable Rate Limiting for APIs

A high-performance **rate-limiting service** built with **Go, Redis, and PostgreSQL** to protect APIs from abuse. Supports the **Fixed Window**

## 🛠️ Features
- ✅ **Fixed Window Rate Limiting** (Prevents excessive requests)
- ✅ **JWT Authentication** (Secure access to protected routes)
- ✅ **Admin API** (Configure rate limits dynamically)
- ✅ **Request Logging** (Track API usage & blocked requests)
- ✅ **Redis for Fast Request Counting** (Low-latency enforcement)

## 🚀 Tech Stack
- **Go (Gin)** – Backend framework
- **Redis** – High-speed request tracking
- **PostgreSQL** – Stores rate limit configs & logs
- **Docker** – Easy deployment

---

## 📌 Installation & Setup

### **1️⃣ Clone the Repository**
```sh
git clone https://github.com/YOUR_GITHUB_USERNAME/limitify.git
cd limitify
```

### **2️⃣ Setup Environment Variables**

Rename ```.env.example``` to ```.env``` and update:
```
REDIS_URL=localhost:6379
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=rate_limiter
JWT_SECRET=your_secret_key
PORT=3000
```

### **3️⃣ Run Locally**
Using Go
```
go mod tidy
go run main.go
```

## **Usage**
**🔹User signup**
```

curl -X POST http://localhost:3000/signup -H "Content-Type: application/json" -d '{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "123456"
}'
```
**🔹User login**
```
curl -X POST http://localhost:3000/login -H "Content-Type: application/json" -d '{
  "email": "john@example.com",
  "password": "123456"
}'

```
*✅ Response: { "token": "YOUR_JWT_TOKEN" }*

**🔹 Access a Protected API (Rate Limited)**

```
curl -X GET http://localhost:3000/protected-resource -H "Authorization: Bearer YOUR_JWT_TOKEN"

```
*✅ Response (Allowed): { "message": "Request successful" }*

*❌ Response (Blocked): { "error": "Too many requests" }*

**🔹 Admin API (Set Global Rate Limit)**
```
curl -X POST http://localhost:3000/admin/set-rate-limit -H "Content-Type: application/json" \
-H "Authorization: Bearer YOUR_ADMIN_JWT" \
-d '{"requests": 100, "time_window": 60}'

```

## **Roadmap (Future Enhancements)**
 - Sliding Window Rate Limiting
 - Token Bucket Algorithm
 - Prometheus Monitoring
 - Web Dashboard for Managing Limits


## **Contributing**
We welcome contributions! To get started:

- Fork the repo & create a feature branch.
- Make changes & test locally.
- Submit a PR with a detailed description
