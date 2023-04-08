import React, { createContext, useContext, useEffect, useState } from 'react'
import { useHttp } from '../hooks/useHttp'
import { AuthContext } from './AuthContext'
import { ProjectsContext } from './ProjectsContext'

export interface Order {
    id: number,
    idCargo: number,
    count: number,
}

interface OrdersProps {
    children: React.ReactNode,
    clientId: number
}

interface IOrdersContext {
    orders: Order[],
    isLoading: boolean,
    delOrder: (id: number) => void,
    changeOrder: (id: number | null, values: object) => void,
    addOrder: (id: number | null, values: object) => void,
}

export const OrdersContext = createContext<IOrdersContext>({
    orders: [],
    isLoading: true,
    delOrder: () => { },
    changeOrder: () => { },
    addOrder: () => { },
})

export const Orders = ({ children, clientId }: OrdersProps) => {
    const { token } = useContext(AuthContext)
    const { currentProjectId } = useContext(ProjectsContext)
    const { request } = useHttp()
    const [orders, setOrders] = useState<Order[]>([])
    const [isLoading, setIsLoading] = useState<boolean>(true)

    useEffect(() => {
        const getOrders = async () => {
            try {
                setIsLoading(true)
                let ordersData = await request(`api/projects/${currentProjectId}/clients/${clientId}/orders`, 'GET', null, {
                    Authorization: `Bearer ${token}`
                })
                setIsLoading(false)
                ordersData = ordersData.data as Order[]
                if (!ordersData) {
                    setOrders([])
                    return
                }
                setOrders(ordersData)
            } catch (e) {
                setIsLoading(false)
            }
        }
        if (currentProjectId === null) return
        getOrders()
    }, [token, request, currentProjectId])

    const delOrder = async (id: number) => {
        try {
            const data = await request(`api/projects/${currentProjectId}/clients/${clientId}/orders/${id}`, 'DELETE', null, {
                Authorization: `Bearer ${token}`
            })
            if (data.status !== 'ok') return
            setOrders(prev => {
                const newOrders: Order[] = []
                for (let i = 0; i < prev.length; ++i) {
                    if (prev[i].id !== id) {
                        newOrders.push({ ...prev[i] })
                    }
                }
                return newOrders
            })
        } catch (e) { }
    }

    const changeOrder = async (id: number | null, values: object) => {
        if (id === null) return
        const orderIndex = orders.findIndex((c) => c.id === id)
        if (orderIndex === -1) return
        try {
            const data = await request(`api/projects/${currentProjectId}/clients/${clientId}/orders/${id}`, 'PUT', values, {
                Authorization: `Bearer ${token}`
            })
            if (data.status !== 'ok') return
            setOrders(prev => {
                const newOrders = [...prev]
                newOrders[orderIndex] = { ...newOrders[orderIndex], ...values }
                return newOrders
            })
        } catch (e) { }
    }

    const addOrder = async (id: number | null, values: object) => {
        try {
            const data = await request(`api/projects/${currentProjectId}/clients/${clientId}/orders`, 'POST', values, {
                Authorization: `Bearer ${token}`
            })
            const car: any = { ...values, id: data.id }
            setOrders(prev => {
                const newOrders: Order[] = [...prev]
                newOrders.push(car)
                return newOrders
            })
        } catch (e) { }
    }

    return (
        <OrdersContext.Provider value={{
            orders,
            isLoading,
            delOrder,
            changeOrder,
            addOrder,
        }}>
            {children}
        </OrdersContext.Provider>
    )
}