import { Box } from '@mui/material'
import React, { useState } from 'react'
import ListRules from './listRules'

export default function Rules(props) {
    const [view, setView] = useState('list')

    const render = () => {
        if (view === 'list') {
            return <ListRules
                listener={props.listener}
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