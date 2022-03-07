import axios from "axios";

const DEFAULT_NODE_URL = "http://localhost:18001/api";
let client = null;

class DesoApi {
  constructor() {
    this.client = null;
    this.baseUrl = DEFAULT_NODE_URL;
  }

  async getSingleProfile(publicKey, username) {
    if (!publicKey) {
      console.log("publicKey or username is required");
      return;
    }

    const path = "/v0/get-single-profile";
    const data = {
      PublicKeyBase58Check: publicKey,
      Username: username,
    };

    try {
      const result = await this.getClient().post(path, data);
      return result.data;
    } catch (error) {
      console.log(error);
      return null;
    }
  }

  async updateProfileDescription(publicKey, newDescription) {
    if (!publicKey) {
      console.log("publicKey or username is required");
      return;
    }

    const path = "/v0/update-profile";
    const data = {
      ProfilePublicKeyBase58Check: "",
      UpdaterPublicKeyBase58Check: publicKey,
      NewDescription: newDescription,
      NewProfilePic: "",
      NewUsername: "",
      NewStakeMultipleBasisPoints: 12500,
      MinFeeRateNanosPerKB: 1000,
      NewCreatorBasisPoints: 1000,
      IsHidden: false,
    };

    try {
      const result = await this.getClient().post(path, data);
      return result.data;
    } catch (error) {
      console.log(error);
      return null;
    }
  }

  async submitTransaction(signedTransactionHex) {
    if (!signedTransactionHex) {
      console.log("signedTransactionHex is required");
      return;
    }

    const path = "/v0/submit-transaction";
    const data = {
      TransactionHex: signedTransactionHex,
    };
    try {
      const result = await this.getClient().post(path, data);
      return result.data;
    } catch (error) {
      console.log(error);
      return null;
    }
  }

  getClient() {
    if (client) return client;
    client = axios.create({
      baseURL: this.baseUrl,
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
    });
    return client;
  }
}

export default DesoApi;
