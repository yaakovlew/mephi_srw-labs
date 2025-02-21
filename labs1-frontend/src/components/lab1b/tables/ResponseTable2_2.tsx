import {Space} from "antd";
import React from "react";
import {renderCell} from "./util.tsx";

interface Props {
    summaryData: any
}

const ResponseTable2_2: React.FC<Props> = ({summaryData}) => {

    return (
        <Space direction="vertical">
            <Space>
                <span>M:</span>
                {renderCell(summaryData.m)}
            </Space>
            <Space>
                <span>~M:</span>
                {renderCell(summaryData.tilda_m)}
            </Space>
            <Space>
                <span>ОСИ:</span>
                {renderCell(summaryData.osi)}
            </Space>
        </Space>
    )
}

export default ResponseTable2_2