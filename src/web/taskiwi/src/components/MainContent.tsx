import React, { useEffect } from 'react'
import {
  makeStyles,
  createStyles,
  Theme,
  Grid,
  Paper,
  MenuItem,
  Checkbox,
  ListItemText,
  Input,
  InputLabel,
  FormControl,
  Select,
  Chip,
} from '@material-ui/core'
import { fetchTags } from '../api/tag'
import { fetchWorkTimes } from '../api/workTimes'
import { AggregateTaskReq } from '../interface/request'
import { Line } from 'react-chartjs-2'
import WorkTimeTable from './tables/WorkTimeTable'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
      padding: '10px',
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: 'center',
      '& > *': {
        margin: theme.spacing(3),
      },
    },
    formControl: {
      margin: theme.spacing(1),
      width: 500,
    },
    chips: {
      display: 'flex',
      flexWrap: 'wrap',
    },
    chip: {
      margin: 2,
    },
  })
)

const ITEM_HEIGHT = 48
const ITEM_PADDING_TOP = 8
const MenuProps = {
  PaperProps: {
    style: {
      maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP,
      width: 250,
    },
  },
}

const data = {
  labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
  datasets: [
    {
      label: 'My First dataset',
      fill: true,
      lineTension: 0.1,
      backgroundColor: 'rgba(75,192,192,0.4)',
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
      data: [3, 10, 21, 31, 34, 40, 48]
    }
  ]
}

function MainContent() {
  const classes = useStyles()
  const [tagName, setTagName] = React.useState<string[]>([])
  const [tagList, setTagList] = React.useState<string[]>([])

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const v = event.target.value
    const values: string[] = []
    if (v instanceof Array && v !== null) {
      v.forEach((tag) => {
        if (typeof tag === 'string') {
          values.push(tag)
        }
      })
      setTagName(values)
    }
    const data = fetchWorkTimesReq({
      tags: values,
      start: '2020-08-03',
      end: '2020-08-04'
    })
    data.then((workTimes) => {
      console.log(workTimes)
    })
  }

  const fetchTagsReq = async () => {
    try {
      const { data } = await fetchTags()
      return data
    } catch (e) {
      console.log(e)
      return []
    }
  }

  const fetchWorkTimesReq = async (params: AggregateTaskReq) => {
    try {
      const { data } = await fetchWorkTimes(params)
      return data
    } catch (e) {
      console.log(e)
      return []
    }
  }

  useEffect(() => {
    const data = fetchTagsReq()
    data.then((tags) => {
      setTagList(tags)
    })
  }, [])
  
  return (
    <Grid container className={classes.root} spacing={3}>
      <Grid item xs={12} justify="center">
        <Paper variant="outlined" elevation={3} className={classes.paper}>
          <FormControl className={classes.formControl}>
            <InputLabel id="tag-multiple-checkbox-label">Tag</InputLabel>
            <Select
              labelId="tag-multiple-checkbox-label"
              id="tag-multiple-checkbox"
              multiple
              value={tagName}
              onChange={handleChange}
              input={<Input />}
              renderValue={(selected: any) => (
                <div className={classes.chips}>
                  {selected.map((value: any) => (
                    <Chip key={value} label={value} className={classes.chip} />
                  ))}
                </div>
              )}
              MenuProps={MenuProps}
            >
              {tagList.map((tag) => (
                <MenuItem key={tag} value={tag}>
                  <Checkbox checked={tagName.indexOf(tag) > -1} />
                  <ListItemText primary={tag} />
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </Paper>
      </Grid>
      <Grid item xs={12} justify="center">
        <Paper variant="outlined" elevation={3}>
          <WorkTimeTable worktimes={null} />
        </Paper>
      </Grid>
      <Grid item xs={6} justify="center">
        <Paper variant="outlined" elevation={3}>
          <Line data={data} />
        </Paper>
      </Grid>
      <Grid item xs={6} justify="center">
        <Paper variant="outlined" elevation={3}>
          <Line data={data} />
        </Paper>
      </Grid>
    </Grid>
  )
}

export default MainContent
