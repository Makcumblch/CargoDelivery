
const Spinner = () => {

    return (
        <div className="h-full w-full flex justify-center items-center">
            <div className="flex items-center gap-2">
                <span className="h-8 w-8 block rounded-full border-4 border-t-gray-400 animate-spin"></span>
            </div>

        </div>
    )
}

export default Spinner