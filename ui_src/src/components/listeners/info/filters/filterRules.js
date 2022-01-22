import React from 'react';
import { IconButton, Paper, Table, TableBody, TableRow, TableHead, TableContainer, TableCell } from '@mui/material';
import { Close, Pause } from '@mui/icons-material';

export default function FilterRules(props) {
    return <TableContainer component={Paper} elevation={0} >
        <Table sx={{ minWidth: 800 }} aria-label="simple table">
            <TableHead>
                <TableRow>
                    <TableCell>Name</TableCell>
                    <TableCell align="right">ID</TableCell>
                    <TableCell align="right">Type</TableCell>
                    <TableCell align="right">Value</TableCell>
                    <TableCell align="right">Subvalue</TableCell>
                    <TableCell align="right">Action</TableCell>
                    <TableCell align="right">Status</TableCell>
                    <TableCell align="right">Options</TableCell>
                </TableRow>
            </TableHead>
            <TableBody>
                {
                    props.rules.map((rule) => (
                        <TableRow
                            key={rule.name}
                            sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
                        >
                            <TableCell component="th" scope="row">
                                {rule.name}
                            </TableCell>
                            <TableCell align="right">{rule.id}</TableCell>
                            <TableCell align="right">{rule.type}</TableCell>
                            <TableCell align="right">{rule.value}</TableCell>
                            <TableCell align="right">{rule.subValue || "-"}</TableCell>
                            <TableCell align="right">{rule.action}</TableCell>
                            <TableCell align="right">{
                                rule.enabled ? 'Enabled' : 'Disabled'
                            }</TableCell>
                            <TableCell align="right">
                                <IconButton color="info">
                                    <Pause />
                                </IconButton>
                                <IconButton color="error">
                                    <Close />
                                </IconButton>
                            </TableCell>
                        </TableRow>
                    ))
                }
            </TableBody>
        </Table>
    </TableContainer>
}
