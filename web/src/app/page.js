import Image from "next/image";
import Link from "next/link";

import styles from "../public/globals.css";

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
			<h1>CROW</h1>
			<Link href="/register">REGISTER</Link><br/>
			<Link href="/">LOGIN</Link>
		</div>
	)
}
