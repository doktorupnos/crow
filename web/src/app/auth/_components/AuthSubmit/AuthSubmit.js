import styles from "./AuthSubmit.module.css";

export default function AuthSubmit({ value }) {
	return (
		<div className="flex justify-center">
			<input className={`${styles.authSubmit}`} type="submit" value={value} />
		</div>
	);
}
