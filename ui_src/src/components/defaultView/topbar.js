import React from 'react';
import { AppBar, Toolbar, Typography } from '@mui/material';

export default function TopBar() {

    return <AppBar position="fixed"
        sx={{
            zIndex: (theme) => theme.zIndex.drawer + 1,
            bgcolor: (theme) => theme.palette.page.background,
            boxShadow: 'none',
            borderBottom: 1,
            // borderBottom: '1px solid',
            borderColor: 'grey.300',
        }}
    >
        <Toolbar>
            <Typography variant="h6" noWrap component="div">
                Clipped drawer
            </Typography>
        </Toolbar>
    </AppBar>
}