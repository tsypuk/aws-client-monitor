import { Box, useTheme } from "@mui/material";
import { tokens } from "../theme";
import React from 'react';

interface ProgressCircleProps {
    sectors?: Array<{ value: number; color: string }>;
    size?: string;
}

const ProgressCircleMulti: React.FC<ProgressCircleProps> = ({ sectors = [{ value: 1, color: "green" }], size = "40" }) => {
    const theme = useTheme();
    const colors = tokens(theme.palette.mode);

    // Calculate the angles for each sector
    let currentAngle = 0;
    const gradientSegments = sectors.map((sector, index) => {
        const sectorAngle = sector.value * 360;
        const segment = `${sector.color} ${currentAngle}deg ${currentAngle + sectorAngle}deg`;
        currentAngle += sectorAngle;
        return segment;
    }).join(", ");

    return (
        <Box
            sx={{
                background: `radial-gradient(${colors.primary[400]} 55%, transparent 56%),
            conic-gradient(${gradientSegments}),
            ${colors.greenAccent[500]}`,
                borderRadius: "50%",
                width: `${size}px`,
                height: `${size}px`,
            }}
        />
    );
};

export default ProgressCircleMulti;
