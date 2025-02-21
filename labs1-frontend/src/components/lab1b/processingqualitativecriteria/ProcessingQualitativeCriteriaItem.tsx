import TableMatrix from "./TableMatrix.tsx";
import TableData from "../pairwisecomparisonmatrices/TableData.tsx";
import React, {useEffect, useState} from "react";
import {Button, message} from "antd";
import {reverseTransformDataLvl1, transformDataLvl1} from "../../../reducers/lvl1Slice.ts";
import {useCheckStepMutation} from "../../../api/lab1BApi.ts";
import {addOrUpdateMark} from "../../../reducers/markSlice.ts";
import {useDispatch, useSelector} from "react-redux";
import {
    getDataQualitative,
    getMatrixQualitative,
    getResponseResultQualitative
} from "../../../reducers/qualitativeSlice.ts";
import {isDataNotEmpty, isMatrixNotEmpty} from "../../../util/isMatrixNotEmpty.ts";
import {setData, setMatrix, setResponseResult} from "../../../reducers/qualitativeSlice.ts";
import ResponseTable from "../tables/ResponseTable.tsx";
import ResultCard from "../results/ResultCard.tsx";

interface Props {
    i: number,
    criteria: any,
    alternatives: any,
    matrix: any,
    updateMatrix: any,
    isEditable: any,
    updateEditable: any,
    data: any,
    updateData: any,
    updateIsSuccess: () => void;

}

const ProcessingQualitativeCriteriaItem: React.FC<Props> = (props) => {
    const {
        i,
        criteria,
        alternatives,
        matrix,
        updateMatrix,
        isEditable,
        updateEditable,
        data,
        updateData,
        updateIsSuccess,
    } = props;

    const markId = `Шаг 5 (${i + 1})`;

    const dispatch = useDispatch();
    const initialMatrix = useSelector(getMatrixQualitative(i));
    const initialData = useSelector(getDataQualitative(i));
    const initialResponseResult = useSelector(getResponseResultQualitative(i));

    const [sendData, {
        data: response,
        isSuccess,
        error
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
    }, [dispatch, response]);

    const onSend = async () => {
        const requestBody = {
            ...transformDataLvl1(data),
            matrix: matrix.map((row) => row.map((cell) => parseFloat(cell) || 0)),
        };

        const result = await sendData({ stepId: "quality", params: { step: i + 1 }, body: requestBody })

        if (result) {
            updateIsSuccess();
            dispatch(setMatrix({index: i, payload: matrix}));
            dispatch(setData({index: i, payload: data}));
            dispatch(setResponseResult({index: i, payload: result.data.result}));
        }
    };

    return (
        <>
            <TableMatrix
                key={`qualitative_m${i}`}
                criteria={criteria}
                alternatives={alternatives}
                matrix={matrix}
                setMatrix={updateMatrix}
                isEditable={isEditable}
                setIsEditable={updateEditable}

            />
            <TableData key={`qualitative_d${i}`}
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

export default ProcessingQualitativeCriteriaItem