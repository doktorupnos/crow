import styles from "./PostHeader.module.scss";

const PostHeader = ({ author, date }) => {
	return (
		<header className={styles.post_header}>
			<span className={styles.post_user}>@{author}</span>
			<span className={styles.post_date}>{date}</span>
		</header>
	);
};

export default PostHeader;
