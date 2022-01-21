import { Box, Button, Tabs, Tab } from '@mui/material'
import React, { useState } from 'react'
import { ArrowBack } from '@mui/icons-material'
import TabPanel from '../../customComponents/tabpanel'
import GeneralInfo from './general'

export default function ListenerInfo(props) {
    const [tab, setTab] = useState(0);

    const changeTab = (event, newValue) => {
        setTab(newValue);
    };

    const infoTabs = [
        {
            label: 'General',
            content: <GeneralInfo />,
        }
    ]

    return <Box >
        <Button
            sx={{ textTransform: 'none', color: 'text.secondary' }}
            onClick={() => props.changeView('list')}
            startIcon={<ArrowBack />}
        >
            Back
        </Button>
        <br />
        <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
            <Tabs value={tab} onChange={changeTab} aria-label="basic tabs example">
                {
                    infoTabs.map((eachTab, index) => {
                        return <Tab label={eachTab.label} style={{ textTransform: 'none' }} />
                    })
                }
            </Tabs>
        </Box>
        {
            infoTabs.map((eachTab, index) => {
                return <TabPanel value={tab} index={index}>
                    {eachTab.content}
                </TabPanel>
            })
        }
    </Box >
}