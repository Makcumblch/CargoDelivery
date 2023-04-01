import { useState } from "react"


interface DropdownMenuProps {
    btnContent?: React.ReactNode | null
    items: React.ReactNode[]
}

export const DropdownMenu = ({ btnContent = null, items }: DropdownMenuProps) => {
    const [isOpen, setIsOpen] = useState<boolean>(false)
    return (
        <>
            <div className="relative">
                <button
                    className="flex items-center rounded text-center p-1.5 px-3 hover:bg-gray-600 hover:shadow-lg"
                    onClick={() => setIsOpen(prev => !prev)}
                >
                    {btnContent ? btnContent :
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 20 20"
                            fill="currentColor"
                            className="h-5 w-5">
                            <path
                                fillRule="evenodd"
                                d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z"
                                clipRule="evenodd" />
                        </svg>}
                </button>

                <ul
                    className={`${isOpen ? '' : 'hidden'} absolute z-[1000] float-left m-0 min-w-max list-none overflow-hidden rounded-lg border-none shadow-lg bg-gray-600`}
                >
                    {items.map((item, index) => {
                        return (
                            <li key={index} className="block w-full p-2 text-sm text-slate-300 font-norma hover:bg-gray-500">
                                {item}
                            </li>
                        )
                    })}
                </ul>
            </div>
            <div className={`${isOpen ? '' : 'hidden'} z-[999] absolute top-0 left-0 w-full h-full`} onClick={() => setIsOpen(false)}/>
        </>
    )
}