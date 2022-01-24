import { Box, Typography } from '@mui/material'
import React from 'react'
import { useSelector } from 'react-redux'
import MapRules from './mapRules'

export default function Rules(props) {
    const listers = useSelector(state => state.listeners)

    for (let i = 0; i < listers.length; i++) {
        if (listers[i].name === props.listener) {
            console.log(listers[i])
            return <Box>
                <MapRules rules={listers[i].filter.rules} />
            </Box>
        }
    }

    return <Box>
        <Typography variant="subtitle1">
            No Such Listener
        </Typography>
    </Box>
}