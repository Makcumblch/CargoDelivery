import { createContext } from 'react'
import { IAuth } from '../hooks/useAuth'

interface IAuthContext extends IAuth {
    isAuthenticated: boolean
}

export const AuthContext = createContext<IAuthContext>({
    login: () => {},
    logout: () => {},
    token: null,
    isAuthenticated: false
})