import React from 'react'
import { Card, CardContent, Typography, Grid, Divider } from '@mui/material'

export default function GeneralInfo(props) {

    return <Card elevation={0} sx={{ border: 1, borderColor: 'divider' }}>
        <CardContent>
            <Grid container spacing={2} >
                <Grid item xs={5} >
                    <Typography variant="h6" gutterBottom >
                        Details
                    </Typography>
                    <Typography variant="body1" color="textSecondary" style={{ paddingLeft: 20 }} >
                        <strong>Name : </strong> {props.cluster.name} <br />
                        <strong>Policy : </strong> {props.cluster.policy} <br />
                        <strong>Timeout : </strong> {props.cluster.timeout} s <br />
                    </Typography>
                </Grid>
                <Grid item style={{ width: 20 }}  >
                    <Divider orientation="vertical" />
                </Grid>
                <Grid item xs={5} >
                    <Typography variant="h6" gutterBottom  >
                        Health
                    </Typography>
                    {/* <Typography variant="body1" color="textSecondary" style={{ paddingLeft: 20 }} >
                        <strong>Status : </strong> {props.cluster.health.status} <br />
                        <strong>Healthy Endpoints : </strong> {props.cluster.health.healthyCount} <br />
                        <strong>Unhealthy Endpoints : </strong> {props.cluster.health.unhealthyCount} <br />
                        <strong>Degraded Endpoints : </strong> {props.cluster.health.degradedCount} <br />
                    </Typography> */}
                </Grid>
            </Grid>
        </CardContent>
    </Card>
}