import React from 'react'
import { WorkTimes } from '../../model/WorkTimes'
import { Bar } from 'react-chartjs-2'
import { createDataFromWorktimes } from './utils'

export interface Props {
  worktimes: WorkTimes
}

const BarChart: React.FC<Props> = (props: Props) => {
  const data = createDataFromWorktimes(props.worktimes, 'bar')

  return <Bar data={data} />
}

export default BarChart
