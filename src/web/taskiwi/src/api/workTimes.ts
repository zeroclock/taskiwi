import client from './client'
import { AGGREGATE_TASK_PATH } from '../constants/api_endpoints'
import { AxiosPromise } from 'axios'
import { WorkTimes } from '../model/WorkTimes'
import { AggregateTaskReq } from '../interface/request'

export const fetchWorkTimes = (params: AggregateTaskReq): AxiosPromise<WorkTimes> => client.post(AGGREGATE_TASK_PATH, params)
