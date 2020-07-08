import React from 'react'
import {
  makeStyles,
  createStyles,
  Theme,
  Grid,
  Paper,
  Button,
  Fab,
  Badge,
} from '@material-ui/core'

import MailIcon from '@material-ui/icons/Mail'
import ShareIcon from '@material-ui/icons/Share'
import ListAlt from '@material-ui/icons/ListAlt'
import PersonAdd from '@material-ui/icons/PersonAdd'
import Lock from '@material-ui/icons/Lock'
import Chat from '@material-ui/icons/Chat'
import Assessment from '@material-ui/icons/Assessment'
import CloudUpload from '@material-ui/icons/CloudUpload'
import AssignmentTurnedIn from '@material-ui/icons/AssignmentTurnedIn'
import AddIcon from '@material-ui/icons/Add'
import EditIcon from '@material-ui/icons/Edit'
import FavoriteIcon from '@material-ui/icons/Favorite'

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
  })
)

function MainContent() {
  const classes = useStyles()
  return (
    <Grid container className={classes.root} spacing={3}>
      <Grid item xs={12} justify="center">
        <Paper variant="outlined" elevation={3} className={classes.paper}>
          <Button variant="contained">Default</Button>
          <Button variant="contained" color="primary">
            Primary
          </Button>
          <Button variant="contained" color="secondary">
            Secondary
          </Button>
          <Button variant="contained" disabled>
            Disabled
          </Button>
          <Button variant="contained" color="primary" href="#contained-buttons">
            Link
          </Button>
        </Paper>
      </Grid>
      <Grid item xs={12} justify="center">
        <Paper variant="outlined" elevation={3} className={classes.paper}>
          <Fab color="primary" aria-label="add">
            <AddIcon />
          </Fab>
          <Fab color="secondary" aria-label="edit">
            <EditIcon />
          </Fab>
          <Fab disabled aria-label="like">
            <FavoriteIcon />
          </Fab>
        </Paper>
      </Grid>
      <Grid item xs={12} justify="center">
        <Paper variant="outlined" elevation={3} className={classes.paper}>
          <Badge badgeContent={4} color="primary">
            <MailIcon fontSize="small" />
          </Badge>
          <Badge badgeContent={3} color="secondary">
            <MailIcon />
          </Badge>
          <Badge badgeContent={2} color="error">
            <MailIcon fontSize="large" />
          </Badge>
        </Paper>
      </Grid>
      <Grid item xs={12} justify="center">
        <Paper variant="outlined" elevation={3} className={classes.paper}>
          <ShareIcon />
          <ListAlt />
          <PersonAdd />
          <Lock />
          <Chat />
          <Assessment />
          <CloudUpload />
          <AssignmentTurnedIn />
        </Paper>
      </Grid>
    </Grid>
  )
}

export default MainContent
