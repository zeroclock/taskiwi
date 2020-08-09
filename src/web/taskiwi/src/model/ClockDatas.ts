import { ClockData, isClockData } from './ClockData'

export type ClockDatas = ClockData[]

export function isClockDatas(arg: any): arg is ClockDatas {
  if (!Array.isArray(arg)) {
    return false
  }
  arg.map((data) => {
    if (!isClockData(data)) {
      return false
    }
  })
  return true
}
