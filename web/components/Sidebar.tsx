import {
    CreditCardIcon,
    LayoutDashboardIcon,
    LogOutIcon,
    RefreshCwIcon,
} from 'lucide-react'
import { Playfair } from 'next/font/google'
import AsideButton from './AsideButton'
import { ReactNode } from 'react'

const playfair = Playfair({
    subsets: ['latin'],
})

export default function Sidebar() {
    return (
        <aside className="h-screen max-w-60 break-all">
            <div className="h-full border-r border-white/10 flex flex-col">
                <div className="px-6 py-4 border-b border-white/10">
                    <div className="flex items-center gap-2">
                        <div className="text-violet-400 bg-violet-400/10 w-fit px-2 py-2 rounded-md">
                            <RefreshCwIcon size="18" />
                        </div>
                        <h1 className={`text-white text-lg font-semibold`}>
                            Subly
                        </h1>
                    </div>
                </div>
                <div className="px-6 py-4 flex flex-col gap-2 flex-1">
                    <AsideButton
                        Icon={LayoutDashboardIcon}
                        content="Dashboard"
                    />
                    <AsideButton
                        Icon={CreditCardIcon}
                        content="Subscriptions"
                    />
                </div>
                <div className="px-6 py-4 text-white border-t border-white/10 flex flex-col gap-4">
                    <div className="flex gap-2 items-center hover:bg-violet-400/20 px-2 py-1 rounded-md">
                        <div className="bg-violet-400 w-8 h-8 flex items-center justify-center rounded-lg">
                            <p className="text-sm font-bold">R</p>
                        </div>
                        <div className="flex flex-col">
                            <p className="text-white font-semibold">
                                Ruhan Freitas
                            </p>
                            <small className="text-xs text-white/60">
                                ruhanfreitas@gmail.com
                            </small>
                        </div>
                    </div>
                    <div className="flex items-center px-3 py-2 hover:bg-red-800/20 rounded-md">
                        <button className="flex gap-2 text-gray-400">
                            <LogOutIcon size="16" />
                            <p className="text-xs">Sign Out</p>
                        </button>
                    </div>
                </div>
            </div>
        </aside>
    )
}
