import { Car } from "../../contexts/CarsContext";
import { FormEvent, useState } from "react";

interface CarFormProps {
    input: Car,
    changeCar: (id: number, value: object) => void
    close: () => void
}

const CarForm = ({ input, changeCar, close }: CarFormProps) => {
    const [inputCar, setInputCar] = useState<Car>(input)

    const onChange = (field: string, value: any): void => {
        setInputCar((prev: Car) => {
            return { ...prev, [field]: value }
        })
    }

    const onSubmit = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        changeCar(inputCar.id, inputCar)
        close()
    }

    return (
        <form id='form' onSubmit={onSubmit} className="flex flex-col space-y-2 w-full mb-4">
            <div>
                <label htmlFor="name" className="block text-sm font-semibold text-slate-300">Название</label>
                <input
                    value={inputCar.name}
                    type="text"
                    name="name"
                    required
                    placeholder="Введите название"
                    className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                    onChange={(e) => onChange('name', e.target.value)}
                />
            </div>
            <div>
                <label htmlFor="loadCapacity" className="block text-sm font-semibold text-slate-300">Грузоподъемность (кг)</label>
                <input
                    value={inputCar.loadCapacity.toString()}
                    type="text"
                    name="loadCapacity"
                    required
                    placeholder="Введите грузоподъемность"
                    className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                    onChange={(e) => onChange('loadCapacity', parseFloat(e.target.value))}
                />
            </div>
            <div>
                <label htmlFor="width" className="block text-sm font-semibold text-slate-300">Ширина (м)</label>
                <input
                    value={inputCar.width.toString()}
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
                    value={inputCar.height.toString()}
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
                    value={inputCar.length.toString()}
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
                <label htmlFor="fuelConsumption" className="block text-sm font-semibold text-slate-300">Расход топлива (л/100км)</label>
                <input
                    value={inputCar.fuelConsumption.toString()}
                    type="number"
                    name="fuelConsumption"
                    required
                    min={0}
                    placeholder="Введите расход топлива"
                    className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                    onChange={(e) => onChange('fuelConsumption', parseFloat(e.target.value))}
                />
            </div>
        </form>
    );
}

export default CarForm