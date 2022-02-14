import React, { useEffect } from 'react'
import { useDispatch } from 'react-redux'
import axios from './axiosConfig'

export default function GetServerLoad() {
    const dispatch = useDispatch()

    const fetchData = () => {
        axios.get('/serverLoad')
            .then(response => {
                dispatch({ type: 'SET_SERVER_LOAD', serverLoad: response.data })
                dispatch({ type: 'SET_SERVER_STATUS', lbStatus: "Active" })
            })
            .catch(error => {
                console.log(error)
                dispatch({ type: 'SET_SERVER_LOAD', serverLoad: { error: true } })
                dispatch({ type: 'SET_SERVER_STATUS', lbStatus: "Offline" })
            })
    }

    useEffect(() => {
        const intervalId = setInterval(() => {
            fetchData()
        }, 3000)

        return () => clearInterval(intervalId);
    }, [])

    return <></>
}