import IconGithub from "./_components/IconGithub/IconGithub";
import IconHome from "./_components/IconHome/IconHome";
import IconProfile from "./_components/IconProfile/IconProfile";
import IconLogout from "./_components/IconLogout/IconLogout";

import styles from "./NavBar.module.scss";

const NavBar = () => {
  return (
    <nav className={styles.nav_grid}>
      <IconGithub />
      <IconHome />
      <div className={styles.nav_grid_profile}>
        <IconProfile />
        <IconLogout />
      </div>
    </nav>
  );
};

export default NavBar;
