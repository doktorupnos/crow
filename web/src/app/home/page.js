"use client";

import axios from "axios";

export default function HomePage() {
	axios
		.post(
			process.env.jwtEndPoint,
			{},
			{
				void: {},
				withCredentials: true,
			}
		)
		.then((response) => {
			if (response.status == 200) {
				console.log("success");
			} else {
				console.log("makaronia");
			}
		})
		.catch((error) => console.log(error));
	return (
		<div>
			<h1>HOME</h1>
			<font>makaronia</font>
		</div>
	);
}
