import Image from "next/image";

import styles from "./PostNone.module.scss";

export default function PostNone() {
	return (
		<div className={styles.post_none_box}>
			<Image
				className={styles.post_none_image}
				src="/images/error/no_posts.jpg"
				width={400}
				height={400}
				alt="error / no posts"
				draggable="false"
			/>
			<font className={styles.post_none_text}>no posts :(</font>
		</div>
	);
}
