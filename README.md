# AquaGuard 🌊🤖

**Protecting Indian sanitation workers through autonomous waste recovery.**

AquaGuard is an autonomous water-cleaning bot system that removes plastic and non-organic waste from Indian water bodies. Our bots work 24/7 to clean rivers, lakes, and urban water channels—reducing health hazards for sanitation workers and protecting aquatic ecosystems.

![Dashboard Preview](https://via.placeholder.com/800x400/0f172a/10b981?text=AquaGuard+Dashboard)

## 🎯 SDG Alignment

| SDG | Goal | Our Contribution |
|-----|------|------------------|
| **11** | Sustainable Cities and Communities | Cleaner urban water bodies improve city livability |
| **13** | Climate Action | Reduced plastic in waterways protects marine ecosystems |

## 🏗️ Tech Stack

```
┌─────────────────────────────────────────────────────────────┐
│                      FRONTEND                                │
│                   Next.js + Tailwind                        │
│              (Real-time Dashboard & Controls)                │
├─────────────────────────────────────────────────────────────┤
│                     WEBSOCKET                                │
│              (Live Telemetry & Commands)                     │
├─────────────────────────────────────────────────────────────┤
│     GO BACKEND           │        RUST SECURITY             │
│   (API + Analytics)      │     (AES-256-GCM Encryption)     │
│   ├── /telemetry        │     ├── /encrypt                  │
│   ├── /detection        │     └── /decrypt                  │
│   ├── /stats            │                                   │
│   └── /command          │                                   │
├─────────────────────────────────────────────────────────────┤
│                        REDIS                                 │
│                (Caching & Persistence)                       │
├─────────────────────────────────────────────────────────────┤
│                    DOCKER + AWS                              │
│              (Containerized Deployment)                      │
└─────────────────────────────────────────────────────────────┘
```

### Why Go + Rust?

- **Go**: Fast HTTP server with excellent concurrency (goroutines) for handling telemetry from hundreds of bots simultaneously
- **Rust**: Memory-safe, zero-cost abstractions for our critical security layer—ensures bot commands can't be hijacked

## ✨ Features

- 📊 **Real-time Dashboard**: Live map and impact metrics
- 🔒 **Encrypted Commands**: AES-256-GCM encryption prevents bot hijacking
- 🤖 **Trash Classification**: AI distinguishes plastic (collect) from organic (ignore)
- 🛑 **Manual Override**: Emergency halt for all bots with one click
- 📈 **Impact Reports**: Track plastic recovered for stakeholder reporting

## 🚀 How to Run Locally

### Prerequisites

- Docker & Docker Compose
- OR: Go 1.21+, Rust 1.74+, Node.js 20+

### Option 1: Docker (Recommended)

```bash
# Clone the repository
git clone https://github.com/your-team/aquaguard.git
cd aquaguard

# Start all services
docker-compose up --build

# Access the dashboard
open http://localhost:3000
```

### Option 2: Manual Setup

```bash
# Terminal 1: Start Go Backend
cd backend
go run .

# Terminal 2: Start Rust Security Service
cd security
cargo run

# Terminal 3: Start Frontend
cd frontend
npm install
npm run dev
```

## 📡 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| POST | `/telemetry` | Receive bot location/battery |
| POST | `/detection` | Report trash detection |
| GET | `/stats` | Get impact statistics |
| POST | `/command` | Send command to bot |
| WS | `/ws` | WebSocket for real-time updates |

## 🔐 Security Service

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/encrypt` | Encrypt bot commands |
| POST | `/decrypt` | Decrypt bot responses |
| GET | `/health` | Service health check |

## 📁 Project Structure

```
AquaGuard/
├── backend/           # Go HTTP server
│   ├── main.go       # Entry point, routes
│   ├── analytics.go  # Trash classification
│   ├── websocket.go  # Real-time updates
│   └── Dockerfile
├── security/          # Rust microservice
│   ├── src/main.rs   # AES-GCM encryption
│   ├── Cargo.toml
│   └── Dockerfile
├── frontend/          # Next.js dashboard
│   ├── src/app/
│   └── Dockerfile
├── docker-compose.yml
└── README.md
```

## 👥 Team

Built with ❤️ for Smart India Hackathon 2024

---

*Empowering sustainable cities, one bot at a time.*
