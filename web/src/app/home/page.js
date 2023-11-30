"use client";

import { redirect } from "next/navigation";
import { useState, useEffect } from "react";
import { fetchPosts } from "../utils/fetchPosts.js";

export default function HomePage() {
	const [session, setSession] = useState(null);

	useEffect(() => {
		fetchPosts()
			.then((response) => {
				setSession(response.auth);
			})
			.catch((error) => {
				console.error("Session validation failed!");
				setSession(false);
			});
	}, []);

	if (session === null) {
		return null;
	} else if (session) {
		return (
			<div>
				<h1>HOME</h1>
				<PostBlock />
			</div>
		);
	} else {
		console.error("Invalid session!\nRedirecting...");
		redirect("/auth");
	}
}

function PostBlock() {
	const [posts, setPosts] = useState(null);
	fetchPosts()
		.then((response) => {
			setPosts(response.payload);
		})
		.catch((error) => {
			console.error("Session validation failed!");
		});

	let postList = [];
	for (let post of posts) {
		postList.push(<li key={post.id}>{post.body}</li>);
	}

	return (
		<ul className="flex flex-col" id={id}>
			{postList}
		</ul>
	);
}
