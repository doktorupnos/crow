import Image from "next/image";

import styles from "../ErrorBlock.module.scss";

const ErrorUser = () => {
	return (
		<div className={styles.error_box}>
			<Image
				src="/images/error/error_user.jpg"
				width={400}
				height={400}
				alt="user not found"
				draggable="false"
				className={styles.error_image}
			/>
			<font className={styles.error_text}>user does not exist..</font>
		</div>
	);
};

export default ErrorUser;
