import Link from "next/link";

import styles from "./AuthButton.module.css";

export default function AuthButton({ name, link }) {
	return (
		<div className="max-w-md mx-auto">
			<Link
				className={`
					${styles.authButton} 
					focus:outline-none 
					font-medium 
					text-md 
					items-center
				`}
				type="button"
				href={link}
			>
				{name}
			</Link>
		</div>
	);
}
