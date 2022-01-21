import React from 'react'
import { Grid } from '@mui/material'
import EachEndpoint from './eachEndpoint'
import NewEndpoint from './newEndpoint'

export default function Endpoints(props) {

    return <Grid container spacing={2}>
        <Grid item xs={5} sx={{ borderRight: 1, borderColor: 'divider' }}>
            <NewEndpoint cluster={props.cluster} />
        </Grid>
        <Grid item xs={7}>
            {
                props.cluster.endpoints.map((endpoint, index) => {
                    return <EachEndpoint endpoint={endpoint} key={index} />
                })
            }
        </Grid>
    </Grid>
}