import styles from "./AuthTitle.module.css";

export default function AuthTitle({ title }) {
	return <font className={`${styles.authTitle} mx-auto`}>{title}</font>;
}
