import { useEffect, useRef, useState } from "react";

interface Values {
    id: number;
    value: string
}

interface SelectProps {
    value: number | undefined,
    values: Values[],
    onChange: (id: number) => void,
}

const Select = ({ value, values, onChange }: SelectProps) => {

    const [input, setInput] = useState<string>('')
    const [isFocused, setInFocused] = useState<boolean>(false)
    const [filteredValues, setFilteredvalues] = useState<Values[]>(values)
    const inputRef = useRef<any>()
    const timer = useRef<NodeJS.Timeout>()

    const changeInput = (val: string) => {
        if (timer.current) clearTimeout(timer.current)
        timer.current = setTimeout(() => {
            const newFilteredValues = []
            for (let v in values) {
                if (values[v].value.includes(val)) newFilteredValues.push(values[v])
            }
            setFilteredvalues(newFilteredValues)
        }, 500)
        setInput(val)
    }

    useEffect(() => {
        if(value === -1) {
            setInput('')
        }
    }, [value])

    useEffect(() => {
        for (let v in values) {
            if (values[v].id === value) {
                setInput(values[v].value)
                break
            }
        }
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [])

    useEffect(() => {
        for (let v in values) {
            if (values[v].value === input) {
                onChange(values[v].id)
                inputRef.current.setCustomValidity('')
                return
            }
        }
        inputRef.current.setCustomValidity('Введите корректное название товара')
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [input, values])

    return (
        <>
            <div className="relative w-full z-[1004]">
                <input
                    ref={inputRef}
                    type="text"
                    value={input}
                    onChange={(e) => changeInput(e.target.value)}
                    className="w-full px-4 py-2 m-x text-white bg-slate-600 rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                    onFocus={() => setInFocused(true)}
                />
                <div
                    className={`${!isFocused ? 'hidden' : ''} max-h-72 overflow-y-auto bg-slate-500 ring-2 ring-slate-800 ring-opacity-40 rounded-md shadow-md absolute left-0 right-0 `}
                >
                    {filteredValues.map(element => {
                        return (
                            <div
                                className="w-full px-4 py-2 m-x text-white bg-slate-600 cursor-pointer hover:bg-slate-700"
                                key={element.id}
                                onClick={() => {
                                    setInput(element.value)
                                    setInFocused(false)
                                }}
                            >
                                {element.value}
                            </div>
                        )
                    })}
                </div>
            </div>
            <div className={`${!isFocused ? 'hidden' : ''} z-[1003] absolute left-0 top-0 w-screen h-screen`} onClick={() => setInFocused(false)}></div>
        </>
    )
}

export default Select