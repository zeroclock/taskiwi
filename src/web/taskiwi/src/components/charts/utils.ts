import { WorkTimes } from '../../model/WorkTimes'

const generateRandomColorHex = () => Math.floor(Math.random()*16777215).toString(16);

export type ChartType = "bar" | "doughnut"

export const createDataFromWorktimes = (worktimes: WorkTimes, type: ChartType) => {
  var labels: string[] = []
  var data: number[] = []
  var colors: string[] = []
  worktimes.map((worktime) => {
    labels.push(worktime.tag)
    data.push(parseInt(worktime.time))
    colors.push((type != 'bar') ? '#' + generateRandomColorHex() : 'rgba(75,192,192,0.4)')
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
  return {labels, datasets}
}
