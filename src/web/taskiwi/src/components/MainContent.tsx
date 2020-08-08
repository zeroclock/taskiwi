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
import {
  MuiPickersUtilsProvider,
  KeyboardDatePicker,
} from '@material-ui/pickers'
import DateFnsUtils from '@date-io/date-fns'
import { fetchTags } from '../api/tag'
import { fetchAggregation } from '../api/aggregation'
import { AggregateTaskReq } from '../interface/request'
import { Line } from 'react-chartjs-2'
import WorkTimeTable from './tables/WorkTimeTable'
import { WorkTimes } from '../model/WorkTimes'
import { Aggregation } from '../model/Aggregation'
import { format } from 'date-fns'
import BarChart from './charts/BarChart'
import DoughnutChart from './charts/DoughnutChart'

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

const emptyAggregation: Aggregation = {
  WorkTimes: [{
    tag: '',
    time: '',
    percent: '',
  }],
  ClockDatas: [{
    task: '',
    parents: '',
    category: '',
    start: '',
    end: '',
    effort: '',
    ishabit: '',
    tags: [],
  }]
}

function MainContent() {
  const classes = useStyles()
  const [tagName, setTagName] = React.useState<string[]>([])
  const [tagList, setTagList] = React.useState<string[]>([])
  const [worktimes, setWorktimes] = React.useState<WorkTimes>([])
  const [from, setFrom] = React.useState<Date | null>(new Date())
  const [to, setTo] = React.useState<Date | null>(new Date())

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
    const data = fetchAggregationReq({
      tags: values.length ? values : tagList,
      start: from != null ? format(from, 'yyyy-MM-dd') : '',
      end: to != null ? format(to, 'yyyy-MM-dd') : '',
    })
    data.then((aggregation) => {
      setWorktimes(aggregation.WorkTimes)
    })
  }

  const handleFromChange = (
    date: Date | null,
    _: string | null | undefined
  ) => {
    setFrom(date)
    const data = fetchAggregationReq({
      tags: tagName.length ? tagName : tagList,
      start: date != null ? format(date, 'yyyy-MM-dd') : '',
      end: to != null ? format(to, 'yyyy-MM-dd') : '',
    })
    data.then((aggregation) => {
      console.log(aggregation)
      setWorktimes(aggregation.WorkTimes)
    })
  }

  const handleToChange = (date: Date | null, _: string | null | undefined) => {
    setTo(date)
    const data = fetchAggregationReq({
      tags: tagName.length ? tagName : tagList,
      start: from != null ? format(from, 'yyyy-MM-dd') : '',
      end: date != null ? format(date, 'yyyy-MM-dd') : '',
    })
    data.then((aggregation) => {
      setWorktimes(aggregation.WorkTimes)
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

  const fetchAggregationReq = async (params: AggregateTaskReq) => {
    try {
      const { data } = await fetchAggregation(params)
      return data
    } catch (e) {
      console.log(e)
      return emptyAggregation
    }
  }

  useEffect(() => {
    const data = fetchTagsReq()
    const todayStr = format(new Date(), 'yyyy-MM-dd')
    data.then((tags) => {
      setTagList(tags)
      const data = fetchAggregationReq({
        tags: tags,
        start: todayStr,
        end: todayStr,
      })
      data.then((aggregation) => {
        setWorktimes(aggregation.WorkTimes)
      })
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
            <MuiPickersUtilsProvider utils={DateFnsUtils}>
              <Grid container xs={12}>
                <Grid item xs={6}>
                  <KeyboardDatePicker
                    disableToolbar
                    variant="inline"
                    format="yyyy/MM/dd"
                    margin="normal"
                    id="date-picker-inline"
                    label="From"
                    value={from}
                    onChange={handleFromChange}
                    KeyboardButtonProps={{
                      'aria-label': 'change from date',
                    }}
                  />
                </Grid>
                <Grid item xs={6}>
                  <KeyboardDatePicker
                    disableToolbar
                    variant="inline"
                    format="yyyy/MM/dd"
                    margin="normal"
                    id="date-picker-inline"
                    label="To"
                    value={to}
                    onChange={handleToChange}
                    KeyboardButtonProps={{
                      'aria-label': 'change to date',
                    }}
                  />
                </Grid>
              </Grid>
            </MuiPickersUtilsProvider>
          </FormControl>
        </Paper>
      </Grid>
      <Grid item xs={12} justify="center">
        <Paper variant="outlined" elevation={3}>
          <WorkTimeTable worktimes={worktimes} />
        </Paper>
      </Grid>
      <Grid item xs={6} justify="center">
        <Paper variant="outlined" elevation={3}>
          <BarChart worktimes={worktimes} />
        </Paper>
      </Grid>
      <Grid item xs={6} justify="center">
        <Paper variant="outlined" elevation={3}>
          <DoughnutChart worktimes={worktimes} />
        </Paper>
      </Grid>
    </Grid>
  )
}

export default MainContent
