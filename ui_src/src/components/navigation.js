import React, { useState } from 'react'
import DefaultView from './defaultView/index';
import { Api, Dashboard as DashboardIcon, Settings as SettingsIcon, Speed, Storage } from '@mui/icons-material';
import Dashboard from './dashboard'
import Listeners from './listeners'
import Clusters from './clusters'
import Settings from './settings'
import Test from './test'
import GetServerLoad from './getServerLoad';
import GetListeners from './getListeners';
import GetClusters from './getClusters';

export default function Navigation() {
    const [activeNav, setActiveNav] = useState(0);

    const navItems = [
        {
            name: 'Dashboard',
            icon: <DashboardIcon />,
            renderContent: <Dashboard />
        },
        {
            name: 'Listeners',
            icon: <Api />,
            renderContent: <Listeners />
        },
        {
            name: 'Clusters',
            icon: <Storage />,
            renderContent: <Clusters />
        },
        {
            name: 'Settings',
            icon: <SettingsIcon />,
            renderContent: <Settings />
        },
        {
            name: 'Test',
            icon: <Speed />,
            renderContent: <Test />
        },
    ]

    const changeNav = (name) => {
        for (let i = 0; i < navItems.length; i++) {
            if (navItems[i].name === name) {
                setActiveNav(i);
                return
            }
        }
        setActiveNav(-1);
    }

    return <>
        <DefaultView
            navItems={navItems}
            activeNav={activeNav}
            changeNav={changeNav}
        />
        <GetServerLoad />
        <GetListeners />
        <GetClusters />
    </>

}