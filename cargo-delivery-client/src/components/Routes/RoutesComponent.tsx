import { useContext } from 'react';
import { Routes, RoutesContext } from '../../contexts/RoutesContext';
import MapComponent from './Components/MapComponent';
import SiderMenuRoutesComponent from './Components/SiderMenuRoutesComponent';
import RouteComponent from './Components/SiderRouteComponent';
import PackingComponent from './Components/PackingComponent';
import { RouteContxComp } from '../../contexts/RouteContext';

const RoutesComponent = () => {
    return (
        <Routes>
            <RouteContxComp>
                <RoutesAndPacking />
            </RouteContxComp>
        </Routes>
    );
}

const RoutesAndPacking = () => {
    const { inDetail, isPacking } = useContext(RoutesContext)
    return (
        <div className='grid grid-cols-[1fr_400px] h-full gap-2'>
            {isPacking ? <PackingComponent /> : <MapComponent />}
            {inDetail ? <RouteComponent /> : <SiderMenuRoutesComponent />}
        </div>
    );
}

export default RoutesComponent