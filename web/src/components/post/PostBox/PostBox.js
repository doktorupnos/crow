import { useState } from "react";

import Image from "next/image";

import PostBoxInfo from "@/components/post/PostBoxInfo/PostBoxInfo";

import { likePost } from "@/app/utils/posts";

import styles from "./PostBox.module.scss";

export default function PostBox({ id, author, message, date, likes, status }) {
	const [postLikes, setPostLikes] = useState(likes);
	const [likeStatus, setLikeStatus] = useState(status);

	const actionLike = async () => {
		try {
			await likePost(id, likeStatus);
		} catch (error) {
			console.error("Failed to (remove) post like!", error);
		}
		setPostLikes(likeStatus ? postLikes - 1 : postLikes + 1);
		setLikeStatus(!likeStatus);
	};

	return (
		<div id={id} className={`${styles.post_block} ${styles.post_fade}`}>
			<PostBoxInfo author={author} date={date} />
			<hr />
			<span className={styles.post_message}>{message}</span>
			<div className={styles.post_info}>
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
				<span className={styles.post_likes_text}>{postLikes}</span>
				<div className={styles.post_info}>
					<button className={styles.post_likes}>
						<Image
							src="/images/bootstrap/comment.svg"
							alt="comment"
							width={20}
							height={20}
							draggable="false"
						/>
						<span className={styles.post_likes_text}>0</span>
					</button>
				</div>
			</div>
		</div>
	);
}
