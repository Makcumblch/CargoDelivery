import { useContext } from "react";
import { Cargo, CargosContext } from "../../contexts/CargosContext";
import { ModalWindowContext } from "../../contexts/ModalWindowContext";
import DelBtn from "../Btns/DelBtn";
import EditBtn from "../Btns/EditBtn";
import CargoForm from "./CargoForm";

const CargoItem = ({ id, name, width, height, length, weight }: Cargo) => {
    const { open, close } = useContext(ModalWindowContext)
    const { delCargo, changeCargo } = useContext(CargosContext)

    const deleteCargo = () => {
        open({
            title: `Удалить товар ${name}`,
            content: null,
            onOk: () => () => {
                delCargo(id)
                close()
            },
            onCancel: () => () => close()
        })
    }

    const editCargo = () => {
        open({
            title: `Редактирование: ${name}`,
            content: <CargoForm input={{ id, name, width, height, length, weight }} changeCargo={changeCargo} close={close} />,
            onOk: () => () => { },
            onCancel: () => () => close()
        })
    }

    return (
        <div className="w-full rounded-md text-slate-300 grid items-center p-2 bg-slate-700 grid-cols-[3fr_1fr_1fr_1fr_1fr_1fr]">
            <p className="truncate">{`Название: ${name}`}</p>
            <p>{`Ширина (м): ${width}`}</p>
            <p>{`Высота (м): ${height}`}</p>
            <p>{`Длина (м): ${length}`}</p>
            <p>{`Вес (кг): ${weight}`}</p>
            <div className="flex space-x-3">
                <EditBtn onClick={editCargo} />
                <DelBtn onClick={deleteCargo} />
            </div>
        </div>
    );
}

export default CargoItem