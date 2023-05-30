import { useContext } from 'react'
import { RoutesContext } from '../../../../contexts/RoutesContext';
import Spinner from '../../../Spinner';
import ListRoutesItem from './ListRoutesItem';

const ListRoutes = () => {
    const { isLoadingRoutesList, routes } = useContext(RoutesContext)
    return (
        <div className="bg-slate-700 relative p-1 overflow-y-auto">
            {routes.length ? routes.map((el, index) => {
                return <ListRoutesItem key={index} index={index} item={el}/>
            }) : <div className='text-white text-center'>Нет маршрутов</div>}
            {isLoadingRoutesList &&
                <div className="absolute z-[1005] w-full h-full left-0 top-0 bg-slate-800 opacity-70">
                    <Spinner />
                </div>}
        </div>
    );
}

export default ListRoutes