import React, {Dispatch, SetStateAction} from "react";
import { useSelector } from "react-redux";
import { getMainCriterias } from "../../../../reducers/criteriasSlice.ts";
import type { ColumnsType } from "antd/es/table";
import { Input, Table, Button, message } from "antd";
import {useCheckMatrixMutation} from "../../../../api/lab1BApi.ts";

interface TableData {
    key: string;
    [key: string]: string | JSX.Element;
}

interface Props {
    matrix: string[][];
    setMatrix: (newMatrix: string[][]) => void;
    isEditable: boolean;
    setIsEditable: Dispatch<SetStateAction<boolean>>;
    disabledInput: boolean;
}

const TableMatrix: React.FC<Props> = (props) => {
    const {
        matrix,
        setMatrix,
        isEditable,
        setIsEditable ,
        disabledInput,
    } = props
    const criteria = useSelector(getMainCriterias);
    const [checkMatrix] = useCheckMatrixMutation();


    const handleInputChange = (rowIndex: number, colIndex: number, value: string) => {
        if (!isEditable) return;

        const newMatrix = matrix.map((row, rIdx) =>
            row.map((cell, cIdx) => (rIdx === rowIndex && cIdx === colIndex ? value : cell))
        );

        setMatrix(newMatrix);
    };

    const handleSendMatrix = async () => {
        try {
            const numberMatrix = matrix.map(row => row.map(cell => parseFloat(cell) || 0));
            const response = await checkMatrix({ body: {matrix: numberMatrix} }).unwrap();

            if (response === true) {
                message.success("Матрица валидная");
                setIsEditable(false);
            } else {
                message.error("Матрица не валидная");
            }
        } catch (error) {
            console.error(error);
            message.error("Ошибка при отправке матрицы!");
        }
    };

    const columns: ColumnsType<TableData> = [
        { title: "", dataIndex: "criteria", key: "criteria", fixed: "left" },
        ...criteria.map((criterion, colIndex) => ({
            title: criterion,
            dataIndex: criterion,
            key: criterion,
            render: (_, record: TableData, rowIndex: number) => (
                <Input
                    value={matrix[rowIndex][colIndex]}
                    onChange={(e) => handleInputChange(rowIndex, colIndex, e.target.value)}
                    disabled={!isEditable}
                    readOnly={disabledInput}
                />
            ),
        })),
    ];

    const dataSource: TableData[] = criteria.map((criterion, rowIndex) => ({
        key: `${criterion}-${rowIndex}`,
        criteria: criterion,
        ...Object.fromEntries(criteria.map((_, colIndex) => [criteria[colIndex], matrix[rowIndex][colIndex]])),
    }));

    return (
        <>
            <Table
                columns={columns}
                dataSource={dataSource}
                pagination={false}
                bordered
            />
            <Button
                type="primary"
                onClick={handleSendMatrix}
                disabled={!isEditable || disabledInput}
                style={{ marginTop: "16px" }}
            >
                Проверить матрицу
            </Button>
        </>
    );
};

export default TableMatrix;
