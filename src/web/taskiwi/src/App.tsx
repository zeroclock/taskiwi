import React from 'react'
// import logo from './logo.svg'
import './App.css'
import green from '@material-ui/core/colors/green'

import MyAppBar from './components/MyAppBar'
import MainContent from './components/MainContent'
import {
  createMuiTheme,
  MuiThemeProvider,
  CssBaseline,
} from '@material-ui/core'

const theme = createMuiTheme({
  palette: {
    primary: green,
  },
  typography: {
    fontFamily: ['Noto Sans', 'sans-serif'].join(','),
    fontSize: 12,
    h1: {
      fontSize: '1.75rem',
    },
    h2: {
      fontSize: '1.5rem',
    },
    h3: {
      fontSize: '1.25rem',
    },
    h4: {
      fontSize: '1.125rem',
    },
    h5: {
      fontSize: '1rem',
    },
    h6: {
      fontSize: '1rem',
    },
  },
})

const App: React.FC = () => {
  return (
    <MuiThemeProvider theme={theme}>
      <CssBaseline />
      <MyAppBar />
      <MainContent />
    </MuiThemeProvider>
  )
}

export default App
