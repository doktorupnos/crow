import { useState, useEffect } from "react";
import { getPostTime } from "@/app/utils/posts";
import { fetchPosts } from "@/utils/posts";

import PostBox from "@/components/post/PostBox/PostBox";
import PostNone from "@/components/post/PostNone/PostNone";

import IconLoad from "./_components/IconLoad/IconLoad";

import styles from "./PostGrid.module.css";

export default function PostGrid() {
	const [postList, setPostList] = useState([null]);
	const [morePosts, setMorePosts] = useState(null);
	const [page, setPage] = useState(0);

	useEffect(() => {
		const getPosts = async (page) => {
			try {
				let response = await fetchPosts(null, page);
				if (!response.length > 0) return setMorePosts(false);
				let newList = response.map((post) => {
					let date = getPostTime(post.created_at);
					return {
						id: post.id,
						timestamp: post.created_at,
						content: (
							<PostBox
								key={post.id}
								id={post.id}
								author={post.user_name}
								message={post.body}
								date={date}
								likes={post.likes}
								status={post.liked_by_user}
							/>
						),
					};
				});
				setPostList((prevList) => {
					let sorted = [...prevList];
					newList.forEach((newItem) => {
						let low = 0;
						let high = sorted.length - 1;
						while (low <= high) {
							let mid = Math.floor((low + high) / 2);
							if (sorted[mid] && sorted[mid].timestamp === newItem.timestamp) {
								sorted.splice(mid + 1, 0, newItem);
								break;
							} else if (
								sorted[mid] &&
								sorted[mid].timestamp < newItem.timestamp
							) {
								high = mid - 1;
							} else {
								low = mid + 1;
							}
							if (low > high) {
								sorted.splice(low, 0, newItem);
							}
						}
					});
					return sorted;
				});
				setMorePosts(true);
			} catch (error) {
				console.error(`Failed to retrieve posts! [${error.message}]`);
			}
		};
		getPosts(page);
	}, [page]);

	useEffect(() => {
		const handleScrollBottom = () => {
			const isScrollAtBottom =
				window.innerHeight + window.scrollY >= document.body.scrollHeight;
			if (isScrollAtBottom && morePosts) setPage((page) => page + 1);
		};
		window.addEventListener("scroll", handleScrollBottom);
		return () => {
			window.removeEventListener("scroll", handleScrollBottom);
		};
	}, [morePosts]);

	const loadMore = () => {
		setPage((page) => page + 1);
	};

	return (
		<>
			<div className={styles.post_grid}>
				{postList.map(
					(post) => post && <div key={post.id}>{post.content}</div>
				)}
			</div>
			{morePosts && (
				<button className={styles.post_load} onClick={loadMore}>
					<IconLoad />
				</button>
			)}
		</>
	);
}
