import React, { useState } from 'react'
// import { BrowserRouter, Routes, Route } from "react-router-dom";
import DefaultView from './defaultView/index';
import { Api, Dashboard as DashboardIcon, PowerSettingsNew, Settings as SettingsIcon, Speed, Storage } from '@mui/icons-material';
import Dashboard from './dashboard'
import Listeners from './listeners'
import Clusters from './clusters'
import Settings from './settings'
import Test from './test'

export default function Navigation() {
    const [activeNav, setActiveNav] = useState(0);

    const navItems = [
        {
            name: 'Dashboard',
            icon: <DashboardIcon />,
            type: 'item',
            renderContent: <Dashboard />
        },
        {
            type: 'divider',
        },
        {
            name: 'Listeners',
            icon: <Api />,
            type: 'item',
            renderContent: <Listeners />
        },
        {
            name: 'Clusters',
            icon: <Storage />,
            type: 'item',
            renderContent: <Clusters />
        },
        {
            type: 'divider',
        },
        {
            name: 'Settings',
            icon: <SettingsIcon />,
            type: 'item',
            renderContent: <Settings />
        },
        {
            name: 'Test',
            icon: <Speed />,
            type: 'item',
            renderContent: <Test />
        },
        {
            type: 'divider',
        },
        {
            name: 'Terminate',
            icon: <PowerSettingsNew />,
            type: 'item',
            color: "#F15741",
            clickHandler: () => console.log("call terminate endpoint here")
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

    return <DefaultView
        navItems={navItems}
        activeNav={activeNav}
        changeNav={changeNav}
    />

    // return <BrowserRouter>
    //     <Routes>
    //         <Route path="/" element={<DefaultView
    //             navItems={navItems}
    //             activeNav={activeNav}
    //         />} />
    //     </Routes>
    // </BrowserRouter>
}