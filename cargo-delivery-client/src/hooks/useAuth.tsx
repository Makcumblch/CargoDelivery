import { useCallback, useEffect, useState } from "react"

const storageName = 'token'

export interface IAuth {
    login: (jwtToken: string) => void,
    logout: () => void,
    token: string | null
}

export const useAuth = (): IAuth => {
    const [token, setToken] = useState<IAuth["token"]>(null)

    const login = useCallback((jwtToken: string) => {
        setToken(jwtToken)
        localStorage.setItem(storageName, jwtToken)
    }, [])

    const logout = useCallback(() => {
        setToken(null)
        localStorage.removeItem(storageName)
    }, [])

    useEffect(() => {
        setToken(localStorage.getItem(storageName))
    }, [])

    return {login, logout, token}
}

