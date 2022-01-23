import { Box, Button, Tabs, Tab } from '@mui/material'
import React, { useEffect, useState } from 'react'
import { ArrowBack } from '@mui/icons-material'
import TabPanel from '../../customComponents/tabpanel'
import GeneralInfo from './general'
import Delete from './delete'
import { useSelector } from 'react-redux'
import Filters from './filters'

export default function ListenerInfo(props) {
    const [tab, setTab] = useState(0);
    const [listener, setListener] = useState({})
    const listeners = useSelector(state => state.listeners)

    const changeTab = (event, newValue) => {
        setTab(newValue);
    };

    useEffect(() => {
        if (listeners.length === 0) {
            return
        } else {
            for (let i = 0; i < listeners.length; i++) {
                if (listeners[i].name === props.listener) {
                    setListener(listeners[i])
                }
            }
        }
    }, [])

    const infoTabs = [
        {
            label: 'General',
            content: <GeneralInfo
                listener={listener.name}
            />,
        },
        {
            label: 'Filters',
            content: <Filters
                listener={listener.name}
            />,
        },
        {
            label: 'Delete',
            content: <Delete
                listener={listener.name}
            />,
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