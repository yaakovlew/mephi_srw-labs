import React, { useState, useEffect } from "react";
import axios from "axios";
import { Tokens, base1AURL } from './const.tsx';

type MatrixTaskProps = {
    matrix: number[][];
    step: number;
    title?: string;
};

type isRight = {
    val: number;
    is_right: boolean;
}

const newTableData = {
    sizes: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
    ss: [0, 0, 0.58, 0.90, 1.12, 1.24, 1.32, 1.41, 1.45, 1.49],
};

type result = {
    x: isRight[];
    w: isRight[];
    mw: isRight[];
    lambda_w: isRight[];
    lambda_max: isRight;
    is: isRight;
    os: isRight;
}

type resp = {
    percentage: number;
    max_mark: number;
    result: result;
}

const getCellStyle = (isRight: boolean) => {
    return isRight ? {
            width: "60px",
            textAlign: "center",
            padding: "8px",
            marginLeft: "5px",
            border: "1px solid #ddd",
            borderRadius: "4px",
            fontSize: "14px",
            color: "green",
            minWidth: "60px",
            maxWidth: "60px",
            height: "60px",
        }
        : { width: "60px",
            textAlign: "center",
            padding: "8px",
            marginLeft: "5px",
            border: "1px solid #ddd",
            borderRadius: "4px",
            fontSize: "14px",
            color: "red",
            minWidth: "60px",
            maxWidth: "60px",
            height: "60px",};
};

const task = [
    "Расчет параметров матрицы важности по цели",
    "Расчет параметров матрицы важности по критерию Доходность",
    "Расчет параметров матрицы важности по критерию Риск",
    "Расчет параметров матрицы важности по критерию Стоимость акций",
    "Расчет параметров матрицы важности по критерию Ликвидность",
];

