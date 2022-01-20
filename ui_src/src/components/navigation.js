import React, { useState } from 'react'
// import { BrowserRouter, Routes, Route } from "react-router-dom";
import DefaultView from './defaultView/index';
import { Api, Dashboard, PowerSettingsNew, Settings, Storage } from '@mui/icons-material';

export default function Navigation() {
    const [activeNav, setActiveNav] = useState(0);

    const navItems = [
        {
            name: 'Dashboard',
            icon: <Dashboard />,
            tooltip: 'Dashboard',
            type: 'item',
        },
        {
            type: 'divider',
        },
        {
            name: 'Listeners',
            icon: <Api />,
            tooltip: 'Listener Elements',
            type: 'item',
        },
        {
            name: 'Clusters',
            icon: <Storage />,
            tooltip: 'Backend Clusters',
            type: 'item',
        },
        {
            type: 'divider',
        },
        {
            name: 'Settings',
            icon: <Settings />,
            tooltip: 'Settings',
            type: 'item',
        },
        {
            name: 'Test',
            icon: <Settings />,
            tooltip: 'Test',
            type: 'item',
        },
        {
            type: 'divider',
        },
        {
            name: 'Terminate',
            icon: <PowerSettingsNew />,
            tooltip: 'Terminate Server',
            type: 'item',
            color: "#F15741"
        },
    ]

    const changeNav = (name) => {
        for (let i = 0; i < navItems.length; i++) {
            if (navItems[i].name === name) {
                setActiveNav(i);
                break;
            }
        }
        setActiveNav(-1);
    }

    return <DefaultView
        navItems={navItems}
        activeNav={activeNav}
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