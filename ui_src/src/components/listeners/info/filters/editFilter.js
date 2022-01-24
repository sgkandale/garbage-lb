import { Add, Close } from '@mui/icons-material'
import { Box, IconButton, Grid, Typography } from '@mui/material'
import React from 'react'
import { useSelector } from 'react-redux'
import FilterRules from './filterRules'

export default function EditFilter(props) {
    const listeners = useSelector(state => state.listeners)

    for (let i = 0; i < listeners.length; i++) {
        if (listeners[i].name === props.listener) {
            for (let j = 0; j < listeners[i].filters.length; j++) {
                if (listeners[i].filters[j].name === props.filterToEdit) {
                    return <Box>
                        <Grid
                            container
                            direction="row"
                            justifyContent="space-between"
                            alignItems="flex-start"
                            style={{ marginTop: 10, marginBottom: 10 }}
                        >
                            <Box >
                                <Typography variant="subtitle1">
                                    {listeners[i].filters[j].name} (id : {listeners[i].filters[j].id})
                                </Typography>
                            </Box>
                            <Box>
                                <IconButton>
                                    <Add />
                                </IconButton>
                                <IconButton onClick={() => props.setView('list')}>
                                    <Close />
                                </IconButton>
                            </Box>
                        </Grid>
                        <Typography variant="subtitle1" style={{ marginTop: 10, marginBottom: 10 }}>
                            Rules :
                        </Typography>
                        <FilterRules rules={listeners[i].filters[j].rules} />
                    </Box>
                }
            }
        }
    }

    return <Box>
        <Grid
            container
            direction="row"
            justifyContent="space-between"
            alignItems="center"
        >
            <Typography variant="subtitle1">
                No Such Filter
            </Typography>
            <IconButton onClick={() => props.setView('list')}>
                <Close />
            </IconButton>
        </Grid>
    </Box>
}