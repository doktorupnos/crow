import { useState, useEffect } from "react";

import Image from "next/image";

import styles from "./ProfileAvatar.module.scss";

export default function ProfileAvatar({ userid, following }) {
	const [followStatus, setFollowStatus] = useState(following);

	return (
		<div className={styles.profile_grid}>
			<Image
				src="images/crow_circle.svg"
				alt="avatar"
				height={120}
				width={120}
				draggable="false"
				className={styles.profile_avatar}
			/>
			{!followStatus ? (
				<button className={styles.profile_follow}>
					<Image
						src="images/bootstrap/user_follow.svg"
						alt="follow user"
						height={28}
						width={28}
						draggable="false"
					/>
				</button>
			) : null}
		</div>
	);
}
