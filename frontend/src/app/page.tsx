'use client';

import { useState, useEffect, useRef } from 'react';

interface ImpactStats {
    total_plastic: number;
    total_metal: number;
    total_other: number;
    total_organic: number;
}

interface Bot {
    bot_id: string;
    latitude: number;
    longitude: number;
    battery_level: number;
    timestamp: string;
}

export default function Dashboard() {
    const [stats, setStats] = useState<ImpactStats>({
        total_plastic: 247,
        total_metal: 89,
        total_other: 34,
        total_organic: 156,
    });
    const [bots, setBots] = useState<Bot[]>([
        { bot_id: 'AQ-001', latitude: 19.0760, longitude: 72.8777, battery_level: 78, timestamp: '' },
        { bot_id: 'AQ-002', latitude: 19.0330, longitude: 73.0297, battery_level: 92, timestamp: '' },
        { bot_id: 'AQ-003', latitude: 18.9220, longitude: 72.8347, battery_level: 45, timestamp: '' },
    ]);
    const [overrideActive, setOverrideActive] = useState(false);
    const wsRef = useRef<WebSocket | null>(null);

    // WebSocket connection
    useEffect(() => {
        const connectWebSocket = () => {
            try {
                const ws = new WebSocket('ws://localhost:8080/ws');
                wsRef.current = ws;

                ws.onmessage = (event) => {
                    if (event.data === 'OVERRIDE_ACTIVE') {
                        setOverrideActive(true);
                    } else if (event.data === 'OVERRIDE_INACTIVE') {
                        setOverrideActive(false);
                    }
                };

                ws.onclose = () => {
                    // Reconnect after 3 seconds
                    setTimeout(connectWebSocket, 3000);
                };
            } catch (e) {
                console.log('WebSocket not available');
            }
        };

        connectWebSocket();

        return () => {
            wsRef.current?.close();
        };
    }, []);

    const handleOverride = () => {
        const command = overrideActive ? 'RESUME' : 'HALT';
        wsRef.current?.send(command);
        setOverrideActive(!overrideActive);
    };

    const avgBattery = bots.length > 0
        ? Math.round(bots.reduce((sum, bot) => sum + bot.battery_level, 0) / bots.length)
        : 0;

    return (
        <div className="min-h-screen p-6">
            {/* Header */}
            <header className="mb-8">
                <div className="flex items-center justify-between">
                    <div>
                        <h1 className="text-4xl font-bold bg-gradient-to-r from-emerald-400 to-cyan-400 bg-clip-text text-transparent">
                            AquaGuard
                        </h1>
                        <p className="text-gray-400 mt-1">Autonomous Water-Cleaning Bot Control Center</p>
                    </div>
                    <div className="flex items-center gap-4">
                        <div className="px-4 py-2 rounded-full bg-emerald-500/20 border border-emerald-500/30">
                            <span className="text-emerald-400 font-medium">SDG 11 & 13</span>
                        </div>
                    </div>
                </div>
            </header>

            {/* Impact Cards */}
            <section className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
                {/* Plastic Recovered */}
                <div className="relative overflow-hidden rounded-2xl bg-gradient-to-br from-emerald-500/20 to-emerald-900/20 border border-emerald-500/30 p-6">
                    <div className="absolute top-0 right-0 w-32 h-32 bg-emerald-500/10 rounded-full blur-3xl"></div>
                    <div className="relative">
                        <div className="flex items-center gap-3 mb-4">
                            <div className="w-12 h-12 rounded-xl bg-emerald-500/30 flex items-center justify-center">
                                <svg className="w-6 h-6 text-emerald-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                                </svg>
                            </div>
                            <h3 className="text-gray-400 font-medium">Plastic Recovered</h3>
                        </div>
                        <p className="text-4xl font-bold text-white">{stats.total_plastic}</p>
                        <p className="text-emerald-400 text-sm mt-1">+{stats.total_metal} metal items</p>
                    </div>
                </div>

                {/* Battery Health */}
                <div className="relative overflow-hidden rounded-2xl bg-gradient-to-br from-cyan-500/20 to-cyan-900/20 border border-cyan-500/30 p-6">
                    <div className="absolute top-0 right-0 w-32 h-32 bg-cyan-500/10 rounded-full blur-3xl"></div>
                    <div className="relative">
                        <div className="flex items-center gap-3 mb-4">
                            <div className="w-12 h-12 rounded-xl bg-cyan-500/30 flex items-center justify-center">
                                <svg className="w-6 h-6 text-cyan-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M13 10V3L4 14h7v7l9-11h-7z" />
                                </svg>
                            </div>
                            <h3 className="text-gray-400 font-medium">Battery Health</h3>
                        </div>
                        <p className="text-4xl font-bold text-white">{avgBattery}%</p>
                        <p className="text-cyan-400 text-sm mt-1">Fleet average</p>
                    </div>
                </div>

                {/* Active Bots */}
                <div className="relative overflow-hidden rounded-2xl bg-gradient-to-br from-violet-500/20 to-violet-900/20 border border-violet-500/30 p-6">
                    <div className="absolute top-0 right-0 w-32 h-32 bg-violet-500/10 rounded-full blur-3xl"></div>
                    <div className="relative">
                        <div className="flex items-center gap-3 mb-4">
                            <div className="w-12 h-12 rounded-xl bg-violet-500/30 flex items-center justify-center">
                                <svg className="w-6 h-6 text-violet-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 3v2m6-2v2M9 19v2m6-2v2M5 9H3m2 6H3m18-6h-2m2 6h-2M7 19h10a2 2 0 002-2V7a2 2 0 00-2-2H7a2 2 0 00-2 2v10a2 2 0 002 2zM9 9h6v6H9V9z" />
                                </svg>
                            </div>
                            <h3 className="text-gray-400 font-medium">Active Bots</h3>
                        </div>
                        <p className="text-4xl font-bold text-white">{bots.length}</p>
                        <p className="text-violet-400 text-sm mt-1">Operating in Mumbai</p>
                    </div>
                </div>
            </section>

            {/* Main Content */}
            <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
                {/* Live Map Placeholder */}
                <div className="lg:col-span-2 rounded-2xl bg-gray-800/50 border border-gray-700/50 overflow-hidden">
                    <div className="p-4 border-b border-gray-700/50">
                        <h2 className="text-xl font-semibold text-white">Live Map</h2>
                    </div>
                    <div className="relative h-96 bg-gradient-to-br from-gray-900 to-gray-800 flex items-center justify-center">
                        {/* Placeholder map visualization */}
                        <div className="absolute inset-0 opacity-30">
                            <svg viewBox="0 0 100 100" className="w-full h-full">
                                <defs>
                                    <pattern id="grid" width="10" height="10" patternUnits="userSpaceOnUse">
                                        <path d="M 10 0 L 0 0 0 10" fill="none" stroke="rgba(16, 185, 129, 0.2)" strokeWidth="0.5" />
                                    </pattern>
                                </defs>
                                <rect width="100" height="100" fill="url(#grid)" />
                            </svg>
                        </div>
                        {/* Bot markers */}
                        {bots.map((bot, index) => (
                            <div
                                key={bot.bot_id}
                                className="absolute transform -translate-x-1/2 -translate-y-1/2 animate-pulse"
                                style={{
                                    left: `${30 + index * 25}%`,
                                    top: `${40 + (index % 2) * 20}%`,
                                }}
                            >
                                <div className={`w-4 h-4 rounded-full ${bot.battery_level < 50 ? 'bg-amber-500' : 'bg-emerald-500'} shadow-lg shadow-emerald-500/50`}></div>
                                <span className="absolute left-5 top-0 text-xs text-gray-400 whitespace-nowrap">{bot.bot_id}</span>
                            </div>
                        ))}
                        <div className="text-gray-500 text-center z-10">
                            <p className="text-lg">📍 Mumbai Water Bodies</p>
                            <p className="text-sm">Real-time bot positions</p>
                        </div>
                    </div>
                </div>

                {/* Control Panel */}
                <div className="rounded-2xl bg-gray-800/50 border border-gray-700/50 p-6">
                    <h2 className="text-xl font-semibold text-white mb-6">Control Panel</h2>

                    {/* Manual Override Button */}
                    <button
                        onClick={handleOverride}
                        className={`w-full py-4 px-6 rounded-xl font-bold text-lg transition-all duration-300 ${overrideActive
                                ? 'bg-red-500 hover:bg-red-600 animate-pulse-red shadow-lg shadow-red-500/50'
                                : 'bg-gray-700 hover:bg-gray-600 text-white'
                            }`}
                    >
                        {overrideActive ? '🛑 OVERRIDE ACTIVE' : '⚡ Manual Override'}
                    </button>

                    <p className="text-gray-500 text-sm mt-3 text-center">
                        {overrideActive
                            ? 'All bots halted. Click to resume.'
                            : 'Emergency stop for all bots'}
                    </p>

                    {/* Bot Status List */}
                    <div className="mt-6 space-y-3">
                        <h3 className="text-gray-400 font-medium mb-3">Bot Status</h3>
                        {bots.map((bot) => (
                            <div key={bot.bot_id} className="flex items-center justify-between p-3 rounded-lg bg-gray-900/50">
                                <div className="flex items-center gap-3">
                                    <div className={`w-2 h-2 rounded-full ${overrideActive ? 'bg-red-500' : 'bg-emerald-500'}`}></div>
                                    <span className="text-white font-medium">{bot.bot_id}</span>
                                </div>
                                <div className="flex items-center gap-2">
                                    <div className="w-16 h-2 rounded-full bg-gray-700 overflow-hidden">
                                        <div
                                            className={`h-full rounded-full ${bot.battery_level < 30 ? 'bg-red-500' : bot.battery_level < 60 ? 'bg-amber-500' : 'bg-emerald-500'}`}
                                            style={{ width: `${bot.battery_level}%` }}
                                        ></div>
                                    </div>
                                    <span className="text-gray-400 text-sm">{bot.battery_level}%</span>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            </div>

            {/* Footer */}
            <footer className="mt-8 text-center text-gray-500 text-sm">
                <p>AquaGuard • Protecting Indian sanitation workers through autonomous waste recovery</p>
                <p className="mt-1">Built with Go, Rust & React for SIH 2024</p>
            </footer>
        </div>
    );
}
