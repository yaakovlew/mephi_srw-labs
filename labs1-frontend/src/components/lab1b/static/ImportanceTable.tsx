import { Table } from 'antd';

const ImportanceTable = () => {
    const data = [
        { key: 1, value: 1, definition: 'Равная важность элементов' },
        { key: 2, value: 3, definition: 'Умеренное превосходство одного элемента над другим' },
        { key: 3, value: 5, definition: 'Существенное или сильное превосходство одного элемента над другим' },
        { key: 4, value: 7, definition: 'Значительное превосходство одного элемента над другим' },
        { key: 5, value: 9, definition: 'Очень сильное превосходство одного элемента над другим' },
        { key: 6, value: '2, 4, 6, 8', definition: 'Промежуточные решения между двумя соседними суждениями, применяются в компромиссном случае' },
        { key: 7, value: '0.14, 0.2, 0.33, 0.5', definition: 'Обратные величины, полученные при сравнении второго элемента с первым, означают ту или иную степень превосходства второго элемента над первым' }
    ];

    const columns = [
        {
            title: 'Значение важности',
            dataIndex: 'value',
            key: 'value',
            width: '30%',
        },
        {
            title: 'Определение',
            dataIndex: 'definition',
            key: 'definition',
            width: '70%',
        },
    ];

    return (
        <Table
            dataSource={data}
            columns={columns}
            bordered
            pagination={false}
            style={{ margin: '20px 0' }}
        />
    );
};

export default ImportanceTable;
