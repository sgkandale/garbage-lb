import React, { useState } from 'react'
import CreateListener from './createListener'
import ListAllListeners from './listAll'

export default function Listeners() {
    const [view, setView] = useState('list')

    const changeView = (targetView) => {
        setTimeout(() => {
            setView(targetView)
        }, 200)
    }

    const renderView = () => {
        if (view === 'list') {
            return <ListAllListeners
                changeView={changeView}
            />
        } else if (view === 'create') {
            return <CreateListener
                changeView={changeView}
            />
        } else {
            return <></>
        }
    }

    return renderView()
}