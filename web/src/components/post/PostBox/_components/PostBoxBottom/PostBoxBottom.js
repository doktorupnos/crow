import PostBoxLike from "./_components/PostBoxLike/PostBoxLike";
import PostBoxCom from "./_components/PostBoxCom/PostBoxCom";

import PostDelete from "./_components/PostDelete/PostDelete";

import styles from "./PostBoxBottom.module.scss";

const PostBoxBottom = ({ id, likes, liked, self }) => {
	return (
		<div className={styles.post_footer}>
			<div className={styles.post_footer_field}>
				<PostBoxLike id={id} likes={likes} status={liked} />
				<PostBoxCom />
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
