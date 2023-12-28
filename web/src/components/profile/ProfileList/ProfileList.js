import Image from "next/image";

import { useState } from "react";

import styles from "./ProfileList.module.scss";

export default function ProfileList({ close }) {
	const [followList, setFollowList] = useState([]);

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
					{followList.length > 0 ? null : null}
				</ul>
			</div>
		</div>
	);
}
