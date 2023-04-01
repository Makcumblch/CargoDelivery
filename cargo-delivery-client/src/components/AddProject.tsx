import { useContext, useState } from "react";
import { ProjectsContext } from "../contexts/ProjectsContext";
import { DropdownMenu } from "./DropdownMenu";


function AddProject() {
    const { addProject } = useContext(ProjectsContext)
    const [projectName, setProjectName] = useState<string>('')
    return (
        <DropdownMenu
            btnContent='+'
            items={[<div className="flex">
                <input
                    value={projectName}
                    type='text'
                    placeholder="Введите название проекта"
                    className="block px-4 py-2 w-60 text-white bg-slate-700  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40 mr-2"
                    onChange={(e) => setProjectName(e.target.value)}
                />
                <button
                    className="px-4 py-2 tracking-wide text-white transition-colors duration-200 transform bg-cyan-700 rounded-md hover:bg-cyan-600 focus:outline-none focus:bg-cyan-600"
                    onClick={() => addProject(projectName)}
                >
                    Добавить
                </button>
            </div>]}
        />
    );
}

export default AddProject;