import { useState, useContext, useEffect } from "react";
import SelectAddress from "../../../SelectAddress";
import { RoutesContext } from "../../../../contexts/RoutesContext";
import Spinner from "../../../Spinner";

const RouteMenu = () => {
    const { depo, updateDepo, createRoute, isLoadingCreate, error } = useContext(RoutesContext)

    const [count, setCount] = useState<number>(100)
    const [TMax, setTMax] = useState<number>(1000)
    const [TMin, setTmin] = useState<number>(1)

    useEffect(() => {
        setCount(Number(localStorage.getItem('count') ?? 100))
        setTMax(Number(localStorage.getItem('TMax') ?? 1000))
        setTmin(Number(localStorage.getItem('TMin') ?? 1))
    }, [])

    const onChange = (key: string, value: string, callback: (value: number) => void) => {
        localStorage.setItem(key, value)
        callback(Number(value))
    }

    return (
        <div className="bg-slate-800 px-1 relative h-full">
            <div className="flex items-center">
                <label htmlFor="name" className="w-20 block text-sm font-semibold text-slate-300 mr-1">Депо</label>
                <SelectAddress
                    address={depo.address}
                    onChange={(address, x, y) => {
                        updateDepo(depo.id, { address: address, coordX: x, coordY: y })
                    }}
                    disabled={isLoadingCreate}
                />
            </div>
            <div className="flex">
                <div className="flex items-center">
                    <label htmlFor="name" className="w-28 block text-sm font-semibold text-slate-300 mr-1">Мин. t</label>
                    <input
                        disabled={isLoadingCreate}
                        type="number"
                        min={0}
                        step={'any'}
                        value={TMin}
                        onChange={(e) => onChange('TMin', e.target.value, setTmin)}
                        className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                    />
                </div>
                <div className="flex items-center ml-2">
                    <label htmlFor="name" className="w-20 block text-sm font-semibold text-slate-300 mr-1">Макс. t</label>
                    <input
                        disabled={isLoadingCreate}
                        type="number"
                        min={0}
                        step={'any'}
                        value={TMax}
                        onChange={(e) => onChange('TMax', e.target.value, setTMax)}
                        className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                    />
                </div>
            </div>
            <div className="flex items-center">
                <label htmlFor="name" className="w-20 block text-sm font-semibold text-slate-300 mr-1">Эволюц.</label>
                <input
                    disabled={isLoadingCreate}
                    type="number" min={0}
                    step={'any'}
                    value={count}
                    onChange={(e) => onChange('count', e.target.value, setCount)}
                    className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                />
            </div>
            <button
                className="mt-2 w-full px-4 py-1.5 tracking-wide text-white transition-colors duration-200 transform bg-cyan-700 rounded-md hover:bg-cyan-600 focus:outline-none focus:bg-cyan-600"
                onClick={() => {
                    createRoute(count, TMax, TMin)
                }}
                disabled={isLoadingCreate}
            >
                Построить решение
            </button>
            {error ? error.status === 400 ? <p className="text-red-700 mb-0 text-center">{error.message}</p> : <p className="text-red-700 mb-0 text-center">Не удалось получить решение</p> : null}
            {isLoadingCreate &&
                <div className="absolute z-[1005] w-full h-full left-0 top-0 bg-slate-800 opacity-70">
                    <Spinner />
                </div>}
        </div>
    );
}

export default RouteMenu