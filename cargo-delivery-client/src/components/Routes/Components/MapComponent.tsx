import { useContext, useEffect, useMemo, useRef } from 'react';
import { MapContainer, Marker, Polyline, Popup, TileLayer, useMap } from 'react-leaflet'
import { Depo, RoutesContext } from '../../../contexts/RoutesContext';
import { LatLngExpression, Map } from 'leaflet';
import { Client, ClientsContext } from '../../../contexts/ClientsContext';

const MapComponent = () => {
    const { depo, routes, selectRouteIndex } = useContext(RoutesContext)
    const { clients } = useContext(ClientsContext)
    const { } = useContext(RoutesContext)

    const map = useRef<Map | null>()

    useEffect(() => {
        if (depo.id === -1 || !map.current) return
        map.current.setView([depo.coordY, depo.coordX], 5)
    }, [depo])

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

    return (
        <MapContainer
            center={[54.7431, 55.9678]}
            zoom={5}
            scrollWheelZoom={true}
            style={{ height: '100%', width: '100%' }}
            ref={(r) => map.current = r}
        >
            <TileLayer
                attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            />
            {depo.id !== -1 &&
                printMarker(depo, 1)
            }
            {clients.map((el, index) => {
                return printMarker(el, index)
            })}
            {route && route.cars_routes.map((el, index) => {
                return (
                    <Polyline key={index} color='red' positions={el.polyline}/>
                )
            })}
        </MapContainer>
    );
}

export default MapComponent