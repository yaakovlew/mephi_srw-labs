import React, { useEffect, useState } from 'react';

interface Result {
    step: number;
    percentage: number;
    maxMark: number;
}

const Lab1AResultsTable: React.FC = () => {
    const [results, setResults] = useState<Result[]>([]);
    const [totalPercentage, setTotalPercentage] = useState(0);
    const [totalMaxMark, setTotalMaxMark] = useState(0);

    useEffect(() => {
        let totalPercentage = 0;
        let totalMaxMark = 0;
        const resultsArr: Result[] = [];

        for (let step = 1; step <= 6; step++) {
            const key = `lab1a_${step}_result`;
            const storedResult = localStorage.getItem(key);

            if (storedResult) {
                try {
                    const parsedResult = JSON.parse(storedResult);
                    const { percentage, max_mark } = parsedResult;

                    if (percentage !== undefined && max_mark !== undefined) {
                        resultsArr.push({
                            step,
                            percentage,
                            maxMark: max_mark
                        });

                        totalPercentage += percentage;
                        totalMaxMark += max_mark;
                    }
                } catch (error) {
                    console.error(`Ошибка при парсинге для ${key}:`, error);
                }
            }
        }

        setResults(resultsArr);
        setTotalPercentage(totalPercentage);
        setTotalMaxMark(totalMaxMark);
    }, []);

    return (
        <div style={containerStyle}>
            <div style={contentStyle}>
                <h2 style={headerStyle}>Результаты Лабораторной Работы</h2>
                <table style={tableStyle}>
                    <thead>
                    <tr>
                        <th style={tableHeaderStyle}>Шаг</th>
                        <th style={tableHeaderStyle}>Процент / Макс. балл</th>
                    </tr>
                    </thead>
                    <tbody>
                    {results.map(({ step, percentage, maxMark }) => (
                        <tr key={step}>
                            <td style={tableCellStyle}>Шаг {step}</td>
                            <td style={tableCellStyle}>{percentage}/{maxMark}</td>
                        </tr>
                    ))}
                    </tbody>
                </table>
                <div style={totalStyle}>
                    <h3>Итого:</h3>
                    <p>{totalPercentage}/{totalMaxMark} ({((totalPercentage / totalMaxMark) * 100).toFixed(2)}%)</p>
                </div>
            </div>
        </div>
    );
};

const containerStyle = {
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    height: '100vh',
    backgroundColor: '#f0f0f0',
};

const contentStyle = {
    textAlign: 'center',
    padding: '20px',
    backgroundColor: '#fff',
    borderRadius: '10px',
    boxShadow: '0 4px 8px rgba(0, 0, 0, 0.1)',
    width: '80%',
    maxWidth: '800px',
};

const headerStyle = {
    fontSize: '24px',
    fontWeight: 'bold',
    marginBottom: '20px',
    color: '#333',
};

const tableStyle = {
    width: '100%',
    borderCollapse: 'collapse',
    marginBottom: '20px',
    textAlign: 'center',
};

const tableHeaderStyle = {
    padding: '10px',
    backgroundColor: '#f4f4f4',
    borderBottom: '1px solid #ccc',
    fontWeight: 'bold',
    color: '#555',
    textAlign: 'center',
};

const tableCellStyle = {
    padding: '10px',
    borderBottom: '1px solid #ccc',
    color: '#555',
};

const totalStyle = {
    marginTop: '20px',
    fontSize: '18px',
    fontWeight: 'bold',
    color: '#333',
};

export default Lab1AResultsTable;
