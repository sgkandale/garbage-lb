import React from 'react'
import { Card, Typography, Grid, IconButton } from '@mui/material'
import { InfoOutlined } from '@mui/icons-material'

export default function ListenerCard(props) {

    return <Card sx={{ width: '100%', maxWidth: 500, padding: 2 }} elevation={3}>
        <Grid
            container
            direction="row"
            justifyContent="space-between"
            alignItems="center"
            sx={{ mb: 1 }}
        >
            <Typography variant="h6" >
                {props.listener.name}
            </Typography>
            <IconButton
                onClick={() => {
                    props.changeView('info')
                    props.setListener(props.listener.id)
                }}
            >
                <InfoOutlined />
            </IconButton>
        </Grid>
        <Typography variant="body1" color="textSecondary" sx={{ paddingLeft: 2 }}>
            <strong>Id : </strong>{props.listener.id}<br />
            <strong>Port : </strong>{props.listener.port}<br />
            <strong>Connection Type : </strong>{props.listener.type}<br />
            <strong>Listening : </strong>{props.listener.listening ? "True" : "False"}<br />
        </Typography>
    </Card>
}