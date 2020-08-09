import { WorkTimes } from '../../model/WorkTimes'

const generateRandomColorHex = (): string =>
  Math.floor(Math.random() * 16777215).toString(16)

export type ChartType = 'bar' | 'doughnut'

export interface DataSet {
  label: string
  fill: boolean
  lineTension: number
  backgroundColor: string[]
  borderColor: string
  borderCapStyle: string
  borderDash: string[]
  borderDashOffset: number
  borderJoinStyle: string
  pointBorderColor: string
  pointBackgroundColor: string
  pointBorderWidth: number
  pointHoverRadius: number
  pointHoverBackgroundColor: string
  pointHoverBorderColor: string
  pointHoverBorderWidth: number
  pointRadius: number
  pointHitRadius: number
  data: number[]
}

export interface ChartData {
  labels: string[]
  datasets: DataSet[]
}

export const createDataFromWorktimes = (
  worktimes: WorkTimes,
  type: ChartType
): ChartData => {
  const labels: string[] = []
  const data: number[] = []
  const colors: string[] = []
  worktimes.map((worktime) => {
    labels.push(worktime.tag)
    data.push(parseInt(worktime.time))
    colors.push(
      type != 'bar' ? '#' + generateRandomColorHex() : 'rgba(75,192,192,0.4)'
    )
  })
  const datasets = [
    {
      label: 'worktimes',
      fill: true,
      lineTension: 0.1,
      backgroundColor: colors,
      borderColor: 'rgba(75,192,192,1)',
      borderCapStyle: 'round',
      borderDash: [],
      borderDashOffset: 0.0,
      borderJoinStyle: 'square',
      pointBorderColor: 'rgba(75,192,192,1)',
      pointBackgroundColor: '#eee',
      pointBorderWidth: 10,
      pointHoverRadius: 5,
      pointHoverBackgroundColor: 'rgba(75,192,192,1)',
      pointHoverBorderColor: 'rgba(220,220,220,1)',
      pointHoverBorderWidth: 1,
      pointRadius: 1,
      pointHitRadius: 10,
      data: data,
    },
  ]
  return { labels, datasets }
}
