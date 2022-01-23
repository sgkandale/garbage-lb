import React from 'react'
import { Grid, Typography } from '@mui/material'

export default function ServerOffline() {

    return <Grid
        container
        direction="row"
        justifyContent="center"
        alignItems="center"
        style={{ width: '100%', minHeight: 200 }}
    >
        <Typography variant="h6" align="center">
            Server Offline
        </Typography>
    </Grid>
}