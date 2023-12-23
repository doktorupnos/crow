import Image from "next/image";

import styles from "./PostBoxCom.module.scss";

const PostBoxCom = () => {
	return (
		<>
			<button className={styles.post_com}>
				<Image
					src="/images/bootstrap/comment.svg"
					alt="comment"
					width={20}
					height={20}
					draggable="false"
				/>
			</button>
			<span className={styles.post_com_num}>0</span>
		</>
	);
};

export default PostBoxCom;
