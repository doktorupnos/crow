import Link from "next/link";
import Image from "next/image";

import styles from "./NavHome.module.scss";

const NavHome = () => {
	return (
		<li className={styles.nav_home}>
			<Link href="/home">
				<Image
					src="/images/crow.svg"
					className={styles.nav_home_img}
					alt="Home"
					width={30}
					height={30}
					draggable="false"
				/>
			</Link>
		</li>
	);
};

export default NavHome;
