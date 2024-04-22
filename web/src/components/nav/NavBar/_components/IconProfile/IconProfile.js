import Image from "next/image";

import styles from "./IconProfile.module.scss";

const handleLink = async () => {
  return (location.href = "/profile");
};

const IconProfile = () => {
  return (
    <button onClick={handleLink}>
      <Image
        src="/images/nav/profile.svg"
        alt="profile"
        width={25}
        height={25}
        draggable="false"
        className={styles.icon_profile}
      />
    </button>
  );
};

export default IconProfile;
