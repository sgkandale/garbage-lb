import React from 'react'
import { Box, Button, Typography } from '@mui/material'
import EachFilter from './listEachFilter'
import { useSelector } from 'react-redux'

export default function ListFilters(props) {
    const listeners = useSelector(state => state.listeners)
    let listenerIndex = -1

    for (let i = 0; i < listeners.length; i++) {
        if (listeners[i].name === props.listener) {
            listenerIndex = i
        }
    }

    if (listenerIndex === -1) {
        return <Box>
            <Typography variant="subtitle1">
                No Such Listener
            </Typography>
        </Box>
    }

    const render = () => {
        return <>
            {
                listeners[listenerIndex].filters.map((filter, index) => {
                    return <EachFilter
                        key={index}
                        filter={filter}
                        setFilterToEdit={props.setFilterToEdit}
                        setView={props.setView}
                    />
                })
            }
        </>
    }

    return <Box>
        <Button
            variant="contained"
            color="primary"
            onClick={() => props.setView('create')}
            sx={{ marginBottom: 2, textTransform: 'none' }}
        >
            Create Filter
        </Button>
        <br />
        {render()}
    </Box>
}