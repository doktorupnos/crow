import styles from "./PostBoxTop.module.scss";

const PostBoxTop = ({ author, date }) => {
	return (
		<div className={styles.post_top}>
			<span className={styles.post_user}>@{author}</span>
			<span className={styles.post_date}>{date}</span>
		</div>
	);
};

export default PostBoxTop;
