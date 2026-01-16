#!/usr/bin/env python3
"""
AquaGuard Telemetry Simulator
Generates realistic telemetry and detection data to populate the dashboard
Run this while the Go backend is running to see live data!
"""

import requests
import random
import time
import json

API_URL = "http://localhost:8080"

# Mumbai water body locations (approximate)
WATER_BODIES = [
    {"name": "Powai Lake", "lat": 19.1273, "lng": 72.9047},
    {"name": "Vihar Lake", "lat": 19.1463, "lng": 72.9298},
    {"name": "Tulsi Lake", "lat": 19.1757, "lng": 72.9198},
    {"name": "Mithi River", "lat": 19.0760, "lng": 72.8650},
    {"name": "Dahisar River", "lat": 19.2590, "lng": 72.8542},
]

# Bot fleet
BOTS = [
    {"id": "AQ-001", "location_idx": 0},
    {"id": "AQ-002", "location_idx": 1},
    {"id": "AQ-003", "location_idx": 2},
    {"id": "AQ-004", "location_idx": 3},
    {"id": "AQ-005", "location_idx": 4},
]


def send_telemetry(bot_id: str, lat: float, lng: float, battery: float):
    """Send a telemetry ping to the backend"""
    payload = {
        "bot_id": bot_id,
        "latitude": lat + random.uniform(-0.001, 0.001),  # Small movement
        "longitude": lng + random.uniform(-0.001, 0.001),
        "battery_level": battery,
    }
    try:
        resp = requests.post(f"{API_URL}/telemetry", json=payload, timeout=2)
        print(f"📍 Telemetry {bot_id}: battery={battery:.1f}% - {resp.status_code}")
    except Exception as e:
        print(f"❌ Failed to send telemetry: {e}")


def send_detection(bot_id: str, trash_type: int, lat: float, lng: float):
    """Send a detection event (trash spotted)"""
    types = {0: "🍂 organic", 1: "🧴 plastic", 2: "🥫 metal", 3: "❓ unknown"}
    payload = {
        "bot_id": bot_id,
        "trash_type": trash_type,
        "latitude": lat,
        "longitude": lng,
    }
    try:
        resp = requests.post(f"{API_URL}/detection", json=payload, timeout=2)
        print(f"🔍 Detection {bot_id}: {types.get(trash_type, 'unknown')} - {resp.status_code}")
    except Exception as e:
        print(f"❌ Failed to send detection: {e}")


def check_health():
    """Verify backend is running"""
    try:
        resp = requests.get(f"{API_URL}/health", timeout=2)
        if resp.status_code == 200:
            print("✅ Backend is healthy!")
            return True
    except:
        pass
    print("❌ Backend not reachable at", API_URL)
    return False


def run_simulation(num_pings: int = 50):
    """Run the simulation"""
    print("=" * 50)
    print("🤖 AquaGuard Telemetry Simulator")
    print("=" * 50)
    
    if not check_health():
        print("Please start the Go backend first: cd backend && go run .")
        return
    
    print(f"\n🚀 Starting simulation with {num_pings} telemetry pings...\n")
    
    # Initialize battery levels
    batteries = {bot["id"]: random.uniform(60, 100) for bot in BOTS}
    
    for i in range(num_pings):
        bot = random.choice(BOTS)
        location = WATER_BODIES[bot["location_idx"]]
        
        # Slowly drain battery
        batteries[bot["id"]] = max(10, batteries[bot["id"]] - random.uniform(0.1, 0.5))
        
        # Send telemetry
        send_telemetry(
            bot["id"],
            location["lat"],
            location["lng"],
            batteries[bot["id"]]
        )
        
        # 30% chance to detect something
        if random.random() < 0.3:
            # Weighted: more plastic than metal
            trash_type = random.choices([0, 1, 2, 3], weights=[3, 5, 2, 1])[0]
            send_detection(bot["id"], trash_type, location["lat"], location["lng"])
        
        # Small delay between pings
        time.sleep(0.2)
    
    print("\n" + "=" * 50)
    print("✅ Simulation complete!")
    print(f"📊 Check stats at: {API_URL}/stats")
    print("=" * 50)
    
    # Fetch and display final stats
    try:
        resp = requests.get(f"{API_URL}/stats", timeout=2)
        stats = resp.json()
        print(f"\n📈 Impact Report:")
        print(f"   Plastic recovered: {stats.get('total_plastic', 0)}")
        print(f"   Metal recovered:   {stats.get('total_metal', 0)}")
        print(f"   Other items:       {stats.get('total_other', 0)}")
        print(f"   Organic (ignored): {stats.get('total_organic', 0)}")
    except:
        pass


if __name__ == "__main__":
    run_simulation(50)
