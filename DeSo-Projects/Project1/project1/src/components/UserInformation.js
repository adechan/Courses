import { TextField, Typography, Button } from "@mui/material";
import React, { useEffect, useState } from "react";

const UserInformation = ({ desoApi, desoIdentity, publicKey }) => {
  const [userInfo, setUserInfo] = useState(null);
  const [newDescription, setNewDescription] = useState("");
  const [deso, setDeso] = useState(null);

  const getUserInfo = async () => {
    const result = await desoApi.getSingleProfile(publicKey, "");

    if (!result) {
      return;
    }

    setUserInfo(result?.Profile);
  };

  // Not sure why this gives me a response with Users as empty list
  const getUserStateless = async () => {
    const result = await desoApi.getUserStateless(publicKey);

    if (!result) {
      return;
    }

    console.log(result);
    setDeso(result);
  };

  useEffect(() => {
    if (!publicKey) {
      return;
    }

    getUserInfo();
    getUserStateless();
  }, []);

  useEffect(() => {
    setNewDescription(userInfo?.Description)
  }, [userInfo?.Description])

  const updateProfileDescription = async () => {
    const result = await desoApi.updateProfileDescription(
      publicKey,
      newDescription
    );
    
    const transactionHex = result.TransactionHex;
    const signedTransactionHex = await desoIdentity.signTxAsync(transactionHex);

    const rtnSubmitTransaction = await desoApi.submitTransaction(
      signedTransactionHex
    );

  };

  return (
    <div>
      <p>Welcome back, <b>@{userInfo?.Username}</b></p>
      <p>You have <b>_ $DESO</b></p>

      <div style={{ display: "flex", flexDirection: "column" }}>
        <Typography>Your Description is: </Typography>
        <div
          style={{
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
          }}
        >
          <TextField
            style={{ marginRight: 20 }}
            variant="outlined"
            placeholder="Description"
            value={newDescription}
            onChange={(e) => setNewDescription(e.target.value)}
          />
          <Button variant="contained" onClick={updateProfileDescription}>
            Save
          </Button>
        </div>
      </div>
    </div>
  );
};

export default UserInformation;
