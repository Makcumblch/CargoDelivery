import { useContext } from "react";
import { ProjectsContext } from "../contexts/ProjectsContext";
import { ProjectItem } from "./Projects/ProjectItem";
import InputProjectName from "./Projects/InputProjectName";
import { AuthContext } from "../contexts/AuthContext";
import Spinner from "./Spinner";


function SideBar() {
    const { isLoading, currentProjectId, projects } = useContext(ProjectsContext)
    const { logout } = useContext(AuthContext)

    return (
        <aside className="flex w-52 h-full flex-col">
            <header className="flex h-12 min-h-[48px] items-center justify-center bg-slate-800">
                <h1 className="text-xl text-white">Cargo Delivery</h1>
            </header>
            <main className="h-[calc(100%_-_96px)] w-full bg-slate-700">
                <div className="flex h-[52px] text-slate-300 items-center justify-between p-2 shadow-xl">
                    <span>Проекты</span>
                    <InputProjectName id={currentProjectId} btnContent={<p className="mx-3 my-1.5">+</p>} mode='add' />
                </div>
                {!isLoading ? <div className="h-[calc(100%_-_52px)] w-screen overflow-y-auto overflow-x-hidden">
                    {!projects || !projects.length ?
                        <div className="text-slate-300 w-52 text-center p-5">Нет проектов</div>
                        :
                        projects?.map((project) => {
                            return <ProjectItem selected={project.id === currentProjectId} key={project.id} {...project} />
                        })}
                </div> : <div className="w-52 h-full"><Spinner /></div>}
            </main>
            <footer className="bg-slate-800 h-12 p-3 min-h-[48px] items-center">
                <button className="flex justify-center w-full h-full text-slate-300 hover:text-white stroke-slate-300 hover:stroke-white" onClick={() => { logout() }}>
                    <p className="mr-4 text-lg">Выход</p>
                    <svg width="30px" height="30px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path d="M10 12H18M18 12L15.5 9.77778M18 12L15.5 14.2222M18 7.11111V5C18 4.44772 17.5523 4 17 4H7C6.44772 4 6 4.44772 6 5V19C6 19.5523 6.44772 20 7 20H17C17.5523 20 18 19.5523 18 19V16.8889" strokeLinecap="round" strokeLinejoin="round" />
                    </svg>
                </button>
            </footer>
        </aside>
    );
}

export default SideBar;