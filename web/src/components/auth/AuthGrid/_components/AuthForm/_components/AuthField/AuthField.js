import styles from "./AuthField.module.scss";

const AuthField = ({ type, handler }) => {
  return (
    <input
      id={type ? "name" : "password"}
      type={type ? "text" : "password"}
      maxLength={type ? 20 : 72}
      placeholder={type ? "Username" : "Password"}
      onChange={handler}
      className={styles.field_input}
      required
    />
  );
};

export default AuthField;
