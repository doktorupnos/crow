import { useState, useEffect } from "react";
import { fetchPosts, getPostTime } from "@/app/utils/posts";

import PostBox from "@/components/post/PostBox/PostBox";

export default function PostBlock() {
	const [posts, setPosts] = useState([]);
	const [postList, setPostList] = useState([]);
	const [page, setPage] = useState([1]);

	useEffect(() => {
		const fetchData = async () => {
			try {
				const response = await fetchPosts(page);
				setPosts((prevPosts) => [...prevPosts, ...response.payload]);
			} catch (error) {
				console.error("Failed to fetch posts!", error);
			}
		};
		fetchData();
	}, [page]);

	useEffect(() => {
		const handleScrollBottom = () => {
			const isScrollAtBottom =
				window.innerHeight + window.scrollY >= document.body.scrollHeight;
			if (isScrollAtBottom) {
				console.log("Bottom");
				setPage((prevPage) => prevPage + 1);
			} else {
				console.log("Not Bottom");
			}
		};
		window.addEventListener("scroll", handleScrollBottom);
		return () => {
			window.removeEventListener("scroll", handleScrollBottom);
		};
	}, []);

	const loadMorePosts = () => {
		setPage((page) => page + 1);
	};

	useEffect(() => {
		const postData = posts.map((post) => {
			const date = getPostTime(post.created_at);
			return (
				<PostBox
					key={post.id}
					id={post.id}
					author={post.user_name}
					message={post.body}
					date={date}
				/>
			);
		});
		setPostList(postData);
	}, [posts]);

	if (!posts.length) {
		return <h1>Posts unavailable!</h1>;
	}
	return (
		<>
			<div className="flex flex-col mx-auto">{postList}</div>
			<button onClick={loadMorePosts}>Load More</button>
		</>
	);
}
