import { useMemo, useState } from "react";
import CargosComponent from "./Cargos/CargosComponent";
import CarsComponent from "./Cars/CarsComponent";
import ClientsComponent from "./Clients/ClientsComponent";
import RoutesComponent from "./Routes/RoutesComponent";

const Tabs = () => {
    const [openTab, setOpenTab] = useState("Товары");

    const tabs = useMemo(() => {
        return [
            { name: "Товары", content: <CargosComponent /> },
            { name: "Клиенты", content: <ClientsComponent /> },
            { name: "Автомобили", content: <CarsComponent /> },
            { name: "Маршруты", content: <RoutesComponent /> },
        ]
    }, [])

    const content = useMemo(() => {
        for (let i = 0; i < tabs.length; ++i) {
            if (tabs[i].name === openTab) {
                return tabs[i].content
            }
        }
        return null
    }, [openTab, tabs])

    return (
        <div className="container flex flex-col h-full shadow-md">
            <div className="h-12 min-h-[48] pt-2 bg-slate-800">
                <ul className="flex space-x-2 h-full overflow-x-auto">
                    {tabs.map((tab) => (
                        <li key={tab.name}>
                            <button
                                onClick={() => setOpenTab(tab.name)}
                                className={`${tab.name === openTab ? 'bg-slate-500' : 'bg-slate-700'} h-full inline-block px-4 py-2 items-center text-slate-300 hover:bg-slate-500 p-2 justify-between rounded-t-xl`}
                            >
                                {tab.name}
                            </button>
                        </li>
                    ))}
                </ul>
            </div>
            <div className="flex-auto h-[calc(100%_-_50px)] w-full bg-slate-500 p-2">
                {content}
            </div>
        </div>
    );
}

export default Tabs