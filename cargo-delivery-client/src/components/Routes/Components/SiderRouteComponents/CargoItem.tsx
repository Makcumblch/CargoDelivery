import { Cargo } from "../../../../contexts/CargosContext"

interface ICargoItemProps {
    item: Cargo
}

const CargoItem = ({ item }: ICargoItemProps) => {
    return (
        <div
            className={`text-white p-1 mb-2 rounded-md w-full cursor-pointer bg-slate-800`}
        >
            <div className='grid grid-flow-col grid-cols-[1fr_1fr]'>
                <p>Название</p>
                <p className="text-right">{item.name}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-[1fr_1fr]'>
                <p>Длина</p>
                <p className="text-right">{item.length}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-[1fr_1fr]'>
                <p>Ширина</p>
                <p className="text-right">{item.width}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-[1fr_1fr]'>
                <p>Высота</p>
                <p className="text-right">{item.height}</p>
            </div>
            <div className='grid grid-flow-col grid-cols-[1fr_1fr]'>
                <p>Вес</p>
                <p className="text-right">{item.weight}</p>
            </div>
        </div>
    )
}

export default CargoItem