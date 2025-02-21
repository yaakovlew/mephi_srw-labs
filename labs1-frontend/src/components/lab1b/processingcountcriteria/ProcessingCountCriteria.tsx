import React, {useEffect, useState} from "react";
import {useDispatch, useSelector} from "react-redux";
import { Button, Input, Space, Table, message } from "antd";
import { useCheckStepMutation } from "../../../api/lab1BApi.ts";
import { getAlternatives } from "../../../reducers/alternativesSlice.ts";
import { getCriterias } from "../../../reducers/criteriasSlice.ts";
import {addOrUpdateMark} from "../../../reducers/markSlice.ts";
import ResponseTable3 from "../tables/ResponseTable3.tsx";
import ResultCard from "../results/ResultCard.tsx";

const procCountTablesLocalStorage = "procCountTablesLocalStorage"
const procCountResponseResultLocalStorage = "procCountResponseResultLocalStorage"

interface Props {
    onCompleteStep: () => void;
}

const ProcessingCountCriteria: React.FC<Props> = ({onCompleteStep}) => {
    const markId = "Шаг 4";
    const dispatch = useDispatch();
    const alternatives = useSelector(getAlternatives); // Альтернативы
    const criterias = useSelector(getCriterias); // Критерии
    const initialTables = JSON.parse(localStorage.getItem(procCountTablesLocalStorage) as string);
    const initialResponseResult = JSON.parse(localStorage.getItem(procCountResponseResultLocalStorage) as string);
    const [disabledInput, setDisabledInput] =  useState(false)


    const [sendData, {
        data: response,
        isLoading, 
        isSuccess, 
        error 
    }] = useCheckStepMutation();


    const [tables, setTables] = useState(() => {
        if (initialTables) {
            setDisabledInput(true);
            return initialTables
        }

        return criterias
            .flatMap((criteria: any) =>
                criteria.extra.filter((sub: any) => sub.isCount).map((sub: any) => sub.id)
            )
            .reduce((acc: any, subId: number) => {
                acc[subId] = {
                    grades: alternatives.map(() => ""), // Строка для оценок
                };
                return acc;
            }, {});
    });


    useEffect(() => {
        if(error) {
            console.error(error);
            message.error(JSON.stringify(error.data.message));
        }
    }, [error]);


    useEffect(() => {
        if (response) {
            dispatch(addOrUpdateMark({id: markId, mark: response.percentage, maxMark: response.max_mark}))
        }
    }, [response, dispatch]);

    const handleInputChange = (subId: number, index: number, value: string) => {
        setTables((prev: any) => ({
            ...prev,
            [subId]: {
                ...prev[subId],
                grades: prev[subId].grades.map((grade: string, i: number) =>
                    i === index ? value : grade
                ),
            },
        }));
    };


    const handleSend = async () => {
        try {
            const sendMatrix = Object.values(tables).map(item =>
                item.grades.map(arrItem =>
                    Number.parseFloat(arrItem) || 0
                )
            );

            const result = await sendData({
                stepId: "6",
                body: {
                    marks: sendMatrix,
                },
            });

            localStorage.setItem(procCountResponseResultLocalStorage, JSON.stringify(result.data.result.marks));
        } catch (error) {
            console.error("Ошибка при отправке данных:", error);
        }
    };

    if (isSuccess) {
        localStorage.setItem(procCountTablesLocalStorage, JSON.stringify(tables));
        onCompleteStep();
    }

    const dataSource = ({sub}) => {
        return (
            [
                {
                    key: "value",
                    rowName: "доп. критерий",
                    ...alternatives.reduce((acc: any, alt: any, index: number) => {
                        acc[`alt${index}`] = sub.value[index]; // Используем value из Redux
                        return acc;
                    }, {}),
                },
                {
                    key: "grades",
                    rowName: "оценка",
                    ...alternatives.reduce((acc: any, alt: any, index: number) => {
                        acc[`alt${index}`] = (
                            <Input
                                value={tables[sub.id]?.grades[index] || ""}
                                onChange={(e) =>
                                    handleInputChange(sub.id, index, e.target.value)
                                }
                                readOnly={disabledInput}
                            />
                        );
                        return acc;
                    }, {}),
                },
            ]
        )
    }

    return (
        <Space direction="vertical" size="large" style={{ width: "100%" }}>
            {criterias.map((criteria: any) =>
                criteria.extra
                    .filter((sub: any) => sub.isCount)
                    .map((sub: any) => (
                        <div
                            key={sub.id}
                            style={{ border: "1px solid #ddd", padding: "16px", marginBottom: "16px" }}
                        >
                            <h3>
                                {criteria.criteria} - {sub.name}
                            </h3>
                            <Table
                                pagination={false}
                                bordered
                                dataSource={dataSource({sub})}
                                columns={[
                                    {
                                        title: "",
                                        dataIndex: "rowName",
                                        key: "rowName",
                                    },
                                    ...alternatives.map((alt: any, index: number) => ({
                                        title: alt,
                                        dataIndex: `alt${index}`,
                                        key: `alt${index}`,
                                    })),
                                ]}
                            />
                        </div>
                    ))
            )}
            <Button
                type="primary"
                onClick={handleSend}
                loading={isLoading}
                disabled={isLoading || disabledInput}
                style={{ marginTop: "16px" }}
                block
            >
                Отправить
            </Button>
            {initialResponseResult && <>
                <ResponseTable3 alternatives={alternatives} tables={initialResponseResult} />
            </>}
            <ResultCard markId={markId}/>
        </Space>
    );
};

export default ProcessingCountCriteria;
