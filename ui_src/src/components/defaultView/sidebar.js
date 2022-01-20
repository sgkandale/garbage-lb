import React from 'react';
import { Box, Drawer, Toolbar, List } from '@mui/material';
import { Divider, ListItem, ListItemIcon, ListItemText, Tooltip } from '@mui/material';

const drawerWidth = 250;

export default function Sidebar(props) {

    return <Drawer
        variant="permanent"
        sx={{
            width: drawerWidth,
            flexShrink: 0,
            [`& .MuiDrawer-paper`]: { width: drawerWidth, boxSizing: 'border-box' },
        }}
    >
        <Toolbar />
        <Box sx={{ overflow: 'auto' }}>
            <List>
                {
                    props.navItems.map((item, index) => {
                        if (item.type === 'divider') {
                            return <Divider key={index} style={{ marginTop: 5 }} />
                        }
                        return <Tooltip title={item.tooltip} placement="right" key={index}>
                            <ListItem
                                button
                                sx={{ color: item.color ? item.color : 'default' }}
                            >
                                <ListItemIcon
                                    sx={{ color: item.color ? item.color : 'inherit' }}
                                >
                                    {item.icon}
                                </ListItemIcon>
                                <ListItemText primary={item.name} />
                            </ListItem>
                        </Tooltip>
                    })
                }
            </List>
        </Box>
    </Drawer>
}