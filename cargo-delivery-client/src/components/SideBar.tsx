import { useContext } from "react";
import { ProjectsContext } from "../contexts/ProjectsContext";
import AddProject from "./AddProject";
import { ProjectItem } from "./ProjectItem";


function SideBar() {
    const { currentProjectId, projects, setProject, delProject, changeProject } = useContext(ProjectsContext)

    return (
        <aside className="flex w-72 h-full flex-col bg-pink-500">
            <header className="flex h-14 items-center justify-center bg-slate-800">
                <h1 className="text-xl text-white">Cargo Delivery</h1>
            </header>
            <main className="h-full bg-slate-700">
                <div className="flex text-slate-300 items-center justify-between p-2 shadow-xl">
                    <span>Проекты</span>
                    <AddProject />
                </div>
                {!projects || !projects.length ?
                    <div className="text-slate-300 text-center p-5">Нет проектов</div>
                    :
                    projects?.map((project) => {
                        return <ProjectItem selected={project.id === currentProjectId} key={project.id} {...project} onChange={changeProject} onSelect={setProject} onDeleteClick={delProject} />
                    })}
            </main>
            <footer className="bg-slate-800 h-14">

            </footer>
        </aside>
    );
}

export default SideBar;