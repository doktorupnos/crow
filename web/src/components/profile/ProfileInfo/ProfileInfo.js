import { useState } from "react";

import styles from "./ProfileInfo.module.scss";

import ProfileList from "@/components/profile/ProfileList/ProfileList";

export default function ProfileInfo({ name, followers, following }) {
	const [listShow, setListShow] = useState(false);

	const showFollowers = async () => {
		setListShow(!listShow);
	};

	return (
		<div className={styles.profile_grid}>
			<span className={styles.profile_name}>@{name}</span>
			<div className={styles.profile_follow}>
				<span>
					<button onClick={showFollowers}>{followers} followers</button>
				</span>
				<span>
					<button>{following} following</button>
				</span>
			</div>
			{listShow ? <ProfileList close={showFollowers} /> : null}
		</div>
	);
}
