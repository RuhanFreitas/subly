import MiniWeb from '@/components/MiniWeb'
import {
    BellIcon,
    ChartColumnIcon,
    LayoutDashboardIcon,
    RefreshCwIcon,
    ShieldCheckIcon,
    TrendingDownIcon,
} from 'lucide-react'
import { Inter, Playfair } from 'next/font/google'

const playfair = Playfair({
    subsets: ['latin'],
})

const inter = Inter({
    subsets: ['latin'],
})

export default function Home() {
    return (
        <div className={`z-10 min-h-screen ${inter.className}`}>
            <div className="flex flex-col justify-center items-center min-w-full my-8">
                <div className="flex items-center justify-center gap-2 mx-auto bg-purple-400/5 border border-gray-800 px-4 py-1 my-8 rounded-full">
                    <span className="bg-violet-400 h-1 w-1 rounded-full"></span>
                    <p className="text-violet-400 text-sm mx-auto">
                        Full control over subscriptions
                    </p>
                </div>
                <div className="max-w-5xl flex flex-col gap-8">
                    <p
                        className={`text-white text-8xl text-center font-extrabold leading-24 tracking-tighter mx-auto ${playfair.className}`}
                    >
                        Every dollar spent on {''}
                        <span className="text-violet-400">
                            subscriptions, visible.
                        </span>
                    </p>
                    <p className="text-slate-300/70 max-w-lg text-lg font-light mx-auto text-center tracking-wide">
                        Subly centralizes all your subscriptions, shows the real
                        impact on your budget, and alerts you before charges
                        catch you off guard.
                    </p>
                </div>
                <div className="flex gap-8 my-8">
                    <button className="text-white text-bold bg-violet-400 py-3 px-6 rounded-md hover:cursor-pointer">
                        Get started for free -&gt;
                    </button>
                    <button className="text-white text-bold bg-white/5 border border-gray-800 py-3 px-6 rounded-md hover:cursor-pointer">
                        I already have an account
                    </button>
                </div>
                <MiniWeb />
                <section className="my-16 w-full flex justify-center items-center border-y border-white/10">
                    <div className="flex items-center justify-between min-w-2xl py-18">
                        <div className="flex-1 py-12 px-8 text-white text-center border-r border-white/10 max-w-56">
                            <p className="text-3xl font-semibold pb-4">
                                $210
                                <span className="text-violet-400 font-light">
                                    /mo
                                </span>
                            </p>
                            <p className="text-xs font-light">
                                Average monthly spend discovered per user
                            </p>
                        </div>
                        <div className="flex-1 py-12 px-8 text-white text-center border-r border-white/10 max-w-56">
                            <p className="text-3xl font-semibold pb-4">4.2x</p>
                            <p className="text-xs font-light">
                                More awereness about recurring charges
                            </p>
                        </div>
                        <div className="flex-1 py-12 px-8 text-white text-center max-w-56">
                            <p className="text-3xl font-semibold pb-4">23%</p>
                            <p className="text-xs font-light">
                                Average reduction in subscriptions after 1 month
                            </p>
                        </div>
                    </div>
                </section>
                <section className="py-12">
                    <div className="flex flex-col gap-6 text-center">
                        <h2 className="text-violet-400 mx-auto text-sm font-semibold tracking-wide">
                            FEATURES
                        </h2>
                        <h1
                            className={`text-white mx-auto text-5xl font-bold tracking-wide max-w-sm ${playfair.className}`}
                        >
                            Everything you need to take back control
                        </h1>
                        <p
                            className={`text-gray-200 mx-auto font-light max-w-sm ${inter.className}`}
                        >
                            From registration to cancellation - Subly tracks
                            every step of your subscriptions.
                        </p>
                    </div>
                    <div className="w-full">
                        <div className="mx-auto my-18 grid grid-cols-3 max-w-6xl border border-white/10 rounded-3xl bg-violet-400/5">
                            <div className="px-8 py-8 flex flex-col gap-4 border-r border-b border-white/10">
                                <div className="text-violet-400 bg-violet-400/10 w-fit px-3 py-3 rounded-lg border border-white/10">
                                    <LayoutDashboardIcon />
                                </div>
                                <p className="text-white font-semibold">
                                    Smart Dashboard
                                </p>
                                <p className="text-gray-200 font-light text-sm max-w-64">
                                    Consolidated view of all your recurring
                                    expenses, with real-time monthly and annual
                                    projections.
                                </p>
                            </div>
                            <div className="px-8 py-8 flex flex-col gap-4 border-r border-b border-white/10">
                                <div className="text-violet-400 bg-violet-400/10 w-fit px-3 py-3 rounded-lg border border-white/10">
                                    <BellIcon />
                                </div>
                                <p className="text-white font-semibold">
                                    Payment Alerts
                                </p>
                                <p className="text-gray-200 font-light text-sm max-w-64">
                                    Know in advance when a charge is coming. No
                                    more surprises on your bank statement.
                                </p>
                            </div>
                            <div className="px-8 py-8 flex flex-col gap-4 border-b border-white/10">
                                <div className="text-violet-400 bg-violet-400/10 w-fit px-3 py-3 rounded-lg border border-white/10">
                                    <TrendingDownIcon />
                                </div>
                                <p className="text-white font-semibold">
                                    Spending Analytics
                                </p>
                                <p className="text-gray-200 font-light text-sm max-w-64">
                                    Clear charts showing your spending evolution
                                    and where money goes every month.
                                </p>
                            </div>
                            <div className="px-8 py-8 flex flex-col gap-4 border-r border-white/10">
                                <div className="text-violet-400 bg-violet-400/10 w-fit px-3 py-3 rounded-lg border border-white/10">
                                    <ShieldCheckIcon />
                                </div>
                                <p className="text-white font-semibold">
                                    Complete History
                                </p>
                                <p className="text-gray-200 font-light text-sm max-w-64">
                                    Cancelled subscriptions stay in your
                                    history. Full control over what came and
                                    went.
                                </p>
                            </div>
                            <div className="px-8 py-8 flex flex-col gap-3 border-r border-white/10">
                                <div className="text-violet-400 bg-violet-400/10 w-fit px-3 py-3 rounded-lg border border-white/10">
                                    <ChartColumnIcon />
                                </div>
                                <p className="text-white font-semibold">
                                    Categories
                                </p>
                                <p className="text-gray-200 font-light text-sm max-w-64">
                                    Organize by entertainments, tools,
                                    productivity - and understand your
                                    consumption profile.
                                </p>
                            </div>
                            <div className="px-8 py-8 flex flex-col gap-3 border-white/10">
                                <div className="text-violet-400 bg-violet-400/10 w-fit px-3 py-3 rounded-lg border border-white/10">
                                    <RefreshCwIcon />
                                </div>
                                <p className="text-white font-semibold">
                                    All Billing Cycles
                                </p>
                                <p className="text-gray-200 font-light text-sm max-w-64">
                                    Daily weekly, monthly, or yearly. Any
                                    billing frequency, all in one place.
                                </p>
                            </div>
                        </div>
                    </div>
                </section>
                <section className="my-12 text-center max-w-md flex flex-col items-center justify-center gap-8">
                    <h2
                        className={`text-6xl text-white font-bold ${playfair.className}`}
                    >
                        Start tracking your spending today.
                    </h2>
                    <p className="text-gray-200 text-light text-sm">
                        Free, no credit card required. Takes less than 2
                        minutes.
                    </p>
                    <button className="text-white text-bold bg-violet-400 py-3 px-6 rounded-md hover:cursor-pointer max-w-56">
                        Create my account -&gt;
                    </button>
                </section>
            </div>
        </div>
    )
}
