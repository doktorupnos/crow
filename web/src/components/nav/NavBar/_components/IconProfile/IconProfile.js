import Image from "next/image";
import Link from "next/link";

import styles from "./IconProfile.module.scss";

const IconProfile = () => {
  return (
    <Link href="/profile">
      <Image
        src="/images/nav/profile.svg"
        alt="profile"
        width={25}
        height={25}
        draggable="false"
        className={styles.icon_profile}
      />
    </Link>
  );
};

export default IconProfile;
