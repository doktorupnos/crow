import styles from "./PostBox.module.scss";

export default function PostBox({ id, author, message, date }) {
	return (
		<ul id={id} className={`${styles.post_block} ${styles.post_fade}`}>
			<li className={styles.post_user}>@{author}</li>
			<hr />
			<li className={styles.post_message}>{message}</li>
			<li className={styles.post_date}>{date}</li>
		</ul>
	);
}
