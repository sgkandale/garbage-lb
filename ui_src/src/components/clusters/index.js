import React, { useState } from 'react'
import CreateCluster from './createCluster'
import ClusterInfo from './info'
import ListAllClusters from './listAll'

export default function Clusters() {
    const [view, setView] = useState('list')
    const [cluster, setCluster] = useState('')

    const changeView = (targetView) => {
        setTimeout(() => {
            setView(targetView)
        }, 200)
    }

    const renderView = () => {
        if (view === 'list') {
            return <ListAllClusters
                changeView={changeView}
                setCluster={setCluster}
            />
        } else if (view === 'create') {
            return <CreateCluster
                changeView={changeView}
            />
        } else if (view === 'info') {
            return <ClusterInfo
                changeView={changeView}
                cluster={cluster}
            />
        } else {
            return <></>
        }
    }

    return renderView()
}