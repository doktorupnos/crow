import IconCrow from "./_components/IconCrow/IconCrow";

import styles from "./AuthHeader.module.scss";

const AuthHeader = ({ method }) => {
  return (
    <header className={styles.header}>
      <IconCrow />
      <h1 className={styles.header_text}>
        {method ? "Welcome Back" : "Become A Crow"}
      </h1>
    </header>
  );
};

export default AuthHeader;
