import { Close } from '@mui/icons-material'
import { Box, IconButton, Grid, Typography } from '@mui/material'
import React from 'react'

export default function CreateFilter(props) {

    return <Box>
        <Grid
            container
            direction="row"
            justifyContent="space-between"
            alignItems="center"
        >
            <Typography variant="subtitle1">
                Create Filter
            </Typography>
            <IconButton onClick={() => props.setView('list')}>
                <Close />
            </IconButton>
        </Grid>
    </Box>
}