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
import { fetchTags } from '../api/tag';

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
      minWidth: 120,
      maxWidth: 300,
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

const ITEM_HEIGHT = 48;
const ITEM_PADDING_TOP = 8;
const MenuProps = {
  PaperProps: {
    style: {
      maxHeight: ITEM_HEIGHT * 4.5 + ITEM_PADDING_TOP,
      width: 250,
    },
  },
}

function MainContent() {
  const classes = useStyles()
  const [tagName, setTagName] = React.useState<string[]>([])
  const [tagList, setTagList] = React.useState<string[]>([])

  const handleChange = (event: React.ChangeEvent<{ name?: string; value: unknown }>) => {
    const v = event.target.value
    var values: string[] = []
    if (v instanceof Array && v !== null) {
      v.forEach((tag) => {
        if (typeof tag === 'string') {
          values.push(tag)
        }
      })
      setTagName(values)
    }
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

  useEffect(() => {
    const data = fetchTagsReq()
    data.then(tags => {
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
    </Grid>
  )
}

export default MainContent
