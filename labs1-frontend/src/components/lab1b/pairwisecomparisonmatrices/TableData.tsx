import React, {Dispatch, SetStateAction} from "react";
import { Table, Input } from "antd";
import type { ColumnsType } from "antd/es/table";

interface Table2Props {
    data: TableDataInterface[];
    setData: Dispatch<SetStateAction<TableDataInterface[]>>;
}

export interface TableDataInterface {
    key: string;
    X: string;
    w: string;
    Mw: string;
    Lw: string;
    Lmax?: string;
    IS?: string;
    OS?: string;
}

const TableData: React.FC<Table2Props> = ({ data, setData }) => {

    const handleInputChange = (rowIndex: number, column: keyof TableDataInterface, value: string) => {
        const newData = [...data];
        newData[rowIndex][column] = value;
        setData(newData);
    };

    const columns: ColumnsType<TableDataInterface> = [
        {
            title: "X",
            dataIndex: "X",
            key: "X",
            render: (_, record, rowIndex) => (
                <Input
                    value={record.X}
                    onChange={(e) => handleInputChange(rowIndex, "X", e.target.value)}
                />
            ),
        },
        {
            title: "w",
            dataIndex: "w",
            key: "w",
            render: (_, record, rowIndex) => (
                <Input
                    value={record.w}
                    onChange={(e) => handleInputChange(rowIndex, "w", e.target.value)}
                />
            ),
        },
        {
            title: "M*w",
            dataIndex: "Mw",
            key: "Mw",
            render: (_, record, rowIndex) => (
                <Input
                    value={record.Mw}
                    onChange={(e) => handleInputChange(rowIndex, "Mw", e.target.value)}
                />
            ),
        },
        {
            title: "λ*w",
            dataIndex: "Lw",
            key: "Lw",
            render: (_, record, rowIndex) => (
                <Input
                    value={record.Lw}
                    onChange={(e) => handleInputChange(rowIndex, "Lw", e.target.value)}
                />
            ),
        },
        {
            title: "λmax",
            dataIndex: "Lmax",
            key: "Lmax",
            render: (_, record, rowIndex) =>
                rowIndex === 0 ? (
                    <Input
                        value={record.Lmax}
                        onChange={(e) => handleInputChange(rowIndex, "Lmax", e.target.value)}
                    />
                ) : null,
        },
        {
            title: "ИС",
            dataIndex: "IS",
            key: "IS",
            render: (_, record, rowIndex) =>
                rowIndex === 0 ? (
                    <Input
                        value={record.IS}
                        onChange={(e) => handleInputChange(rowIndex, "IS", e.target.value)}
                    />
                ) : null,
        },
        {
            title: "ОС",
            dataIndex: "OS",
            key: "OS",
            render: (_, record, rowIndex) =>
                rowIndex === 0 ? (
                    <Input
                        value={record.OS}
                        onChange={(e) => handleInputChange(rowIndex, "OS", e.target.value)}
                    />
                ) : null,
        },
    ];

    return (
        <Table
            columns={columns}
            dataSource={data}
            pagination={false}
            bordered
        />
    );
};

export default TableData;
