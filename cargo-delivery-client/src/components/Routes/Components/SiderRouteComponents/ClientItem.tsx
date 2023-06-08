import { useContext, useState } from "react"
import { ClientRoute, Items, RouteContext } from "../../../../contexts/RouteContext"
import CargoItem from "./CargoItem"

interface IClientItemProps {
    index: number
    item: ClientRoute
    indexCar: number
    cargos: Items[]
}

const ClientItem = ({ index, indexCar, item, cargos }: IClientItemProps) => {
    const { setVisibleClient } = useContext(RouteContext)
    const [isOpen, setIsOpen] = useState<boolean>(false)
    return (
        <div
            className={`text-white p-1 mb-2 rounded-md w-full cursor-pointer bg-slate-800`}
        >
            <div className='grid grid-flow-col grid-cols-[1fr_1fr]'>
                <p>Название</p>
                <p className="text-right">{item.name}</p>
            </div>
            <div>
                <p>{`Адрес: ${item.address}`}</p>
            </div>
            {cargos.length ? <>
                <div className='grid grid-flow-col grid-cols-[1fr_20px] gap-1 h-7'>
                    <div
                        className="w-full bg-slate-700 hover:bg-slate-500 flex align-middle justify-center rounded-md cursor-pointer"
                        onClick={(e) => {
                            e.stopPropagation()
                            setIsOpen(prev => !prev)
                        }}
                    >
                        {isOpen ? '∧ Товары' : '∨ Товары'}
                    </div>
                    <input type={'checkbox'} checked={item.isVisible} onChange={(e) => setVisibleClient(indexCar, index, e.target.checked)}/>
                </div>
                <div className={`${!isOpen ? 'hidden' : ''} mt-1 p-1 bg-slate-600 rounded-md`}>
                    {cargos && cargos.map((el, index) => {
                        return <CargoItem key={index} item={el.cargo} position={el.position} />
                    })}
                </div>
            </> : null}
        </div>
    )
}

export default ClientItem