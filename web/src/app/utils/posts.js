import axios from "axios";

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
