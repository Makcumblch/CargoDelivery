import { useCallback, useContext, useMemo } from "react";
import { Project, ProjectsContext } from "../../contexts/ProjectsContext";
import { DropdownMenu } from "../DropdownMenu";
import { ModalWindowContext } from "../../contexts/ModalWindowContext";
import InputProjectNameProps from "./InputProjectName";

interface IProjectItemPriops extends Project {
    selected: boolean,
}

export const ProjectItem = ({ id, selected, name, access }: IProjectItemPriops) => {
    const { open, close } = useContext(ModalWindowContext)
    const { setProject, delProject, getProjectById } = useContext(ProjectsContext)

    const del = useCallback(() => {
        open({
            title: `Удалить проект ${getProjectById(id)?.name}?`,
            content: null,
            onOk: () => () => {
                delProject(id)
                close()
            },
            onCancel: () => () => close()
        })
    }, [delProject, getProjectById, id, open])

    const itemsMenu = useMemo(() => {
        const items = []
        switch (access) {
            case 'owner':
                items.push(
                    <InputProjectNameProps id={id} btnContent={'Переименовать'} mode="change" />,
                    <button className="w-full" onClick={del}>Удалить</button>
                )
                break
            default:
                items.push(
                    <button onClick={() => { }}>Покинуть</button>
                )
                break
        }
        return items
    }, [access, id, del])

    return (
        <div
            className={`flex items-center w-52 text-slate-300 hover:bg-slate-600 p-2 justify-between ${selected && 'bg-slate-600'}`}
            onClick={() => setProject(id)}
        >
            <div className="flex-1 truncate">{name}</div>
            <div className="w-9">
                <DropdownMenu items={itemsMenu} />
            </div>
        </div>
    );
}