/** @type {import('next').NextConfig} */

module.exports = {
  // output: "standalone",
  env: {
    authValidEndPoint: "http//api:8000/admin/jwt",
    authRegEndPoint: "http//api:8000/users",
    authLoginEndPoint: "http//api:8000/login",
    postGetEndPoint: "http//api:8000/posts",
    postLikeEndPoint: "http//api:8000/post_likes",
    profileEndPoint: "http//api:8000/profile",
    followEndPoint: "http//api:8000/follow",
    unfollowEndPoint: "http//api:8000/unfollow",
    followersEndPoint: "http//api:8000/followers",
    followingEndPoint: "http//api:8000/following",
  },
  reactStrictMode: false,
};
