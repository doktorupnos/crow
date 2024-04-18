import Image from "next/image";
import Link from "next/link";

import styles from "./IconHome.module.scss";

const IconHome = () => {
  return (
    <Link href="/home" className={styles.icon_home}>
      <Image
        src="/images/nav/home.svg"
        alt="home"
        width={30}
        height={30}
        draggable="false"
      />
    </Link>
  );
};

export default IconHome;
