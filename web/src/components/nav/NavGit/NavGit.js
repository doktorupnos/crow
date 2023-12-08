import Link from "next/link";
import Image from "next/image";

import styles from "./NavGit.module.scss";

const NavGit = () => {
	return (
		<li className={styles.nav_git}>
			<Link href="https://github.com/doktorupnos/crow/">
				<Image
					src="/images/bootstrap/github.svg"
					alt="Github"
					width={25}
					height={25}
					draggable="false"
				/>
			</Link>
		</li>
	);
};

export default NavGit;
