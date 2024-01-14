import Image from "next/image";

import { useState } from "react";

import { addLike, remLike } from "@/app/utils/posts";

import styles from "./PostBoxLike.module.scss";

export default function PostBoxLike({ id, likes, liked }) {
	const [postLikes, setPostLikes] = useState(likes);
	const [likeStatus, setLikeStatus] = useState(liked);

	const actionLike = async () => {
		if (!likeStatus) {
			try {
				let response = await addLike(id);
				if (response) {
					setPostLikes(postLikes + 1);
					setLikeStatus(!likeStatus);
				}
			} catch (error) {
				return console.error("Failed to like post!", error);
			}
		} else {
			try {
				let response = await remLike(id);
				if (response) {
					setPostLikes(postLikes - 1);
					setLikeStatus(!likeStatus);
				}
			} catch (error) {
				return console.error("Failed to remove like!", error);
			}
		}
	};

	return (
		<>
			<button onClick={() => actionLike()} className={styles.post_likes}>
				<Image
					src={
						likeStatus
							? "/images/bootstrap/like_true.svg"
							: "/images/bootstrap/like_false.svg"
					}
					alt="heart"
					width={20}
					height={20}
					draggable="false"
				/>
			</button>
			<span className={styles.post_likes_num}>{postLikes}</span>
		</>
	);
}
