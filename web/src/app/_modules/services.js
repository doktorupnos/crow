const axios = require("axios");
module.exports = {
	async jwtCheck() {
		let valid = false;
		await axios
			.post(process.env.jwtEndPoint, {}, { withCredentials: true })
			.then((response) => {
				if (response.status == 200) {
					valid = true;
				}
			})
			.catch((error) => {
				console.log("Connection error!");
			});
		return valid;
	},
	async getPosts() {
		let posts = null;
		await axios
			.get(process.env.postGetEndPoint, {}, { withCredentials: true })
			.then((response) => {
				if (response.status == 200) {
					posts = response.body;
				}
			})
			.catch((error) => {
				console.log("Connection error!");
			});
		return posts;
	},
};
