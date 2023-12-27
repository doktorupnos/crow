import axios from "axios";

// Fetch profile data.
export const getProfile = async (username) => {
	try {
		if (!username) {
			const response = await axios.post(
				process.env.postGetEndPoint,
				{},
				{ withCredentials: true }
			);
			console.log(response);
			return response.payload;
		} else {
			const response = await axios.post(
				`${process.env.postGetEndPoint}?u=${username}`,
				{},
				{ withCredentials: true }
			);
			console.log(response.payload);
			return response.payload;
		}
	} catch (error) {
		throw error;
	}
};
