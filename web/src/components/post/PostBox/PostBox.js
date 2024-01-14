import PostBoxTop from "./_components/PostBoxTop/PostBoxTop";
import PostBoxBottom from "./_components/PostBoxBottom/PostBoxBottom";

import { postTime } from "@/utils/posts";

import styles from "./PostBox.module.scss";

export default function PostBox({ post }) {
	let date = postTime(post.created_at);
	return (
		<div className={`${styles.post_block} ${styles.post_fade}`}>
			<PostBoxTop author={post.user_name} date={date} />
			<hr />
			<span className={styles.post_message}>{post.body}</span>
			<PostBoxBottom
				id={post.id}
				likes={post.likes}
				status={post.liked_by_user}
				self={post.self}
			/>
		</div>
	);
}
