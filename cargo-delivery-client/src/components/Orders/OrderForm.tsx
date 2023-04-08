import { FormEvent, useCallback, useEffect, useMemo, useState } from "react";
import { Order } from "../../contexts/OrdersContext";
import { Cargo } from "../../contexts/CargosContext";
import Select from "../Select";

interface OrderFormProps {
    input: Order,
    changeOrder: (id: number, value: object) => void
    close: () => void
    cargos: Cargo[]
}

const OrderForm = ({ input, changeOrder, close, cargos }: OrderFormProps) => {
    const [inputOrder, setInputOrder] = useState<Order>({id: -1, idCargo: -1, count: 1})


    useEffect(() => {
        setInputOrder(input)
    }, [input])

    const onChange = useCallback((field: string, value: any): void => {
        setInputOrder((prev: Order) => {
            return { ...prev, [field]: value }
        })
    }, [])

    const onSubmit = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        changeOrder(inputOrder.id, inputOrder)
        close()
    }

    const values = useMemo(() => {
        return cargos.map(element => ({ id: element.id, value: element.name }))
    }, [cargos])

    return (
        <form id='form' onSubmit={onSubmit} className="flex flex-col space-y-2 w-full mb-4">
            <div>
                <label htmlFor="name" className="block text-sm font-semibold text-slate-300">Название</label>
                <Select value={inputOrder.idCargo} values={values} onChange={(id) => onChange('idCargo', id)} />
            </div>
            <div>
                <label htmlFor="loadCapacity" className="block text-sm font-semibold text-slate-300">Количество (шт)</label>
                <input
                    value={inputOrder.count.toString()}
                    type="number"
                    min={1}
                    step={1}
                    name="count"
                    required
                    placeholder="Введите количество"
                    className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                    onChange={(e) => onChange('count', parseInt(e.target.value))}
                />
            </div>
        </form>
    );
}

export default OrderForm