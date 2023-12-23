import PostBoxLike from "./_components/PostBoxLike/PostBoxLike";
import PostBoxCom from "./_components/PostBoxCom/PostBoxCom";

import styles from "./PostBoxBottom.module.scss";

const PostBoxBottom = ({ id, likes, status }) => {
	return (
		<div className={styles.post_bottom}>
			<PostBoxLike id={id} likes={likes} status={status} />
			<PostBoxCom />
		</div>
	);
};

export default PostBoxBottom;
