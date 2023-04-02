import { useContext, useMemo, useState } from "react";
import { ProjectsContext } from "../contexts/ProjectsContext";
import { DropdownMenu } from "./DropdownMenu";

interface InputProjectNameProps {
    btnContent: React.ReactNode | null,
    mode: 'add' | 'change',
    id: number | null,
}

function InputProjectName({ btnContent, mode, id }: InputProjectNameProps) {
    const { addProject, changeProject } = useContext(ProjectsContext)
    const [projectName, setProjectName] = useState<string>('')

    const props = useMemo(() => {
        switch (mode) {
            case 'add':
                return { title: 'Добавить', func: (name: string) => addProject(name) }
            case 'change':
                return { title: 'Переименовать', func: (name: string) => changeProject(id, { name: name }) }
            default: return null
        }
    }, [mode, id, addProject, changeProject])

    return (
        <DropdownMenu
            btnContent={btnContent}
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
                    onClick={() => props?.func(projectName)}
                >
                    {props?.title}
                </button>
            </div>]}
        />
    );
}

export default InputProjectName;