import React, { useContext, useState } from "react";
import { Link } from "react-router-dom";

import { useHttp } from "../hooks/useHttp";
import { AuthContext } from "../contexts/AuthContext";

interface ISinginForm {
    username: string,
    password: string
}

function LoginPage() {
    const { login } = useContext(AuthContext)
    const { loading, error, request } = useHttp()
    const [form, setForm] = useState<ISinginForm>({
        username: "",
        password: ""
    })

    const changeHandler = (event: any): void => {
        setForm({ ...form, [event.target.name]: event.target.value })
    }

    const loginHandler = async () => {
        try {
            const data = await request('/auth/sign-in', 'POST', { ...form })
            login(data.token)
        } catch (e) { }
    }

    return (
        <div className="relative flex flex-col justify-center bg-slate-800 min-h-screen overflow-hidden">
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
                            onChange={changeHandler}
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
                            onChange={changeHandler}
                        />
                    </div>
                    {error && <p className="text-red-700 mb-0">
                        {error}
                    </p>}
                    <div className="mt-8">
                        <button
                            className="w-full px-4 py-2 tracking-wide text-white transition-colors duration-200 transform bg-cyan-700 rounded-md hover:bg-cyan-600 focus:outline-none focus:bg-cyan-600"
                            onClick={loginHandler}
                            disabled={loading}
                        >
                            Войти
                        </button>
                    </div>
                </form>

                <p className="mt-8 text-xs font-light text-center text-slate-300">
                    {" "}
                    У вас нет учетной записи?{" "}
                    <Link to='/signup' className="font-medium text-cyan-700 hover:underline">
                        Зарегистрироваться
                    </Link>
                </p>
            </div>
        </div>
    );
}

export default LoginPage;
