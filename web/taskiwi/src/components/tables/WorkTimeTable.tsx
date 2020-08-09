import React from 'react'
import { WorkTimes } from '../../model/WorkTimes'
import {
  TableHead,
  TableRow,
  TableCell,
  TableContainer,
  TableBody,
  TablePagination,
  makeStyles,
  Table,
} from '@material-ui/core'

interface Props {
  worktimes: WorkTimes | null
}

interface Row {
  tag: string
  worktime: string
  percent: string
}

const headers = [
  { id: 'tag', numeric: false, disablePadding: true, label: 'Tag' },
  {
    id: 'worktime',
    numeric: true,
    disablePadding: false,
    label: 'Work Time (H:m)',
  },
  { id: 'percent', numeric: true, disablePadding: false, label: 'Percent (%)' },
]

const minuteToHM = (minute: number): string => {
  const hours = Math.floor(minute / 60)
  const min = minute - hours * 60
  return hours + ':' + ('00' + min.toString()).slice(-2)
}

const createRowsFromProps = (props: Props): Row[] => {
  if (props.worktimes != null) {
    let totalTime = 0
    let totalPercent = 0.0
    const rows = props.worktimes.map((worktime) => {
      totalTime += parseInt(worktime.time)
      totalPercent += parseFloat(worktime.percent)
      return {
        tag: worktime.tag,
        worktime: minuteToHM(parseInt(worktime.time)),
        percent: worktime.percent,
      }
    })
    rows.push({
      tag: 'TOTAL',
      worktime: minuteToHM(totalTime).toString(),
      percent: totalPercent.toString(),
    })
    return rows
  } else {
    return []
  }
}

const useStyles = makeStyles(() => ({
  root: {
    maxWidth: '750px',
    margin: 'auto',
  },
  table: {},
}))

const WorkTimeTable: React.FC<Props> = (props: Props) => {
  const classes = useStyles()
  const rows = createRowsFromProps(props)

  return (
    <div className={classes.root}>
      <TableContainer>
        <Table
          className={classes.table}
          aria-labelledby="worktimeTable"
          size={'medium'}
          aria-label="worktime table"
        >
          <TableHead>
            <TableRow>
              {headers.map((header) => (
                <TableCell
                  key={header.id}
                  align={header.numeric ? 'right' : 'left'}
                  padding={header.disablePadding ? 'none' : 'default'}
                  sortDirection={false}
                >
                  {header.label}
                </TableCell>
              ))}
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map((row: any) => {
              return (
                <TableRow hover key={row.tag}>
                  <TableCell component="th" scope="row" padding="none">
                    {row.tag}
                  </TableCell>
                  <TableCell align="right">{row.worktime}</TableCell>
                  <TableCell align="right">
                    {Math.round(row.percent * 100)}
                  </TableCell>
                </TableRow>
              )
            })}
          </TableBody>
        </Table>
      </TableContainer>
      <TablePagination
        rowsPerPageOptions={[5, 10, 25]}
        component="div"
        count={rows.length}
        rowsPerPage={5}
        page={1}
        onChangePage={(): void => {
          console.log('TODO: implement table pagination')
        }}
        onChangeRowsPerPage={(): void => {
          console.log('TODO: implement table pagination')
        }}
      />
    </div>
  )
}

export default WorkTimeTable
