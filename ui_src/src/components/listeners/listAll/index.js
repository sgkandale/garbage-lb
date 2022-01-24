import { Box, Grid } from '@mui/material'
import React from 'react'
import { useSelector } from 'react-redux'
import ServerOffline from '../../customComponents/serverOffline'
import ListenerCard from './listenerCard'

export default function ListAllListeners(props) {
    const listeners = useSelector(state => state.listeners)
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
                listeners.map((eachListener) => {
                    return <ListenerCard
                        listener={eachListener}
                        key={eachListener.name}
                        setListener={props.setListener}
                        changeView={props.changeView}
                    />
                })
            }
        </Grid>
    </Box >
}