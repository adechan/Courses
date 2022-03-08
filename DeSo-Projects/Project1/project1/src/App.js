import React, { useState, useEffect } from "react";
import { Box, Grid, Stack } from "@mui/material";
import DesoIdentity from "./helpers/desoIdentity";
import DesoApi from "./helpers/desoApi";
import UserInformation from "./components/UserInformation";
import LogInOut from "./components/LogInOut";

const IdentityUsersKey = "identityUsersV2";

function App() {
  const [loggedIn, setLoggedIn] = useState(false);

  const [publicKey, setSetPublicKey] = useState(null);
  const [desoIdentity, setDesoIdentity] = useState(null);
  const [desoApi, setDesoApi] = useState(null);

  useEffect(() => {
    const di = new DesoIdentity();
    setDesoIdentity(di);
    const da = new DesoApi();
    setDesoApi(da);

    let user = {};
    if (localStorage.getItem(IdentityUsersKey) === "undefined") {
      user = {};
    } else if (localStorage.getItem(IdentityUsersKey)) {
      user = JSON.parse(localStorage.getItem(IdentityUsersKey) || "{}");
    }

    if (user.publicKey) {
      setLoggedIn(true);
      setSetPublicKey(user.publicKey);
    }
  }, []);

  const login = async () => {
    const user = await desoIdentity.loginAsync(4);
    setSetPublicKey(user.publicKey);
    setLoggedIn(true);
  };

  const logout = async () => {
    const result = await desoIdentity.logout(publicKey);
    setSetPublicKey(null);
    setLoggedIn(false);
  };

  return (
    <>
      <iframe
        title="desoidentity"
        id="identity"
        frameBorder="0"
        src="https://identity.deso.org/embed?v=2"
        style={{
          height: "100vh",
          width: "100vw",
          display: "none",
          position: "fixed",
          zIndex: 1000,
          left: 0,
          top: 0,
        }}
      ></iframe>
      <Grid container>
        <Grid
          item
          sm={12}
          lg={4}
          sx={{
            alignItems: "center",
            justifyContent: "center",
            display: "flex",
          }}
        >
          <Stack>
            <LogInOut loggedIn={loggedIn} login={login} logout={logout} />

            {loggedIn && (
              <Box sx={{ mb: 2 }}>
                <UserInformation
                  desoApi={desoApi}
                  desoIdentity={desoIdentity}
                  publicKey={publicKey}
                />
              </Box>
            )}
          </Stack>
        </Grid>
        <Grid item sm={0} lg={4}></Grid>
      </Grid>
    </>
  );
}

export default App;
