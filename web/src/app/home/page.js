"use client";

import { redirect } from "next/navigation";
import { useState, useEffect } from "react";
import { fetchPosts } from "../utils/posts";

import PostBlock from "../_components/PostBlock/PostBlock";

export default function HomePage() {
	const [session, setSession] = useState(null);

	useEffect(() => {
		const fetchPostData = async () => {
			try {
				let response = await fetchPosts();
				setSession(response.auth);
			} catch (error) {
				console.error(error);
				setSession(false);
			}
		};
		fetchPostData();
	}, []);

	if (session === null) {
		return null;
	} else if (session) {
		return (
			<div>
				<div className="mb-5"></div>
				<PostBlock />
			</div>
		);
	} else {
		console.error("Invalid session!");
		redirect("/auth");
	}
}
