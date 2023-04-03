import { useContext } from "react";
import { CargosContext } from "../../contexts/CargosContext";
import CargoItem from "./CargoItem";
import Spinner from "../Spinner";
import AddBtn from "../Btns/AddBtn";
import { ModalWindowContext } from "../../contexts/ModalWindowContext";
import CargoForm from "./CargoForm";

const CargosComponent = () => {
    const { open, close } = useContext(ModalWindowContext)
    const { isLoading, cargos, addCargo } = useContext(CargosContext)

    const AddCargo = () => {
        open({
            title: `Новый товар`,
            content: <CargoForm input={{ id: -1, name: '', width: 0, height: 0, length: 0, weight: 0 }} changeCargo={addCargo} close={close} />,
            onOk: () => () => { },
            onCancel: () => () => close()
        })
    }

    return (
        <>
            {!isLoading ?
                <>
                    <div className="flex items-center space-x-3 mb-2">
                        <span className="text-white text-lg">Добавить новый товар</span>
                        <AddBtn className="bg-slate-700 hover:bg-slate-600" onClick={AddCargo} />
                    </div>
                    <div className="space-y-2">
                        {!cargos || !cargos.length ? <div className="text-white text-center p-5 mt-6">Нет товаров</div> : cargos.map((cargo) => {
                            return <CargoItem key={cargo.id} {...cargo} />
                        })}
                    </div>
                </>
                : <Spinner />}
        </>
    );
}

export default CargosComponent