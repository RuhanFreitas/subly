import { RefreshCwIcon } from 'lucide-react'
import { Inter, Playfair } from 'next/font/google'

const playfair = Playfair({
    subsets: ['latin'],
})

const inter = Inter({
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
            <div
                className={`flex gap-4 items-center ${inter.className} text-sm`}
            >
                <button className="py-2 px-5 rounded-lg text-center font-medium text-white border border-white/20 hover:border-violet-400 cursor-pointer">
                    Sign In
                </button>
                <button className="py-2 px-5 rounded-lg text-center font-medium text-white bg-violet-400 cursor-pointer">
                    Get started
                </button>
            </div>
        </nav>
    )
}
