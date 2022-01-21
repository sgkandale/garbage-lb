import { Close } from '@mui/icons-material'
import { IconButton } from '@mui/material'
import React, { useState } from 'react'
import ButtonStat from '../../../customComponents/buttonStat'

export default function DeleteEndpoint(props) {
    const [stat, setStat] = useState(0)

    const handleDelete = () => {
        setStat(2)
    }

    return <IconButton
        color="error"
        onClick={handleDelete}
        disabled={stat !== 0}
    >
        <ButtonStat
            stat={stat}
            zeroStat={<Close />}
        />
    </IconButton>
}