import React from 'react'
import { Card, CardContent, Typography, Grid, Box } from '@mui/material'
// import DeleteEndpoint from './deleteButton'

const getColor = (status) => {
    if (status === 'Healthy') {
        return '#42D948'
    } else if (status === 'Unhealthy') {
        return '#FF0000'
    } else {
        return '#000000'
    }
}

export default function EachEndpoint(props) {

    const renderHealth = (statusValue) => {
        return <Grid
            container
            direction="row"
            justifyContent="flex-start"
            alignItems="center"
            style={{ margin: 5 }}
        >
            <div style={{ width: 10, height: 10, backgroundColor: getColor(statusValue), borderRadius: '50%', marginRight: 10 }} />
            <Typography variant="subtitle2" color="textSecondary">
                {statusValue}
            </Typography>
        </Grid>
    }

    return <Card style={{ marginBottom: 20 }}>
        <CardContent>
            <Grid
                container
                direction="row"
                justifyContent="space-between"
                alignItems="center"
            >
                <Box>
                    <Typography variant="body2" color="textSecondary">
                        <strong>Name : </strong>{props.endpoint.name} <br />
                        <strong>Address : </strong>{props.endpoint.address} <br />
                        <strong>Port : </strong>{props.endpoint.port} <br />
                    </Typography>
                </Box>
                {/* <DeleteEndpoint endpoint={props.endpoint} /> */}
            </Grid>
            {renderHealth(props.endpoint.health)}
        </CardContent>
    </Card>
}