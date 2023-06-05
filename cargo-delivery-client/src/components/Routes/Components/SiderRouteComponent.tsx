import { useContext, useMemo } from "react";
import { RoutesContext } from "../../../contexts/RoutesContext";
import { RouteContext } from "../../../contexts/RouteContext";
import Spinner from "../../Spinner";
import RouteCarItem from "./SiderRouteComponents/RouteCarItem";

const RouteComponent = () => {
    const { setInDetail, setIsPacking, isPacking, routes, selectRouteIndex } = useContext(RoutesContext)
    const { isLoading, routeSolution } = useContext(RouteContext)

    const date = useMemo(() => {
        const newDate = new Date(routes[selectRouteIndex].date)
        return newDate.toLocaleString()
    }, [routes, selectRouteIndex])

    return (
        <div className="grid grid-rows-[32px_32px_105px_1fr] gap-1 h-[calc(100vh_-_64px)]">
            <div className="bg-slate-700 p-1 text-white">
                <button
                    className='bg-slate-500 hover:bg-slate-400 w-full h-full'
                    onClick={() => {
                        setIsPacking(false)
                        setInDetail(false)
                    }}
                >{'< Назад'}</button>
            </div>
            <div className="bg-slate-700 p-1 grid grid-cols-[1fr_1fr] gap-1 text-white">
                <button
                    className={`${isPacking ? 'bg-slate-500' : 'bg-slate-400'} hover:bg-slate-400 w-full h-full`}
                    onClick={() => setIsPacking(false)}
                >Маршрут</button>
                <button
                    className={`${isPacking ? 'bg-slate-400' : 'bg-slate-500'} hover:bg-slate-400 w-full h-full`}
                    onClick={() => setIsPacking(true)}
                >Упаковка</button>
            </div>
            <div className="bg-slate-800 p-1 text-white">
                <div className='grid grid-flow-col grid-cols-1'>
                    <p>Дата</p>
                    <p>{date}</p>
                </div>
                <div className='grid grid-flow-col grid-cols-1'>
                    <p>Дистанция (км)</p>
                    <p>{routeSolution?.distance}</p>
                </div>
                <div className='grid grid-flow-col grid-cols-1'>
                    <p>Топливо (л)</p>
                    <p>{routeSolution?.fuel}</p>
                </div>
                <div className='grid grid-flow-col grid-cols-1'>
                    <p>Качество упаковки</p>
                    <p>{routeSolution?.packingCost}</p>
                </div>
            </div>
            <div className="bg-slate-700 overflow-y-scroll p-2">
                {routeSolution && routeSolution.carsRouteSolution.map((element, index) => {
                    return <RouteCarItem key={index} item={element} index={index} />
                })}
                {isLoading &&
                    <div className="z-[1005] w-full h-full left-0 top-0 bg-slate-800 opacity-70">
                        <Spinner />
                    </div>
                }
            </div>
        </div>
    );
}

export default RouteComponent