import SideBar from "../components/SideBar";
import { ModalWindow } from "../contexts/ModalWindowContext";

function AppPage() {
    return (
        <ModalWindow>
            <div className="flex h-screen w-screen flex-row">
                <SideBar />
                <main className="w-full bg-white">

                </main>
            </div>
        </ModalWindow>
    );
}

export default AppPage;