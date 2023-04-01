import { Routes, Route, Navigate } from 'react-router-dom'
import { Projects } from '../contexts/ProjectsContext'

import AppPage from "../pages/AppPage"
import LoginPage from "../pages/SigninPage"
import SigninPage from "../pages/SignupPage"

const useRoutes = (isAuthenticated: boolean) => {
    if (isAuthenticated) {
        return (
            <Routes>
                <Route path='/' element={<Projects><AppPage /></Projects>}></Route>
                <Route path='*' element={<Navigate to='/' replace />} />
            </Routes>
        )
    }

    return (
        <Routes>
            <Route path='/signin' element={<LoginPage />} />
            <Route path='/signup' element={<SigninPage />} />
            <Route path='*' element={<Navigate to='/signin' replace />} />
        </Routes>
    )
}

export default useRoutes;
