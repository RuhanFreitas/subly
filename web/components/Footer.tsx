import { RefreshCwIcon } from 'lucide-react'

export default function Footer() {
    return (
        <footer className="px-12 py-6 flex justify-between items-center border-t border-white/10">
            <div className="flex items-center gap-2">
                <div className="text-violet-400 bg-violet-400/10 w-fit px-1 py-1 rounded-md">
                    <RefreshCwIcon />
                </div>
                <p className="text-md font-semibold text-white">Subly</p>
            </div>
            <div>
                <p className="text-xs text-gray-200">
                    © {new Date().getFullYear()} Subly. All rights reserved.
                </p>
            </div>
        </footer>
    )
}
