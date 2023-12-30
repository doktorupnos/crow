import { useState, useEffect } from "react";

import IconClose from "./_components/IconClose/IconClose";
import IconLoad from "./_components/IconLoad/IconLoad";
import { fetchFollow } from "@/utils/profile";

import styles from "./ProfileList.module.scss";

export default function ProfileList({ name, close, type }) {
	const [followList, setFollowList] = useState([null]);
	const [moreList, setMoreList] = useState(null);
	const [page, setPage] = useState(0);

	useEffect(() => {
		const getList = async (name, page, type) => {
			try {
				let response = await fetchFollow(name, page, type);
				if (!response.length > 0) return setMoreList(false);
				let newList = response.map((user) => (
					<li key={user.id}>@{user.name}</li>
				));
				setFollowList((prevList) => [...prevList, ...newList]);
				setMoreList(true);
			} catch (error) {
				console.error(`Failed to retrieve user list! [${error.message}]`);
			}
		};
		getList(name, page, type);
	}, [name, page, type]);

	const loadMore = () => {
		setPage(page + 1);
	};

	return (
		<div className={styles.list_back}>
			<div className={styles.list_inner}>
				<header className={styles.list_header}>
					<span className={styles.list_title}>
						{type ? "Followers" : "Following"}
					</span>
					<button onClick={close} className={styles.list_close}>
						<IconClose />
					</button>
				</header>
				<hr />
				<ul className={styles.list_ul}>{followList}</ul>
				{moreList ? (
					<button onClick={loadMore} className={styles.list_load}>
						<IconLoad />
					</button>
				) : null}
			</div>
		</div>
	);
}
