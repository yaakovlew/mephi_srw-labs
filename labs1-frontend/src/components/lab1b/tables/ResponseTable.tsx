import React from "react";
import { Table } from "antd";
import type { ColumnsType } from "antd/es/table";
import {renderCell} from "./util.tsx";

interface TableDataInterface {
    key: string;
    X: string;
    w: string;
    Mw: string;
    Lw: string;
    Lmax?: string;
    IS?: string;
    OS?: string;
    is_right: boolean;
}

const ResponseTable: React.FC<{ data: TableDataInterface[] }> = ({ data }) => {


    const columns: ColumnsType<TableDataInterface> = [
        {
            title: "X",
            dataIndex: "X",
            key: "X",
            render: renderCell,
        },
        {
            title: "w",
            dataIndex: "w",
            key: "w",
            render: renderCell,
        },
        {
            title: "M*w",
            dataIndex: "Mw",
            key: "Mw",
            render: renderCell,
        },
        {
            title: "λ*w",
            dataIndex: "Lw",
            key: "Lw",
            render: renderCell,
        },
        {
            title: "λmax",
            dataIndex: "Lmax",
            key: "Lmax",
            render: renderCell,
        },
        {
            title: "ИС",
            dataIndex: "IS",
            key: "IS",
            render: renderCell,
        },
        {
            title: "ОС",
            dataIndex: "OS",
            key: "OS",
            render: renderCell,
        },
    ];

    return (
        <Table columns={columns}
               dataSource={data}
               pagination={false}
               bordered
               title={() => <div style={{ fontWeight: "bold", fontSize: "16px", textAlign: "center" }}>Ответы</div>}
        />
    );
};

export default ResponseTable;
