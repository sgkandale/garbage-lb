import React from 'react'
import { TableRow, TableCell } from '@mui/material';
// import { Close, Pause } from '@mui/icons-material';

export default function EachRule(props) {
    const { rule } = props;

    return <TableRow
        key={rule.name}
        sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
    >
        <TableCell component="th" scope="row">
            {rule.name}
        </TableCell>
        <TableCell align="right">{rule.type}</TableCell>
        <TableCell align="right">{rule.key || "-"}</TableCell>
        <TableCell align="right">{rule.value}</TableCell>
        <TableCell align="right">{rule.action}</TableCell>
        <TableCell align="right">{
            rule.enabled ? 'Enabled' : 'Disabled'
        }</TableCell>
        <TableCell align="right">{rule.cluster}</TableCell>
        {/* <TableCell align="right">
            <IconButton color="info">
                <Pause />
            </IconButton>
            <IconButton color="error">
                <Close />
            </IconButton>
        </TableCell> */}
    </TableRow>
}