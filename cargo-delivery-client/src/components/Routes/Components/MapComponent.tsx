import { useContext, useEffect, useMemo, useRef } from 'react';
import { MapContainer, Marker, Polyline, Popup, TileLayer } from 'react-leaflet'
import { Depo, RoutesContext } from '../../../contexts/RoutesContext';
import { LatLngExpression, Map } from 'leaflet';
import { Client } from '../../../contexts/ClientsContext';
import { RouteContext } from '../../../contexts/RouteContext';

const stringToHex = (str: string) => {
    var result = '';
    for (var i = 0; i < str.length; i++) {
        result += str.charCodeAt(i).toString(16);
    }
    return result;
}

const getColorByName = (name: string) => {
    const hexStr = stringToHex(`${name}name`)
    return `#${hexStr.slice(0, 2)}${hexStr.slice(2, 4)}${hexStr.slice(4, 6)}`
}

const MapComponent = () => {
    const { routes, selectRouteIndex, inDetail } = useContext(RoutesContext)
    const { routeSolution } = useContext(RouteContext)

    const map = useRef<Map | null>()

    const printMarker = (client: Depo | Client, index: number) => {
        return (
            <Marker title={client.name} key={index} position={[client.coordY, client.coordX]}>
                <Popup>
                    <p>{client.name}</p>
                    <p>{client.address}</p>
                </Popup>
            </Marker>
        )
    }

    const route = useMemo(() => {
        if (selectRouteIndex === -1) return
        return routes[selectRouteIndex]
    }, [routes, selectRouteIndex])

    useEffect(() => {
        if (!route || !map.current) return
        const depo = route.clients[0]
        if (!depo) return
        map.current.setView([depo.coordY, depo.coordX], 5)
    }, [route])

    return (
        <MapContainer
            className='h-full w-full'
            center={[54.7431, 55.9678]}
            zoom={5}
            scrollWheelZoom={true}
            ref={(r) => map.current = r}
        >
            <TileLayer
                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            />
            {route && route.clients.map((el, index) => {
                return printMarker(el, index)
            })}
            {inDetail
                ?
                <>
                    {routeSolution && routeSolution.carsRouteSolution.map((el, index) => {
                        if (!el.isVisible) return null
                        return <Polyline key={index} color={getColorByName(el.name)} positions={el.route.polyline as LatLngExpression[] | LatLngExpression[][]} />
                    })}
                </>
                :
                <>
                    {route && route.cars_routes.map((el, index) => {
                        return (
                            <Polyline key={index} color={getColorByName(el.name)} positions={el.polyline} />
                        )
                    }
                    )}
                </>
            }
        </MapContainer>
    );
}

export default MapComponent