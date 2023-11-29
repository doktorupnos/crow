"use client";

import { redirect } from "next/navigation";
import { useState, useEffect } from "react";
import { jwtCheck, getPosts } from "../_modules/services.js";

export default function HomePage() {
	const [session, setSession] = useState(null);

	useEffect(() => {
		jwtCheck()
			.then((response) => {
				setSession(response);
			})
			.catch((error) => {
				console.error("Token validation error!");
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

function PostBlock(id, username, date, body) {
	const [posts, setPosts] = useState(null);
	getPosts().then((response) => {
		setPosts(response);
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
