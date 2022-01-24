import axios from 'axios'

const instance = axios.create({
    // development
    // baseURL: 'http://localhost:8081'
    // production
    baseURL: ''
})

export default instance;