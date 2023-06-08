import { useContext, useRef, useState, useMemo } from 'react'
import { Canvas } from '@react-three/fiber'
import {  RouteContext } from '../../../contexts/RouteContext'
import { getColorByName } from './MapComponent'
import { OrbitControls } from '@react-three/drei'

interface IBoxProps {
    x: number
    y: number
    z: number
    w: number
    l: number
    h: number
    color: string
    wireframe?: boolean
}

interface IDoorProps{
    x: number
    y: number
    z: number
    w: number
    h: number
    color: string
    wireframe?: boolean
    angle: number
}
const Door = ({ x, y, z, w,  h, color, wireframe, angle}: IDoorProps) => {

    const ref = useRef<THREE.Mesh>(null!)
    const [hovered, hover] = useState(false)
    // useFrame((state, delta) => (ref.current.rotation.x += delta))
    return (
        <mesh 
            //rotation={[-Math.PI / 2, 0, 0]}>
            
            rotation={[0,angle,0]}
            position={[x, y, z]}
            ref={ref}
            onPointerOver={() => hover(true)}
            onPointerOut={() => hover(false)}>
            <planeGeometry args={[w/2, h]} /> 
            <meshStandardMaterial wireframe = {wireframe} opacity={hovered ? 1 : 0.5} color={color} />


        </mesh>
    )

}
const Box = ({ x, y, z, w, l, h, color, wireframe}: IBoxProps) => { 
    const ref = useRef<THREE.Mesh>(null!)
    const [hovered, hover] = useState(false)
    // useFrame((state, delta) => (ref.current.rotation.x += delta))
    return (
        <mesh
            position={[x, y, z]}
            ref={ref}
            onPointerOver={() => hover(true)}
            onPointerOut={() => hover(false)}>
            <boxGeometry args={[w, h, l]} />
            <meshStandardMaterial wireframe = {wireframe} opacity={hovered ? 1 : 0.5} color={color} />
            
           
            

        </mesh>
    )
}

const PackingComponent = () => {
    const { routeSolution, selectedCarIndex } = useContext(RouteContext)

    const car = useMemo(() => {
        if (!routeSolution) return null
        const solution = routeSolution.carsRouteSolution[selectedCarIndex]
        if (!solution) return null
        return routeSolution.carsRouteSolution[selectedCarIndex]
    }, [routeSolution, selectedCarIndex])

    const items = useMemo(() => {
        if (!routeSolution) return []
        const solution = routeSolution.carsRouteSolution[selectedCarIndex]
        if (!solution) return []
        const newItems = []
        const items = routeSolution.carsRouteSolution[selectedCarIndex].items
        const clients = routeSolution.carsRouteSolution[selectedCarIndex].route.clients
        for (let i in items) {
            if (!clients[i].isVisible) continue
            for (let j in items[i]) {
                newItems.push(items[i][j])
            }
        }
        return newItems
    }, [routeSolution, selectedCarIndex])

    return (
        <div className='h-full w-full bg-slate-200'>
            <Canvas>
                <ambientLight />
                <pointLight position={[10, 10, 10]} />
                <OrbitControls />
                {car && <Box x={0} y={0} z={0} w={car.width} h={car.height} l={car.length} wireframe = {true} color={"black"} />}
                {car&&<Door x={0 - car.width/2 -car.width/4} y={0} z={car.length/2} w={car.width} h={car.height} wireframe = {true} color={"black"} angle = {-Math.PI} />}
                {car&&<Door x={0 +car.width/2 +car.width/4} y={0} z={car.length/2} w={car.width} h={car.height} wireframe = {true} color={"black"} angle = {Math.PI}/>}
                {car && items && items.map((item, index) => {
                    return <Box key={index} x={item.position.x_pos+item.cargo.width/2-car.width/2 } y={item.position.z_pos+item.cargo.height/2-car.height/2} z={item.position.y_pos+item.cargo.length/2-car.length/2} w={item.cargo.width} h={item.cargo.height} l={item.cargo.length} color={getColorByName(item.client.name)}  />
                })}
                {car && items && items.map((item, index) => {
                    return <Box key={index} x={item.position.x_pos+item.cargo.width/2 -car.width/2 } y={item.position.z_pos+item.cargo.height/2-car.height/2} z={item.position.y_pos+item.cargo.length/2 -car.length/2} w={item.cargo.width} h={item.cargo.height} l={item.cargo.length} wireframe = {true} color={"black"} />
                })}
            </Canvas>
        </div>
    );
}

export default PackingComponent