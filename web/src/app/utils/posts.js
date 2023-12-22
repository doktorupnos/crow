import axios from "axios";

// Fetch posts. (re-do)
export const fetchPosts = async (page) => {
	var fetchPosts = {
		auth: false,
		payload: [],
	};
	await axios
		.get(`${process.env.postGetEndPoint}?page=${page}`, {
			withCredentials: true,
		})
		.then((response) => {
			if (response.status == 200) {
				fetchPosts.auth = true;
				fetchPosts.payload = response.data;
			} else if (response.status == 401) {
				console.error("Session expired!");
			}
		})
		.catch((error) => {
			console.error("Failed to fetch posts!", error);
		});
	return fetchPosts;
};

// Create post.
export const createPost = async (body) => {
	try {
		const response = await axios.post(
			process.env.postGetEndPoint,
			{ body: body },
			{ withCredentials: true }
		);
		if (response.status == 201) {
			return true;
		} else if (response.status == 401) {
			console.error("Session expired!");
			return false;
		}
	} catch (error) {
		throw error;
	}
};

// Add like to post.
export const addLike = async (id) => {
	var addLike = false;
	await axios
		.post(
			process.env.postLikeEndPoint,
			{ post_id: id },
			{ withCredentials: true }
		)
		.then((response) => {
			if (response.status == 201) {
				addLike = true;
			} else if (response.status == 401) {
				console.error("Session expired!");
			}
		})
		.catch((error) => {
			console.error("Failed to like post!", error);
		});
	return addLike;
};

// Remove like from post.
export const remLike = async (id) => {
	var remLike = false;
	await axios
		.delete(process.env.postLikeEndPoint, {
			data: { post_id: id },
			withCredentials: true,
		})
		.then((response) => {
			if (response.status == 200) {
				remLike = true;
			} else if (response.status == 401) {
				console.error("Session expired!");
			}
		})
		.catch((error) => {
			console.error("Failed to remove like!", error);
		});
	return remLike;
};

export const getPostTime = (timestamp) => {
	let getPostTime = "";

	let timeDiff = Math.floor(Date.now() / 1000) - timestamp;
	if (timeDiff < 60) {
		// Seconds format.
		getPostTime = `${timeDiff} seconds ago.`;
	} else if (timeDiff < 3600) {
		// Minutes format.
		getPostTime = `${Math.floor(timeDiff / 60)} minutes ago.`;
	} else if (timeDiff < 86400) {
		// Hours format.
		getPostTime = `${Math.floor(timeDiff / 3600)} hours ago.`;
	} else {
		// Date format.
		let date = new Date(timestamp * 1000);
		let dateObj = {
			year: date.getFullYear(),
			month: ("0" + (date.getMonth() + 1)).slice(-2),
			day: ("0" + date.getDate()).slice(-2),
			hours: ("0" + date.getHours()).slice(-2),
			minutes: ("0" + date.getMinutes()).slice(-2),
		};
		getPostTime = `${dateObj.year}-${dateObj.month}-${dateObj.day} ${dateObj.hours}:${dateObj.minutes}`;
	}

	return getPostTime;
};
