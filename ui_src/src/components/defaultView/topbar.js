import React from 'react';
import { AppBar, Toolbar, Typography, Grid } from '@mui/material';
import TopbarStatus from './topbarStatus';
import logo from '../../statics/logo.png';

export default function TopBar() {

    return <AppBar position="fixed"
        sx={{
            zIndex: (theme) => theme.zIndex.drawer + 1,
            bgcolor: (theme) => theme.palette.page.background,
            boxShadow: 'none',
            borderBottom: 1,
            borderColor: 'grey.300',
        }}
    >
        <Toolbar>
            <img
                src={logo}
                alt='garbage-lb logo'
                style={{ width: 50, height: 50, marginRight: 20 }}
            />
            <Typography variant="h6" noWrap color="textPrimary" sx={{ flexGrow: 1 }}>
                Garbage-LB
            </Typography>
            <TopbarStatus />
        </Toolbar>
    </AppBar >
}