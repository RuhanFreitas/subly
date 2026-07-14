import React from 'react'

export default function AsideButton({
    Icon,
    content,
}: Readonly<{ Icon: React.ForwardRefExoticComponent<any>; content: string }>) {
    return (
        <button className="text-violet-400 flex gap-2 items-center bg-violet-400/20 w-full px-4 py-3 rounded-md hover:bg-violet-400/10">
            <Icon size={20} />
            <p className="text-white text-sm font-semibold">{content}</p>
        </button>
    )
}
