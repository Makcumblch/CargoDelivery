import { useCallback, useContext, useState } from "react";
import { AuthContext } from "../contexts/AuthContext";

export interface ReqError {
    status: number
    message: string
}

export const useHttp = () => {
    const { logout } = useContext(AuthContext)
    const [loading, setLoading] = useState<boolean>(false)
    const [error, setError] = useState<ReqError | null>(null)

    const request = useCallback(async (url: string, method: string = 'GET', body: object | string | null = null, headers: HeadersInit | undefined = {}) => {
        setLoading(true)
        cleareError()
        let response = null
        try {
            if (body && typeof body !== 'string') {
                body = JSON.stringify(body)
            }
            response = await fetch(url, { method, body, headers })
            const data = await response.json()
            if (!response.ok) {
                if (response.status === 401) {
                    logout()
                    throw new Error(data.message || 'Пользователь не авторизован')
                }
                throw new Error(data.message || 'Что-то пошло не так')
            }
            setLoading(false)
            return data
        } catch (e) {
            setLoading(false)
            setError({ status: response?.status ?? -1, message: (e as Error).message })
            throw e
        }
    }, [logout])

    const cleareError = () => setError(null)

    return { loading, request, error }
}