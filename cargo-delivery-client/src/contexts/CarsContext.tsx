import React, { createContext, useContext, useEffect, useState } from 'react'
import { useHttp } from '../hooks/useHttp'
import { AuthContext } from './AuthContext'
import { ProjectsContext } from './ProjectsContext'

export interface Car {
    id: number,
    name: string,
    loadCapacity: number,
    width: number,
    height: number,
    length: number,
    fuelConsumption: number,
}

interface CarsProps {
    children: React.ReactNode
}

interface ICarsContext {
    cars: Car[],
    isLoading: boolean,
    delCar: (id: number) => void,
    changeCar: (id: number | null, values: object) => void,
    addCar: (id: number | null, values: object) => void,
}

export const CarsContext = createContext<ICarsContext>({
    cars: [],
    isLoading: true,
    delCar: () => { },
    changeCar: () => { },
    addCar: () => { },
})

export const Cars = ({ children }: CarsProps) => {
    const { token } = useContext(AuthContext)
    const { currentProjectId } = useContext(ProjectsContext)
    const { request } = useHttp()
    const [cars, setCars] = useState<Car[]>([])
    const [isLoading, setIsLoading] = useState<boolean>(true)

    useEffect(() => {
        const getCars = async () => {
            try {
                setIsLoading(true)
                let carsData = await request(`api/projects/${currentProjectId}/cars`, 'GET', null, {
                    Authorization: `Bearer ${token}`
                })
                setIsLoading(false)
                carsData = carsData.data as Car[]
                if (!carsData) {
                    setCars([])
                    return
                }
                setCars(carsData)
            } catch (e) {
                setIsLoading(false)
            }
        }
        if (currentProjectId === null) return
        getCars()
    }, [token, request, currentProjectId])

    const delCar = async (id: number) => {
        try {
            const data = await request(`api/projects/${currentProjectId}/cars/${id}`, 'DELETE', null, {
                Authorization: `Bearer ${token}`
            })
            if (data.status !== 'ok') return
            setCars(prev => {
                const newCars: Car[] = []
                for (let i = 0; i < prev.length; ++i) {
                    if (prev[i].id !== id) {
                        newCars.push({ ...prev[i] })
                    }
                }
                return newCars
            })
        } catch (e) { }
    }

    const changeCar = async (id: number | null, values: object) => {
        if (id === null) return
        const carIndex = cars.findIndex((c) => c.id === id)
        if (carIndex === -1) return
        try {
            const data = await request(`api/projects/${currentProjectId}/cars/${id}`, 'PUT', values, {
                Authorization: `Bearer ${token}`
            })
            if (data.status !== 'ok') return
            setCars(prev => {
                const newCars = [...prev]
                newCars[carIndex] = { ...newCars[carIndex], ...values }
                return newCars
            })
        } catch (e) { }
    }

    const addCar = async (id: number | null, values: object) => {
        try {
            const data = await request(`api/projects/${currentProjectId}/cars`, 'POST', values, {
                Authorization: `Bearer ${token}`
            })
            const car: any = { ...values, id: data.id }
            setCars(prev => {
                const newCars: Car[] = [...prev]
                newCars.push(car)
                return newCars
            })
        } catch (e) { }
    }

    return (
        <CarsContext.Provider value={{
            cars,
            isLoading,
            delCar,
            changeCar,
            addCar,
        }}>
            {children}
        </CarsContext.Provider>
    )
}