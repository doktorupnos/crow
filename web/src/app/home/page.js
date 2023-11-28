"use client";

import { redirect } from "next/navigation";
import { useState, useEffect } from "react";
import { jwtCheck } from "../_modules/services.js";

export default function HomePage() {
	const [session, setSession] = useState(null);

	useEffect(() => {
		jwtCheck().then((response) => {
			setSession(response);
		});
	}, []);

	if (session) {
		return (
			<div>
				<h1>HOME</h1>
				<font>makaronia</font>
			</div>
		);
	} else {
		redirect("/auth");
	}
}
