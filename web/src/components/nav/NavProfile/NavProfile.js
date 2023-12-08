import Link from "next/link";
import Image from "next/image";

import styles from "./NavProfile.module.scss";

const NavProfile = () => {
	return (
		<li className={styles.nav_profile}>
			<Link href="/profile">
				<Image
					src="/images/bootstrap/person-circle.svg"
					alt="Profile"
					width={25}
					height={25}
					draggable="false"
				/>
			</Link>
		</li>
	);
};

export default NavProfile;
