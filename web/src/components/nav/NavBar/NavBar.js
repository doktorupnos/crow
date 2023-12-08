"use client";

import NavGit from "@/components/nav/NavGit/NavGit";
import NavHome from "@/components/nav/NavHome/NavHome";
import NavProfile from "@/components/nav/NavProfile/NavProfile";

import styles from "./NavBar.module.scss";

const NavBar = () => {
	return (
		<nav className={styles.nav_bar}>
			<ul className={styles.nav_grid}>
				<NavGit />
				<NavHome />
				<NavProfile />
			</ul>
		</nav>
	);
};

export default NavBar;
