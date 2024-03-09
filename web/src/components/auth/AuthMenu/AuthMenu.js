import IconTitle from "./_components/IconTitle/IconTitle";
import AuthOption from "./_components/AuthOption/AuthOption";

import styles from "./AuthMenu.module.scss";

const AuthMenu = () => {
  return (
    <div className={styles.auth_menu}>
      <IconTitle />
      <AuthOption name="Sign In" link="/auth/login" />
      <AuthOption name="Register" link="/auth/register" />
    </div>
  );
};

export default AuthMenu;
