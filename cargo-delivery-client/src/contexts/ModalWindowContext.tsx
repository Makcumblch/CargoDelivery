import React, { createContext, useState } from 'react'

interface ModalWindowProps {
    children: React.ReactNode,
}

interface OpenProps {
    title: string,
    content: string | React.ReactNode,
    onOk: () => void,
    onCancel: () => void,
}

interface ModalWindowContextProps {
    open: (props: OpenProps) => void
    close: () => void,
}

export const ModalWindowContext = createContext<ModalWindowContextProps>({
    open: () => { },
    close: () => { },
})

export const ModalWindow = ({ children }: ModalWindowProps) => {
    const [isOpen, setIsOpen] = useState<boolean>(false)
    const [title, setTitle] = useState<string>('')
    const [content, setContent] = useState<string | React.ReactNode>('')
    const [onOk, setOnOk] = useState<() => void>(() => { })
    const [onCancel, setOnCancel] = useState<() => void>(() => { })

    const open = ({ title: _title, content: _content, onOk: _onOk, onCancel: _onCancel }: OpenProps) => {
        setTitle(_title)
        setContent(_content)
        setOnOk(_onOk)
        setOnCancel(_onCancel)
        setIsOpen(true)
    }

    return (
        <>
            <ModalWindowContext.Provider value={{
                open: open,
                close: () => setIsOpen(false),
            }}>
                {children}
            </ModalWindowContext.Provider>
            <div
                className={`${isOpen ? '' : 'hidden'} absolute top-0 z-[1001] left-0 w-full h-full flex flex-col justify-center min-h-screen backdrop-blur-sm overflow-hidden shadow-lg`}
                onClick={() => setIsOpen(false)}
            >
                <div className="w-full p-4 m-auto z-[1002] bg-slate-700 rounded-md shadow-md md:max-w-md" onClick={e => e.stopPropagation()}>
                    <h1 className="text-2xl font-semibold text-center text-white truncate">{title}</h1>
                    <div className='text-center text-slate-300 mt-2 truncate'>{content}</div>
                    <div className='flex justify-around mt-2'>
                        <button
                            form='form'
                            type='submit'
                            className="w-36 px-4 py-1.5 tracking-wide text-white transition-colors duration-200 transform bg-cyan-700 rounded-md hover:bg-cyan-600 focus:outline-none focus:bg-cyan-600"
                            onClick={() => {
                                onOk()
                            }}>Ок</button>
                        <button
                            className="w-36 px-4 py-1.5 tracking-wide text-white transition-colors duration-200 transform bg-cyan-700 rounded-md hover:bg-cyan-600 focus:outline-none focus:bg-cyan-600"
                            onClick={() => {
                                onCancel()
                            }}>Отмена</button>
                    </div>
                </div>
            </div>
        </>
    )
}