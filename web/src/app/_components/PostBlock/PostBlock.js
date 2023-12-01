import { useState, useEffect } from "react";
import { fetchPosts, getPostTime } from "@/app/utils/posts";

import PostBox from "@/app/_components/PostBox/PostBox";

export default function PostBlock() {
	const [posts, setPosts] = useState([]);

	useEffect(() => {
		const fetchData = async () => {
			try {
				const response = await fetchPosts();
				setPosts(response.payload);
			} catch (error) {
				console.error("Failed to fetch posts!", error);
			}
		};
		fetchData();
	}, []);

	if (!posts.length) {
		return <h1>Posts unavailable!</h1>;
	} else {
		var postList = [];
		posts.map((post) => {
			let date = getPostTime(post.timestamp);
			postList.push(
				<PostBox
					id={post.id}
					author={post.user_name}
					message={post.body}
					date={date}
				/>
			);
		});
	}
	return <div className="flex flex-col mx-auto">{postList}</div>;
}
