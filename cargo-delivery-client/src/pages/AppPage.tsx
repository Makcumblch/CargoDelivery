import SideBar from "../components/SideBar";
import Tabs from "../components/Tabs";
import { Cargos } from "../contexts/CargosContext";
import { Cars } from "../contexts/CarsContext";
import { ModalWindow } from "../contexts/ModalWindowContext";

function AppPage() {
    return (
        <ModalWindow>
            <div className="flex h-screen w-screen flex-row">
                <SideBar />
                <main className="w-full bg-white">
                    <Cargos>
                        <Cars>
                            <Tabs />
                        </Cars>
                    </Cargos>
                </main>
            </div>
        </ModalWindow>
    );
}

export default AppPage;