import PostLike from "./_components/PostLike/PostLike";
import PostComment from "./_components/PostComment/PostComment";
import PostDelete from "./_components/PostDelete/PostDelete";

import styles from "./PostBoxBottom.module.scss";

const PostBoxBottom = ({ id, likes, liked, self }) => {
	return (
		<div className={styles.post_footer}>
			<div className={styles.post_footer_field}>
				<PostLike id={id} likes={likes} liked={liked} />
				<PostComment />
			</div>
			{self && (
				<div className={styles.post_footer_field}>
					<PostDelete id={id} />
				</div>
			)}
		</div>
	);
};

export default PostBoxBottom;
