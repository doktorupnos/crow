import axios from "axios";

// Validate user session.
export const validSession = async () => {
	try {
		let response = await axios.post(
			process.env.authValidEndPoint,
			{},
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
