export const renderCell = (cell) => {
    if (!cell) {
        return null;
    }

    return (
        <div style={{backgroundColor: cell?.is_right ? "#d4edda" : "#f8d7da"}}>
            {cell?.val}
        </div>
    )
}