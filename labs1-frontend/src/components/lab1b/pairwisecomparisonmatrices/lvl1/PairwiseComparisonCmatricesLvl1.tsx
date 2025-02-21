import React, {useEffect, useState} from "react";
import {useDispatch, useSelector} from "react-redux";
import {getMainCriterias} from "../../../../reducers/criteriasSlice.ts";
import TableMatrix from "./TableMatrix.tsx";
import TableData, {TableDataInterface} from "../TableData.tsx";
import {Button, message} from "antd";
import {useCheckStepMutation} from "../../../../api/lab1BApi.ts";
import {
    getDataLvl1,
    getMatrixLvl1, getResponseResultLvl1,
    reverseTransformDataLvl1,
    setMatrix as setMatrixRedux, setResponseResult
} from "../../../../reducers/lvl1Slice.ts";
import {setData as setDataRedux} from "../../../../reducers/lvl1Slice.ts";
import {addOrUpdateMark} from "../../../../reducers/markSlice.ts";
import {isDataNotEmpty, isMatrixNotEmpty} from "../../../../util/isMatrixNotEmpty.ts";
import ResponseTable from "../../tables/ResponseTable.tsx";
import ResultCard from "../../results/ResultCard.tsx";



function createZeroMatrix(n: number) {
    const matrix = [];
    for (let i = 0; i < n; i++) {
        const row = new Array(n).fill(0);
        matrix.push(row);
    }
    return matrix;
}

interface Props {
    onCompleteStep: () => void;
}

const PairwiseComparisonCmatricesLvl1: React.FC<Props> = ({onCompleteStep}) => {
    const markId = "Шаг 1"
    const dispatch = useDispatch();
    const initialMatrix = useSelector(getMatrixLvl1);
    const initialData = useSelector(getDataLvl1);
    const responseResult = useSelector(getResponseResultLvl1);
    const [sendData, {
        data: response,
        isSuccess,
        error,
    }] = useCheckStepMutation()
    const [isEditableTable1, setIsEditableTable1] = useState(true);
    const [disabledInput, setDisabledInput] =  useState(false)

    const criteria = useSelector(getMainCriterias)
    const [matrix, setMatrix] = useState<string[][]>(
        createZeroMatrix(criteria.length)
    );

    const [data, setData] = useState<TableDataInterface[]>(
        Array.from({ length: criteria.length }, (_, i) => ({
            key: i.toString(),
            X: "",
            w: "",
            Mw: "",
            Lw: "",
            Lmax: i === 0 ? "" : undefined, // Только для первой строки
            IS: i === 0 ? "" : undefined,  // Только для первой строки
            OS: i === 0 ? "" : undefined,  // Только для первой строки
        }))
    );


    useEffect(() => {
        if (isMatrixNotEmpty(initialMatrix)) {
            setMatrix(initialMatrix);
        }
        if (isDataNotEmpty(initialData)) {
            setData(initialData);
        }
        if (isMatrixNotEmpty(initialMatrix) && isDataNotEmpty(initialData)) {
            setDisabledInput(true)
        }
    }, []);

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


    if (isSuccess) {
        dispatch(setMatrixRedux(matrix));
        dispatch(setDataRedux(data));
        dispatch(setResponseResult(response?.result))
        onCompleteStep();
    }


    const onSend = () => {
        const requestBody = {
            matrix: matrix.map(row => row.map(cell => parseFloat(cell) || 0)),
            X: data.map((row) => parseFloat(row.X) || 0),
            w: data.map((row) => parseFloat(row.w) || 0),
            mw: data.map((row) => parseFloat(row.Mw) || 0),
            lambda_w: data.map((row) => parseFloat(row.Lw) || 0),
            lambda_max: parseFloat(data[0]?.Lmax || "0"),
            is: parseFloat(data[0]?.IS || "0"),
            os: parseFloat((data[0]?.OS || "0").replace("%", "")),
        };

        sendData({stepId: 1, body: requestBody});
    }


    return (
        <>
            <TableMatrix matrix={matrix}
                         setMatrix={setMatrix}
                         isEditable={isEditableTable1}
                         setIsEditable={setIsEditableTable1}
                         disabledInput={disabledInput}
            />
            <TableData data={data} setData={setData} />
            <Button type="primary"
                    onClick={onSend}
                    disabled={isEditableTable1 || disabledInput}
                    block
            >
                Отправить
            </Button>
            {responseResult && <ResponseTable data={reverseTransformDataLvl1(responseResult) }/>}
            <ResultCard markId={markId}/>
        </>
    )
};

export default PairwiseComparisonCmatricesLvl1;
