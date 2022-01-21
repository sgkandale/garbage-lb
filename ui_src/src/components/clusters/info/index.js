import { Box, Button } from '@mui/material'
import React from 'react'
import { useSelector } from 'react-redux'
import { ArrowBack } from '@mui/icons-material'

export default function ClusterInfo(props) {
    const clusters = useSelector(state => state.clusters)

    return <Box >
        <Button
            sx={{ textTransform: 'none', color: 'text.secondary' }}
            onClick={() => props.changeView('list')}
            startIcon={<ArrowBack />}
        >
            Back
        </Button>
        <br />
    </Box >
}