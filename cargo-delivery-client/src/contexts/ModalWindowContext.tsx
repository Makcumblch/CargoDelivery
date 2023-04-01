import React, { createContext, useState } from 'react'



interface ModalWindowProps {
    children: React.ReactNode
}

export const ModalWindowContext = createContext({})

export const ModalWindow = ({ children }: ModalWindowProps) => {
    const [isOpen, setIsopen] = useState<boolean>(true)

    return (
        <>
            <div className='absolute top-0 left-0 w-full h-full justify-center align-middle'>
                <div className="w-full p-6 m-auto bg-slate-700 rounded-md shadow-md lg:max-w-xl">
                    <h1 className="text-3xl font-semibold text-center text-white">
                        Авторизация
                    </h1>
                    <form className="mt-6">
                        <div className="mb-2">
                            <label
                                htmlFor="username"
                                className="block text-sm font-semibold text-slate-300"
                            >
                                Логин
                            </label>
                            <input
                                type="username"
                                name="username"
                                placeholder="Введите логин"
                                className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                            />
                        </div>
                        <div className="mb-2">
                            <label
                                htmlFor="password"
                                className="block text-sm font-semibold text-slate-300"
                            >
                                Password
                            </label>
                            <input
                                type="password"
                                name="password"
                                placeholder="Введите пароль"
                                className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                            />
                        </div>
                        <div className="mt-8">
                            <button
                                className="w-full px-4 py-2 tracking-wide text-white transition-colors duration-200 transform bg-cyan-700 rounded-md hover:bg-cyan-600 focus:outline-none focus:bg-cyan-600"
                            >
                                Войти
                            </button>
                        </div>
                    </form>

                    <p className="mt-8 text-xs font-light text-center text-slate-300">
                        {" "}
                        У вас нет учетной записи?{" "}
                    </p>
                </div>
            </div>
            <ModalWindowContext.Provider value={{

            }}>
                {children}
            </ModalWindowContext.Provider>
        </>
    )
}