import React, {useEffect, useState} from "react";
import { Table, Select, Space, Input, Button, message } from "antd";
import {useDispatch, useSelector} from "react-redux";
import { getAlternatives } from "../../../reducers/alternativesSlice.ts";
import {useOpenNextStepMutation, useSendResultMutation} from "../../../api/lab1BApi.ts";
import {addOrUpdateMark} from "../../../reducers/markSlice.ts";
import ResponseTable4 from "../tables/ResponseTable4.tsx";
import {renderCell} from "../tables/util.tsx";
import ResultCard from "../results/ResultCard.tsx";

const { Option } = Select;

const prioritiesLocalStorage = "prioritiesLocalStorage";
const selectedAlternativeLocalStorage = "prioritiesSelectedAlternativeLocalStorage";
const prioritiesResponseResultLocalStorage = "prioritiesResponseResultLocalStorage";
const prioritiesResponseIndexLocalStorage = "prioritiesResponseIndexLocalStorage";

interface Props {
    onCompleteStep: () => void;
}

const Priorities: React.FC<Props> = ({ onCompleteStep }) => {
    const markId = "Приоритеты";
    const dispatch = useDispatch();
    const alternatives = useSelector(getAlternatives);
    const initialSelectedAlternative = JSON.parse(localStorage.getItem(selectedAlternativeLocalStorage) as string)
    const initialPriorities = JSON.parse(localStorage.getItem(prioritiesLocalStorage) as string)
    const initialResponseResult = JSON.parse(localStorage.getItem(prioritiesResponseResultLocalStorage) as string)
    const initialResponseIndex = JSON.parse(localStorage.getItem(prioritiesResponseIndexLocalStorage) as string)
    const [selectedAlternative, setSelectedAlternative] = useState<string | null>(initialSelectedAlternative);
    const [priorities, setPriorities] = useState<Record<string, number | string>>(initialPriorities || {});
    const [disabledInput, setDisabledInput] =  useState(false)


    const [sendResult, {
        data: response,
        error,
        isSuccess
    }] = useSendResultMutation();

    const [openNextStep] = useOpenNextStepMutation()

    useEffect(() => {
        if (initialResponseResult && initialResponseIndex) {
            setDisabledInput(true);
        }
    }, [initialResponseIndex, initialResponseResult]);

    useEffect(() => {
        openNextStep(5)
    }, []);

    useEffect(() => {
        if (response) {
            dispatch(addOrUpdateMark({id: markId, mark: response.percentage, maxMark: response.max_mark}))
        }
    }, [dispatch, response]);

    useEffect(() => {
        if (error) {
            message.error(JSON.stringify(error));
            console.error(error)
        }
    }, [error]);


    if (response) {
        localStorage.setItem(prioritiesResponseResultLocalStorage, JSON.stringify(response.result));
        localStorage.setItem(prioritiesResponseIndexLocalStorage, JSON.stringify(response.index));
    }
    if (isSuccess) {
        localStorage.setItem(selectedAlternativeLocalStorage, JSON.stringify(selectedAlternative));
        localStorage.setItem(prioritiesLocalStorage, JSON.stringify(priorities));
        localStorage.setItem(`isDone`, "1");
        onCompleteStep()
    }

    const handlePriorityChange = (alt: string, value: string) => {
        setPriorities(prev => ({ ...prev, [alt]: value }));
    };

    const handleSave = async () => {
        if (Object.values(priorities).some(value => value === "")) {
            message.error("Заполните все приоритеты перед сохранением!");
            return;
        }

        const requestBody = {
            set: Object.values(priorities).map(Number),
            chosen_index: alternatives.indexOf(selectedAlternative) + 1,
        }

        await sendResult(requestBody)
    };


    const columns = [
        {
            title: "Альтернативы",
            dataIndex: "alternative",
            key: "alternative",
        },
        {
            title: "Приоритет",
            dataIndex: "priority",
            key: "priority",
            render: (_: any, record: { alternative: string }) => (
                <Input
                    type="number"
                    value={priorities[record.alternative] || ""}
                    onChange={e => handlePriorityChange(record.alternative, e.target.value)}
                    readOnly={disabledInput}
                />
            ),
        },
    ];


    const dataSource = alternatives.map((alt: string, index: number) => ({
        key: index,
        alternative: alt,
    }));

    return (
        <Space direction="vertical" size="large" style={{ width: "100%" }}>
            <Table
                pagination={false}
                bordered
                dataSource={dataSource}
                columns={columns}
                style={{ maxWidth: "400px" }}
            />

            <Select
                style={{ width: "300px" }}
                placeholder="Выберите альтернативу"
                onChange={setSelectedAlternative}
                value={selectedAlternative}
                disabled={disabledInput}
            >
                {alternatives.map((alt: string, index: number) => (
                    <Option key={index} value={alt}>
                        {alt}
                    </Option>
                ))}
            </Select>

            <Button type="primary" onClick={handleSave} disabled={disabledInput}>
                Отправить
            </Button>
            {initialResponseResult && <>
                <div style={{fontWeight: "bold", fontSize: 16}}>Ответы</div>
                <div style={{display: "flex", gap: 10}}>
                    Альтернатива: <span style={{width: 30, textAlign: "center"}}>{renderCell(initialResponseIndex)}</span>
                </div>
                <ResponseTable4 data={initialResponseResult} alternatives={alternatives} />
            </>}
            <ResultCard markId={markId}/>
        </Space>
    );
};

export default Priorities;
