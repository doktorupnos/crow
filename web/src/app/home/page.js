"use client";

import { redirect } from "next/navigation";
import { useState, useEffect } from "react";
import { fetchPosts } from "../utils/fetchPosts";

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
				<PostBlock />
			</div>
		);
	} else {
		console.error("Invalid session!\nRedirecting...");
		redirect("/auth");
	}
}

function PostBlock() {
	const [posts, setPosts] = useState([]);

	useEffect(() => {
		const fetchData = async () => {
			try {
				const response = await fetchPosts();
				console.log(response.payload);
				setPosts(response.payload);
			} catch (error) {
				console.error("Failed to fetch posts!");
			}
		};
		fetchData();
	}, []);

	// Check received posts.
	if (!posts.length) {
		return <h1>Posts unavailable!</h1>;
	} else {
		let postList = [];
		for (let post of posts) {
			postList.push(<li key={post.id}>{post.body}</li>);
		}
		return <ul className="flex flex-col">{postList}</ul>;
	}
}
