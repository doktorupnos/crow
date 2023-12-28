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
		const response = await axios.post(
			process.env.profileEndPoint,
			{
				user_id: uuid,
			},
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
