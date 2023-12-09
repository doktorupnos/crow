import PostBoxInfo from "./_components/PostBoxTop/PostBoxTop";
import PostBoxBottom from "./_components/PostBoxBottom/PostBoxBottom";

import styles from "./PostBox.module.scss";

export default function PostBox({ id, author, message, date, likes, status }) {
	return (
		<div id={id} className={`${styles.post_block} ${styles.post_fade}`}>
			<PostBoxInfo author={author} date={date} />
			<hr />
			<span className={styles.post_message}>{message}</span>
			<PostBoxBottom id={id} likes={likes} status={status} />
		</div>
	);
}
