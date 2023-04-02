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
    changeProject: (id: number | null, values: object) => void,
    addProject: (name: string) => void,
    currentProject: Project | null,
    getProjectById: (id: number) => Project | undefined,
    isLoading: boolean,
}

export const ProjectsContext = createContext<IProjectsContext>({
    projects: [],
    currentProjectId: null,
    setProject: () => { },
    delProject: () => { },
    changeProject: () => { },
    addProject: () => { },
    currentProject: null,
    getProjectById: () => undefined,
    isLoading: true
})

const projectStorage = 'project'

export const Projects = ({ children }: ProjectsProps) => {
    const { token } = useContext(AuthContext)
    const { request } = useHttp()
    const [projects, setProjects] = useState<Project[]>([])
    const [currentProjectId, setCurrentProjectId] = useState<number | null>(null)
    const [currentProject, setCurrentProject] = useState<Project | null>(null)
    const [isLoading, setIsLoading] = useState<boolean>(true)

    useEffect(() => {
        const getProjects = async () => {
            try {
                setIsLoading(true)
                let projectsData = await request('api/projects', 'GET', null, {
                    Authorization: `Bearer ${token}`
                })
                setIsLoading(false)
                projectsData = projectsData.data as Project[]
                if (!projectsData) {
                    setProjects([])
                    setCurrentProjectId(null)
                    setCurrentProject(null)
                    return
                }
                setProjects(projectsData)
                setCurrentProjectId(projectsData[0]?.id ?? null)
                setCurrentProject(projectsData[0] ?? null)
                const strId = localStorage.getItem(projectStorage)
                if (!strId) return
                const id = parseInt(strId)
                if (isNaN(id)) return
                const project = projectsData.find((pr: Project) => pr.id === id)
                if (!project) return
                setCurrentProjectId(project.id)
                setCurrentProject(project)
            } catch (e) {
                setIsLoading(false)
            }
        }
        getProjects()
    }, [token, request])

    const setProject = (id: number): void => {
        const project = projects.find((pr) => pr.id === id)
        if (!project) return
        setCurrentProjectId(project.id)
        setCurrentProject(project)
        localStorage.setItem(projectStorage, project.id.toString())
    }

    const delProject = async (id: number) => {
        try {
            const data = await request(`api/projects/${id}`, 'DELETE', null, {
                Authorization: `Bearer ${token}`
            })
            if (data.status !== 'ok') return
            let index = 0
            setProjects(prev => {
                const newProjects: Project[] = []
                for (let i = 0; i < prev.length; ++i) {
                    if (prev[i].id !== id) {
                        newProjects.push({ ...prev[i] })
                    } else {
                        index = i - 1
                    }
                }
                const project = newProjects[index < 0 ? 0 : index]
                setCurrentProjectId(project?.id ?? null)
                setCurrentProject(project ?? null)
                return newProjects
            })
        } catch (e) { }
    }

    const changeProject = async (id: number | null, values: object) => {
        if (id === null) return
        const projectIndex = projects.findIndex((pr) => pr.id === id)
        if (projectIndex === -1) return
        try {
            const data = await request(`api/projects/${id}`, 'PUT', values, {
                Authorization: `Bearer ${token}`
            })
            if (data.status !== 'ok') return
            setProjects(prev => {
                const newProjects = [...prev]
                newProjects[projectIndex] = { ...newProjects[projectIndex], ...values }
                return newProjects
            })
        } catch (e) { }
    }

    const addProject = async (name: string) => {
        try {
            const data = await request(`api/projects/`, 'POST', { name }, {
                Authorization: `Bearer ${token}`
            })
            const project = {
                id: data.id,
                name: name,
                access: 'owner'
            }
            setProjects(prev => {
                const newProjects = [...prev]
                newProjects.push(project)
                return newProjects
            })
            setCurrentProjectId(data.id)
            setCurrentProject(project)
        } catch (e) { }
    }

    const getProjectById = (id: number) => {
        const project = projects.find((pr: Project) => pr.id === id)
        return project
    }

    return (
        <ProjectsContext.Provider value={{
            projects,
            currentProjectId,
            currentProject,
            setProject,
            delProject,
            changeProject,
            addProject,
            getProjectById,
            isLoading
        }}>
            {children}
        </ProjectsContext.Provider>
    )
}