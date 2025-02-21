import React from "react";
import { Table, Space, Typography } from "antd";
import { useSelector } from "react-redux";
import {getMarks} from "../../../reducers/markSlice.ts";

const { Title } = Typography;

const Results: React.FC = () => {
    const marks = useSelector(getMarks);

    const totalSum = marks.reduce((acc, item) => acc + item.mark, 0);

    
    const columns = [
        {
            title: "ID",
            dataIndex: "id",
            key: "id",
        },
        {
            title: "Оценка",
            dataIndex: "mark",
            key: "mark",
        },
        {
            title: "Макс оценка",
            dataIndex: "maxMark",
            key: "maxMark",
        },
    ];

    const mergedMarks = Object.values(
        marks.reduce((acc, step) => {
            if (step.id === "Приоритеты") {
                acc["Приоритеты"] = step;
                return acc;
            }

            const baseId = step.id.match(/Шаг \d+/)[0];
            if (!acc[baseId]) {
                acc[baseId] = { id: baseId, mark: 0, maxMark: 0 };
            }

            acc[baseId].mark += step.mark;
            acc[baseId].maxMark += step.maxMark;

            return acc;
        }, {})
    );

    const dataSource = mergedMarks.map((mark) => ({
        key: mark.id,
        id: mark.id,
        mark: mark.mark,
        maxMark: mark.maxMark,
    }));

    return (
        <Space direction="vertical" style={{ width: "100%" }}>
            <Title level={3}>Результаты</Title>
            <Table
                dataSource={dataSource}
                columns={columns}
                pagination={false}
                bordered
                summary={() => (
                    <Table.Summary fixed>
                        <Table.Summary.Row>
                            <Table.Summary.Cell colSpan={2} style={{ textAlign: "right" }} index={0}>
                                <strong>Итого: {totalSum}</strong>
                            </Table.Summary.Cell>
                        </Table.Summary.Row>
                    </Table.Summary>
                )}
            />
        </Space>
    );
};

export default Results;
