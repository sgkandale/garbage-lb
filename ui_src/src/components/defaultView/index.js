import React from 'react';
import { Box, CssBaseline, Toolbar } from '@mui/material';
import TopBar from './topbar';
import Sidebar from './sidebar';

export default function DefaultView(props) {

    return <Box sx={{ display: 'flex' }}>
        <CssBaseline />
        <TopBar />
        <Sidebar
            navItems={props.navItems}
            activeNav={props.activeNav}
            changeNav={props.changeNav}
        />
        <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
            <Toolbar />
            <main>

            </main>
        </Box>
    </Box >
}
