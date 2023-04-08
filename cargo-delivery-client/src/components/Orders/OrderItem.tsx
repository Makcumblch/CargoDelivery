import { useContext, useMemo } from "react";
import { ModalWindowContext } from "../../contexts/ModalWindowContext";
import DelBtn from "../Btns/DelBtn";
import EditBtn from "../Btns/EditBtn";
import { Order, OrdersContext } from "../../contexts/OrdersContext";
import { CargosContext } from "../../contexts/CargosContext";
import OrderForm from "./OrderForm";

const OrderItem = ({ id, idCargo, count }: Order) => {
    const { open, close } = useContext(ModalWindowContext)
    const { delOrder, changeOrder } = useContext(OrdersContext)
    const { getCargoById } = useContext(CargosContext)
    const { cargos } = useContext(CargosContext)

    const cargo = useMemo(() => {
        return getCargoById(idCargo)
    }, [idCargo, getCargoById])

    const deleteOrder = () => {
        open({
            title: `Удалить заказ ${cargo?.name} - ${count}шт.`,
            content: null,
            onOk: () => () => {
                delOrder(id)
                close()
            },
            onCancel: () => () => close()
        })
    }

    const editOrder = () => {
        open({
            title: `Редактирование: ${cargo?.name} - ${count}шт.`,
            content: <OrderForm input={{ id, idCargo, count }} changeOrder={changeOrder} close={close} cargos={cargos} />,
            onOk: () => () => { },
            onCancel: () => () => close()
        })
    }

    return (
        <div className="w-full rounded-md text-slate-300 grid items-center p-2 bg-slate-700 grid-cols-[4fr_6fr_1fr]">
            <p className="truncate">{`Товар: ${cargo?.name}`}</p>
            <p>{`Количество (шт): ${count}`}</p>
            <div className="flex space-x-3">
                <EditBtn onClick={editOrder} />
                <DelBtn onClick={deleteOrder} />
            </div>
        </div>
    );
}

export default OrderItem