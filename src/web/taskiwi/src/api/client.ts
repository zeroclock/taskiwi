import axios, { AxiosInstance } from 'axios'
import { API_ROOT_PATH } from '../constants/api_endpoints'

let client: AxiosInstance

export default client = axios.create({
  baseURL: API_ROOT_PATH,
  headers: {
    'Content-Type': 'application/json',
  },
})
