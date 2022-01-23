import React, { useEffect } from 'react'
import { useDispatch } from 'react-redux'
import axios from './axiosConfig'

export default function GetListeners() {
    const dispatch = useDispatch()

    const fetchData = () => {
        axios.get('/listener')
            .then(response => {
                dispatch({ type: 'SET_LISTENERS', serverLoad: response.data })
                dispatch({ type: 'SET_SERVER_STATUS', lbStatus: "Active" })
            })
            .catch(error => {
                console.log(error)
                dispatch({ type: 'SET_SERVER_STATUS', lbStatus: "Offline" })
            })
    }

    useEffect(() => {
        setInterval(() => {
            fetchData()
        }, 3000)
    }, [])

    return <></>
}