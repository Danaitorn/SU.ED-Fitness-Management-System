# 🏋️ SU.ED Fitness Management System

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21-00ADD8?style=for-the-badge&logo=go&logoColor=white"/>
  <img src="https://img.shields.io/badge/React-18-61DAFB?style=for-the-badge&logo=react&logoColor=black"/>
  <img src="https://img.shields.io/badge/PostgreSQL-16-4169E1?style=for-the-badge&logo=postgresql&logoColor=white"/>
  <img src="https://img.shields.io/badge/License-MIT-green?style=for-the-badge"/>
</p>

<p align="center">
  A full-stack web-based fitness center management system developed as a <strong>Senior Project</strong> for the Bachelor of Science in Information Technology program at <strong>Silpakorn University</strong>.
</p>

<p align="center">
  🏆 <strong>AUCC 2026 Award — Very Good Level</strong> (Academic Undergraduate Conference Competition)
</p>

---

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Contributors](#contributors)

---

## Overview

SU.ED is a comprehensive fitness center management platform designed to streamline operations for both members and administrators. The system handles the full membership lifecycle — from registration and health screening to QR-code check-in and analytics.

---

## ✨ Features

### 👤 Member Features
- Register and login with **JWT authentication**
- Apply for fitness memberships online
- Fill in personal information and **PAR-Q health questionnaires**
- View membership status and details
- Display a personal **QR Code** for gym check-in

### 🛠️ Admin Features
- Manage users (Create / Edit / Delete)
- Approve or reject membership applications
- Record fitness visits manually
- **Scan QR Code** for automatic check-in
- View real-time dashboard statistics
- Manage news and announcements

### 📊 Dashboard Analytics
| Metric | Description |
|---|---|
| Active Members | Total currently active memberships |
| Daily Visits | Number of check-ins today |
| Membership Stats | Daily / Monthly / Yearly breakdown |
| Visit History | Recent visit log with timestamps |
| Activity Logs | System-wide audit trail |

---

## 🛠️ Tech Stack

### Frontend (`my_project/`)
| Technology | Purpose |
|---|---|
| React.js | UI framework |
| React Bootstrap | Component library |
| Recharts | Dashboard charts & graphs |
| React Router | Client-side routing |
| html5-qrcode | QR code scanning |
| js-cookie | Session management |

### Backend (`my_backend/`)
| Technology | Purpose |
|---|---|
| Go (Golang) | Server language |
| Gin Framework | HTTP web framework |
| GORM | ORM for database access |
| JWT | Authentication tokens |
| Casbin | Role-based access control (RBAC) |

### Database (`my_database/`)
| Technology | Purpose |
|---|---|
| PostgreSQL | Primary relational database |

### Cloud Services
| Service | Purpose |
|---|---|
| Cloudinary | Image upload & storage |

---

## 📁 Project Structure

```
SU.ED-Fitness-Management-System/
├── my_backend/          # Go API server (Gin + GORM + JWT + Casbin)
├── my_project/          # React web client
└── my_database/         # PostgreSQL schema & migrations
```

**3-Tier Architecture:**
```
[Web Client (React)]  ←→  [App Server (Go/Gin)]  ←→  [Database (PostgreSQL)]
     my_project               my_backend                  my_database
```

---

## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- Node.js 18+
- PostgreSQL 15+
- Cloudinary account (for image upload)

### 1. Clone the repository
```bash
git clone https://github.com/Danaitorn/SU.ED-Fitness-Management-System.git
cd SU.ED-Fitness-Management-System
```

### 2. Setup Database
```bash
cd my_database
# Run the SQL schema file in your PostgreSQL instance
psql -U postgres -d your_database_name -f schema.sql
```

### 3. Start the Backend
```bash
cd my_backend
# Copy and configure environment variables
cp .env.example .env
# Edit .env with your DB credentials and Cloudinary keys

go mod tidy
go run main.go
# Server runs on http://localhost:8080
```

### 4. Start the Frontend
```bash
cd my_project
npm install
npm start
# App runs on http://localhost:3000
```

---

## 👥 Contributors

| Name | Role |
|---|---|
| Danaitorn Saibiew (Tan) | Full-Stack Developer |

---

## 📄 License

This project is licensed under the MIT License.

---

<p align="center">
  Developed as a Senior Project · Bachelor of Science in Information Technology · Silpakorn University · 2026
</p>
