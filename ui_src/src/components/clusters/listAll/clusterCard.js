import React from 'react'
import { Card, Typography, Grid, IconButton } from '@mui/material'
import { InfoOutlined } from '@mui/icons-material'

export default function ClusterCard(props) {

    return <Card sx={{ width: '100%', maxWidth: 500, padding: 2 }} elevation={3}>
        <Grid
            container
            direction="row"
            justifyContent="space-between"
            alignItems="center"
            sx={{ mb: 1 }}
        >
            <Typography variant="h6" >
                {props.cluster.name}
            </Typography>
            <IconButton
                onClick={() => {
                    props.changeView('info')
                    props.setCluster(props.cluster.id)
                }}
            >
                <InfoOutlined />
            </IconButton>
        </Grid>
        <Typography variant="body1" color="textSecondary" sx={{ paddingLeft: 2 }}>
            <strong>Id : </strong>{props.cluster.id}<br />
            <strong>Port : </strong>{props.cluster.port}<br />
            <strong>Connection Type : </strong>{props.cluster.type}<br />
        </Typography>
    </Card>
}