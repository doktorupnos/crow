import Image from "next/image";

import styles from "./ProfileGrid.module.scss";

export default function ProfileGrid({ userData }) {
	return (
		<>
			<header className={styles.profile_grid}>
				<Image
					src="images/crow_circle.svg"
					alt="avatar"
					height={120}
					width={120}
					className={styles.profile_avatar}
					draggable="false"
				/>
				<div className={styles.profile_info}>
					<span className={styles.profile_name}>@{userData.name}</span>
					<div className={styles.profile_follow_grid}>
						<span className={styles.profile_follow_text}>
							{userData.follower_count} followers
						</span>
						<span className={styles.profile_follow_text}>
							{userData.following_count} following
						</span>
					</div>
				</div>
			</header>
			<hr className={styles.profile_grid_line} />
		</>
	);
}
