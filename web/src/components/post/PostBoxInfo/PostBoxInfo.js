import styles from "./PostBoxInfo.module.scss";

const PostInfo = ({ author, date }) => {
	return (
		<div className={styles.post_info}>
			<span className={styles.post_user}>@{author}</span>
			<span className={styles.post_date}>{date}</span>
		</div>
	);
};

export default PostInfo;
