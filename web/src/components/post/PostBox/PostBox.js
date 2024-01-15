import PostHeader from "./_components/PostHeader/PostHeader";
import PostFooter from "./_components/PostFooter/PostFooter";

import { postTime } from "@/utils/posts";

import styles from "./PostBox.module.scss";

const PostBox = ({ post }) => {
	let date = postTime(post.created_at);
	return (
		<article className={styles.post_box}>
			<PostHeader author={post.user_name} date={date} />
			<hr />
			<p className={styles.post_message}>{post.body}</p>
			<PostFooter
				id={post.id}
				likes={post.likes}
				liked={post.liked_by_user}
				self={post.self}
			/>
		</article>
	);
};

export default PostBox;
