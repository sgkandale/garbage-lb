import React, { useState } from 'react'
import { Box, Button, TextField, Typography } from '@mui/material'
import { ArrowBack } from '@mui/icons-material'
import ButtonStat from '../../customComponents/buttonStat'

export default function CreateCluster(props) {

    const [cluster, setCluster] = useState({
        name: "",
    })
    const [submitStat, setSubmitStat] = useState(0)
    const [error, setError] = useState({
        msg: "",
    })

    const changeClusterValue = (event) => {
        setCluster({ ...cluster, [event.target.name]: event.target.value })
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
                id="cluster-name"
                label="Cluster Name"
                variant="standard"
                name='name'
                sx={{ width: '100%', maxWidth: 500, marginBottom: 4 }}
                value={cluster.name}
                onChange={changeClusterValue}
                type='name'
            />
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