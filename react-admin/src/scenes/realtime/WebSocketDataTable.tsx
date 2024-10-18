// WebSocketDataTable.tsx
import React, { useState, useEffect } from 'react';
import websocketService, { WebSocketData } from '../../service/websocketService';

const WebSocketDataTable: React.FC = () => {
    const [tableData, setTableData] = useState<WebSocketData[]>([]);

    useEffect(() => {
        const handleData = (data: WebSocketData) => {
            setTableData(prevData => [...prevData, data]);
        };

        websocketService.addDataListener(handleData);

        return () => {
            websocketService.removeDataListener(handleData);
        };
    }, []);

    return (
        <table>
            <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Status</th>
                <th>Created At</th>
            </tr>
            </thead>
            <tbody>
            {tableData.map((row, index) => (
                <tr key={index}>
                    <td>{row.id}</td>
                    <td>{row.name}</td>
                    <td>{row.status}</td>
                    <td>{row.createdAt}</td>
                </tr>
            ))}
            </tbody>
        </table>
    );
};

export default WebSocketDataTable;
