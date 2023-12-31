import { useState, useEffect } from "react";
import { fetchPosts, getPostTime } from "@/app/utils/posts";

import Image from "next/image";

import PostBox from "@/components/post/PostBox/PostBox";
import PostNone from "@/components/post/PostNone/PostNone";

import styles from "./PostGrid.module.css";

export default function PostGrid() {
	const [posts, setPosts] = useState([]);
	const [postList, setPostList] = useState([]);
	const [page, setPage] = useState(1);
	const [morePosts, setMorePosts] = useState(true);

	useEffect(() => {
		const fetchData = async () => {
			if (morePosts) {
				try {
					const response = await fetchPosts(page);
					if (!response.payload.length) return setMorePosts(false);
					setPosts((prevPosts) => [...prevPosts, ...response.payload]);
				} catch (error) {
					console.error("Failed to fetch posts!", error);
					return setMorePosts(false);
				}
			}
		};
		fetchData();
	}, [page, morePosts]);

	useEffect(() => {
		const handleScrollBottom = () => {
			const isScrollAtBottom =
				window.innerHeight + window.scrollY >= document.body.scrollHeight;
			if (isScrollAtBottom) setPage((page) => page + 1);
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
					likes={post.likes}
					status={post.liked_by_user}
				/>
			);
		});
		setPostList(postData);
	}, [posts]);

	if (!posts.length) return <PostNone />;

	return (
		<>
			<div className="flex flex-col mx-auto">{postList}</div>
			<button className={styles.button_load} onClick={loadMorePosts}>
				<Image
					src="/images/bootstrap/load_posts.svg"
					alt="load_more"
					width={30}
					height={30}
					draggable="false"
				/>
			</button>
		</>
	);
}
