import React, {useEffect, useState} from "react";
import { Button, Input, Space, Table, message } from "antd";
import {useCheckStepMutation, useOpenNextStepMutation} from "../../../api/lab1BApi.ts";
import {useDispatch, useSelector} from "react-redux";
import {getMainCriterias} from "../../../reducers/criteriasSlice.ts";
import {addOrUpdateMark} from "../../../reducers/markSlice.ts";
import {
    getResponseResultEvaluation,
    getSummaryDataEvaluation,
    getTableDataEvaluation,
    setTableData as dispatchTableData
} from "../../../reducers/evaluationSlice.ts";
import {setSummaryData as dispatchSummaryData} from "../../../reducers/evaluationSlice.ts";
import {setResponseResult as dispatchResponseResult} from "../../../reducers/evaluationSlice.ts";
import ResponseTable2 from "../tables/ResponseTable2.tsx";
import ResponseTable2_2 from "../tables/ResponseTable2_2.tsx";
import ResultCard from "../results/ResultCard.tsx";

interface RowData {
    w: string;
    IS: string;
    wIS: string;
    CC: string;
    wCC: string;
}

interface Props {
    onCompleteStep: () => void
}

const EvaluationHierarchyConsistency: React.FC<Props> = ({onCompleteStep}) => {
    const markId = "Шаг 3"
    const dispatch = useDispatch();
    const initialTableData = useSelector(getTableDataEvaluation);
    const initialSummaryData = useSelector(getSummaryDataEvaluation);
    const initialResponseResult = useSelector(getResponseResultEvaluation);
    const criterias = useSelector(getMainCriterias)
    const [disabledInput, setDisabledInput] =  useState(false)
    const rows = [...criterias];

    const [tableData, setTableData] = useState<{ [key: string]: RowData }>(
        rows.reduce((acc, row) => {
            acc[row] = { w: "", IS: "", wIS: "", CC: "", wCC: "" };
            return acc;
        }, {} as { [key: string]: RowData })
    );

    const [summaryData, setSummaryData] = useState({
        m: "",
        tilda_m: "",
        osi: "",
    });

    const [sendData, {
        data: response,
        isLoading,
        isSuccess,
        error
    }] = useCheckStepMutation();
    
    const [openNextStep] = useOpenNextStepMutation()

    useEffect(() => {
        openNextStep(2)
    }, []);

    useEffect(() => {
        if (initialTableData) {
            setTableData(initialTableData);
        }
        if (initialSummaryData) {
            setSummaryData(initialSummaryData);
        }
        if(initialTableData && initialSummaryData) {
            setDisabledInput(true)
        }
    }, []);

    useEffect(() => {
        if (response) {
            dispatch(addOrUpdateMark({id: markId, mark: response.percentage, maxMark: response.max_mark}))
        }
    }, [dispatch, response]);
    
    const handleInputChange = (rowName: string, field: keyof RowData, value: string) => {
        setTableData((prev) => ({
            ...prev,
            [rowName]: { ...prev[rowName], [field]: value },
        }));
    };

    const handleSummaryChange = (field: string, value: string) => {
        setSummaryData((prev) => ({
            ...prev,
            [field]: value,
        }));
    };

    const handleSend = async () => {
        try {
            const transformedData = {
                w_is: rows.map((row) => parseFloat(tableData[row].wIS) || 0),
                w_cc: rows.map((row) => parseFloat(tableData[row].wCC) || 0),
                m: parseFloat(summaryData.m) || 0,
                tilda_m: parseFloat(summaryData.tilda_m) || 0,
                osi: parseFloat(summaryData.osi) || 0,
            };


            const result = await sendData({
                stepId: "5",
                params: {},
                body: transformedData,
            });

            dispatch(dispatchResponseResult(result.data.result))

            message.success("Данные успешно отправлены!");
        } catch (error) {
            console.error("Ошибка при отправке данных:", error);
            message.error("Ошибка при отправке данных!");
        }
    };

    useEffect(() => {
        if (error) {
            console.error(error)
            message.error(JSON.stringify(error));
        }
    }, [error]);
    
    if (isSuccess) {
        dispatch(dispatchTableData(tableData));
        dispatch(dispatchSummaryData(summaryData));
        onCompleteStep();
    }

    const columns = [
        {
            title: "",
            dataIndex: "rowName",
            key: "rowName",
        },
        {
            title: "w*ИС",
            dataIndex: "wIS",
            key: "wIS",
            render: (_: any, record: any) => (
                <Input
                    readOnly={disabledInput}
                    value={tableData[record.rowName]?.wIS || ""}
                    onChange={(e) => handleInputChange(record.rowName, "wIS", e.target.value)}
                />
            ),
        },
        {
            title: "w*CC",
            dataIndex: "wCC",
            key: "wCC",
            render: (_: any, record: any) => (
                <Input
                    readOnly={disabledInput}
                    value={tableData[record.rowName]?.wCC || ""}
                    onChange={(e) => handleInputChange(record.rowName, "wCC", e.target.value)}
                />
            ),
        },
    ];

    const dataSource = rows.map((rowName) => ({ key: rowName, rowName }));

    return (
        <Space direction="vertical" style={{ width: "100%" }}>
            <Table columns={columns} dataSource={dataSource} pagination={false} />
            <Space direction="vertical">
                <Space>
                    <span>M:</span>
                    <Input
                        readOnly={disabledInput}
                        value={summaryData.m}
                        onChange={(e) => handleSummaryChange("m", e.target.value)}
                    />
                </Space>
                <Space>
                    <span>~M:</span>
                    <Input
                        readOnly={disabledInput}
                        value={summaryData.tilda_m}
                        onChange={(e) => handleSummaryChange("tilda_m", e.target.value)}
                    />
                </Space>
                <Space>
                    <span>ОСИ:</span>
                    <Input
                        readOnly={disabledInput}
                        value={summaryData.osi}
                        onChange={(e) => handleSummaryChange("osi", e.target.value)}
                    />
                </Space>
            </Space>
            <Button
                type="primary"
                onClick={handleSend}
                loading={isLoading}
                disabled={isLoading || disabledInput}
                block
            >
                Отправить
            </Button>
            {initialResponseResult && <>
                <ResponseTable2 data={initialResponseResult} />
                <ResponseTable2_2 summaryData={initialResponseResult} />
            </>}
            <ResultCard markId={markId}/>
        </Space>
    );
};

export default EvaluationHierarchyConsistency;
