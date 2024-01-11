import { useState } from "react";

import styles from "./ProfileInfo.module.scss";

import ProfileList from "@/components/profile/ProfileList/ProfileList";

export default function ProfileInfo({ name, followers, following }) {
	const [followersList, setFollowersList] = useState(false);
	const [followingList, setFollowingList] = useState(false);

	const showFollowers = async () => {
		setFollowersList(!followersList);
	};

	const showFollowing = async () => {
		setFollowingList(!followingList);
	};

	return (
		<div className={styles.profile_grid}>
			<span className={styles.profile_name}>@{name}</span>
			<div className={styles.profile_follow}>
				<span>
					<button onClick={showFollowers}>{followers} followers</button>
				</span>
				<span>
					<button onClick={showFollowing}>{following} following</button>
				</span>
			</div>
			{followersList ? (
				<ProfileList name={name} close={showFollowers} type={1} />
			) : null}
			{followingList ? (
				<ProfileList name={name} close={showFollowing} type={0} />
			) : null}
		</div>
	);
}
