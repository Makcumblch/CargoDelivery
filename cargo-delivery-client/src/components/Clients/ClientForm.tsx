import { FormEvent, useState } from "react";
import { Client } from "../../contexts/ClientsContext";

interface ClientFormProps {
    input: Client,
    changeClient: (id: number, value: object) => void
    close: () => void
}

const ClientForm = ({ input, changeClient, close }: ClientFormProps) => {
    const [inputClient, setInputClient] = useState<Client>(input)

    const onChange = (field: string, value: any): void => {
        setInputClient((prev: Client) => {
            return { ...prev, [field]: value }
        })
    }

    const onSubmit = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()
        changeClient(inputClient.id, inputClient)
        close()
    }

    return (
        <form id='form' onSubmit={onSubmit} className="flex flex-col space-y-2 w-full mb-4">
            <div>
                <label htmlFor="name" className="block text-sm font-semibold text-slate-300">Название</label>
                <input
                    value={inputClient.name}
                    type="text"
                    name="name"
                    required
                    placeholder="Введите название"
                    className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                    onChange={(e) => onChange('name', e.target.value)}
                />
            </div>
            <div>
                <label htmlFor="loadCapacity" className="block text-sm font-semibold text-slate-300">Адрес</label>
                <input
                    value={inputClient.address}
                    type="text"
                    name="address"
                    required
                    placeholder="Введите адрес"
                    className="block w-full px-4 py-2 mt-2 text-white bg-slate-600  rounded-md focus:ring-slate-800 focus:outline-none focus:ring focus:ring-opacity-40"
                    onChange={(e) => onChange('address', e.target.value)}
                />
            </div>
        </form>
    );
}

export default ClientForm