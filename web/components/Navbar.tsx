import { RefreshCwIcon } from 'lucide-react'
import { Playfair } from 'next/font/google'

const playfair = Playfair({
    subsets: ['latin'],
})

export default function Navbar() {
    return (
        <nav className="z-20 bg-gray-950 flex justify-between px-8 py-4 border-b border-white/10">
            <div className="flex items-center gap-2">
                <div className="text-violet-400 bg-violet-400/10 w-fit px-2 py-2 rounded-md">
                    <RefreshCwIcon />
                </div>
                <h1
                    className={`text-white text-xl font-semibold ${playfair.className}`}
                >
                    Subly
                </h1>
            </div>
            <div className="flex gap-4">
                <button className="py-1.5 px-5 rounded-lg text-center font-semibold text-white border border-white/60 ">
                    Sign In
                </button>
                <button className="py-1.5 px-5 rounded-lg text-center font-semibold text-white border border-white/60 bg-violet-400">
                    Get started
                </button>
            </div>
        </nav>
    )
}
