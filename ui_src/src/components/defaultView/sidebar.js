import React from 'react';
import { Box, Drawer, Toolbar, List } from '@mui/material';
import { Divider, ListItem, ListItemIcon, ListItemText } from '@mui/material';

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
                            return <Divider key={index} style={{ marginTop: 5, marginBottom: 5 }} />
                        }
                        return <ListItem
                            button
                            sx={{
                                color: item.color ?
                                    item.color : 'grey.700',
                                bgcolor: props.activeNav === index ?
                                    'primary.ultraLight' : 'none',
                                borderTopRightRadius: 20,
                                borderBottomRightRadius: 20,
                            }}
                            onClick={
                                () => {
                                    if (item.clickHandler) {
                                        item.clickHandler()
                                    } else {
                                        props.changeNav(item.name)
                                    }
                                }
                            }
                            key={index}
                        >
                            <ListItemIcon
                                sx={{ color: item.color ? item.color : 'inherit' }}
                            >
                                {item.icon}
                            </ListItemIcon>
                            <ListItemText primary={item.name} />
                        </ListItem>
                    })
                }
            </List>
        </Box>
    </Drawer>
}