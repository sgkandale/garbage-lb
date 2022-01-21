import React, { useState } from 'react'
import CreateListener from './createListener'
import ListenerInfo from './info'
import ListAllListeners from './listAll'

export default function Listeners() {
    const [view, setView] = useState('list')
    const [listener, setListener] = useState('')

    const changeView = (targetView) => {
        setTimeout(() => {
            setView(targetView)
        }, 200)
    }

    const renderView = () => {
        if (view === 'list') {
            return <ListAllListeners
                changeView={changeView}
                setListener={setListener}
            />
        } else if (view === 'create') {
            return <CreateListener
                changeView={changeView}
            />
        } else if (view === 'info') {
            return <ListenerInfo
                changeView={changeView}
                listener={listener}
            />
        } else {
            return <></>
        }
    }

    return renderView()
}