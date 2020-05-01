import axios from 'axios'
import constants from './constants'

const localApi = axios.create({
    baseURL: constants.API_BASE_URL,
    withCredentials: false
});

export default localApi;
