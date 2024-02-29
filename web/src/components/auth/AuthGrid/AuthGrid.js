import AuthHeader from "./_components/AuthHeader/AuthHeader";
import AuthForm from "./_components/AuthForm/AuthForm";

import styles from "./AuthGrid.module.scss";

export default function AuthGrid({ method }) {
  return (
    <div className={styles.auth_grid}>
      <AuthHeader method={method} />
      <AuthForm method={method} />
    </div>
  );
}
