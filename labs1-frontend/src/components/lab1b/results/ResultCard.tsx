import React from 'react';
import { useSelector } from 'react-redux';
import {Table} from 'antd';
import {getMark} from "../../../reducers/markSlice.ts";


interface Props {
    markId: string
}

const ResultCard: React.FC<Props> = ({markId}) => {
    const mark = useSelector(getMark(markId));

    if (!mark) {
        return null;
    }
    const columns = [
        {
            title: 'Полученная оценка',
            dataIndex: 'mark',
            key: 'mark',
            align: 'center',
        },
        {
            title: 'Максимальная оценка',
            dataIndex: 'maxMark',
            key: 'maxMark',
            align: 'center',
        },
    ];


    const data = [
        {
            key: `mark_${mark.id}`,
            mark: mark.mark,
            maxMark: mark.maxMark,
        },
    ];

    return (
        <Table
            dataSource={data}
            columns={columns}
            bordered
            pagination={false}
            size="small"
            style={{ margin: '20px 0', width: '60%' }}
        />
    );
};

export default ResultCard;
