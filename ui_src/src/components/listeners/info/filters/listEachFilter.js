import React from 'react'
import { Typography, Card, CardContent, Grid, Box, IconButton, Tooltip } from '@mui/material'
import { Edit } from '@mui/icons-material'
import FilterRules from './filterRules'

export default function EachFilter(props) {

    const changeToEditView = () => {
        props.setFilterToEdit(props.filter.id)
        props.setView('edit')
    }

    return <Card elevation={2}>
        <CardContent>
            <Grid
                container
                direction="row"
                justifyContent="space-between"
                alignItems="flex-start"
            >
                <Box >
                    <Typography variant="subtitle1">
                        {props.filter.name} (id : {props.filter.id})
                    </Typography>
                </Box>
                <Tooltip title="Add Rule">
                    <IconButton onClick={changeToEditView}>
                        <Edit />
                    </IconButton>
                </Tooltip>
            </Grid>
            <Typography variant="subtitle1" style={{ marginTop: 10, marginBottom: 10 }}>
                Rules :
            </Typography>
            <FilterRules rules={props.filter.rules} />
        </CardContent>
    </Card>
}