/** @type {import('next').NextConfig} */
module.exports = {
	env: {
		authValidEndPoint: "//localhost:8000/admin/jwt",
		authRegEndPoint: "//localhost:8000/users",
		authLoginEndPoint: "//localhost:8000/login",
		postGetEndPoint: "//localhost:8000/posts",
		postLikeEndPoint: "//localhost:8000/post_likes",
		profileEndPoint: "//localhost:8000/profile",
		followEndPoint: "//localhost:8000/follow",
		unfollowEndPoint: "//localhost:8000/unfollow",
		followingEndPoint: "//localhost:8000/following",
		followersEndPoint: "//localhost:8000/followers",
	},
	reactStrictMode: false,
};
