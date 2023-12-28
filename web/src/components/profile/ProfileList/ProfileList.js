import Image from "next/image";
import { useState, useEffect } from "react";

import { fetchFollowers } from "@/utils/profile";

import styles from "./ProfileList.module.scss";

export default function ProfileList({ name, close }) {
	const [followList, setFollowList] = useState([]);
	const [moreList, setMoreList] = useState(true);
	const [page, setPage] = useState(0);

	useEffect(() => {
		const getFollowers = async (name) => {
			try {
				let response = await fetchFollowers(name);
				if (!response.length > 0) return setMoreList(false);
				setFollowList((prevList) => [...prevList, ...response]);
			} catch (error) {
				console.error(`Failed to retrieve user list! [${error.message}]`);
			}
		};
		getFollowers(name);
	}, [page, name]);

	return (
		<div className={styles.follow_list_back}>
			<div className={styles.follow_list_inner}>
				<header className={styles.list_header}>
					<span className={styles.list_title}>Followers</span>
					<button onClick={close} className={styles.list_close}>
						<Image
							src="/images/bootstrap/button_close.svg"
							alt="close"
							height={20}
							width={20}
							draggable="false"
						/>
					</button>
				</header>
				<hr />
				<ul className={styles.follow_list_ul}>
					{followList.length > 0
						? followList.map((user) => <li key={user.id}>@{user.name}</li>)
						: null}
				</ul>
			</div>
		</div>
	);
}
