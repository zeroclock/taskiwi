import axios from 'axios'
import { API_ROOT_PATH } from '../constants/api_endpoints'

const client = axios.create({
  baseURL: API_ROOT_PATH,
  headers: {
    'Content-Type': 'application/json',
  },
})

export default client
