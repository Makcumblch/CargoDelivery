import { useContext } from "react";
import Spinner from "../Spinner";
import AddBtn from "../Btns/AddBtn";
import { ModalWindowContext } from "../../contexts/ModalWindowContext";
import { CarsContext } from "../../contexts/CarsContext";
import CarForm from "./CarForm";
import CarItem from "./CarItem";

const CarsComponent = () => {
    const { open, close } = useContext(ModalWindowContext)
    const { isLoading, cars, addCar } = useContext(CarsContext)

    const AddCar = () => {
        open({
            title: `Новый автомобиль`,
            content: <CarForm input={{ id: -1, name: '', loadCapacity: 0, width: 0, height: 0, length: 0, fuelConsumption: 0 }} changeCar={addCar} close={close} />,
            onOk: () => () => { },
            onCancel: () => () => close()
        })
    }

    return (
        <>
            {!isLoading ?
                <>
                    <div className="flex items-center space-x-3 mb-2">
                        <span className="text-white text-lg">Добавить новый автомобиль</span>
                        <AddBtn className="bg-slate-700 hover:bg-slate-600" onClick={AddCar} />
                    </div>
                    <div className="space-y-2 h-[calc(100%_-_45px)] overflow-y-auto">
                        {!cars || !cars.length ? <div className="text-white text-center p-5 mt-6">Нет автомобилей</div> : cars.map((car) => {
                            return <CarItem key={car.id} {...car} />
                        })}
                    </div>
                </>
                : <Spinner />}
        </>
    );
}

export default CarsComponent