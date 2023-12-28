import ProfileAvatar from "@/components/profile/ProfileAvatar/ProfileAvatar";
import ProfileInfo from "@/components/profile/ProfileInfo/ProfileInfo";

import styles from "./ProfileGrid.module.scss";

export default function ProfileGrid({ userData }) {
	return (
		<>
			<header className={styles.profile_grid}>
				<ProfileAvatar
					uuid={userData.id}
					self={userData.self}
					following={userData.following}
				/>
				<ProfileInfo
					name={userData.name}
					followers={userData.follower_count}
					following={userData.following_count}
				/>
			</header>
			<hr className={styles.profile_grid_line} />
		</>
	);
}
