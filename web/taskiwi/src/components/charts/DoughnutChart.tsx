import React from 'react'
import { WorkTimes } from '../../model/WorkTimes'
import { Doughnut } from 'react-chartjs-2'
import { createDataFromWorktimes } from './utils'

export interface Props {
  worktimes: WorkTimes
}

const DoughnutChartt: React.FC<Props> = (props: Props) => {
  const data = createDataFromWorktimes(props.worktimes, 'doughnut')

  return <Doughnut data={data} />
}

export default DoughnutChartt
