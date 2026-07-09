import {
    BellIcon,
    ChartColumnIcon,
    LayoutDashboard,
    LayoutDashboardIcon,
    RefreshCwIcon,
    ShieldCheckIcon,
    TrendingDownIcon,
} from 'lucide-react'
import { Inter } from 'next/font/google'

const inter = Inter({
    subsets: ['latin'],
})

export default function MiniWeb() {
    return (
        <div className={`w-full my-6 ${inter.className}`}>
            <div className="max-w-2xl mx-auto bg-slate-950 border border-gray-400/20 rounded-xl flex flex-col gap-8 py-8 px-8">
                <div className="flex gap-4 items-center">
                    <div className="flex gap-2">
                        <div className="w-3 h-3 bg-gray-600/30 rounded-full"></div>
                        <div className="w-3 h-3 bg-gray-600/30 rounded-full"></div>
                        <div className="w-3 h-3 bg-gray-600/30 rounded-full"></div>
                    </div>
                    <div className="flex-1 w-full h-4 bg-gray-600/30 rounded-xs"></div>
                </div>

                <div className="flex gap-4 opacity-60">
                    <div className="flex-1 bg-slate-900 py-4 px-6 rounded-md text-center flex flex-col items-center justify-center gap-2">
                        <p className="text-xs font-medium text-white/40">
                            Monthly spend
                        </p>
                        <p className="font-bold text-violet-400">$124.80</p>
                    </div>
                    <div className="flex-1 bg-slate-900 py-4 px-6 rounded-md text-center flex flex-col items-center justify-center gap-2">
                        <p className="text-xs font-medium text-white/40">
                            Annual estimate
                        </p>
                        <p className="font-bold text-green-400">$1,497.60</p>
                    </div>
                    <div className="flex-1 bg-slate-900 py-4 px-6 rounded-md text-center flex flex-col items-center justify-center gap-2">
                        <p className="text-xs font-medium text-white/40">
                            Active Subscriptions
                        </p>
                        <p className="font-bold text-orange-400">11</p>
                    </div>
                </div>

                <div className="rounded-md bg-slate-900 opacity-60 border-t border-white/5 [mask-image:linear-gradient(to_bottom,black_40%,transparent)] [-webkit-mask-image:linear-gradient(to_bottom,black_40%,transparent)]">
                    <div>
                        <div className="flex justify-between items-center px-6 py-4 bg-slate-9">
                            <div>
                                <div className="w-8 h-8 rounded-lg bg-green-400/20 border border-green-400/40 flex items-center justify-center">
                                    <div className="w-2 h-2 rounded-full bg-green-400"></div>
                                </div>
                            </div>
                            <div className="flex flex-col justify-center items-center gap-1">
                                <p className="text-white text-sm">Spotify</p>
                                <span className="text-white/60 text-xs">
                                    Next: Jul 10
                                </span>
                            </div>
                            <div>
                                <p className="text-md text-white">$9.99</p>
                            </div>
                        </div>
                    </div>
                    <div className="border-t border-white/10">
                        <div className="flex justify-between items-center px-6 py-4">
                            <div>
                                <div className="w-8 h-8 rounded-lg bg-red-400/20 border border-red-400/40 flex items-center justify-center">
                                    <div className="w-2 h-2 rounded-full bg-red-400"></div>
                                </div>
                            </div>
                            <div className="flex flex-col justify-center items-center gap-1">
                                <p className="text-white text-sm">Netflix</p>
                                <span className="text-white/60 text-xs">
                                    Next: Jul 15
                                </span>
                            </div>
                            <div>
                                <p className="text-md text-white">$15.99</p>
                            </div>
                        </div>
                    </div>
                    <div className="border-t border-white/10">
                        <div className="flex justify-between items-center px-6 py-4">
                            <div>
                                <div className="w-8 h-8 rounded-lg bg-red-400/20 border border-red-400/40 flex items-center justify-center">
                                    <div className="w-2 h-2 rounded-full bg-red-400"></div>
                                </div>
                            </div>
                            <div className="flex flex-col justify-center items-center gap-1">
                                <p className="text-white text-sm">Adobe CC</p>
                                <span className="text-white/60 text-xs">
                                    Next: Jul 20
                                </span>
                            </div>
                            <div>
                                <p className="text-md text-white">$54.99</p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}
