import React, {useEffect, useState} from 'react';
import websocketService, {WebSocketCounter} from '../service/websocketService';
import StatBox from "./StatBox";
import {tokens} from "../theme";
import {useTheme} from "@mui/material";
import {PointOfSale} from "@mui/icons-material";
import PersonAddIcon from "@mui/icons-material/PersonAdd";
import TrafficIcon from "@mui/icons-material/Traffic";

const Regions: React.FC = () => {
    const [counter, setCounter] = useState<number>(0);
    const theme = useTheme();
    const colors = tokens(theme.palette.mode);

    useEffect(() => {
        const handleCountChange = (newCounter: number) => {
            setCounter(newCounter)
            console.log(newCounter)
        };

        websocketService.addRegionListener(handleCountChange);
        // websocketService.connect();

        return () => {
            websocketService.removeRegionListener(handleCountChange);
        };
    }, []);

    return (
        <StatBox
            title={counter}
            subtitle="Region Traffic"
            progress="0.80"
            increase="+43%"
            icon={
                <TrafficIcon
                    sx={{color: colors.greenAccent[600], fontSize: "26px"}}
                />
            }
        />
    );
};

export default Regions;
