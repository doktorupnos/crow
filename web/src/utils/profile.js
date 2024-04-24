import axios from "axios";

// Fetch profile data.
export const fetchProfile = async (user) => {
  const endpoint = "http://api:8000/profile";
  try {
    if (!user) {
      const response = await axios.get(endpoint, {
        withCredentials: true,
      });
      return response.data;
    } else {
      const response = await axios.get(`${endpoint}?u=${user}`, {
        withCredentials: true,
      });
      return response.data;
    }
  } catch (error) {
    if (error.response.status == 401) {
      location.href = "/auth";
    }
    throw error;
  }
};

// Follow user.
export const followUser = async (uuid) => {
  try {
    let response = await axios.post(
      "http://api:8000/follow",
      { user_id: uuid },
      { withCredentials: true },
    );
    if (response.status == 200) {
      return true;
    } else {
      return false;
    }
  } catch (error) {
    if (error.response.status == 401) {
      location.href = "/auth";
    }
    throw error;
  }
};

// Unfollow user.
export const unfollowUser = async (uuid) => {
  try {
    let response = await axios.post(
      "http://api:8000/unfollow",
      { user_id: uuid },
      { withCredentials: true },
    );
    if (response.status == 200) {
      return true;
    } else {
      return false;
    }
  } catch (error) {
    if (error.response.status == 401) {
      location.href = "/auth";
    }
    throw error;
  }
};

// Fetch user follow list.
export const fetchFollow = async (name, page, type) => {
  const endpoint = "http://api:8000/followers";
  try {
    let response = await axios.get(
      type
        ? `${endpoint}?u=${name}&page=${page}`
        : `${endpoint}?u=${name}&page=${page}`,
      { withCredentials: true },
    );
    if (response.status == 200) {
      return response.data;
    } else {
      return null;
    }
  } catch (error) {
    throw error;
  }
};
