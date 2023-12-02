import axios from "axios";

// Login form function.
export const userLogin = async (post) => {
	var userLogin = false;
	await axios
		.post(
			process.env.authLoginEndPoint,
			{},
			{
				auth: {
					username: `${post.name}`,
					password: `${post.password}`,
				},
				withCredentials: true,
			}
		)
		.then((response) => {
			if (response.status == 200) userLogin = true;
		})
		.catch((error) => console.error("Log-in error!", error));
	return userLogin;
};

// Register form function.
export const userRegister = async (post) => {
	var userRegister = false;
	await axios
		.post(process.env.authRegEndPoint, post)
		.then((response) => {
			if (response.status == 201) userRegister = true;
		})
		.catch((error) => console.error("Register error!", error));
	return userRegister;
};

// Check if session token is valid.
export const userValid = async () => {
	var userValid = false;
	await axios
		.post(process.env.authValidEndpoint, {}, { withCredentials: true })
		.then((response) => {
			if (response.status == 200) {
				userValid = true;
			} else {
				console.error("Invalid session!", error);
			}
		})
		.catch((error) => {
			console.error("Failed to validate session!", error);
		});
	return userValid;
};
