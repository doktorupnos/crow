/** @type {import('next').NextConfig} */
module.exports = {
	env: {
		authRegEndPoint: "//0.0.0.0:8000/users",
		authLoginEndPoint: "//0.0.0.0:8000/login",
		postGetEndPoint: "//0.0.0.0:8000/posts",
		jwtEndPoint: "//0.0.0.0:8000/admin/jwt",
	},
	reactStrictMode: false,
};
