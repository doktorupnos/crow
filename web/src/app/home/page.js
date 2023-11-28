"use client";

import { jwtTest } from "../_modules/jwt.js";
import { useRouter } from "next/navigation";

export default function HomePage() {
	const router = useRouter();
	jwtTest(router).catch((error) => console.log(error));
	return (
		<div>
			<h1>HOME</h1>
			<font>makaronia</font>
		</div>
	);
}
