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
