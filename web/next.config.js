/** @type {import('next').NextConfig} */

module.exports = {
  // output: "standalone",
  env: {
    authValidEndPoint: "//api:8000/admin/jwt",
    authRegEndPoint: "//api:8000/users",
    authLoginEndPoint: "//api:8000/login",
    postGetEndPoint: "//api:8000/posts",
    postLikeEndPoint: "//api:8000/post_likes",
    profileEndPoint: "//api:8000/profile",
    followEndPoint: "//api:8000/follow",
    unfollowEndPoint: "//api:8000/unfollow",
    followersEndPoint: "//api:8000/followers",
    followingEndPoint: "//api:8000/following",
  },
  reactStrictMode: false,
};
