import React, {useEffect, useState} from 'react';
import websocketService, {WebSocketCounter} from '../service/websocketService';
import StatBox from "./StatBox";
import {tokens} from "../theme";
import {useTheme} from "@mui/material";
import {PointOfSale} from "@mui/icons-material";

const ApiCalls: React.FC = () => {
    const [counter200, setCounter200] = useState<number>(0);
    const [counter400, setCounter400] = useState<number>(0);
    const theme = useTheme();
    const colors = tokens(theme.palette.mode);

    useEffect(() => {
        const handleCountChange = (newCounter: WebSocketCounter) => {
            setCounter200(newCounter.counter200);
            setCounter400(newCounter.counter400);
            console.log(newCounter)
        };

        websocketService.addCounterAttemptListener(handleCountChange);
        websocketService.connect();

        return () => {
            websocketService.removeCounterAttemptListener(handleCountChange);
        };
    }, []);

    // Calculate progress as a decimal value (0-1)
    const progressValue = counter200 > 0 ? (counter200 *100)/ (counter200 + counter400)  : 0;
    const increase = progressValue / 100


    return (

        <StatBox
            title={counter200 + counter400}
            subtitle="API Call Attempts"
            progress={increase.toFixed(2)}
            increase={`${progressValue.toFixed(0)}%`}
            icon={
                <PointOfSale
                    sx={{color: colors.greenAccent[600], fontSize: "26px"}}
                />
            }
        />
    );
};

export default ApiCalls;
