"use client";

import { redirect } from "next/navigation";
import { useState, useEffect } from "react";
import { userValid } from "@/app/utils/auth";

import NavBar from "@/components/nav/NavBar/NavBar";
import PostBlock from "../_components/PostBlock/PostBlock";

export default function HomePage() {
	const [session, setSession] = useState(null);

	useEffect(() => {
		const sessionValid = async () => {
			try {
				let response = await userValid();
				setSession(response);
			} catch (error) {
				console.error("Failed to validate session!", error);
				setSession(false);
			}
		};
		sessionValid();
	}, []);

	if (session === null) {
		return null;
	} else if (session) {
		return (
			<>
				<NavBar />
				<PostBlock />
			</>
		);
	} else {
		redirect("/auth");
	}
}
