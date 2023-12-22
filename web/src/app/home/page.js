"use client";

import { useState, useEffect } from "react";
import { userValid } from "@/app/utils/auth";
import { redirect } from "next/navigation";

import NavBar from "@/components/nav/NavBar/NavBar";
import PostGrid from "@/components/post/PostGrid/PostGrid";
import PostCreate from "@/components/post/PostCreate/PostCreate";

const Home = () => {
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
				<PostCreate />
				<PostGrid />
			</>
		);
	} else {
		redirect("/auth");
	}
};

export default Home;
