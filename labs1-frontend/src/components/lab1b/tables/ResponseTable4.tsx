import React from "react";
import { Table } from "antd";
import {renderCell} from "./util.tsx";

interface ResponseTable4Props {
    data: { val: number; is_right: boolean }[];  // Массив объектов с val и is_right
    alternatives: string[];
}

const ResponseTable4: React.FC<ResponseTable4Props> = ({ data, alternatives }) => {
    // Маппинг для данных таблицы
    const columns = [
        {
            title: "Альтернатива",
            dataIndex: "alternative",
            key: "alternative",
        },
        {
            title: "Значение",
            dataIndex: "value",
            key: "value",
            render: renderCell,  // Отображаем только поле "val"
        },
    ];

    // Преобразуем данные в формат, подходящий для отображения в таблице
    const dataSource = data.map((row, rowIndex) => ({
        key: rowIndex,
        alternative: alternatives[rowIndex] || `Альтернатива ${rowIndex + 1}`,  // Используем альтернативы
        value: row,  // Передаем весь объект для поля "Значение"
    }));

    return <Table dataSource={dataSource}
                  columns={columns}
                  pagination={false}
                  bordered
                  style={{ maxWidth: "400px" }}


    />;
};

export default ResponseTable4;
