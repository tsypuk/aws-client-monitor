import React, { useEffect, useState } from 'react';
import './StatusIndicator.css';

const StatusIndicator: React.FC = () => {
    const [status, setStatus] = useState<'green' | 'red'>('red');

    const checkStatus = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/v1/status'); // Replace with your endpoint
            console.log(response)
            if (response.status === 200) {
                setStatus('green');
            } else {
                setStatus('red');
            }
        } catch (error) {
            console.error('Error fetching status:', error);
            setStatus('red'); // Set red if the request fails
        }
    };

    useEffect(() => {
        // Initial check
        checkStatus();

        // Poll the API every 5 seconds
        const intervalId = setInterval(() => {
            checkStatus();
        }, 5000); // 5 seconds

        // Clean up the interval when the component is unmounted
        return () => clearInterval(intervalId);
    }, []);

    return (
        <div className="status-container">
            <div className={`status-indicator ${status}`}></div>
            <p>Status: {status === 'green' ? 'Operational' : 'Down'}</p>
        </div>
    );
};

export default StatusIndicator;
