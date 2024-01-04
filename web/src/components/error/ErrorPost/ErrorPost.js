import Image from "next/image";

import styles from "../ErrorBlock.module.scss";

const ErrorPost = () => {
	return (
		<div className={styles.error_box}>
			<Image
				src="/images/error/error_post.jpg"
				width={400}
				height={400}
				alt="no posts"
				draggable="false"
				className={styles.error_image}
			/>
			<font className={styles.error_text}>no posts :(</font>
		</div>
	);
};

export default ErrorPost;
