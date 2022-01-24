import React from 'react'
import { Paper, Table, TableBody, TableRow, TableHead, TableContainer, TableCell, Box, Typography } from '@mui/material'
import EachRule from './eachRule'
import { useSelector } from 'react-redux'

export default function ListRules(props) {
    const listeners = useSelector(state => state.listeners)

    for (let i = 0; i < listeners.length; i++) {
        if (listeners[i].name === props.listener) {
            return <TableContainer component={Paper} elevation={0} >
                <Table sx={{ minWidth: 800 }} aria-label="simple table">
                    <TableHead>
                        <TableRow>
                            <TableCell>Name</TableCell>
                            <TableCell align="right">Type</TableCell>
                            <TableCell align="right">Key</TableCell>
                            <TableCell align="right">Value</TableCell>
                            <TableCell align="right">Action</TableCell>
                            <TableCell align="right">Status</TableCell>
                            <TableCell align="right">Target Cluster</TableCell>
                            {/* <TableCell align="right">Options</TableCell> */}
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {
                            listeners[i].filter.rules.map((rule, index) => {
                                return <EachRule rule={rule} key={index} />
                            })
                        }
                    </TableBody>
                </Table>
            </TableContainer>
        }
    }

    return <Box>
        <Typography variant="subtitle1">
            No Such Listener
        </Typography>
    </Box>

    return
}