import { Box, Button } from '@mui/material'
import React from 'react'

export default function ListAllListeners(props) {

    return <Box sx={{ display: 'flex' }}>
        <Button
            variant="contained"
            color="primary"
            sx={{ textTransform: 'none' }}
            onClick={() => props.changeView('create')}
        >
            Create New
        </Button>
        <br />
    </Box >
}