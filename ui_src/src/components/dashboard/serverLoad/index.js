import { Paper, Typography, Grid, Box, Divider } from '@mui/material'
import React from 'react'
import { useSelector } from 'react-redux'

export default function ServerLoad() {
    const serverLoad = useSelector(state => state.serverLoad)

    return <Paper elevation={0} sx={{ border: '1px solid', borderColor: 'divider', padding: 1 }}>
        <Typography variant="h6" >
            Server Load
        </Typography>
        <br />
        <Grid
            container
            direction="row"
            justifyContent="space-around"
            alignItems="flex-start"
        >
            <Box>
                <Typography variant="body1" style={{ marginBottom: 10 }}>
                    CPU
                </Typography>
                <Typography variant="body1" color="textSecondary" style={{ paddingLeft: 12 }}>
                    <strong>Total : </strong>{serverLoad.cpu.total}
                    <br />
                    <strong>In Use : </strong>{serverLoad.cpu.inUse}
                    <br />
                    <strong>Available : </strong>{serverLoad.cpu.available}
                    <br />
                    <strong>Percentage : </strong>{serverLoad.cpu.inUsePercentage}
                    <br />
                </Typography>
            </Box>
            <Divider orientation="vertical" flexItem />
            <Box>
                <Typography variant="body1" style={{ marginBottom: 10 }}>
                    RAM
                </Typography>
                <Typography variant="body1" color="textSecondary" style={{ paddingLeft: 12 }}>
                    <strong>Total : </strong>{serverLoad.ram.total}
                    <br />
                    <strong>In Use : </strong>{serverLoad.ram.inUse}
                    <br />
                    <strong>Available : </strong>{serverLoad.ram.available}
                    <br />
                    <strong>Percentage : </strong>{serverLoad.ram.inUsePercentage}
                    <br />
                </Typography>
            </Box>
            <Divider orientation="vertical" flexItem />
            <Box>
                <Typography variant="body1" style={{ marginBottom: 10 }}>
                    Storage
                </Typography>
                <Typography variant="body1" color="textSecondary" style={{ paddingLeft: 12 }}>
                    <strong>Total : </strong>{serverLoad.storage.total}
                    <br />
                    <strong>In Use : </strong>{serverLoad.storage.inUse}
                    <br />
                    <strong>Available : </strong>{serverLoad.storage.available}
                    <br />
                    <strong>Percentage : </strong>{serverLoad.storage.inUsePercentage}
                    <br />
                </Typography>
            </Box>
            <Divider orientation="vertical" flexItem />
            <Box>
                <Typography variant="body1" style={{ marginBottom: 10 }}>
                    Network
                </Typography>
                <Typography variant="body1" color="textSecondary" style={{ paddingLeft: 12 }}>
                    <strong>In Use : </strong>{serverLoad.network.inUse}
                    <br />
                    <strong>Up : </strong>{serverLoad.network.up}
                    <br />
                    <strong>Down : </strong>{serverLoad.network.down}
                    <br />
                </Typography>
            </Box>
        </Grid>
    </Paper >
}