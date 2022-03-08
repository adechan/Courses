import React from 'react'
import { Box, Button } from "@mui/material";

const LogInOut = ({loggedIn, login, logout}) => {
  return (
    <Box sx={{ mb: 2, mt: 2 }}>
    {loggedIn ? (
      <Button variant="contained" onClick={logout}>
        Log Out
      </Button>
    ) : (
      <Button variant="contained" onClick={login}>
        Login
      </Button>
    )}
  </Box>
  )
}

export default LogInOut