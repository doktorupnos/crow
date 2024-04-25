import axios from "axios";

// Fetch profile data.
export const fetchProfile = async (user) => {
  try {
    if (!user) {
      const response = await axios.get("//crow.zapto.org/api/profile", {
        withCredentials: true,
      });
      return response.data;
    } else {
      const response = await axios.get(
        `//crow.zapto.org/api/profile?u=${user}`,
        { withCredentials: true },
      );
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
      "//crow.zapto.org/api/follow",
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
      "//crow.zapto.org/api/unfollow",
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
  try {
    let response = await axios.get(
      type
        ? `//crow.zapto.org/api/followers?u=${name}&page=${page}`
        : `//crow.zapto.org/api/following?u=${name}&page=${page}`,
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
