import React, { createContext, useContext, useEffect, useState } from 'react'
import { useHttp } from '../hooks/useHttp'
import { AuthContext } from './AuthContext'

export interface Project {
    id: number,
    name: string,
    access: string
}

interface ProjectsProps {
    children: React.ReactNode
}

interface IProjectsContext {
    projects: Project[],
    currentProjectId: number | null,
    setProject: (id: number) => void,
    delProject: (id: number) => void,
    changeProject: (id: number, values: object) => void,
    addProject: (name: string) => void,
}

export const ProjectsContext = createContext<IProjectsContext>({
    projects: [],
    currentProjectId: null,
    setProject: () => { },
    delProject: () => { },
    changeProject: () => { },
    addProject: () => { },
})

const projectStorage = 'project'

export const Projects = ({ children }: ProjectsProps) => {
    const { token } = useContext(AuthContext)
    const { request } = useHttp()
    const [projects, setProjects] = useState<Project[]>([])
    const [currentProjectId, setCurrentProjectId] = useState<number | null>(null)

    useEffect(() => {
        const getProjects = async () => {
            try {
                let projectsData = await request('api/projects', 'GET', null, {
                    Authorization: `Bearer ${token}`
                })
                projectsData = projectsData.data as Project[]
                if (!projectsData) return
                setProjects(projectsData)
                setCurrentProjectId(projectsData[0]?.id ?? null)
                const strId = localStorage.getItem(projectStorage)
                if (!strId) return
                const id = parseInt(strId)
                if (isNaN(id)) return
                const project = projectsData.find((pr: Project) => pr.id === id)
                if (!project) return
                setCurrentProjectId(project.id)
            } catch (e) { }
        }
        getProjects()
    }, [token, request])

    const setProject = (id: number): void => {
        const project = projects.find((pr) => pr.id === id)
        if (!project) return
        setCurrentProjectId(project.id)
        localStorage.setItem(projectStorage, project.id.toString())
    }

    const delProject = async (id: number) => {
        try {
            const data = await request(`api/projects/${id}`, 'DELETE', null, {
                Authorization: `Bearer ${token}`
            })
            if (data.status !== 'ok') return
            let index = 0
            const newProjects = projects.filter((pr, i) => {
                if (pr.id !== id) return true
                index = i - 1
                return false
            })
            setProjects(newProjects)
            const project = newProjects[index < 0 ? 0 : index]
            setCurrentProjectId(project?.id ?? null)
        } catch (e) { }
    }

    const changeProject = (id: number, values: object) => {
        const projectIndex = projects.findIndex((pr) => pr.id === id)
        if (projectIndex === -1) return
        const newProjects = [...projects]
        newProjects[projectIndex] = { ...newProjects[projectIndex], ...values }
        setProjects(newProjects)
    }

    const addProject = async (name: string) => {
        try {
            const data = await request(`api/projects/`, 'POST', { name }, {
                Authorization: `Bearer ${token}`
            })
            const newProjects = [...projects]
            newProjects.push({
                id: data.id,
                name: name,
                access: 'owner'
            })
            setProjects(newProjects)
            setCurrentProjectId(data.id)
        } catch (e) { }
    }

    return (
        <ProjectsContext.Provider value={{
            projects,
            currentProjectId,
            setProject,
            delProject,
            changeProject,
            addProject,
        }}>
            {children}
        </ProjectsContext.Provider>
    )
}