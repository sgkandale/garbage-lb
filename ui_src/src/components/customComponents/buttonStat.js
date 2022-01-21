import { ErrorOutline } from '@mui/icons-material'
import { CircularProgress } from '@mui/material'
import React from 'react'

export default function ButtonStat(props) {

    const render = () => {
        if (props.stat === 2) {
            return <CircularProgress
                size={25}
                color="inherit"
            />
        } else if (props.stat === 0) {
            return <>{props.zeroStat}</>
        } else if (props.stat === -1) {
            return <ErrorOutline color="error" />
        } else {
            return <></>
        }
    }

    return render()
}