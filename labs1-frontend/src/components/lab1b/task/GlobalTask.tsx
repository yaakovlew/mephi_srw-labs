import React from "react";
import {Card, Descriptions} from "antd";

interface Props {
    data: string
}

const GlobalTask: React.FC<Props> = ({data}) => {

    return (
        <Descriptions bordered  style={{ margin: '20px auto', width: '60%' }}>
            <Descriptions.Item label="Вариант">
                {data}
            </Descriptions.Item>
        </Descriptions>
        // <div style={{display: 'flex', justifyContent: "center", marginTop: "2rem"}}>
        //     <Card title="Вариант" style={{ width: 300 }}>
        //         <p style={{}}>{data}</p>
        //     </Card>
        // </div>
    )
}

export default GlobalTask