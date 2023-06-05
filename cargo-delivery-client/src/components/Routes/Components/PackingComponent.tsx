import { useContext, useRef, useState, useMemo } from 'react'
import { Canvas, useFrame } from '@react-three/fiber'
import { Position, RouteContext } from '../../../contexts/RouteContext'
import { getColorByName } from './MapComponent'

interface IBoxProps {
    x: number
    y: number
    z: number
    w: number
    l: number
    h: number
    color: string
}

const Box = ({ x, y, z, w, l, h, color }: IBoxProps) => {
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
            <meshStandardMaterial opacity={hovered ? 1 : 0.5} color={color} />
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
                {car && <Box x={0} y={0} z={0} w={car.width} h={car.height} l={car.length} color={getColorByName(car.name)} />}
                {/* {items && items.map((item, index) => {
                    return <Box key={index} position={[item.position.x_pos, item.position.y_pos, item.position.z_pos]} />
                })} */}
            </Canvas>
        </div>
    );
}

export default PackingComponent