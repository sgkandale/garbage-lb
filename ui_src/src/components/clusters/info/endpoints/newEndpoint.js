import React, { useState } from 'react'
import { Box, Button, TextField, Typography } from '@mui/material'
import ButtonStat from '../../../customComponents/buttonStat'

export default function NewEndpoint(props) {
    const [endpoint, setEndpoint] = useState({
        name: '',
        address: '',
    })
    const [stat, setStat] = useState(0)

    const changeClusterValue = (event) => {
        setEndpoint({ ...endpoint, [event.target.name]: event.target.value })
    }

    const handleSubmit = (event) => {
        event.preventDefault()
        setStat(2)
        if (stat === 0) {
            // handle
        }
    }

    return <Box
        component="form"
        sx={{ width: '100%' }}
        noValidate
        autoComplete="off"
        onSubmit={handleSubmit}
    >
        <Typography variant="h6" >
            Register new endpoint
        </Typography>
        <TextField
            id="endpoint-name"
            label="Endpoint Name"
            variant="standard"
            name='name'
            sx={{ width: '100%', maxWidth: 450, marginBottom: 3, marginTop: 2 }}
            value={endpoint.name}
            onChange={changeClusterValue}
            type='name'
        />
        <br />
        <TextField
            id="endpoint-address"
            label="Endpoint Address"
            variant="standard"
            name='address'
            sx={{ width: '100%', maxWidth: 450, marginBottom: 3 }}
            value={endpoint.address}
            onChange={changeClusterValue}
        />
        <br />
        <TextField
            id="endpoint-port"
            label="Port"
            variant="standard"
            name='port'
            sx={{ width: '100%', maxWidth: 450, marginBottom: 3 }}
            value={endpoint.port}
            onChange={changeClusterValue}
            type='number'
            inputProps={{
                min: "11",
                max: "65535"
            }}
        />
        <br />
        <Button
            variant="contained"
            color="primary"
            type='submit'
            onClick={handleSubmit}
            style={{ textTransform: 'none' }}
        >
            <ButtonStat
                stat={stat}
                zeroStat="Register"
            />
        </Button>
    </Box>
}