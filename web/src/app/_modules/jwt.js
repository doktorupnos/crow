const axios = require("axios");
module.exports = {
	async jwtValidMain(router) {
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
				if (response.status != 200) router.push("/auth");
			})
			.catch(router.push("/auth"));
	},
	async jwtValidHome(router) {
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
					router.push("/home");
				} else {
					router.push("/auth");
				}
			})
			.catch(router.push("/auth"));
	},
};
