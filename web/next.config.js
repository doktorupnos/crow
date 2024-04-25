/** @type {import('next').NextConfig} */

module.exports = {
  // output: "standalone",
  env: {
    authValidEndPoint: "//crow.zapto.org/api/admin/jwt",
    authRegEndPoint: "//crow.zapto.org/api/users",
    authLoginEndPoint: "//crow.zapto.org/api/login",
    authLogoutEndPoint: "//crow.zapto.org/api/logout",
    postGetEndPoint: "//crow.zapto.org/api/posts",
    postLikeEndPoint: "//crow.zapto.org/api/post_likes",
    profileEndPoint: "//crow.zapto.org/api/profile",
    followEndPoint: "//crow.zapto.org/api/follow",
    unfollowEndPoint: "//crow.zapto.org/api/unfollow",
    followersEndPoint: "//crow.zapto.org/api/followers",
    followingEndPoint: "//crow.zapto.org/api/following",
  },
  reactStrictMode: false,
};
