// WebSocketStatus.tsx
import React, {useEffect, useState} from 'react';
import websocketService, {Status} from '../service/websocketService';

const WebSocketStatus: React.FC = () => {
    const [status, setStatus] = useState<Status>('Down');

    useEffect(() => {
        const handleStatusChange = (newStatus: Status) => {
            setStatus(newStatus);
        };

        websocketService.addStatusListener(handleStatusChange);
        websocketService.connect();

        return () => {
            websocketService.removeStatusListener(handleStatusChange);
        };
    }, []);

    return (

        <div className="status-container">
            <div className={`status-indicator ${status === 'Operational' ? 'green' : 'red'}`}></div>
            <p>WebSocket: {status}</p>
        </div>

    );
};

export default WebSocketStatus;
