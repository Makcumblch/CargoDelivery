import { useContext } from "react";
import Spinner from "../Spinner";
import AddBtn from "../Btns/AddBtn";
import { ModalWindowContext } from "../../contexts/ModalWindowContext";
import { ClientsContext } from "../../contexts/ClientsContext";
import ClientForm from "./ClientForm";
import ClientItem from "./ClientItem";

const ClientsComponent = () => {
    const { open, close } = useContext(ModalWindowContext)
    const { isLoading, clients, addClient } = useContext(ClientsContext)

    const AddClient = () => {
        open({
            title: `Новый клиент`,
            content: <ClientForm input={{ id: -1, name: '', address: '' }} changeClient={addClient} close={close} />,
            onOk: () => () => { },
            onCancel: () => () => close()
        })
    }

    return (
        <>
            {!isLoading ?
                <>
                    <div className="flex items-center space-x-3 mb-2">
                        <span className="text-white text-lg">Добавить нового клиента</span>
                        <AddBtn className="bg-slate-700 hover:bg-slate-600" onClick={AddClient} />
                    </div>
                    <div className="space-y-2 h-[calc(100%_-_45px)] overflow-y-auto">
                        {!clients || !clients.length ? <div className="text-white text-center p-5 mt-6">Нет клиентов</div> : clients.map((client) => {
                            return <ClientItem key={client.id} {...client} />
                        })}
                    </div>
                </>
                : <Spinner />}
        </>
    );
}

export default ClientsComponent