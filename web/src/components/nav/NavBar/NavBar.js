import NavHome from "@/components/nav/NavHome/NavHome";

import IconGithub from "./_components/IconGithub/IconGithub";
import IconProfile from "./_components/IconProfile/IconProfile";
import IconLogout from "./_components/IconLogout/IconLogout";

import styles from "./NavBar.module.scss";

const NavBar = () => {
  return (
    <nav className={styles.nav_bar}>
      <ul className={styles.nav_grid}>
        <IconGithub />
        <NavHome />
        <ul className={styles.nav_grid_profile}>
          <IconProfile />
          <IconLogout />
        </ul>
      </ul>
    </nav>
  );
};

export default NavBar;
