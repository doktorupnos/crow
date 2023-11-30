import axios from "axios";

export const fetchPosts = async () => {
	var fetchPosts = {
		auth: false,
		payload: [],
	};
	await axios
		.get(process.env.postGetEndPoint, { withCredentials: true })
		.then((response) => {
			if (response.status == 200) {
				fetchPosts.auth = true;
				fetchPosts.payload = response.data;
			} else if (response.status == 401) {
				console.error("Invalid session!");
			}
		})
		.catch((error) => {
			console.error("Could not fetch posts!");
		});
	return fetchPosts;
};
