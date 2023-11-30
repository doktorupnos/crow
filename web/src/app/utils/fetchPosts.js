import axios from "axios";

const fetchPosts = async () => {
	var fetchPosts = {
		auth: false,
		payload: null,
	};
	await axios
		.get(process.env.postGetEndPoint, { withCredentials: true })
		.then((response) => {
			if (response.status == 200) {
				fetchPosts.auth = true;
				fetchPosts.payload = response.body;
			} else if (response.status == 401) {
				console.error("Invalid session!");
			}
		})
		.catch((error) => {
			console.log("Something went wrong!");
		});
	return fetchPosts;
};

export default fetchPosts;
