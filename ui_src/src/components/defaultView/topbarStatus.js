import { Paper, MenuItem, Menu, Button, Typography, Grid } from '@mui/material'
import React, { useState } from 'react'
import { useSelector } from 'react-redux'

export default function TopbarStatus() {
    const lbStatus = useSelector(state => state.lbStatus)

    const statusOptions = [
        {
            name: 'Active',
            dotColor: '#42D948',
        },
        {
            name: 'Rejecting',
            dotColor: '#F8FF00',
        },
        {
            name: 'Terminated',
            dotColor: '#FF0000',
        },
        {
            name: 'Unknown',
            dotColor: '#000000',
        }
    ]

    const statusComponent = (status) => {
        return <Grid
            container
            direction="row"
            justifyContent="center"
            alignItems="center"
            sx={{ maxWidth: 200 }}
        >
            <Typography variant="subtitle2" color="textSecondary" style={{ marginRight: 5 }} >
                Status :
            </Typography>
            <div style={{ width: 10, height: 10, backgroundColor: status.dotColor, borderRadius: '50%', marginRight: 5 }} />
            <Typography variant="subtitle2" color="textSecondary">
                {status.name}
            </Typography>
        </Grid>
    }

    const renderStatus = () => {
        for (let i = 0; i < statusOptions.length; i++) {
            if (statusOptions[i].name === lbStatus) {
                return statusComponent(statusOptions[i])
            }
        }
        return statusComponent(statusOptions[statusOptions.length - 1])
    }

    return renderStatus()
}