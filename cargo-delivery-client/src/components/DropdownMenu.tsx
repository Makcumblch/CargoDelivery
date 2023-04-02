import { useState } from "react"


interface DropdownMenuProps {
    btnContent?: React.ReactNode | null
    items: React.ReactNode[]
}

export const DropdownMenu = ({ btnContent = null, items }: DropdownMenuProps) => {
    const [isOpen, setIsOpen] = useState<boolean>(false)
    return (
        <>
            <div className="relative flex items-start">
                <button
                    className="flex items-center rounded text-center bg-slate-500 hover:bg-slate-400"
                    onClick={() => setIsOpen(prev => !prev)}
                >
                    {btnContent ? btnContent :
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="5 5 10 10"
                            fill="currentColor"
                            className="h-5 w-3 mx-3 my-2">
                            <path
                                fillRule="evenodd"
                                d="M5.23 7.21a.75.75 0 011.06.02L10 11.168l3.71-3.938a.75.75 0 111.08 1.04l-4.25 4.5a.75.75 0 01-1.08 0l-4.25-4.5a.75.75 0 01.02-1.06z"
                                clipRule="evenodd" />
                        </svg>}
                </button>

                <ul
                    className={`${isOpen ? '' : 'hidden'} absolute z-[1000] ml-2.5 left-full min-w-max list-none rounded shadow-sm bg-gray-600 overflow-visible`}
                >
                    {items.map((item, index) => {
                        return (
                            <li key={index} className="block w-full p-2 text-sm text-slate-300 font-norma rounded hover:bg-gray-500 shadow-2xl">
                                {item}
                            </li>
                        )
                    })}
                </ul>
            </div>
            <div className={`${isOpen ? '' : 'hidden'} z-[999] fixed top-0 left-0 w-screen h-screen overflow-visible`} onClick={() => {
                setIsOpen(false)
            }} />
        </>
    )
}