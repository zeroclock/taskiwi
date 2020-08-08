import client from './client'
import { AGGREGATE_TASK_PATH } from '../constants/api_endpoints'
import { AxiosPromise } from 'axios'
import { Aggregation } from '../model/Aggregation'
import { AggregateTaskReq } from '../interface/request'

export const fetchAggregation = (
  params: AggregateTaskReq
): AxiosPromise<Aggregation> => client.post(AGGREGATE_TASK_PATH, params)
