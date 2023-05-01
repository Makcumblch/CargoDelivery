import { useContext, useState } from "react";
import { ModalWindowContext } from "../../contexts/ModalWindowContext";
import DelBtn from "../Btns/DelBtn";
import EditBtn from "../Btns/EditBtn";
import { Client, ClientsContext } from "../../contexts/ClientsContext";
import ClientForm from "./ClientForm";
import { Orders } from "../../contexts/OrdersContext";
import OrdersComponent from "../Orders/OrdersComponent";

const ClientItem = ({ id, name, address, coordX, coordY }: Client) => {
    const { open, close } = useContext(ModalWindowContext)
    const { delClient, changeClient } = useContext(ClientsContext)
    const [isOpenOrders, setIsOpenOrders] = useState<boolean>(false)

    const deleteClient = (e: any) => {
        e.stopPropagation()
        open({
            title: `Удалить клиента ${name}`,
            content: null,
            onOk: () => () => {
                delClient(id)
                close()
            },
            onCancel: () => () => close()
        })
    }

    const editClient = (e: any) => {
        e.stopPropagation()
        open({
            title: `Редактирование: ${name}`,
            content: <ClientForm input={{ id, name, address, coordX, coordY }} changeClient={changeClient} close={close} />,
            onOk: () => () => { },
            onCancel: () => () => close()
        })
    }

    return (
        <div>
            <div
                className={`w-full rounded-md ${isOpenOrders ? 'rounded-b-none' : ''} text-slate-300 grid items-center p-2 bg-slate-700 grid-cols-[24px_4fr_6fr_1fr] cursor-pointer`}
                onClick={() => setIsOpenOrders(prev => !prev)}
            >
                <span className="text-lg text-center pb-0.5">{isOpenOrders ? '∧' : '∨'}</span>
                <p className="truncate">{`Название: ${name}`}</p>
                <p className="truncate">{`Адрес: ${address}`}</p>
                <div className="flex space-x-3">
                    <EditBtn onClick={editClient} />
                    <DelBtn onClick={deleteClient} />
                </div>
            </div>
            <div className={`${!isOpenOrders ? 'hidden': ''} p-2 bg-slate-600 rounded-b-md`}>
                <Orders clientId={id}>
                    <OrdersComponent />
                </Orders>
            </div>
        </div>
    );
}

export default ClientItem