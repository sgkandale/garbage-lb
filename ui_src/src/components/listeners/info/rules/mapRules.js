import React from 'react'
import { Paper, Table, TableBody, TableRow, TableHead, TableContainer, TableCell } from '@mui/material'
import EachRule from './eachRule'

export default function MapRules(props) {

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
                    props.rules.map((rule, index) => {
                        return <EachRule rule={rule} key={index} />
                    })
                }
            </TableBody>
        </Table>
    </TableContainer>
}