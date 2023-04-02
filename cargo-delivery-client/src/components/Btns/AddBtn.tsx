interface AddBtnProps {
    onClick: () => void
    className?: string
}

const AddBtn = ({ onClick, className }: AddBtnProps) => {
    return (
        <button
            className={`bg-slate-500 rounded hover:bg-slate-400 ${className}`}
            onClick={onClick}
        >
            <svg width="30px" height="30px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" className="stroke-slate-300">
                <path d="M7 12L12 12M12 12L17 12M12 12V7M12 12L12 17" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round" />
            </svg>
        </button>
    );
}

export default AddBtn