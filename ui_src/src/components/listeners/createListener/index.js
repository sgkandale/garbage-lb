import React, { useState } from 'react'
import { Box, Button, TextField, Typography, FormControl, InputLabel, Select, MenuItem } from '@mui/material'
import { ArrowBack } from '@mui/icons-material'
import ButtonStat from '../../customComponents/buttonStat'

export default function CreateListener(props) {

    const [listener, setListener] = useState({
        name: "",
        port: 80,
        connectionType: "",
    })
    const [submitStat, setSubmitStat] = useState(0)
    const [error, setError] = useState({
        msg: "",
    })

    const changeListenerValue = (event) => {
        setListener({ ...listener, [event.target.name]: event.target.value })
        setSubmitStat(0)
    }

    const handleSubmit = (event) => {
        event.preventDefault()
        setSubmitStat(2)
    }

    return <Box >
        <Button
            sx={{ textTransform: 'none', color: 'text.secondary' }}
            onClick={() => props.changeView('list')}
            startIcon={<ArrowBack />}
        >
            Back
        </Button>
        <br />
        <Box
            component="form"
            sx={{ width: '100%', marginTop: 3, marginBottom: 3, paddingLeft: 4 }}
            noValidate
            autoComplete="off"
            onSubmit={handleSubmit}
        >
            <TextField
                id="listener-name"
                label="Listener Name"
                variant="standard"
                name='name'
                sx={{ width: '100%', maxWidth: 500, marginBottom: 4 }}
                value={listener.name}
                onChange={changeListenerValue}
                type='name'
            />
            <br />
            <TextField
                id="listener-port"
                label="Listen on Port"
                variant="standard"
                name='port'
                sx={{ width: '100%', maxWidth: 500, marginBottom: 3 }}
                value={listener.port}
                onChange={changeListenerValue}
                type='number'
                inputProps={{
                    min: "11",
                    max: "65535"
                }}
            />
            <br />
            <FormControl variant="standard" sx={{ width: '100%', maxWidth: 500, marginBottom: 4 }}>
                <InputLabel id="connection-type-select-label">Connection Type</InputLabel>
                <Select
                    labelId="connection-type-select-label"
                    id="connection-type-select"
                    value={listener.connectionType}
                    onChange={changeListenerValue}
                    label="Age"
                    name='connectionType'
                >
                    <MenuItem value="">
                        <em>None</em>
                    </MenuItem>
                    <MenuItem value='http'>HTTP</MenuItem>
                </Select>
            </FormControl>
            <br />
            {
                error.msg !== "" ? <>
                    <Typography variant="caption" color="error" >
                        {error.msg}
                    </Typography>
                    <br />
                    <br />
                </> : <></>
            }
            <Button
                type="submit"
                variant="contained"
                color="primary"
                sx={{ textTransform: 'none' }}
                onClick={handleSubmit}
            >
                <ButtonStat
                    stat={submitStat}
                    zeroStat="Create"
                />
            </Button>
        </Box>
    </Box>
}