import Image from "next/image";

import styles from "./ProfileGrid.module.scss";

export default function ProfilePic() {
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
					<span className={styles.profile_name}>@stefanos</span>
					<div className={styles.profile_follow_grid}>
						<span className={styles.profile_follow_text}>0 followers</span>
						<span className={styles.profile_follow_text}>0 following</span>
					</div>
				</div>
			</header>
			<hr className={styles.profile_grid_line} />
		</>
	);
}
