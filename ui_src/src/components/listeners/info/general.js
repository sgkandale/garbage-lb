import React from 'react'
import { useSelector } from 'react-redux'
import { Box, Card, CardContent, Typography, Grid, Divider } from '@mui/material'

export default function GeneralInfo(props) {
    const listeners = useSelector(state => state.listeners)

    for (let i = 0; i < listeners.length; i++) {
        if (listeners[i].name === props.listener) {
            return <Card elevation={0} sx={{ border: 1, borderColor: 'divider' }}>
                <CardContent>
                    <Grid container spacing={2} >
                        <Grid item xs={5} >
                            <Typography variant="h6" gutterBottom >
                                Details
                            </Typography>
                            <Typography variant="body1" color="textSecondary" style={{ paddingLeft: 20 }}>
                                <strong>Name : </strong> {listeners[i].name} <br />
                                <strong>Port : </strong>{listeners[i].port} <br />
                                <strong>Type : </strong>{listeners[i].type} <br />
                                <strong>Listening : </strong>{listeners[i].listening ? "True" : "False"} <br />
                            </Typography>
                        </Grid>
                        <Grid item style={{ width: 20 }}  >
                            <Divider orientation="vertical" />
                        </Grid>
                        <Grid item style={{ width: 20 }} xs={6} >
                            <Typography variant="h6" gutterBottom  >
                                Filter
                            </Typography>
                            <Typography variant="body1" color="textSecondary" style={{ paddingLeft: 20 }} >
                                <strong>Name : </strong> {listeners[i].filter.name} <br />
                            </Typography>
                        </Grid>
                    </Grid>
                </CardContent>
            </Card>
        }
    }

    return <Box>
        <Typography variant="subtitle1">
            No Such Listener
        </Typography>
    </Box>
}