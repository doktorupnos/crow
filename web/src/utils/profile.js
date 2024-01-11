import axios from "axios";

// Fetch profile data.
export const fetchProfile = async (user) => {
	try {
		if (!user) {
			const response = await axios.get(process.env.profileEndPoint, {
				withCredentials: true,
			});
			return response.data;
		} else {
			const response = await axios.get(
				`${process.env.profileEndPoint}?u=${user}`,
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

// Fetch user follow list.
export const fetchFollow = async (name, page, type) => {
	try {
		let response = await axios.get(
			type
				? `${process.env.followersEndPoint}?u=${name}&page=${page}`
				: `${process.env.followingEndPoint}?u=${name}&page=${page}`,
			{ withCredentials: true }
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
