import { useContext } from "react";
import Spinner from "../Spinner";
import AddBtn from "../Btns/AddBtn";
import { ModalWindowContext } from "../../contexts/ModalWindowContext";
import { OrdersContext } from "../../contexts/OrdersContext";
import OrderItem from "./OrderItem";
import OrderForm from "./OrderForm";
import { CargosContext } from "../../contexts/CargosContext";

const OrdersComponent = () => {
    const { open, close } = useContext(ModalWindowContext)
    const { cargos } = useContext(CargosContext)
    const { isLoading, orders, addOrder } = useContext(OrdersContext)

    const AddOrder = () => {
        open({
            title: `Новый заказ`,
            content: <OrderForm input={{ id: -1, idCargo: -1, count: 1 }} changeOrder={addOrder} close={close} cargos={cargos}/>,
            onOk: () => () => { },
            onCancel: () => () => close()
        })
    }

    return (
        <>
            {!isLoading ?
                <>
                    <div className="flex items-center space-x-3 mb-2">
                        <span className="text-white text-lg">Добавить новый заказ</span>
                        <AddBtn className="bg-slate-700 hover:bg-slate-600" onClick={AddOrder} />
                    </div>
                    <div className="space-y-2 h-[calc(100%_-_45px)] overflow-y-auto">
                        {!orders || !orders.length ? <div className="text-white text-center p-2">Нет заказов</div> : orders.map((order) => {
                            return <OrderItem key={order.id} {...order} />
                        })}
                    </div>
                </>
                : <Spinner />}
        </>
    );
}

export default OrdersComponent