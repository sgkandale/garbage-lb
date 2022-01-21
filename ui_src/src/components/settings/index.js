import React, { useState } from 'react'
import { Typography, Tab, Tabs, Box } from '@mui/material'

function TabPanel(props) {
    const { children, value, index, ...other } = props;

    return (
        <div
            role="tabpanel"
            hidden={value !== index}
            id={`vertical-tabpanel-${index}`}
            aria-labelledby={`vertical-tab-${index}`}
            {...other}
        >
            {value === index && (
                <Box sx={{ p: 3 }}>
                    <Typography>{children}</Typography>
                </Box>
            )}
        </div>
    );
}


export default function Settings() {
    const [tab, setTab] = useState(0);

    const changeTab = (event, newValue) => {
        setTab(newValue);
    };

    const settingsTabs = [
        {
            label: 'General',
            content: <>General</>,
        },
        {
            label: 'Data',
            content: <>Data</>,
        },
    ]

    return <Box
        sx={{ flexGrow: 1, bgcolor: 'background.paper', display: 'flex' }}
    >
        <Tabs
            orientation="vertical"
            variant="scrollable"
            value={tab}
            onChange={changeTab}
            aria-label="settings-tabs"
            sx={{ borderRight: 1, marginLeft: 5, borderColor: 'divider' }}
        >
            {
                settingsTabs.map((eachTab, index) => {
                    return <Tab
                        label={eachTab.label}
                        style={{ textTransform: 'none' }}
                        key={index}
                    />
                })
            }
        </Tabs>
        {
            settingsTabs.map((eachTab, index) => {
                return <TabPanel value={tab} index={index}>
                    {eachTab.content}
                </TabPanel>
            })
        }
    </Box>
}