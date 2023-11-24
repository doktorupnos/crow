import Link from "next/link";

import styles from "./AuthButton.module.css";

export default function AuthButton({ name, link }) {
	return (
		<div className="max-w-md mx-auto">
			<Link
				type="button"
				className={`${styles.authButton} hover:bg-[#D22B2B] focus:outline-none font-medium text-md items-center`}
				href={link}
			>
				{name}
			</Link>
		</div>
	);
}
