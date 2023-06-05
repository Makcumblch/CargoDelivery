import { useContext, useState } from "react"
import { CarRoute, RouteContext } from "../../../../contexts/RouteContext"
import ClientItem from "./ClientItem"

interface IRouteCarItemProps {
    index: number
    item: CarRoute
}

const RouteCarItem = ({ index, item }: IRouteCarItemProps) => {
    const { selectedCarIndex, setSelectedCarIndex, setVisibleCar } = useContext(RouteContext)
    const [isOpen, setIsOpen] = useState<boolean>(false)
    return (
        <div
            className={`text-white p-1 mb-2 rounded-md ${isOpen ? 'rounded-b-none' : ''} w-full cursor-pointer hover:bg-slate-600 ${index === selectedCarIndex ? 'bg-slate-600' : 'bg-slate-800'}`}
            onClick={() => setSelectedCarIndex(index)}
        >
            <div className='grid grid-flow-col grid-cols-[1fr_1fr]'>
                <p>Название ТС</p>
                <p className="truncate text-right">{item.name}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-[1fr_1fr]'>
                <p>Ширина кузова (м)</p>
                <p className="text-right">{item.width}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-[1fr_1fr]'>
                <p>Длина кузова (м)</p>
                <p className="text-right">{item.length}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-[1fr_1fr]'>
                <p>Высота кузова (м)</p>
                <p className="text-right">{item.height}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-[1fr_1fr]'>
                <p>Грузоподъемность (кг)</p>
                <p className="text-right">{item.loadCapacity}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-[1fr_1fr]'>
                <p>Расход (л/100км)</p>
                <p className="text-right">{item.fuelConsumption}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-[1fr_1fr]'>
                <p>Дистанция (км)</p>
                <p className="text-right">{'123'}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-[1fr_1fr]'>
                <p>Топливо (л)</p>
                <p className="text-right">{'123'}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-[1fr_1fr]'>
                <p>Качество упаковки</p>
                <p className="text-right">{'123'}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-[1fr_20px] gap-1 h-7'> 
                <div 
                className="w-full bg-slate-700 hover:bg-slate-500 flex align-middle justify-center rounded-md cursor-pointer"
                onClick={(e) => {
                    e.stopPropagation()
                    setIsOpen(prev => !prev)
                }}
                >
                    {isOpen ? '∧ Маршрут' : '∨ Маршрут'}
                </div>
                <input type={'checkbox'} checked={item.isVisible} onChange={(e) => setVisibleCar(index, e.target.checked)}/>
            </div>
            <div className={`${!isOpen ? 'hidden': ''} mt-1 p-1 bg-slate-600 rounded-md`}>
                {item.route.clients.map((el, i) => {
                    return <ClientItem key={i} indexCar={index} index={i} item={el} cargos={item.items[i]}/>
                })}
            </div>
        </div>
    )
}

export default RouteCarItem