import Image from "next/image";
import Link from "next/link";

import styles from "./page.module.css"

export default function Home() {
	return (
		<div>
			<WelcomeForm />
		</div>
	)
}

function WelcomeForm() {
	return (
		<div>
			<AppTitle />
			<br /><br />
			<Link className={styles.testLink} href="/register">REGISTER</Link><br/><br />
			<Link className={styles.testLink} href="/">LOGIN</Link>
		</div>
	)
}

function AppTitle() {
	return (
		<Image
			src="/images/title.svg"
			width={300}
			height={150}
			alt="app title"
			draggable="false"
		/>
	)
}
