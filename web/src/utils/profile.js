import axios from "axios";

// Fetch profile data.
export const getProfile = async (username) => {
	try {
		if (!username) {
			const response = await axios.get(process.env.profileEndPoint, {
				withCredentials: true,
			});
			console.log(response);
			return response.payload;
		} else {
			const response = await axios.get(
				`${process.env.profileEndPoint}?u=${username}`,
				{ withCredentials: true }
			);
			console.log(response.payload);
			return response.payload;
		}
	} catch (error) {
		throw error;
	}
};
