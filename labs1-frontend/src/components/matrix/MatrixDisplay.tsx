import React from "react";

type MatrixDisplayProps = {
    matrix: number[][];
    title?: string;
    isFirstTable?: boolean;
};

const MatrixDisplay: React.FC<MatrixDisplayProps> = ({ matrix, index}) => {
    const firstTableHeaders = ['Д', 'Р', 'С', 'Л'];
    const otherTableHeaders = ['А', 'Б', 'В'];

    const renderTable = () => {
        return (
            <table style={styles.table}>
                <thead>
                <tr>
                    <th style={styles.verticalHeader}></th>
                    {index === 1
                        ? firstTableHeaders.map((header, index) => (
                            <th key={index} style={styles.cell}>
                                {header}
                            </th>
                        ))
                        : otherTableHeaders.map((header, index) => (
                            <th key={index} style={styles.cell}>
                                {header}
                            </th>
                        ))}
                </tr>
                </thead>
                <tbody>
                {matrix.map((row, rowIndex) => (
                    <tr key={rowIndex}>
                        <td style={styles.verticalHeader}>{
                            index === 1
                                ? firstTableHeaders[rowIndex]
                                : otherTableHeaders[rowIndex]
                        }</td>
                        {row.map((cell, cellIndex) => (
                            <td key={cellIndex} style={styles.cell}>
                                {console.log(cellIndex)}
                                {cell}
                            </td>
                        ))}
                    </tr>
                ))}
                </tbody>
            </table>
        );
    };

    return (
        <div style={styles.container}>
            {renderTable()}
        </div>
    );
};

const styles: { [key: string]: React.CSSProperties } = {
    container: {
        background: "#f9f9f9",
        borderRadius: "8px",
        boxShadow: "0 4px 6px rgba(0, 0, 0, 0.1)",
        padding: "15px",
        textAlign: "center",
        margin: "10px 0",
    },
    title: {
        marginBottom: "10px",
        fontSize: "18px",
        fontWeight: "bold",
    },
    table: {
        width: "100%",
        borderCollapse: "collapse",
    },
    cell: {
        border: "1px solid #ccc",
        padding: "10px",
        textAlign: "center",
        minWidth: "50px",
        maxWidth: "50px",
        height: "50px",
    },
    verticalHeader: {
        border: "1px solid #ccc",
        padding: "10px",
        textAlign: "center",
        width: "50px",
        height: "50px",
        writingMode: "vertical-rl", // Вертикальное написание
        transform: "rotate(180deg)", // Поворот текста для нормального отображения
        fontWeight: "bold", // Жирный шрифт для вертикальных заголовков
    },
};

export default MatrixDisplay;
