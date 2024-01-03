import axios from "axios";

export const fetchPosts = async (user, page) => {
	try {
		let response;
		if (user) {
			response = await axios.get(
				`${process.env.postGetEndPoint}?u=${user}&page=${page}`,
				{ withCredentials: true }
			);
			return response.data;
		} else {
			response = await axios.get(
				`${process.env.postGetEndPoint}?page=${page}`,
				{ withCredentials: true }
			);
		}
		if (response.status == 200) {
			return response.data;
		} else {
			return null;
		}
	} catch (error) {
		throw error;
	}
};
