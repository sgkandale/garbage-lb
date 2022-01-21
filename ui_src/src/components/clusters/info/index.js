import { Box, Button, Tab, Tabs } from '@mui/material'
import React, { useState } from 'react'
import { ArrowBack } from '@mui/icons-material'
import TabPanel from '../../customComponents/tabpanel'
import GeneralInfo from './general'
import Delete from './delete'

export default function ClusterInfo(props) {
    const [tab, setTab] = useState(0);

    const changeTab = (event, newValue) => {
        setTab(newValue);
    };

    const infoTabs = [
        {
            label: 'General',
            content: <GeneralInfo />,
        },
        {
            label: 'Delete',
            content: <Delete />,
            color: '#F15741'
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
                        return <Tab
                            label={eachTab.label}
                            sx={{
                                textTransform: 'none',
                                color: eachTab.color || 'text.secondary'
                            }}
                            key={index}
                        />
                    })
                }
            </Tabs>
        </Box>
        {
            infoTabs.map((eachTab, index) => {
                return <TabPanel value={tab} index={index} key={index}>
                    {eachTab.content}
                </TabPanel>
            })
        }
    </Box >
}