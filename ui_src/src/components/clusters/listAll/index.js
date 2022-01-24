import { Box, Grid } from '@mui/material'
import React from 'react'
import { useSelector } from 'react-redux'
import ServerOffline from '../../customComponents/serverOffline'
import ClusterCard from './clusterCard'

export default function ListAllClusters(props) {
    const clusters = useSelector(state => state.clusters)
    const lbStatus = useSelector(state => state.lbStatus)

    if (lbStatus !== "Active") {
        return <ServerOffline />
    }

    return <Box >
        {/* <Button
            variant="contained"
            color="primary"
            sx={{ textTransform: 'none' }}
            onClick={() => props.changeView('create')}
        >
            Create New
        </Button> */}
        <br />
        <Grid
            container
            direction="row"
            justifyContent="flex-start"
            alignItems="center"
            sx={{ width: '100%', marginTop: 3, marginBottom: 3 }}
        >
            {
                clusters.map((eachCluster) => {
                    return <ClusterCard
                        cluster={eachCluster}
                        key={eachCluster.name}
                        setCluster={props.setCluster}
                        changeView={props.changeView}
                    />
                })
            }
        </Grid>
    </Box >
}