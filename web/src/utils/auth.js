import axios from "axios";

// Validate user session.
export const validSession = async () => {
  try {
    let response = await axios.post(
      "http://api:8000/admin/jwt",
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
    let response = await axios.post("http://api:8000/users", fields, {
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
      "http://api:8000/login",
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
