import axios from "axios";

// Fetch user posts.
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
