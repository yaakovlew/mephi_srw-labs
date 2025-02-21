import { useState, useEffect } from "react";
import axios from "axios";
import { Tokens, base1AURL } from './const.tsx';

type isRight = {
    val: number;
    is_right: boolean;
};

type reqRes = {
    max_mark: number;
    percentage: number;
    result: isRight[];
    index: isRight;
};

type reqInput = {
    set: number[];
    index: number;
};

const ElementSelector = ({changeStepFunc}) => {
    const [elements, setElements] = useState(['', '', '']);
    const [selectedElement, setSelectedElement] = useState<string | null>(null);
    const [chosenIndex, setChosenIndex] = useState(0);
    const [reqResp, setReqResp] = useState<reqRes | null>(null);
    const [isSubmitted, setIsSubmitted] = useState(false);
    const [reqIn, setReqIn] = useState<reqInput | null>(null);

    const handleChange = (index: number, value: string) => {
        if (/^\d*\.?\d*$/.test(value)) {
            const newElements = [...elements];
            newElements[index] = value;
            setElements(newElements);
        }
    };

    useEffect(() => {
        const savedData = localStorage.getItem(`lab1a_6_data`);
        const resData = localStorage.getItem(`lab1a_6_result`);
        if (savedData && resData) {
            const data = JSON.parse(savedData);
            setChosenIndex(data.chosen_index);
            setElements(data.set);

            const resDataL = JSON.parse(resData);
            setReqResp(resDataL);
            setIsSubmitted(true);
        } else {
            setIsSubmitted(false);
        }
    }, []);

    const handleSelect = (index: number) => {
        setSelectedElement(elements[index]);
        setChosenIndex(index);
    };

    const roundNumbers = (inputElements: string[]) => {
        return inputElements.map((el) => {
            const num = parseFloat(el);
            if (!isNaN(num)) {
                return num.toFixed(2);
            }
            return el;
        });
    };

    const handleSubmit = async () => {
        const numericElements = elements.map((el) => parseFloat(el)).filter((el) => !isNaN(el));

        if (numericElements.length !== elements.length) {
            return;
        }

        const requestData: reqInput = {
            set: numericElements,
            index: chosenIndex+1,
        };
        setReqIn(requestData);

        let token = localStorage.getItem("token");

        try {
            const response = await axios.post<reqRes>(
                `${base1AURL}/lab1a/variant/result`,
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
            let currStepString = localStorage.getItem("step");
            let currStep = currStepString ? parseInt(currStepString, 10) : 0;
            if (currStep !== 0) {
                localStorage.setItem("step", `${currStep+1}`);
            }
            changeStepFunc();

            if (response) {
                localStorage.setItem(`lab1a_6_data`, JSON.stringify(requestData));
                localStorage.setItem(`lab1a_6_result`, JSON.stringify(response.data));
                setIsSubmitted(true);
                localStorage.setItem(`isDone`, "1");
            } else {
                alert("Произошла ошибка при отправке данных.");
            }
        } catch (error) {
            alert("Произошла ошибка при отправке запроса.");
        }
    };

    const getResultStyle = (isRight: boolean) => ({
        color: isRight ? 'green' : 'red',
        textDecoration: 'underline',
    });

    return (
        <div style={styles.container}>
            <h2>Приоритеты альтернатив:</h2>
            <div style={styles.inputContainer}>
                {['A', 'B', 'C'].map((label, index) => (
                    <div key={index} style={styles.inputWrapper}>
                        <label style={styles.label}>{label}</label>
                        <input
                            type="text"
                            value={elements[index]}
                            onChange={(e) => handleChange(index, e.target.value)}
                            placeholder={`Альтернатива ${label}`}
                            style={styles.input}
                        />
                    </div>
                ))}
            </div>

            {elements.every((el) => el) && !isSubmitted && (
                <div style={styles.selectContainer}>
                    <h3>Выберите альтернативу:</h3>
                    <div style={styles.buttonGroup}>
                        {['A', 'B', 'C'].map((label, index) => (
                            <button
                                key={index}
                                onClick={() => handleSelect(index)}
                                style={{
                                    ...styles.button,
                                    backgroundColor: chosenIndex === index ? '#4CAF50' : '#2196F3',
                                    border: chosenIndex === index ? '2px solid #fff' : 'none',
                                    color: 'white',
                                    paddingLeft: "20px",
                                }}
                            >
                                {label}
                            </button>
                        ))}
                    </div>
                </div>
            )}

            {selectedElement && (
                <div style={styles.selectedElementContainer}>
                </div>
            )}

            {selectedElement !== null && !isSubmitted && (
                <button onClick={handleSubmit} style={styles.submitButton}>
                    Отправить
                </button>
            )}

            {isSubmitted && reqResp && (
                <div style={styles.resultsContainer}>
                    <div style={{margin: "0 auto"}}>
                        <h4 style={{display: "flex", justifyContent: "center"}}>Результаты:</h4>
                    </div>
                    <div style={styles.inputContainer}>
                        {['A', 'B', 'C'].map((label, index) => (
                            <div key={index} style={styles.inputWrapper}>
                                <label style={styles.label}>{label}</label>
                                <p
                                    type="text"
                                    style={reqResp.result[index].is_right ? styles.pGreen : styles.pRed}>
                                    {reqResp.result[index].val}
                                </p>
                            </div>

                        ))}
                    </div>
                    <br></br>
                    <p
                        type="text"
                        style={reqResp.index.is_right ? styles.pAlternativeGreen : styles.pAlternativeRed}>
                        Наилучшая альтернатива: {['A', 'B', 'C'][reqResp.index.val-1]}</p>
                </div>
            )}
        </div>
    );
};

const styles = {
    container: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        padding: '20px',
        fontFamily: 'Arial, sans-serif',
    },
    inputContainer: {
        display: 'flex',
        gap: '10px',
        marginBottom: '20px',
    },
    inputWrapper: {
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
    },
    label: {
        fontSize: '18px',
        marginBottom: '5px',
    },
    input: {
        padding: '10px',
        fontSize: '16px',
        borderRadius: '5px',
        border: '1px solid #ccc',
        width: '150px',
    },
    pRed: {
        padding: '10px',
        fontSize: '16px',
        borderRadius: '5px',
        border: '1px solid #ccc',
        color: "red",
        width: '150px',
    },
    pGreen: {
        padding: '10px',
        fontSize: '16px',
        borderRadius: '5px',
        border: '1px solid #ccc',
        color: "green",
        width: '150px',
    },
    pAlternativeRed: {
        margin: "0 auto",
        color: "red",
        width: '300px',
        paddingLeft: '80px',
    },
    pAlternativeGreen: {
        margin: "0 auto",
        paddingLeft: '80px',
        width: '300px',
        color: "green",
    },
    selectContainer: {
        marginTop: '20px',
    },
    buttonGroup: {
        display: 'flex',
        gap: '10px',
    },
    button: {
        padding: '10px 20px',
        fontSize: '16px',
        borderRadius: '5px',
        border: 'none',
        cursor: 'pointer',
        transition: 'background-color 0.3s, border 0.3s',
    },
    selectedElementContainer: {
        marginTop: '20px',
        fontSize: '18px',
        fontWeight: 'bold',
    },
    submitButton: {
        marginTop: '20px',
        padding: '10px 20px',
        backgroundColor: '#4CAF50',
        color: 'white',
        border: 'none',
        borderRadius: '5px',
        fontSize: '16px',
        cursor: 'pointer',
        transition: 'background-color 0.3s',
    },
    resultsContainer: {
        marginTop: '30px',
        fontSize: '18px',
        fontWeight: 'bold',
        textAlign: 'left',
    },
    resultItem: {
        marginBottom: '10px',
    },
};

export default ElementSelector;