import { Box, Grid, Button } from '@mui/material'
import React from 'react'
import { useSelector } from 'react-redux'
import ListenerCard from './listenerCard'

export default function ListAllListeners(props) {
    const listeners = useSelector(state => state.listeners)

    return <Box >
        <Button
            variant="contained"
            color="primary"
            sx={{ textTransform: 'none' }}
            onClick={() => props.changeView('create')}
        >
            Create New
        </Button>
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
                        key={eachListener.id}
                        setListener={props.setListener}
                        changeView={props.changeView}
                    />
                })
            }
        </Grid>
    </Box >
}