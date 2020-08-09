import client from './client'
import { TASK_BY_DATE } from '../constants/api_endpoints'
import { ClockDatas, isClockDatas } from '../model/ClockDatas'

export const fetchTasksByDate = async (date: string): Promise<ClockDatas> => {
  const data = await client.post(TASK_BY_DATE, { date })
  if (isClockDatas(data.data)) {
    return data.data
  } else {
    return []
  }
}
