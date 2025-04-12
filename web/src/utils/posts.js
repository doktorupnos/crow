import axios from "axios";
import { formatDistanceToNow } from "date-fns";

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
  const date = new Date(timestamp);
  return trimPrefix(formatDistanceToNow(date) + " ago", "about");
};

const trimPrefix = (str, prefix) => {
  return str.startsWith(prefix) ? str.slice(prefix.length) : str;
};
