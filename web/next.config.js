/** @type {import('next').NextConfig} */
module.exports = {
	env: {
		authRegEndPoint: "//localhost:8000/users",
		authLoginEndPoint: "//localhost:8000/login",
		postGetEndPoint: "//localhost:8000/posts",
		postLikeEndPoint: "//localhost:8000/post_likes",
		authValidEndPoint: "//localhost:8000/admin/jwt",
	},
	reactStrictMode: false,
};
