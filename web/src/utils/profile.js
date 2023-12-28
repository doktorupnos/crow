import axios from "axios";

// Fetch profile data.
export const getProfile = async (username) => {
	try {
		if (!username) {
			const response = await axios.get(process.env.profileEndPoint, {
				withCredentials: true,
			});
			return response.data;
		} else {
			const response = await axios.get(
				`${process.env.profileEndPoint}?u=${username}`,
				{ withCredentials: true }
			);
			return response.data;
		}
	} catch (error) {
		throw error;
	}
};

// Follow user.
export const followUser = async (uuid) => {
	try {
		let response = await axios.post(
			process.env.followEndPoint,
			{ user_id: uuid },
			{ withCredentials: true }
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

// Unfollow user.
export const unfollowUser = async (uuid) => {
	try {
		let response = await axios.post(
			process.env.unfollowEndPoint,
			{ user_id: uuid },
			{ withCredentials: true }
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

// Fetch user followers.
export const fetchFollowers = async (name) => {
	try {
		let response = await axios.get(
			`${process.env.followersEndPoint}?u=${name}`,
			{ withCredentials: true }
		);
		if (response.status == 200) {
			return response.data;
		} else {
			return false;
		}
	} catch (error) {
		throw error;
	}
};
