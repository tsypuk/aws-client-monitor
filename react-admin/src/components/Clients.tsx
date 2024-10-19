import React, {useEffect, useState} from 'react';
import websocketService, {WebSocketCounter} from '../service/websocketService';
import StatBox from "./StatBox";
import {tokens} from "../theme";
import {useTheme} from "@mui/material";
import {PointOfSale} from "@mui/icons-material";
import PersonAddIcon from "@mui/icons-material/PersonAdd";

const Clients: React.FC = () => {
    const [counter, setCounter] = useState<number>(0);
    const theme = useTheme();
    const colors = tokens(theme.palette.mode);

    useEffect(() => {
        const handleCountChange = (newCounter: number) => {
            setCounter(newCounter)
            console.log(newCounter)
        };

        websocketService.addClientListener(handleCountChange);
        websocketService.connect();

        return () => {
            websocketService.removeClientListener(handleCountChange);
        };
    }, []);

    return (
        <StatBox
            title={counter}
            subtitle="Clients"
            progress="0.30"
            increase="+5%"
            icon={
                <PersonAddIcon
                    sx={{color: colors.greenAccent[600], fontSize: "26px"}}
                />
            }
        />
    );
};

export default Clients;
