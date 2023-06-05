import React, { createContext, useContext, useEffect, useState } from 'react'
import { ReqError, useHttp } from '../hooks/useHttp'
import { AuthContext } from './AuthContext'
import { ProjectsContext } from './ProjectsContext'
import { Car } from './CarsContext'
import { LatLngExpression } from 'leaflet'
import { Client } from './ClientsContext'

export interface Depo {
    id: number,
    address: string,
    coordX: number,
    coordY: number
    name: string,
}

export interface CarsRoutes extends Car  {
    polyline: LatLngExpression[][]
    checked: boolean
}

export interface ItemRoutesList {
    id: number
    date: string
    fuel: number
    distance: number
    packing: number
    clients: Client[]
    cars_routes: CarsRoutes[]
}

interface RoutesProps {
    children: React.ReactNode
}

interface IRoutesContext {
    isLoadingCreate: boolean,
    isLoadingRoutesList: boolean,
    depo: Depo,
    updateDepo: (id: number, values: any) => void,
    routes: ItemRoutesList[],
    createRoute: (count: number, TMax: number, TMin: number, packingType: boolean) => void,
    deleteRoute: (id: number) => void,
    selectRouteIndex: number,
    setSelectRouteIndex: (index: number) => void,
    error: ReqError | null,
    isPacking: boolean,
    setIsPacking: (value: boolean) => void,
    inDetail: boolean,
    setInDetail: (value: boolean) => void,
}

export const RoutesContext = createContext<IRoutesContext>({
    isLoadingCreate: false,
    isLoadingRoutesList: false,
    depo: { id: -1, address: '', coordX: -1, coordY: -1, name: 'Депо' },
    updateDepo: () => { },
    routes: [],
    createRoute: () => { },
    deleteRoute: () => { },
    selectRouteIndex: -1,
    setSelectRouteIndex: () => {},
    error: null,
    inDetail: false,
    isPacking: false,
    setIsPacking: () => {},
    setInDetail: () => {},
})

export const Routes = ({ children }: RoutesProps) => {
    const { token } = useContext(AuthContext)
    const { currentProjectId } = useContext(ProjectsContext)
    const { request, error } = useHttp()

    const [isLoadingCreate, setIsLoadingCreate] = useState<boolean>(false)
    const [isLoadingRoutesList, setIsLoadingRoutesList] = useState<boolean>(false)

    const [depo, setDepo] = useState<Depo>({ id: -1, address: '', coordX: -1, coordY: -1, name: 'Депо' })
    const [routes, setRoutes] = useState<ItemRoutesList[]>([])
    const [selectRouteIndex, setSelectRouteIndex] = useState<number>(-1)

    const [inDetail, setInDetail] = useState<boolean>(false)
    
    const [isPacking, setIsPacking] = useState<boolean>(false)

    useEffect(() => {
        const getDepo = async () => {
            try {
                let depoData = await request(`api/projects/${currentProjectId}/depo`, 'GET', null, {
                    Authorization: `Bearer ${token}`
                })
                depoData = depoData.data as Depo
                if (!depoData) return
                setDepo(depoData)
            } catch (e) { }
        }

        const getListRoutes = async () => {
            try {
                let routesData = await request(`api/projects/${currentProjectId}/routes`, 'GET', null, {
                    Authorization: `Bearer ${token}`
                })
                routesData = routesData.data as ItemRoutesList[]
                if (!routesData) return
                setRoutes(routesData)
                if(routesData.length){
                    setSelectRouteIndex(0)
                }
            } catch (e) { }
        }

        getListRoutes()
        getDepo()
    }, [token, request, currentProjectId])

    const updateDepo = async (id: number, values: any) => {
        try {
            if (id === -1) {
                const data = await request(`api/projects/${currentProjectId}/depo`, 'POST', values, {
                    Authorization: `Bearer ${token}`
                })
                setDepo({ id: data.id, ...values })
            } else {
                const data = await request(`api/projects/${currentProjectId}/depo`, 'PUT', values, {
                    Authorization: `Bearer ${token}`
                })
                if (data.status !== 'ok') return
                setDepo(prev => ({ ...prev, ...values }))
            }
        } catch (e) { }
    }

    const createRoute = async (count: number, TMax: number, TMin: number, packingType: boolean) => {
        console.log('packingType', packingType)
        setIsLoadingCreate(true)
        try {
            let newRoute = await request(`api/projects/${currentProjectId}/routes`, 'POST', {
                evCount: count,
                tMax: TMax,
                tMin: TMin,
                packingType: packingType,
            }, {
                Authorization: `Bearer ${token}`
            })
            const newRoutes = [...routes]
            newRoutes.unshift(newRoute.data)
            setRoutes(newRoutes)
            setSelectRouteIndex(prev => prev + 1)
        } catch (e) { }
        setIsLoadingCreate(false)
    }

    const deleteRoute = async (id: number) => {
        try {
            const data = await request(`api/projects/${currentProjectId}/routes/${id}`, 'DELETE', null, {
                Authorization: `Bearer ${token}`
            })
            if (data.status !== 'ok') return
            setRoutes(prev => {
                const newRoutes: ItemRoutesList[] = []
                for (let i = 0; i < prev.length; ++i) {
                    if (prev[i].id !== id) {
                        newRoutes.push({ ...prev[i] })
                    }else{
                        if(i <= selectRouteIndex){
                            setSelectRouteIndex(selectRouteIndex - 1)
                        }
                    }
                }
                if(newRoutes.length === 0){
                    setSelectRouteIndex(-1)
                }
                return newRoutes
            })
        } catch (e) { }
    }

    return (
        <RoutesContext.Provider value={{
            isLoadingCreate,
            isLoadingRoutesList,
            depo,
            updateDepo,
            routes,
            createRoute,
            deleteRoute,
            selectRouteIndex,
            setSelectRouteIndex,
            error,
            inDetail,
            setInDetail,
            isPacking,
            setIsPacking,
        }}>
            {children}
        </RoutesContext.Provider>
    )
}