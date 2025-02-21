import React from 'react';
import { Table } from 'antd';

const MatrixSizeTable = () => {
    const data = [
        { key: 1, value1: 'Размер матрицы', value2: '1', value3: '2', value4: '3', value5: '4', value6: '5', value7: '6', value8: '7', value9: '8', value10: '9', value11: '10' },
        { key: 2, value1: 'СС', value2: '0', value3: '0', value4: '0,58', value5: '0,90', value6: '1,12', value7: '1,24', value8: '1,32', value9: '1,41', value10: '1,45', value11: '1,49' },
    ];

    const columns = [
        { title: '', dataIndex: 'value1', key: 'value1', align: 'center' },
        { title: '', dataIndex: 'value2', key: 'value2', align: 'center' },
        { title: '', dataIndex: 'value3', key: 'value3', align: 'center' },
        { title: '', dataIndex: 'value4', key: 'value4', align: 'center' },
        { title: '', dataIndex: 'value5', key: 'value5', align: 'center' },
        { title: '', dataIndex: 'value6', key: 'value6', align: 'center' },
        { title: '', dataIndex: 'value7', key: 'value7', align: 'center' },
        { title: '', dataIndex: 'value8', key: 'value8', align: 'center' },
        { title: '', dataIndex: 'value9', key: 'value9', align: 'center' },
        { title: '', dataIndex: 'value10', key: 'value10', align: 'center' },
        { title: '', dataIndex: 'value11', key: 'value11', align: 'center' },
    ];

    return (
        <Table
            dataSource={data}
            columns={columns}
            bordered
            pagination={false}
            style={{ margin: '20px 0', textAlign: 'center' }}
            showHeader={false}
        />
    );
};

export default MatrixSizeTable;
