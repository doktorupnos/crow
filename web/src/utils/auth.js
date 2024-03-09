import axios from "axios";

// Validate user session.
export const validSession = async () => {
  try {
    let response = await axios.post(
      process.env.authValidEndPoint,
      {},
      { withCredentials: true },
    );
    if (response.status == 200) {
      return true;
    } else {
      return false;
    }
  } catch (error) {
    throw error;
  }
};

// Register user.
export const userRegister = async (fields) => {
  try {
    let response = await axios.post(process.env.authRegEndPoint, fields, {
      withCredentials: true,
    });
    if (response.status == 201) {
      return true;
    } else {
      return false;
    }
  } catch (error) {
    throw error;
  }
};

// Login user.
export const userLogin = async (creds) => {
  try {
    let response = await axios.post(
      process.env.authLoginEndPoint,
      {},
      {
        auth: {
          username: `${creds.name}`,
          password: `${creds.password}`,
        },
        withCredentials: true,
      },
    );
    if (response.status == 200) {
      return true;
    } else {
      return false;
    }
  } catch (error) {
    throw error;
  }
};
