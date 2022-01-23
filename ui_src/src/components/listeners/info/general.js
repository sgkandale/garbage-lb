import React from 'react'
import { useSelector } from 'react-redux'
import { Box, Card, CardContent, Typography } from '@mui/material'

export default function GeneralInfo(props) {
    const listeners = useSelector(state => state.listeners)

    for (let i = 0; i < listeners.length; i++) {
        if (listeners[i].id === props.listener) {
            return <Card elevation={2}>
                <CardContent>
                    <Typography variant="subtitle1">
                        <strong>Name : </strong>{listeners[i].name}
                    </Typography>
                    <Typography variant="body2" color="textSecondary" style={{ paddingLeft: 20 }}>
                        <strong>ID : </strong>{listeners[i].id} <br />
                        <strong>Port : </strong>{listeners[i].port} <br />
                        <strong>Type : </strong>{listeners[i].type} <br />
                        <strong>Listening : </strong>{listeners[i].listening ? "True" : "False"} <br />
                    </Typography>
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