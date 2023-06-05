import { MouseEvent, useContext, useMemo } from 'react'
import { ItemRoutesList, RoutesContext } from '../../../../contexts/RoutesContext';
import DelBtn from '../../../Btns/DelBtn';

interface ListRoutesItemProps {
    item: ItemRoutesList
    index: number
}

const ListRoutesItem = ({ item, index }: ListRoutesItemProps) => {
    const { deleteRoute, selectRouteIndex, setSelectRouteIndex, setInDetail } = useContext(RoutesContext)
    const date = useMemo(() => {
        const newDate = new Date(item.date)
        return newDate.toLocaleString()
    }, [item])

    const inDetail = (e: any) => {
        e.stopPropagation()
        setSelectRouteIndex(index)
        setInDetail(true)
    }

    return (
        <div
            className={`text-white p-1 mb-1 rounded-sm hover:bg-slate-600 ${index === selectRouteIndex ? 'bg-slate-600' : 'bg-slate-800'}`}
            onClick={() => setSelectRouteIndex(index)}
        >
            <div className='grid grid-flow-col grid-cols-1'>
                <p>Дата</p>
                <p>{date}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-1'>
                <p>Дистанция (км)</p>
                <p>{item.distance}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-1'>
                <p>Топливо (л)</p>
                <p>{item.fuel}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-1'>
                <p>Качество упаковки</p>
                <p>{item.packing}</p>
            </div>
            <div className='flex justify-between mt-1'>
                <button className='bg-slate-500 rounded hover:bg-slate-400 px-2' onClick={inDetail}>Подробно</button>
                <DelBtn onClick={(e) => {
                    e.stopPropagation()
                    deleteRoute(item.id)
                }} />
            </div>
        </div>
    );
}

export default ListRoutesItem