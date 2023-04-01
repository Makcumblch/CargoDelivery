import { useMemo } from "react";
import { Project } from "../contexts/ProjectsContext";
import { DropdownMenu } from "./DropdownMenu";

interface IProjectItemPriops extends Project {
    selected: boolean,
    onDeleteClick: (id: number) => void,
    onChange: (id: number, values: object) => void,
    onSelect: (id: number) => void,
}

export const ProjectItem = ({ id, selected, name, access, onChange, onSelect, onDeleteClick }: IProjectItemPriops) => {

    const renameProject = () => {

    }

    const delProject = () => {

    }

    const changeProjects = () => {

    }

    const itemsMenu = useMemo(() => {
        const items = []
        switch (access) {
            case 'owner':
                items.push(
                    <button onClick={renameProject}>Переименовать</button>,
                    <button onClick={delProject}>Удалит</button>
                )
                break
            default:
                items.push(
                    <button onClick={() => { }}>Покинуть</button>
                )
                break
        }
        return items
    }, [access])

    return (
        <div
            className={`flex items-center text-slate-300 hover:bg-slate-600 p-2 justify-between ${selected && 'bg-slate-600'}`}
            onClick={() => onSelect(id)}
        >
            <div>{name}</div>
            <DropdownMenu items={itemsMenu} />
        </div>
    );
}