const Lab1AMatrixTaskStep: React.FC<MatrixTaskProps> = ({ matrix, step, changeStepFunc}) => {
    const [lambdaMax, setLambdaMax] = useState<number | string>("");
    const [IS, setIS] = useState<number | string>("");
    const [OS, setOS] = useState<number | string>("");
    const [reqResp, setReqResp] = useState<resp>({});
    const [matrix4, setMatrix4] = useState(
        matrix.map(() => Array(4).fill(''))
    );
    const [isSubmitted, setIsSubmitted] = useState(false);

    useEffect(() => {
        const savedData = localStorage.getItem(`lab1a_${step}_data`);
        const resData = localStorage.getItem(`lab1a_${step}_result`);
        if (savedData && resData) {
            let data =JSON.parse(savedData);
            setMatrix4(transposeMatrix(
                [data.x, data.w, data.mw, data.lambda_w]
            ));

            setLambdaMax(data.lambda_max);
            setIS(data.is);
            setOS(data.os);

            let resDataL =JSON.parse(resData);
            setReqResp(resDataL)
            setIsSubmitted(true);
        } else {
            setIsSubmitted(false);
            setMatrix4(matrix.map(() => Array(4).fill('')));

						setLambdaMax(0);
            setIS(0);
            setOS(0);
        }
    }, [matrix, step, changeStepFunc]);

    function transposeMatrix(matrix) {
        if (!matrix || matrix.length === 0) return [];

        const rows = matrix.length;
        const cols = matrix[0].length;
        const transposed = Array.from({ length: cols }, () => Array(rows).fill(0));

        for (let i = 0; i < rows; i++) {
            for (let j = 0; j < cols; j++) {
                transposed[j][i] = matrix[i][j];
            }
        }

        return transposed;
    }

    const handleMatrixChange = (rowIndex: number, colIndex: number, value: string) => {
        if (isSubmitted) return;
        const newMatrix = [...matrix4];
        newMatrix[rowIndex][colIndex] = value;
        setMatrix4(newMatrix);
    };

    const roundToTwoDecimals = (value: string) => {
        if (value === "") {
            return "0.00";
        }
        return parseFloat(value).toFixed(2);
    };

    const handleSubmit = async () => {
        const x = matrix4.map(row => roundToTwoDecimals(row[0] || ""));
        const w = matrix4.map(row => roundToTwoDecimals(row[1] || ""));
        const mw = matrix4.map(row => roundToTwoDecimals(row[2] || ""));
        const lambdaW = matrix4.map(row => roundToTwoDecimals(row[3] || ""));

        const requestData = {
            x: x.map((val) => parseFloat(val)),
            w: w.map((val) => parseFloat(val)),
            mw: mw.map((val) => parseFloat(val)),
            lambda_w: lambdaW.map((val) => parseFloat(val)),
            lambda_max: parseFloat(roundToTwoDecimals(lambdaMax)),
            is: parseFloat(roundToTwoDecimals(IS)),
            os: parseFloat(roundToTwoDecimals(OS)),
        };

        let token = localStorage.getItem("token");
        try {
            const response = await axios.post<result>(
                `${base1AURL}/lab1a/variant/${step}`,
                requestData,
                {
                    headers: {
                        Authorization: `Bearer ${token}`,
                        'lab-token': Tokens.lab1AToken,
                        'Content-Type': 'application/json',
                    }
                }
            );

            setReqResp(response.data);

            if (response) {
                localStorage.setItem(`lab1a_${step}_data`, JSON.stringify(requestData));

                localStorage.setItem(`lab1a_${step}_result`, JSON.stringify(response.data));

                setIsSubmitted(true);
                let currStepString = localStorage.getItem("step");
                let currStep = currStepString ? parseInt(currStepString, 10) : 0;
                if (currStep !== 0) {
                    localStorage.setItem("step", `${step+1}`);
                }
                changeStepFunc();
            } else {
                alert("Произошла ошибка при отправке данных.");
            }
        } catch (error) {
            alert("Произошла ошибка при отправке запроса.");
        }
    };

    return (
        <div style={styles.container}>
            <div style={styles.matrixContainer}>
                <h4>{task[step-1]}</h4>
                <h5>(необходимо дать ответ для скорректированных значений матрицы)</h5>
                <div style={styles.matrixContainer}>
                    <table style={styles.smallTable}>
                        <thead>
                        <tr>
                            <th style={{ ...styles.smallCell, fontWeight: "bold" }}>Размер матрицы</th>
                            {newTableData.sizes.map((size, index) => (
                                <th key={index} style={styles.smallCell}>{size}</th>
                            ))}
                        </tr>
                        </thead>
                        <tbody>
                        <tr>
                            <th style={{ ...styles.smallCell, fontWeight: "bold" }}>СС</th>
                            {newTableData.ss.map((value, index) => (
                                <td key={index} style={styles.smallCell}>{value}</td>
                            ))}
                        </tr>
                        </tbody>
                    </table>
                </div>
                <table style={styles.table}>
                    <thead>
                    <tr>
                        <th style={styles.headerCell}></th> {/* Пустая ячейка для заголовков строк */}
                        {step === 1
                            ?	["Д", "Р", "С", "Л"].map((header, index) => (
                                <th key={index} style={styles.headerCell}>{header}</th>
                            ))
                            : ["А", "Б", "В"].map((header, index) => (
                                <th key={index} style={styles.headerCell}>{header}</th>
                            ))
                        }
                    </tr>
                    </thead>
                    <tbody>
                    {matrix.map((row, rowIndex) => (
                        <tr key={rowIndex}>
                            <th style={styles.headerCell}>{
                                step === 1
                                    ? ["Д", "Р", "С", "Л"][rowIndex]
                                    : ["А", "Б", "В"][rowIndex]
                            }</th>
                            {row.map((cell, cellIndex) => (
                                <td key={cellIndex} style={styles.cell}>
                                    {cell}
                                </td>
                            ))}
                        </tr>
                    ))}
                    </tbody>
                </table>
            </div>

            <div style={styles.matrixContainer}>
                <h4>Параметры матрицы:</h4>
                <table style={styles.table}>
                    <thead>
                    <tr>
                        {['x', 'w', 'M * w', 'λ * w'].map((header, index) => (
                            <th key={index} style={styles.headerCell}>{header}</th>
                        ))}
                    </tr>
                    </thead>
                    <tbody>

                    {matrix4.map((row, rowIndex) => (
                        <tr key={rowIndex}>
                            {matrix4[rowIndex].map((cell, cellIndex) => (
                                <td key={cellIndex} style={styles.cell}>
                                    <input
                                        type="number"
                                        step="0.01"
                                        value={cell || ""}
                                        onChange={(e) =>
                                            handleMatrixChange(rowIndex, cellIndex, e.target.value)
                                        }
                                        disabled={isSubmitted}
                                        style={styles.inputCell}
                                    />
                                </td>
                            ))}
                        </tr>
                    ))}
                    </tbody>
                </table>
            </div>

            <div style={styles.inputContainer}>
                <div style={styles.inputGroupRow}>
                    <div style={styles.inputGroup}>
                        <label style={styles.inputLabel}>λmax:</label>
                        <input
                            type="number"
                            value={lambdaMax || ""}
                            onChange={(e) => setLambdaMax(e.target.value)}
                            disabled={isSubmitted}
                            style={styles.input}
                        />
                    </div>

                    <div style={styles.inputGroup}>
                        <label style={styles.inputLabel}>ИС:</label>
                        <input
                            type="number"
                            value={IS || ""}
                            onChange={(e) => setIS(e.target.value)}
                            disabled={isSubmitted}
                            style={styles.input}
                        />
                    </div>

                    <div style={styles.inputGroup}>
                        <label style={styles.inputLabel}>ОС:</label>
                        <input
                            type="number"
                            value={OS || ""}
                            onChange={(e) => setOS(e.target.value)}
                            disabled={isSubmitted}
                            style={styles.input}
                        />
                    </div>
                </div>
            </div>

            {!isSubmitted && <button onClick={handleSubmit} style={styles.button}>Отправить</button>}
            {isSubmitted && (
                <div>
                    <div style={styles.matrixContainer}>
                        <h4>Результаты расчетов</h4>
                        <table style={styles.table}>
                            <thead>
                            <tr>
                                {['x', 'w', 'M * w', 'λ * w'].map((header, index) => (
                                    <th key={index} style={styles.headerCell}>{header}</th>
                                ))}
                            </tr>
                            </thead>
                            <tbody>
                            {reqResp.result.x.map((_, rowIndex) => (
                                <tr key={rowIndex}>
                                    <td style={getCellStyle(reqResp.result.x[rowIndex].is_right)}>{reqResp.result.x[rowIndex].val}</td>
                                    <td style={getCellStyle(reqResp.result.w[rowIndex].is_right)}>{reqResp.result.w[rowIndex].val}</td>
                                    <td style={getCellStyle(reqResp.result.mw[rowIndex].is_right)}>{reqResp.result.mw[rowIndex].val}</td>
                                    <td style={getCellStyle(reqResp.result.lambda_w[rowIndex].is_right)}>{reqResp.result.lambda_w[rowIndex].val}</td>
                                </tr>
                            ))}
                            </tbody>
                        </table>
                    </div>

                    <div style={styles.resultSummary}>
                        <span style={getCellStyle(reqResp.result.lambda_max.is_right)}>{`λmax: ${reqResp.result.lambda_max.val}`}</span>
                        <span style={getCellStyle(reqResp.result.is.is_right)}>{`ИС: ${reqResp.result.is.val}`}</span>
                        <span style={getCellStyle(reqResp.result.os.is_right)}>{`ОС: ${reqResp.result.os.val}`}</span>
                    </div>
                    <br></br>
                    <b>{`Получено баллов: ${reqResp.percentage}/${reqResp.max_mark}`}</b>
                </div>
            ) }
        </div>
    );
};

