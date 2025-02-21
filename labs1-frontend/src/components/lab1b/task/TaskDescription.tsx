import React from 'react';
import { Descriptions } from 'antd';

interface TaskDescriptionProps {
    taskName: string;
}

const TaskDescription: React.FC<TaskDescriptionProps> = ({ taskName }) => {
    return (
        <Descriptions bordered size="small" style={{ margin: '20px auto', width: '60%' }}>
            <Descriptions.Item label="Задание">{taskName}</Descriptions.Item>
        </Descriptions>
    );
};

export default TaskDescription;
