import IconLoad from "./_components/IconLoad/IconLoad";
import PostBox from "@/components/post/PostBox/PostBox";
import ErrorPost from "@/components/error/ErrorPost/ErrorPost";

import { useState, useEffect } from "react";

import { fetchPosts, postTime } from "@/utils/posts";

import styles from "./PostGrid.module.scss";

const PostGrid = ({ user }) => {
	const [postList, setPostList] = useState([]);
	const [postLoad, setPostLoad] = useState(false);
	const [morePosts, setMorePosts] = useState(null);
	const [page, setPage] = useState(0);

	useEffect(() => {
		const getPosts = async (page) => {
			try {
				let response;
				if (user) {
					response = await fetchPosts(user, page);
				} else {
					response = await fetchPosts(null, page);
				}
				if (!response.length > 0) return setMorePosts(false);
				let newList = response.map((post) => {
					return <PostBox key={post.id} post={post} />;
				});
				setPostList((prevList) => [...prevList, newList]);
				setMorePosts(true);
				setPostLoad(false);
			} catch (error) {
				console.error(`Failed to retrieve posts! [${error.message}]`);
			}
		};
		getPosts(page);
	}, [page, user]);

	useEffect(() => {
		const handleScrollBottom = () => {
			const isScrollAtBottom =
				window.innerHeight + window.scrollY >= document.body.scrollHeight;
			if (isScrollAtBottom && !postLoad) {
				setPage((page) => page + 1);
				setPostLoad(true);
			}
		};
		window.addEventListener("scroll", handleScrollBottom);
		return () => {
			window.removeEventListener("scroll", handleScrollBottom);
		};
	}, [postLoad]);

	const handleLoad = () => {
		if (!postLoad) {
			setPage((page) => page + 1);
			setPostLoad(true);
		}
	};

	return (
		<>
			{postList.length > 0 && (
				<>
					<div className={styles.post_grid}>{postList}</div>
					{morePosts && (
						<button className={styles.post_load} onClick={handleLoad}>
							<IconLoad />
						</button>
					)}
				</>
			)}
			{morePosts == false && postList.length == 0 && <ErrorPost />}
		</>
	);
};

export default PostGrid;
