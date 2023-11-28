"use client";

import { redirect } from "next/navigation";
import { useState } from "react";
import { jwtCheck } from "./_modules/services.js";

export default function Home() {
	const [session, setSession] = useState(null);
	jwtCheck().then((response) => {
		setSession(response);
	});
	if (session) {
		redirect("/home");
	} else {
		redirect("/auth");
	}
}
