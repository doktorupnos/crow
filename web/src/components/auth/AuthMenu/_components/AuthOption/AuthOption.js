import Link from "next/link";

import styles from "./AuthOption.module.scss";

const AuthOption = ({ name, link }) => {
  return (
    <Link className={styles.auth_option} href={link}>
      {name}
    </Link>
  );
};

export default AuthOption;
