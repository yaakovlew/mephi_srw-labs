import React from 'react';
import { Table } from 'antd';
import {renderCell} from "./util.tsx";


// const data = [
//     {
//         key: '1',
//         w_is: { val: -0.24, is_right: false },
//         w_cc: { val: 0.19, is_right: false }
//     },
//     {
//         key: '2',
//         w_is: { val: -0.22, is_right: false },
//         w_cc: { val: 0.19, is_right: false }
//     },
//     {
//         key: '3',
//         w_is: { val: -0.24, is_right: false },
//         w_cc: { val: 0.19, is_right: false }
//     }
// ];

interface Props {
    data: any;
}


const ResponseTable2: React.FC<Props> = ({data}) => {
    const transformedData = data.w_is.map((item, index) => ({
        key: `${index + 1}`,
        w_is: item,
        w_cc: data.w_cc[index],
        m: data.m,
        tilda_m: data.tilda_m,
        osi: data.osi,
    }));

    const columns = [
        {
            title: 'w*ИС',
            dataIndex: 'w_is',
            key: 'w_is',
            render: renderCell
        },
        {
            title: 'w*CC',
            dataIndex: 'w_cc',
            key: 'w_cc',
            render: renderCell
        }
    ];

    return (
        <Table
            columns={columns}
            dataSource={transformedData}
            pagination={false}
            bordered
        />
    )
}


export default ResponseTable2;
