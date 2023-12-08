import axios from "axios";

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
				console.error("Invalid session!");
			}
		})
		.catch((error) => {
			console.error("Failed to fetch posts!", error);
		});
	return fetchPosts;
};

export const likePost = async (id, likeStatus) => {
	var likePosts = false;

	if (!likeStatus) {
		await axios
			.post(
				process.env.postLikeEndPoint,
				{ post_id: id },
				{ withCredentials: true }
			)
			.then((response) => {
				if (response.status == 201) likePosts = true;
			})
			.catch((error) => {
				console.error("Failed to like post!", error);
			});
	} else {
		await axios
			.delete(process.env.postLikeEndPoint, {
				data: { post_id: id },
				withCredentials: true,
			})
			.then((response) => {
				console.log(response);
			})
			.catch((error) => {
				console.error(error);
			});
	}

	return likePosts;
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
			seconds: ("0" + date.getSeconds()).slice(-2),
		};
		getPostTime = `${dateObj.year}-${dateObj.month}-${dateObj.day} ${dateObj.hours}:${dateObj.minutes}:${dateObj.seconds}`;
	}

	return getPostTime;
};
