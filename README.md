# AquaGuard 🌊🤖

**Protecting Indian sanitation workers through autonomous waste recovery.**

AquaGuard is an autonomous water-cleaning bot system that removes plastic and non-organic waste from Indian water bodies. Our bots work 24/7 to clean rivers, lakes, and urban water channels—reducing health hazards for sanitation workers and protecting aquatic ecosystems.


##  SDG Alignment

| SDG | Goal | Our Contribution |
|-----|------|------------------|
| **11** | Sustainable Cities and Communities | Cleaner urban water bodies improve city livability |
| **13** | Climate Action | Reduced plastic in waterways protects marine ecosystems |

##  Tech Stack

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

##  Features

-  **Real-time Dashboard**: Live map and impact metrics
-  **Encrypted Commands**: AES-256-GCM encryption prevents bot hijacking
-  **Trash Classification**: AI distinguishes plastic (collect) from organic (ignore)
-  **Manual Override**: Emergency halt for all bots with one click
-  **Impact Reports**: Track plastic recovered for stakeholder reporting

##  Quick Start (2 minutes)

```bash
# 1. Clone and enter the project
git clone https://github.com/pradhyum6144/AquaGuard.git
cd AquaGuard

# 2. Start all services (choose one option)

# Option A: If you have Go + Rust + Node installed
cd backend && go run . &
cd ../security && cargo run &
cd ../frontend && npm install && npm run dev &

# Option B: Docker (if available)
docker-compose up --build
```

**Access the dashboard at http://localhost:3000** 

### Generate Demo Data

```bash
# Run the simulator to populate the dashboard with realistic data
pip install requests  # if needed
python scripts/simulate.py
```

This sends 50 telemetry pings from 5 virtual bots, making the dashboard come alive!

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| POST | `/telemetry` | Receive bot location/battery |
| POST | `/detection` | Report trash detection |
| GET | `/stats` | Get impact statistics |
| POST | `/command` | Send command to bot |
| WS | `/ws` | WebSocket for real-time updates |

##  Security Service

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/encrypt` | Encrypt bot commands |
| POST | `/decrypt` | Decrypt bot responses |
| GET | `/health` | Service health check |

##  Project Structure

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
├── scripts/
│   └── simulate.py   # Demo data generator
├── docker-compose.yml
└── README.md
```



