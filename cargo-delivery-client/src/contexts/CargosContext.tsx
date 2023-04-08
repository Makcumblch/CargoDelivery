import React, { createContext, useContext, useEffect, useState } from 'react'
import { useHttp } from '../hooks/useHttp'
import { AuthContext } from './AuthContext'
import { ProjectsContext } from './ProjectsContext'

export interface Cargo {
    id: number,
    name: string,
    width: number,
    height: number,
    length: number,
    weight: number,
}

interface CargosProps {
    children: React.ReactNode
}

interface IProjectsContext {
    cargos: Cargo[],
    isLoading: boolean,
    delCargo: (id: number) => void,
    changeCargo: (id: number | null, values: object) => void,
    addCargo: (id: number | null, values: object) => void,
    getCargoById: (id: number) => Cargo | undefined,
}

export const CargosContext = createContext<IProjectsContext>({
    cargos: [],
    isLoading: true,
    delCargo: () => { },
    changeCargo: () => { },
    addCargo: () => { },
    getCargoById: () => undefined,
})

export const Cargos = ({ children }: CargosProps) => {
    const { token } = useContext(AuthContext)
    const { currentProjectId } = useContext(ProjectsContext)
    const { request } = useHttp()
    const [cargos, setCargos] = useState<Cargo[]>([])
    const [isLoading, setIsLoading] = useState<boolean>(true)

    useEffect(() => {
        const getCargos = async () => {
            try {
                setIsLoading(true)
                let cargosData = await request(`api/projects/${currentProjectId}/cargos`, 'GET', null, {
                    Authorization: `Bearer ${token}`
                })
                setIsLoading(false)
                cargosData = cargosData.data as Cargo[]
                if (!cargosData) {
                    setCargos([])
                    return
                }
                setCargos(cargosData)
            } catch (e) {
                setIsLoading(false)
            }
        }
        if (currentProjectId === null) return
        getCargos()
    }, [token, request, currentProjectId])

    const delCargo = async (id: number) => {
        try {
            const data = await request(`api/projects/${currentProjectId}/cargos/${id}`, 'DELETE', null, {
                Authorization: `Bearer ${token}`
            })
            if (data.status !== 'ok') return
            setCargos(prev => {
                const newCargos: Cargo[] = []
                for (let i = 0; i < prev.length; ++i) {
                    if (prev[i].id !== id) {
                        newCargos.push({ ...prev[i] })
                    }
                }
                return newCargos
            })
        } catch (e) { }
    }

    const changeCargo = async (id: number | null, values: object) => {
        if (id === null) return
        const cargoIndex = cargos.findIndex((c) => c.id === id)
        if (cargoIndex === -1) return
        try {
            const data = await request(`api/projects/${currentProjectId}/cargos/${id}`, 'PUT', values, {
                Authorization: `Bearer ${token}`
            })
            if (data.status !== 'ok') return
            setCargos(prev => {
                const newCargos = [...prev]
                newCargos[cargoIndex] = { ...newCargos[cargoIndex], ...values }
                return newCargos
            })
        } catch (e) { }
    }

    const addCargo = async (id: number | null, values: object) => {
        try {
            const data = await request(`api/projects/${currentProjectId}/cargos`, 'POST', values, {
                Authorization: `Bearer ${token}`
            })
            const cargo: any = { ...values, id: data.id }
            setCargos(prev => {
                const newCargos: Cargo[] = [...prev]
                newCargos.push(cargo)
                return newCargos
            })
        } catch (e) { }
    }

    const getCargoById = (id: number) => {
        const cargo = cargos.find((pr: Cargo) => pr.id === id)
        return cargo
    }

    return (
        <CargosContext.Provider value={{
            cargos,
            isLoading,
            delCargo,
            changeCargo,
            addCargo,
            getCargoById,
        }}>
            {children}
        </CargosContext.Provider>
    )
}