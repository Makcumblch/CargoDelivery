import React, { createContext, useContext, useEffect, useState } from 'react'
import { useHttp } from '../hooks/useHttp'
import { AuthContext } from './AuthContext'
import { ProjectsContext } from './ProjectsContext'

export interface Client {
    id: number,
    name: string,
    address: string,
}

interface ClientsProps {
    children: React.ReactNode
}

interface IClientsContext {
    clients: Client[],
    isLoading: boolean,
    delClient: (id: number) => void,
    changeClient: (id: number | null, values: object) => void,
    addClient: (id: number | null, values: object) => void,
}

export const ClientsContext = createContext<IClientsContext>({
    clients: [],
    isLoading: true,
    delClient: () => { },
    changeClient: () => { },
    addClient: () => { },
})

export const Clients = ({ children }: ClientsProps) => {
    const { token } = useContext(AuthContext)
    const { currentProjectId } = useContext(ProjectsContext)
    const { request } = useHttp()
    const [clients, setClients] = useState<Client[]>([])
    const [isLoading, setIsLoading] = useState<boolean>(true)

    useEffect(() => {
        const getClients = async () => {
            try {
                setIsLoading(true)
                let clientsData = await request(`api/projects/${currentProjectId}/clients`, 'GET', null, {
                    Authorization: `Bearer ${token}`
                })
                setIsLoading(false)
                clientsData = clientsData.data as Client[]
                if (!clientsData) {
                    setClients([])
                    return
                }
                setClients(clientsData)
            } catch (e) {
                setIsLoading(false)
            }
        }
        if (currentProjectId === null) return
        getClients()
    }, [token, request, currentProjectId])

    const delClient = async (id: number) => {
        try {
            const data = await request(`api/projects/${currentProjectId}/clients/${id}`, 'DELETE', null, {
                Authorization: `Bearer ${token}`
            })
            if (data.status !== 'ok') return
            setClients(prev => {
                const newClients: Client[] = []
                for (let i = 0; i < prev.length; ++i) {
                    if (prev[i].id !== id) {
                        newClients.push({ ...prev[i] })
                    }
                }
                return newClients
            })
        } catch (e) { }
    }

    const changeClient = async (id: number | null, values: object) => {
        if (id === null) return
        const clientIndex = clients.findIndex((c) => c.id === id)
        if (clientIndex === -1) return
        try {
            const data = await request(`api/projects/${currentProjectId}/clients/${id}`, 'PUT', values, {
                Authorization: `Bearer ${token}`
            })
            if (data.status !== 'ok') return
            setClients(prev => {
                const newClients = [...prev]
                newClients[clientIndex] = { ...newClients[clientIndex], ...values }
                return newClients
            })
        } catch (e) { }
    }

    const addClient = async (id: number | null, values: object) => {
        try {
            const data = await request(`api/projects/${currentProjectId}/clients`, 'POST', values, {
                Authorization: `Bearer ${token}`
            })
            const client: any = { ...values, id: data.id }
            setClients(prev => {
                const newClients: Client[] = [...prev]
                newClients.push(client)
                return newClients
            })
        } catch (e) { }
    }

    return (
        <ClientsContext.Provider value={{
            clients,
            isLoading,
            delClient,
            changeClient,
            addClient,
        }}>
            {children}
        </ClientsContext.Provider>
    )
}