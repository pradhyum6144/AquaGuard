interface OverrideButtonProps {
    isActive: boolean;
    onClick: () => void;
}

export function OverrideButton({ isActive, onClick }: OverrideButtonProps) {
    return (
        <div className="text-center">
            <button
                onClick={onClick}
                className={`w-full py-4 px-6 rounded-xl font-bold text-lg transition-all duration-300 transform hover:scale-105 ${isActive
                        ? 'bg-red-500 hover:bg-red-600 animate-pulse shadow-lg shadow-red-500/50'
                        : 'bg-gradient-to-r from-gray-700 to-gray-600 hover:from-gray-600 hover:to-gray-500 text-white'
                    }`}
            >
                {isActive ? (
                    <span className="flex items-center justify-center gap-2">
                        <span className="animate-ping inline-flex h-3 w-3 rounded-full bg-white opacity-75"></span>
                        🛑 OVERRIDE ACTIVE
                    </span>
                ) : (
                    '⚡ Manual Override'
                )}
            </button>
            <p className="text-gray-500 text-sm mt-3">
                {isActive
                    ? 'All bots halted. Click to resume operations.'
                    : 'Emergency stop for all bots in the fleet'}
            </p>
        </div>
    );
}
