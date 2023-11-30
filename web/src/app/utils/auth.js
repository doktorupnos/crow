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
			if (response.status == 200) userRegister = true;
		})
		.catch((error) => console.error("Register error!", error));
	return userRegister;
};
