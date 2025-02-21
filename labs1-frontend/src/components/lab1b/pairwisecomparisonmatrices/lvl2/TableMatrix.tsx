import React, { Dispatch, SetStateAction } from "react";
import { Input, Table, Button, message } from "antd";
import type { ColumnsType } from "antd/es/table";
import {useCheckMatrixMutation} from "../../../../api/lab1BApi.ts";

interface ExtraItem {
    id: number;
    name: string;
    isCount: boolean;
    isReverse: boolean;
    value: number[];
}

interface Criteria {
    id: number;
    criteria: string;
    extra: ExtraItem[];
}

interface TableData {
    key: string;
    [key: string]: string | JSX.Element;
}

interface Props {
    criteria: Criteria;
    matrix: string[][];
    setMatrix: (newMatrix: string[][]) => void;
    isEditable: boolean;
    setIsEditable: Dispatch<SetStateAction<boolean>>;
}

const Table1: React.FC<Props> = ({ criteria, matrix, setMatrix, isEditable, setIsEditable }) => {
    const extra = criteria.extra;
    const [checkMatrix] = useCheckMatrixMutation();


    const handleInputChange = (rowIndex: number, colIndex: number, value: string) => {
        if (!isEditable) return; // Если таблица неизменяемая, игнорируем ввод
        const newMatrix = matrix.map((row, rIdx) =>
            row.map((cell, cIdx) => (rIdx === rowIndex && cIdx === colIndex ? value : cell))
        );
        setMatrix(newMatrix);
    };

    const handleSendMatrix = async () => {
        try {
            const numberMatrix = matrix.map((row) => row.map((cell) => parseFloat(cell) || 0));
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
        { title: criteria.criteria, dataIndex: "criteria", key: "criteria", fixed: "left" },
        ...extra.map((extraItem, colIndex) => ({
            title: extraItem.name,
            dataIndex: extraItem.name,
            key: extraItem.name,
            render: (_, record: TableData, rowIndex: number) => (
                <Input
                    value={matrix[rowIndex][colIndex]}
                    onChange={(e) => handleInputChange(rowIndex, colIndex, e.target.value)}
                    readOnly={!isEditable}
                />
            ),
        })),
    ];

    const dataSource: TableData[] = extra.map((extraItem, rowIndex) => ({
        key: `${extraItem.name}-${rowIndex}`,
        criteria: extraItem.name,
        ...Object.fromEntries(extra.map((_, colIndex) => [extra[colIndex].name, matrix[rowIndex][colIndex]])),
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
                disabled={!isEditable}
                style={{ marginTop: "16px" }}
            >
                Проверить матрицу
            </Button>
        </>
    );
};

export default Table1;
