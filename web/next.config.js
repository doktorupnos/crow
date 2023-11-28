/** @type {import('next').NextConfig} */
const axios = require("axios");
module.exports = {
	env: {
		authRegEndPoint: "//localhost:8000/users",
		authLoginEndPoint: "//localhost:8000/login",
		jwtEndPoint: "//localhost:8000/admin/jwt",
	},
};
