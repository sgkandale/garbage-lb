import { Paper, Typography, Grid, Box, Divider, CircularProgress } from '@mui/material'
import React from 'react'
import { useSelector } from 'react-redux'

export default function ServerLoad() {
    const serverLoad = useSelector(state => state.serverLoad)

    const putInGrid = (children) => {
        return <Grid
            container
            direction="row"
            justifyContent="center"
            alignItems="center"
            style={{ width: '100%', minHeight: 200 }}
        >
            {children}
        </Grid>
    }

    if (serverLoad.loading) {
        return putInGrid(<CircularProgress size={25} />)
    }

    if (serverLoad.error) {
        return putInGrid(<Typography variant="h6" align="center" >
            Server Offline
        </Typography>)
    }

    return <Paper elevation={0} sx={{ border: '1px solid', borderColor: 'divider', padding: 1, paddingBottom: 2 }}>
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
                    Memory
                </Typography>
                <Typography variant="body1" color="textSecondary" style={{ paddingLeft: 12 }}>
                    <strong>Total : </strong>{serverLoad.memory.total}
                    <br />
                    <strong>In Use : </strong>{serverLoad.memory.inUse}
                    <br />
                    <strong>Available : </strong>{serverLoad.memory.available}
                    <br />
                    <strong>Percentage : </strong>{serverLoad.memory.inUsePercentage}
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
                    <strong>Total Use : </strong>{serverLoad.network.sentData + serverLoad.network.receivedData}
                    <br />
                    <strong>Sent Data : </strong>{serverLoad.network.sentData}
                    <br />
                    <strong>Received Data : </strong>{serverLoad.network.receivedData}
                    <br />
                </Typography>
            </Box>
        </Grid>
    </Paper >
}