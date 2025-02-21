import React, { useEffect, useState } from "react";
import { useSelector } from "react-redux";
import {Space} from "antd";
import {getCriterias} from "../../../../reducers/criteriasSlice.ts";
import {TableDataInterface} from "../TableData.tsx";
import PairwiseComparisonCmatricesLvl2Item from "./PairwiseComparisonCmatricesLvl2Item.tsx";


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

const PairwiseComparisonCmatricesLvl2: React.FC<Props> = ({onCompleteStep}) => {
    const criterias = useSelector(getCriterias);

    const [matrices, setMatrices] = useState<string[][][]>(criterias.map((criteria) =>
        createZeroMatrix(criteria.extra.length)
    ));
    const [isEditable, setIsEditable] = useState<boolean[]>(Array(criterias.length).fill(true));
    const [isSuccess, setIsSuccess] = useState(new Set())

    const [data, setData] = useState<TableDataInterface[][]>(
        criterias.map((criteria) =>(
            Array.from({ length: criteria.extra.length }, (_, i) => ({
                key: i.toString(),
                X: "",
                w: "",
                Mw: "",
                Lw: "",
                Lmax: i === 0 ? "" : undefined, // Только для первой строки
                IS: i === 0 ? "" : undefined,  // Только для первой строки
                OS: i === 0 ? "" : undefined,  // Только для первой строки
            }))
        ))
    );

    useEffect(() => {
        if(isSuccess.size === isEditable.length) {
            onCompleteStep()
        }
    }, [isEditable.length, isSuccess.size, onCompleteStep]);


    const updateMatrix = (index: number) => (newMatrix: string[][]) => {
        setMatrices((prevMatrices) =>
            prevMatrices.map((prevMatrix, i) => (i === index ? newMatrix : prevMatrix))
        );
    }

    const updateEditable = (index: number) => (newEditable: boolean) => {
        setIsEditable((prevArr) =>
            prevArr.map((prevEditable, i) => (i === index ? newEditable : prevEditable))
        );
    }

    const updateData = (index: number) => (newData: TableDataInterface[]) => {
        setData((prevData) =>
            prevData.map((prevDataItem, i) => (i === index ? newData : prevDataItem))
        );
    }

    const updateIsSuccess = (index: number) => {
        setIsSuccess(prevSet => {
            const newSet = new Set(prevSet);
            newSet.add(index);
            return newSet;
        } )
    }
    

    return (
        <Space direction="vertical" size="large">
            {
                matrices.map((matrix, i) => (
                    <>
                        <PairwiseComparisonCmatricesLvl2Item
                            i={i}
                            matrix={matrix}
                            criteria={criterias[i]}
                            updateMatrix={updateMatrix(i)}
                            isEditable={isEditable[i]}
                            updateEditable={updateEditable(i)}
                            data={data[i]}
                            updateData={updateData(i)}
                            updateIsSuccess={() => updateIsSuccess(i)}
                        />
                    </>

                ))
            }
        </Space>
    );
};

export default PairwiseComparisonCmatricesLvl2;
