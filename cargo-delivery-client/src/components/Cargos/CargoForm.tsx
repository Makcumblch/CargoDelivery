import { FormEvent, useState } from "react";
import { Cargo } from "../../contexts/CargosContext";

interface CargoFormProps {
    input: Cargo,
    changeCargo: (id: number, value: object) => void
    close: () => void
}

const CargoForm = ({ input, changeCargo, close }: CargoFormProps) => {
    const [inputCargo, setInputCargo] = useState<Cargo>(input)

    const onChange = (field: string, value: any): void => {
        setInputCargo((prev: Cargo) => {
            return { ...prev, [field]: value }
        })
    }

    const onSubmit = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        changeCargo(inputCargo.id, inputCargo)
        close()
    }

    return (
        <form id='form' onSubmit={onSubmit} className="flex flex-col space-y-2 w-full mb-4">
            <div>
                <label htmlFor="name" className="block text-sm font-semibold text-slate-300">Название</label>
                <input
                    value={inputCargo.name}
                    type="text"
                    name="name"
                    required
                    placeholder="Введите название"
                    className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                    onChange={(e) => onChange('name', e.target.value)}
                />
            </div>
            <div>
                <label htmlFor="width" className="block text-sm font-semibold text-slate-300">Ширина (м)</label>
                <input
                    value={inputCargo.width.toString()}
                    type="number"
                    name="width"
                    required
                    min={0}
                    placeholder="Введите ширину"
                    className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                    onChange={(e) => onChange('width', parseFloat(e.target.value))}
                />
            </div>
            <div>
                <label htmlFor="height" className="block text-sm font-semibold text-slate-300">Высота (м)</label>
                <input
                    value={inputCargo.height.toString()}
                    type="number"
                    name="height"
                    required
                    min={0}
                    placeholder="Введите высоту"
                    className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                    onChange={(e) => onChange('height', parseFloat(e.target.value))}
                />
            </div>
            <div>
                <label htmlFor="length" className="block text-sm font-semibold text-slate-300">Длина (м)</label>
                <input
                    value={inputCargo.length.toString()}
                    type="number"
                    name="length"
                    required
                    min={0}
                    placeholder="Введите длину"
                    className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                    onChange={(e) => onChange('length', parseFloat(e.target.value))}
                />
            </div>
            <div>
                <label htmlFor="weight" className="block text-sm font-semibold text-slate-300">Вес (кг)</label>
                <input
                    value={inputCargo.weight.toString()}
                    type="number"
                    name="weight"
                    required
                    min={0}
                    placeholder="Введите вес"
                    className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                    onChange={(e) => onChange('weight', parseFloat(e.target.value))}
                />
            </div>
        </form>
    );
}

export default CargoForm