/** @type {import('next').NextConfig} */

module.exports = {
  // output: "standalone",
  env: {
    authValidEndPoint: "//api-gateway-service:8000/admin/jwt",
    authRegEndPoint: "//api-gateway-service:8000/users",
    authLoginEndPoint: "//api-gateway-service:8000/login",
    postGetEndPoint: "//api-gateway-service:8000/posts",
    postLikeEndPoint: "//api-gateway-service:8000/post_likes",
    profileEndPoint: "//api-gateway-service:8000/profile",
    followEndPoint: "//api-gateway-service:8000/follow",
    unfollowEndPoint: "//api-gateway-service:8000/unfollow",
    followersEndPoint: "//api-gateway-service:8000/followers",
    followingEndPoint: "//api-gateway-service:8000/following",
  },
  reactStrictMode: false,
};
