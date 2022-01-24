import React, { useEffect, useState } from 'react'
// import { BrowserRouter, Routes, Route } from "react-router-dom";
import DefaultView from './defaultView/index';
import { Api, Dashboard as DashboardIcon, PowerSettingsNew, Settings as SettingsIcon, Speed, Storage } from '@mui/icons-material';
import Dashboard from './dashboard'
import Listeners from './listeners'
import Clusters from './clusters'
import Settings from './settings'
import Test from './test'
import GetServerLoad from './getServerLoad';
import GetListeners from './getListeners';
import { useSelector } from 'react-redux';
import GetClusters from './getClusters';

export default function Navigation() {

    const defaultNavItems = [
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
    const [activeNav, setActiveNav] = useState(0);
    const [navItems, setNavItems] = useState(defaultNavItems.slice(0, defaultNavItems.length - 2))
    const lbStatus = useSelector(state => state.lbStatus)

    useEffect(() => {
        if (lbStatus === "Active") {
            setNavItems(defaultNavItems.slice(0, defaultNavItems.length))
        } else {
            setNavItems(defaultNavItems.slice(0, defaultNavItems.length - 2))
        }
    }, [lbStatus])

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

    // return <BrowserRouter>
    //     <Routes>
    //         <Route path="/" element={<DefaultView
    //             navItems={navItems}
    //             activeNav={activeNav}
    //         />} />
    //     </Routes>
    // </BrowserRouter>
}