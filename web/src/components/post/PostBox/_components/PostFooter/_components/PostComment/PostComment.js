import IconComment from "./IconComment/IconComment";

import styles from "./PostComment.module.scss";

const PostComment = () => {
	return (
		<>
			<button className={styles.post_com}>
				<IconComment />
			</button>
			<span className={styles.post_com_count}>0</span>
		</>
	);
};

export default PostComment;
