import IconLike from "./_components/IconLike/IconLike";

import { useState } from "react";

import { postLike, postUnlike } from "@/utils/posts";

import styles from "./PostLike.module.scss";

const PostLike = ({ id, likes, liked }) => {
	const [likeCount, setLikeCount] = useState(likes);
	const [likeStatus, setLikeStatus] = useState(liked);
	const [likeLoad, setLikeLoad] = useState(false);

	const handleLike = async () => {
		if (!likeLoad) {
			try {
				setLikeLoad(true);
				let response;
				if (!likeStatus) {
					response = await postLike(id);
				} else {
					response = await postUnlike(id);
				}
				if (response) {
					setLikeCount((prevLikes) =>
						likeStatus ? prevLikes - 1 : prevLikes + 1
					);
					setLikeStatus((prevStatus) => !prevStatus);
					return setLikeLoad(false);
				}
			} catch (error) {
				return console.error(`Failed to like/unlike post! [${error.message}]]`);
			}
		}
	};

	return (
		<>
			<button onClick={handleLike} className={styles.post_likecom}>
				<IconLike likeStatus={likeStatus} />
			</button>
			<span className={styles.post_like_count}>{likeCount}</span>
		</>
	);
};

export default PostLike;
