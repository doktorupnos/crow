import { useState, useEffect } from "react";

import Image from "next/image";

import { followUser } from "@/utils/profile";

import styles from "./ProfileAvatar.module.scss";

export default function ProfileAvatar({ uuid, self, following }) {
	const [followStatus, setFollowStatus] = useState(following);

	const handleFollow = async () => {
		try {
			let response = await followUser(uuid);
			if (response) {
				setFollowStatus(!followStatus);
			} else {
				console.error("Failed to follow user!");
			}
		} catch (error) {
			console.error(`Failed to follow user! [${error.message}]`);
		}
	};

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
			{!followStatus || !self ? (
				<button onClick={handleFollow} className={styles.profile_follow}>
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
