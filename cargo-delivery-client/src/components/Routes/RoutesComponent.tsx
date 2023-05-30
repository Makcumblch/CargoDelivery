import { Routes } from '../../contexts/RoutesContext';
import MapComponent from './Components/MapComponent';
import SiderMenuComponent from './Components/SiderMenuComponent';

const RoutesComponent = () => {

    return (
        <Routes>
            <div className='grid grid-cols-[1fr_400px] h-full gap-2'>
                <MapComponent />
                <SiderMenuComponent />
            </div>
        </Routes>
    );
}

export default RoutesComponent