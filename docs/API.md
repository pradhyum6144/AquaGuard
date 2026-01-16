# AquaGuard API Documentation

## Base URL
```
http://localhost:8080
```

## Authentication
Currently no authentication required for demo purposes.
In production, use the security service for encrypted commands.

## Endpoints

### Health Check
```
GET /health
```
Response:
```json
{
  "status": "healthy",
  "service": "aquaguard-api"
}
```

### Telemetry (Bot → Server)
```
POST /telemetry
```
Request:
```json
{
  "bot_id": "AQ-001",
  "latitude": 19.0760,
  "longitude": 72.8777,
  "battery_level": 78.5
}
```

### Detection Event
```
POST /detection
```
Request:
```json
{
  "bot_id": "AQ-001",
  "trash_type": 1,
  "latitude": 19.0760,
  "longitude": 72.8777
}
```
Trash types: 0=organic, 1=plastic, 2=metal, 3=other

### Statistics
```
GET /stats
```
Response:
```json
{
  "total_plastic": 247,
  "total_metal": 89,
  "total_other": 34,
  "total_organic": 156
}
```

### Bot Command
```
POST /command
```
Request:
```json
{
  "bot_id": "AQ-001",
  "command": "halt"
}
```
Commands: "halt", "resume", "return_home"

### List Bots
```
GET /bots
```
Response:
```json
{
  "bots": [...],
  "override_active": false,
  "total_bots": 3
}
```

## WebSocket
```
WS /ws
```
Messages: "HALT", "RESUME"
