import React, { useEffect, useState } from "react";
import { useSelector } from "react-redux";
import {Space} from "antd";
import {getAlternatives} from "../../../reducers/alternativesSlice.ts";
import {getCriterias} from "../../../reducers/criteriasSlice.ts";
import {TableDataInterface} from "../pairwisecomparisonmatrices/TableData.tsx";
import ProcessingQualitativeCriteriaItem from "./ProcessingQualitativeCriteriaItem.tsx";


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

const ProcessingQualitativeCriteria: React.FC<Props> = ({onCompleteStep}) => {
    const alternatives = useSelector(getAlternatives);
    const criteriasAll = useSelector(getCriterias);

    const criterias = criteriasAll.flatMap(section =>
        section.extra.filter(sub => !sub.isCount).map(sub => sub.name)
    )

    const [matrices, setMatrices] = useState<string[][][]>(criterias.map((criteria) =>
        createZeroMatrix(alternatives.length)
    ));
    const [isEditable, setIsEditable] = useState<boolean[]>(Array(criterias.length).fill(true));
    const [isSuccess, setIsSuccess] = useState(new Set());


    const [data, setData] = useState<TableDataInterface[][]>(
        criterias.map((criteria) =>(
            Array.from({ length: alternatives.length }, (_, i) => ({
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


    const updateMatrix = (index) => (newMatrix) => {
        setMatrices((prevMatrices) =>
            prevMatrices.map((prevMatrix, i) => (i === index ? newMatrix : prevMatrix))
        );
    }

    const updateEditable = (index) => (newEditable) => {
        setIsEditable((prevArr) =>
            prevArr.map((prevEditable, i) => (i === index ? newEditable : prevEditable))
        );
    }

    const updateData = (index) => (newData) => {
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
                    <ProcessingQualitativeCriteriaItem
                        i={i}
                        criteria={criterias[i]}
                        alternatives={alternatives}
                        matrix={matrix}
                        updateMatrix={updateMatrix(i)}
                        isEditable={isEditable[i]}
                        updateEditable={updateEditable(i)}
                        data={data[i]}
                        updateData={updateData(i)}
                        updateIsSuccess={() => updateIsSuccess(i)}
                    />
                ))
            }
        </Space>
    );
};

export default ProcessingQualitativeCriteria;
