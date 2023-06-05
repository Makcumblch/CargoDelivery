import { useEffect, createContext, useContext, useState } from 'react'
import { useHttp } from '../hooks/useHttp'
import { AuthContext } from './AuthContext'
import { ProjectsContext } from './ProjectsContext'
import { Car } from './CarsContext'
import { Client } from './ClientsContext'
import { Cargo } from './CargosContext'
import { RoutesContext } from './RoutesContext'

export interface ClientRoute extends Client {
    index: number
    isVisible: boolean
}

export interface Route {
    clients: ClientRoute[]
    polyline: number[][]
}

export interface Position {
    x_pos: number
    y_pos: number
    z_pos: number
}

export interface Items {
    cargo: Cargo
    client: Client
    position: Position
}

export interface CarRoute extends Car {
    freeLoadCapacity: number
    freeVolume: number
    items: Items[][]
    route: Route
    isVisible: boolean
}

export interface RouteSolution {
    packingCost: number
    fuel: number
    distance: number
    carsRouteSolution: CarRoute[]
}

interface IRouteContext {
    isLoading: boolean
    routeSolution: RouteSolution | undefined
    selectedCarIndex: number
    setSelectedCarIndex: (index: number) => void
    setVisibleCar: (index: number, visible: boolean) => void
    setVisibleClient: (indexCar: number, indexClient: number, visible: boolean) => void
}

interface RouteProps {
    children: React.ReactNode
}

export const RouteContext = createContext<IRouteContext>({
    isLoading: false,
    routeSolution: undefined,
    selectedCarIndex: -1,
    setSelectedCarIndex: () => { },
    setVisibleCar: () => {},
    setVisibleClient: () => {},
})

export const RouteContxComp = ({ children }: RouteProps) => {
    const { token } = useContext(AuthContext)
    const { currentProjectId } = useContext(ProjectsContext)
    const { request } = useHttp()
    const { inDetail, routes, selectRouteIndex } = useContext(RoutesContext)

    const [isLoading, setIsLoading] = useState<boolean>(false)
    const [routeSolution, setRouteSolution] = useState<RouteSolution | undefined>()
    const [selectedCarIndex, setSelectedCarIndex] = useState<number>(-1)

    useEffect(() => {
        if (!inDetail || selectRouteIndex === -1) return
        const id = routes[selectRouteIndex].id
        const getRoute = async () => {
            setIsLoading(true)
            try {
                let routesData = await request(`api/projects/${currentProjectId}/routes/${id}`, 'GET', null, {
                    Authorization: `Bearer ${token}`
                })
                routesData = routesData as RouteSolution
                if (!routesData) return
                for (let i in routesData.carsRouteSolution) {
                    routesData.carsRouteSolution[i].isVisible = true
                    for (let j in routesData.carsRouteSolution[i].route.clients) {
                        routesData.carsRouteSolution[i].route.clients[j].isVisible = true
                    }
                }
                setRouteSolution(routesData)
            } catch (e) { }
            setIsLoading(false)
        }
        getRoute()
    }, [inDetail, routes, selectRouteIndex, currentProjectId, request, token])

    const setVisibleCar = (index: number, visible: boolean) => {
        if (!routeSolution) return
        const newRouteSolution = { ...routeSolution }
        newRouteSolution.carsRouteSolution[index].isVisible = visible
        setRouteSolution(newRouteSolution)
    }

    const setVisibleClient = (indexCar: number, indexClient: number, visible: boolean) => {
        if (!routeSolution) return
        const newRouteSolution = { ...routeSolution }
        newRouteSolution.carsRouteSolution[indexCar].route.clients[indexClient].isVisible = visible
        setRouteSolution(newRouteSolution)
    }

    return (
        <RouteContext.Provider value={{
            isLoading,
            routeSolution,
            selectedCarIndex,
            setSelectedCarIndex,
            setVisibleCar,
            setVisibleClient,
        }}>
            {children}
        </RouteContext.Provider>
    )
}