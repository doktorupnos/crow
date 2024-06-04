/** @type {import('next').NextConfig} */

module.exports = {
  // output: "standalone",
  env: {
    authValidEndPoint: "//localhost:8000/api/admin/jwt",
    authRegEndPoint: "//localhost:8000/api/users",
    authLoginEndPoint: "//localhost:8000/api/login",
    authLogoutEndPoint: "//localhost:8000/api/logout",
    postGetEndPoint: "//localhost:8000/api/posts",
    postLikeEndPoint: "//localhost:8000/api/post_likes",
    profileEndPoint: "//localhost:8000/api/profile",
    followEndPoint: "//localhost:8000/api/follow",
    unfollowEndPoint: "//localhost:8000/api/unfollow",
    followersEndPoint: "//localhost:8000/api/followers",
    followingEndPoint: "//localhost:8000/api/following",
    messageEchoEndPoint: "//localhost:8000/api/ws/echo",
  },
  reactStrictMode: false,
};
