import { useEffect, useMemo, useRef, useState } from "react";
import { OpenStreetMapProvider } from "leaflet-geosearch"
import { SearchResult } from "leaflet-geosearch/dist/providers/provider";
import { RawResult } from "leaflet-geosearch/dist/providers/openStreetMapProvider";

interface SelectAddressProps {
    address: string,
    onChange: (address: string, coordX: number, coordY: number) => void,
}

const SelectAddress = ({ address, onChange }: SelectAddressProps) => {

    const [input, setInput] = useState<string>(address)
    const [isFocused, setInFocused] = useState<boolean>(false)
    const [filteredValues, setFilteredvalues] = useState<SearchResult<RawResult>[]>([])
    const inputRef = useRef<any>()
    const timer = useRef<NodeJS.Timeout>()

    const provider = useMemo(() => {
        return new OpenStreetMapProvider()
    }, [])

    const changeInput = (val: string) => {
        if (timer.current) clearTimeout(timer.current)
        timer.current = setTimeout(async () => {
            const res = await provider.search({ query: val })
            setFilteredvalues(res)
        }, 500)
        setInput(val)
    }

    useEffect(() => {
        for (let v in filteredValues) {
            if (filteredValues[v].label === input) {
                onChange(filteredValues[v].label, filteredValues[v].x, filteredValues[v].y)
                inputRef.current.setCustomValidity('')
                return
            }
        }
        inputRef.current.setCustomValidity('Введите корректный адрес')
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [input, filteredValues])

    return (
        <>
            <div className="relative w-full z-[1004]">
                <input
                    ref={inputRef}
                    type="text"
                    value={input}
                    placeholder="Введите адрес"
                    onChange={(e) => changeInput(e.target.value)}
                    className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                    onFocus={() => setInFocused(true)}
                />
                <div
                    className={`${!isFocused ? 'hidden' : ''} max-h-72 overflow-y-auto bg-slate-500 ring-2 ring-slate-800 ring-opacity-40 rounded-md shadow-md absolute left-0 right-0 `}
                >
                    {filteredValues.map(element => {
                        return (
                            <div
                                className="w-full px-4 py-2 m-x text-white bg-slate-600 cursor-pointer hover:bg-slate-700"
                                key={element.raw.place_id}
                                onClick={() => {
                                    setInput(element.label)
                                    setInFocused(false)
                                }}
                            >
                                {element.label}
                            </div>
                        )
                    })}
                </div>
            </div>
            <div className={`${!isFocused ? 'hidden' : ''} z-[1003] absolute left-0 top-0 w-screen h-screen`} onClick={() => setInFocused(false)}></div>
        </>
    )
}

export default SelectAddress