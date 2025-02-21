import React from "react";
import { Table } from "antd";

interface Props {
    alternatives: any[];
    tables: any[];
}

const renderCell = (cell) => {
    if (!cell) {
        return null;
    }

    return (
        <div style={{ backgroundColor: cell?.is_right ? "#d4edda" : "#f8d7da" }}>
            {cell?.val}
        </div>
    );
};

const ResponseTable3: React.FC<Props> = ({ alternatives, tables }) => {
    // Функция для формирования источника данных для таблицы
    const dataSource = () => {
        return tables.map((grades, index) => ({
            key: `row-${index}`,
            rowName: `Оценки ${index + 1}`,
            ...alternatives.reduce((acc: any, alt: any, idx: number) => {
                acc[`alt${idx}`] = renderCell(grades[idx] || null);
                return acc;
            }, {}),
        }));
    };

    return (
        <Table
            pagination={false}
            bordered
            dataSource={dataSource()}
            columns={[
                {
                    title: "Оценки",
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
    );
};

export default ResponseTable3;
