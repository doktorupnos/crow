/** @type {import('next').NextConfig} */
module.exports = {
	env: {
		authRegEndPoint: "//localhost:8000/users",
		authLoginEndPoint: "//localhost:8000/login",
		postGetEndPoint: "//localhost:8000/posts",
		authValidEndPoint: "//localhost:8000/admin/jwt",
        followEndPoint: "//localhost:8000/follow",
        unfollowEndPoint: "//localhost:8000/unfollow",
        followingEndPoint: "//localhost:8000/following",
        followersEndPoint: "//localhost:8000/followers",
        followingCountEndpoint: "//localhost:8000/following_count",
        followersCountEndpoint: "//localhost:8000/followers_count",
	},
	reactStrictMode: false,
};
