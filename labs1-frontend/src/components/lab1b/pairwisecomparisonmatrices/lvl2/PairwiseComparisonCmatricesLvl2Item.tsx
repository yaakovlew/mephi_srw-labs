import Table2 from "./TableMatrix.tsx";
import TableData, {TableDataInterface} from "../TableData.tsx";
import React, {Dispatch, SetStateAction, useEffect, useState} from "react";
import {Button, message} from "antd";
import {useCheckStepMutation} from "../../../../api/lab1BApi.ts";
import {reverseTransformDataLvl1, transformDataLvl1} from "../../../../reducers/lvl1Slice.ts";
import {addOrUpdateMark} from "../../../../reducers/markSlice.ts";
import {useDispatch, useSelector} from "react-redux";
import ResponseTable from "../../tables/ResponseTable.tsx";
import {
    getDataLvl2,
    getMatrixLvl2,
    getResponseResultLvl2, lvl2SliceInterface,
    setData,
    setMatrix,
    setResponseResult
} from "../../../../reducers/lvl2Slice.ts";
import {isDataNotEmpty, isMatrixNotEmpty} from "../../../../util/isMatrixNotEmpty.ts";
import ResultCard from "../../results/ResultCard.tsx";

interface Props  {
    i: number;
    matrix: string[][];
    criteria: string;
    updateMatrix: (newMatrix: string[][]) => void;
    isEditable: boolean;
    updateEditable: Dispatch<SetStateAction<boolean>>;
    data: TableDataInterface[];
    updateData: Dispatch<SetStateAction<TableDataInterface[]>>;
    updateIsSuccess: () => void;
}

const PairwiseComparisonCmatricesLvl2Item: React.FC<Props> = (props) => {
    const {
        i,
        matrix,
        criteria,
        updateMatrix,
        isEditable,
        updateEditable,
        data,
        updateData,
        updateIsSuccess
    } = props;

    const markId =  `Шаг 2 (${i + 1})`;

    const dispatch = useDispatch();
    const initialMatrix = useSelector(getMatrixLvl2(i));
    const initialData = useSelector(getDataLvl2(i));
    const initialResponseResult = useSelector(getResponseResultLvl2(i));

    const [sendData, {
        isSuccess,
        error,
        data: response
    }] = useCheckStepMutation();
    const [disabledInput, setDisabledInput] =  useState(false)


    useEffect(() => {
        if (isMatrixNotEmpty(initialMatrix)) {
            updateMatrix(initialMatrix);
        }
        if (isDataNotEmpty(initialData)) {
            updateData(initialData);
        }
        if (isMatrixNotEmpty(initialMatrix) && isDataNotEmpty(initialData)) {
            setDisabledInput(true);
            updateEditable(false);
            updateIsSuccess();
        }
    }, [initialData, initialMatrix]);

    useEffect(() => {
        if (error) {
            console.error(error);
            message.error(JSON.stringify(error));
        }
    }, [error]);

    useEffect(() => {
        if (response) {
            dispatch(addOrUpdateMark({id: markId, mark: response.percentage, maxMark: response.max_mark}))
        }
    }, [dispatch, i, response]);


    const onSend = async () => {
        const requestBody = {
            ...transformDataLvl1(data),
            matrix: matrix.map((row) => row.map((cell) => parseFloat(cell) || 0)),
        };

        const result = await sendData({ stepId: "2-4", params: { step: i + 1 }, body: requestBody })

        if (result) {
            updateIsSuccess();
            dispatch(setMatrix({index: i, payload: matrix}));
            dispatch(setData({index: i, payload: data}));
            dispatch(setResponseResult({index: i, payload: result.data.result}));
        }
    };


    return (
        <>
            <Table2
                key={`lvl2_m${i}`}
                criteria={criteria}
                matrix={matrix}
                setMatrix={updateMatrix}
                isEditable={isEditable}
                setIsEditable={updateEditable}
            />
            <TableData key={`lvl2_d${i}`}
                       data={data}
                       setData={updateData}
            />
            <Button
                type="primary"
                onClick={onSend}
                disabled={isEditable || isSuccess || disabledInput}
                style={{ marginTop: "16px" }}
                block
            >
                Отправить
            </Button>
            {initialResponseResult && (
                <ResponseTable data={reverseTransformDataLvl1(initialResponseResult)} />
            )}
            <ResultCard markId={markId}/>
        </>
    )
}

export default PairwiseComparisonCmatricesLvl2Item