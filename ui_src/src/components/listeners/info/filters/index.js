import { Box } from '@mui/material'
import React, { useState } from 'react'
import CreateFilter from './createFilter'
import EditFilter from './editFilter'
import ListFilters from './listFilters'

export default function Filters(props) {
    const [view, setView] = useState('list')
    const [filterToEdit, setFilterToEdit] = useState('')

    const render = () => {
        if (view === 'list') {
            return <ListFilters
                listener={props.listener}
                setView={setView}
                setFilterToEdit={setFilterToEdit}
            />
        } else if (view === 'create') {
            return <CreateFilter
                listener={props.listener}
                setView={setView}
            />
        } else if (view === 'edit') {
            return <EditFilter
                listener={props.listener}
                filterToEdit={filterToEdit}
                setView={setView}
            />
        } else {
            return <></>
        }
    }

    return <Box>
        {render()}
    </Box>
}