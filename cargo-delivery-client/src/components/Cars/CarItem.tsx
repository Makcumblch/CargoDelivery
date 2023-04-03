import { useContext } from "react";
import { ModalWindowContext } from "../../contexts/ModalWindowContext";
import DelBtn from "../Btns/DelBtn";
import EditBtn from "../Btns/EditBtn";
import { Car, CarsContext } from "../../contexts/CarsContext";
import CarForm from "./CarForm";

const CarItem = ({ id, name, loadCapacity, width, height, length, fuelConsumption }: Car) => {
    const { open, close } = useContext(ModalWindowContext)
    const { delCar, changeCar } = useContext(CarsContext)

    const deleteCar = () => {
        open({
            title: `Удалить автомобиль ${name}`,
            content: null,
            onOk: () => () => {
                delCar(id)
                close()
            },
            onCancel: () => () => close()
        })
    }

    const editCar = () => {
        open({
            title: `Редактирование: ${name}`,
            content: <CarForm input={{ id, name, loadCapacity, width, height, length, fuelConsumption }} changeCar={changeCar} close={close} />,
            onOk: () => () => { },
            onCancel: () => () => close()
        })
    }

    return (
        <div className="w-full rounded-md text-slate-300 grid items-center p-2 bg-slate-700 grid-cols-[2fr_2fr_1fr_1fr_1fr_2fr_1fr]">
            <p className="truncate">{`Название: ${name}`}</p>
            <p>{`Грузоподъемность (кг): ${loadCapacity}`}</p>
            <p>{`Ширина (м): ${width}`}</p>
            <p>{`Высота (м): ${height}`}</p>
            <p>{`Длина (м): ${length}`}</p>
            <p>{`Расход топлива (л/100км): ${fuelConsumption}`}</p>
            <div className="flex space-x-3">
                <EditBtn onClick={editCar} />
                <DelBtn onClick={deleteCar} />
            </div>
        </div>
    );
}

export default CarItem