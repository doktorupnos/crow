import axios from "axios";

// Fetch user posts.
export const fetchPosts = async (user, page, limit) => {
  try {
    let response;
    if (user) {
      response = await axios.get(
        `${process.env.postGetEndPoint}?u=${user}&page=${page}`,
        { withCredentials: true },
      );
      return response.data;
    } else {
      response = await axios.get(
        `${process.env.postGetEndPoint}?page=${page}&limit=${limit}`,
        { withCredentials: true },
      );
    }
    if (response.status == 200) {
      return response.data;
    } else {
      return null;
    }
  } catch (error) {
    if (error.response.status == 401) {
      location.href = "/auth";
    }
    throw error;
  }
};

// Like user post.
export const postLike = async (id) => {
  try {
    let response = await axios.post(
      process.env.postLikeEndPoint,
      { post_id: id },
      { withCredentials: true },
    );
    if (response.status == 201) {
      return true;
    } else {
      return null;
    }
  } catch (error) {
    if (error.response.status == 401) {
      location.href = "/auth";
    }
    throw error;
  }
};

// Remove like from user post.
export const postUnlike = async (id) => {
  try {
    let response = await axios.delete(process.env.postLikeEndPoint, {
      data: { post_id: id },
      withCredentials: true,
    });
    if (response.status == 200) {
      return true;
    } else {
      return null;
    }
  } catch (error) {
    if (error.response.status == 401) {
      location.href = "/auth";
    }
    throw error;
  }
};

// Create user post.
export const postCreate = async (body) => {
  try {
    let response = await axios.post(
      process.env.postGetEndPoint,
      { body: body },
      { withCredentials: true },
    );
    if (response.status == 201) {
      return true;
    } else {
      return null;
    }
  } catch (error) {
    if (error.response.status == 401) {
      location.href = "/auth";
    }
    throw error;
  }
};

// Delete user post.
export const postDelete = async (id) => {
  try {
    let response = await axios.delete(`${process.env.postGetEndPoint}/${id}`, {
      withCredentials: true,
    });
    if (response.status == 200) {
      return true;
    } else {
      return null;
    }
  } catch (error) {
    if (error.response.status == 401) {
      location.href = "/auth";
    }
    throw error;
  }
};

// Format post timestamp.
export const postTime = (timestamp) => {
  let timeDiff = Math.floor(Date.now() / 1000) - timestamp;
  if (timeDiff < 60) {
    return `${timeDiff} seconds ago.`;
  } else if (timeDiff < 3600) {
    return `${Math.floor(timeDiff / 60)} minutes ago.`;
  } else if (timeDiff < 86400) {
    return `${Math.floor(timeDiff / 3600)} hours ago.`;
  }
  let date = new Date(timestamp * 1000);
  let dateObj = {
    year: date.getFullYear(),
    month: ("0" + (date.getMonth() + 1)).slice(-2),
    day: ("0" + date.getDate()).slice(-2),
    hours: ("0" + date.getHours()).slice(-2),
    minutes: ("0" + date.getMinutes()).slice(-2),
  };
  return `${dateObj.year}-${dateObj.month}-${dateObj.day} ${dateObj.hours}:${dateObj.minutes}`;
};
