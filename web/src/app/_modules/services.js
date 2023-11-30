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
				console.error("Could not validate session!");
			});
		return valid;
	},
};
