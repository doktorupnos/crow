/** @type {import('next').NextConfig} */

module.exports = {
  // output: "standalone",
  env: {
    authValidEndPoint: "https://crow.zapto.org/api/admin/jwt",
    authRegEndPoint: "https://crow.zapto.org/api/users",
    authLoginEndPoint: "https://crow.zapto.org/api/login",
    authLogoutEndPoint: "https://crow.zapto.org/api/logout",
    postGetEndPoint: "https://crow.zapto.org/api/posts",
    postLikeEndPoint: "https://crow.zapto.org/api/post_likes",
    profileEndPoint: "https://crow.zapto.org/api/profile",
    followEndPoint: "https://crow.zapto.org/api/follow",
    unfollowEndPoint: "https://crow.zapto.org/api/unfollow",
    followersEndPoint: "https://crow.zapto.org/api/followers",
    followingEndPoint: "https://crow.zapto.org/api/following",
  },
  reactStrictMode: false,
};
