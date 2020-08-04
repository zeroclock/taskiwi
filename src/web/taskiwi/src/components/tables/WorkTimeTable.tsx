import React from 'react'
import { WorkTimes } from '../../model/WorkTimes'
import { TableHead, TableRow, TableCell, Checkbox, TableContainer, TableBody, TablePagination, makeStyles, Table } from '@material-ui/core'

interface Props {
  worktimes: WorkTimes | null
}

const headers = [
  { id: 'tag', numeric: false, disablePadding: true, label: 'Tag' },
  { id: 'worktime', numeric: true, disablePadding: false, label: 'Work Time (minutes)' },
  { id: 'percent', numeric: true, disablePadding: false, label: 'Percent (%)' },
]

const createData = (tag: string, worktime: number, percent: number) => {
  return { tag, worktime, percent }
}

const rows = [
  createData('#1234', 15, 25),
  createData('maintenance', 100, 50),
  createData('#1234', 50, 25),
]

const useStyles = makeStyles((theme) => ({
  root: {
    maxWidth: '750px',
    margin: 'auto',
  },
  table: {
  },
}))

const WorkTimeTable: React.FC<Props> = (props: Props) => {
  console.log(props)
  const classes = useStyles()

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
            {rows.map((row) => {
              return (
                <TableRow
                  hover
                  key={row.tag}
                >
                  <TableCell component="th" scope="row" padding="none">
                    {row.tag}
                  </TableCell>
                  <TableCell align="right">{row.worktime}</TableCell>
                  <TableCell align="right">{row.percent}</TableCell>
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
        onChangePage={() => {}}
        onChangeRowsPerPage={() => {}}
      />
    </div>
  )
}

export default WorkTimeTable
