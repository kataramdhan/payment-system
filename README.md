# 🚀 Payment Transaction System (Golang + React)

A production-style transaction system built with Golang and React, featuring asynchronous processing, retry mechanisms, and real-time dashboard monitoring.

---

## 📌 Background

This project was built to simulate a real-world transaction system commonly used in fintech and e-commerce platforms.

The goal is to demonstrate how to design and implement a **scalable backend system** with:

* asynchronous processing (queue & worker)
* retry mechanism for failed transactions
* real-time monitoring via dashboard

---

## 🧠 Key Features

### 🔐 Authentication

* JWT-based login system
* Secure API access

### 💳 Transaction System

* Create transaction
* Status lifecycle:

  * pending
  * success
  * failed

### ⚙️ Background Processing

* Queue-based processing using Redis
* Worker handles transaction processing
* Automatic retry for failed jobs

### 📊 Dashboard (React)

* View transaction list
* Auto refresh (real-time polling)
* Status visualization

---

## 🏗️ Tech Stack

### Backend

* Golang (Gin)
* PostgreSQL
* Redis (Queue)
* Asynq (Worker)
* Docker

### Frontend

* React + TypeScript (Vite)

---

## 🧱 System Architecture

```text
Client (React)
     ↓
API Server (Golang)
     ↓
Redis Queue
     ↓
Worker (Asynq)
     ↓
PostgreSQL
```

---

## ⚡ How It Works

1. User creates a transaction → status = `pending`
2. Transaction is pushed to Redis queue
3. Worker processes the transaction asynchronously
4. If success → status updated to `success`
5. If failed → retry mechanism kicks in
6. Dashboard auto-refresh shows updated status

---

## 🖥️ Demo Preview

### 🔐 Login Page

screenshots/login.png

### 📊 Dashboard

screenshots/dashboard.png

---

## 📦 How to Run

### 1. Clone repository

```bash
git clone https://github.com/kataramdhan/payment-system.git
cd payment-system
```

---

### 2. Run infrastructure (PostgreSQL + Redis)

```bash
docker-compose up -d
```

---

### 3. Run Backend API

```bash
cd backend
go run cmd/api/main.go
```

---

### 4. Run Worker

```bash
go run cmd/worker/main.go
```

---

### 5. Run Frontend

```bash
cd ../frontend
npm install
npm run dev
```

---

## 🔑 API Endpoints

### Auth

* `POST /login`

### Transactions

* `POST /transactions`
* `GET /transactions`

---

## 💡 Why This Project Matters

This project demonstrates:

* Designing **asynchronous systems**
* Handling **real-world failure scenarios (retry)**
* Building **scalable backend architecture**
* Integrating backend with frontend dashboard

---

## 🚀 Future Improvements

* WebSocket for real-time updates
* Transaction filtering & search
* Pagination
* Role-based access control

---

## 👨‍💻 Author

Backend Developer with 9+ years experience in building scalable web applications across multiple industries.

---

⭐ If you find this project useful, feel free to give it a star!
