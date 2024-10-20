import React, {useEffect, useState} from 'react';
import websocketService, {ClientsCounter} from '../service/websocketService';
import {tokens} from "../theme";
import {useTheme} from "@mui/material";
import PersonAddIcon from "@mui/icons-material/PersonAdd";
import StatBoxMulti from "./StatBoxMulti";

const Clients: React.FC = () => {
    const [counter, setCounter] = useState<Map<string, number>>(new Map<string, number>());

    const theme = useTheme();
    const colors = tokens(theme.palette.mode);

    useEffect(() => {
        const handleCountChange = (newCounter: ClientsCounter) => {
            setCounter(newCounter)
            // console.log(newCounter)
        };

        websocketService.addClientListener(handleCountChange);
        // websocketService.connect();

        return () => {
            websocketService.removeClientListener(handleCountChange);
        };
    }, []);

    const getSectors = (): Array<{ value: number; color: string }> => {
        const total = Array.from(counter.values()).reduce((acc, count) => acc + count, 0); // Calculate the total count
        const colorsArray = [colors.greenAccent[500], colors.blueAccent[500], "orange", "purple", "yellow"]; // Define colors for each client

        let colorIndex = 0; // Track the index of the colors array

        // Generate sectors based on the counter map
        if (counter.size == 0) {

            return [{value: 1, color: colors.redAccent[500]}]
        }
        const sectors = Array.from(counter.entries()).map(([clientId, count]) => {
            const value = count / total; // Calculate the fraction of the total for this client
            const color = colorsArray[colorIndex % colorsArray.length]; // Assign a color from the list, looping if needed
            colorIndex++;

            return {value, color};
        });

        return sectors;
    };
    const sectors = getSectors();
    // console.log(counter.size)

    return (
        <StatBoxMulti
            title={counter.size}
            subtitle="Clients"
            progress={sectors}
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
