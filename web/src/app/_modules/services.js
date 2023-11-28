const axios = require("axios");
module.exports = {
	async jwtCheck() {
		await axios
			.post(
				process.env.jwtEndPoint,
				{},
				{
					void: {},
					withCredentials: true,
				}
			)
			.then((response) => {
				if (response.status == 200) {
					return true;
				} else {
					return false;
				}
			})
			.catch((error) => {
				console.log("Connection error!");
				return false;
			});
	},
	async getPosts() {
		await axios
			.get(process.env.postGetEndPoint, {}, { withCredentials: true })
			.then((response) => {
				if (response.status == 200) {
					return response.body;
				} else {
					return null;
				}
			})
			.catch((error) => {
				console.log("Connection error!");
				return null;
			});
	},
};