const styles: { [key: string]: React.CSSProperties } = {
    container: {
        background: "#f9f9f9",
        borderRadius: "8px",
        boxShadow: "0 4px 6px rgba(0, 0, 0, 0.1)",
        padding: "30px",
        textAlign: "center",
        margin: "30px 0",
        width: "80%",
        maxWidth: "1000px",
        marginLeft: "auto",
        marginRight: "auto",
    },
    title: {
        marginBottom: "20px",
        fontSize: "24px",
        fontWeight: "bold",
        color: "#333",
    },
    matrixContainer: {
        marginBottom: "30px",
    },
    table: {
        width: "100%",
        borderCollapse: "collapse",
        marginBottom: "20px",
        borderRadius: "8px",
        overflow: "hidden",
        boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)",
    },
    headerCell: {
        backgroundColor: "#f1f1f1",
        fontWeight: "bold",
        padding: "10px",
        border: "1px solid #ddd",
        textAlign: "center",
        minWidth: "100px",
    },
    cell: {
        border: "1px solid #ddd",
        padding: "12px",
        textAlign: "center",
        minWidth: "60px",
        maxWidth: "60px",
        height: "60px",
        fontSize: "14px",
    },
    inputCell: {
        width: "60px",
        textAlign: "center",
        padding: "8px",
        border: "1px solid #ddd",
        borderRadius: "4px",
        fontSize: "14px",
    },
    inputContainer: {
        marginTop: "20px",
        textAlign: "center",
    },
    inputGroupRow: {
        display: "flex",
        justifyContent: "space-around",
        marginBottom: "20px",
    },
    inputGroup: {
        width: "30%",
    },
    inputLabel: {
        fontSize: "16px",
        marginBottom: "10px",
        display: "block",
        fontWeight: "bold",
        color: "#333",
    },
    input: {
        width: "100%",
        padding: "10px",
        fontSize: "16px",
        borderRadius: "4px",
        border: "1px solid #ddd",
    },
    button: {
        backgroundColor: "#4CAF50",
        color: "white",
        padding: "14px 20px",
        fontSize: "16px",
        borderRadius: "4px",
        cursor: "pointer",
        border: "none",
        transition: "background-color 0.3s ease",
    },
    buttonHover: {
        backgroundColor: "#45a049",
    },

    smallTable: {
        width: "60%",
        margin: "0 auto",
        fontSize: "12px",
        borderCollapse: "collapse",
    },

    smallCell: {
        border: "1px solid #ddd",
        padding: "6px", //
        textAlign: "center",
        minWidth: "30px",
        maxWidth: "50px",
        height: "30px",
    },
};

export default Lab1AMatrixTaskStep;