interface ImpactCardProps {
    title: string;
    value: number | string;
    subtitle: string;
    icon: React.ReactNode;
    gradient: string;
    borderColor: string;
    accentColor: string;
}

export function ImpactCard({
    title,
    value,
    subtitle,
    icon,
    gradient,
    borderColor,
    accentColor,
}: ImpactCardProps) {
    return (
        <div
            className={`relative overflow-hidden rounded-2xl ${gradient} ${borderColor} border p-6 transform hover:scale-105 transition-all duration-300`}
        >
            <div className={`absolute top-0 right-0 w-32 h-32 ${accentColor} rounded-full blur-3xl`}></div>
            <div className="relative">
                <div className="flex items-center gap-3 mb-4">
                    {icon}
                    <h3 className="text-gray-400 font-medium">{title}</h3>
                </div>
                <p className="text-4xl font-bold text-white">{value}</p>
                <p className={`${accentColor.replace('/10', '')} text-sm mt-1`}>{subtitle}</p>
            </div>
        </div>
    );
}
