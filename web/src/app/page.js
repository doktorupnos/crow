"use client";

import { useRouter } from "next/navigation";
import { jwtValidHome } from "./_modules/jwt";

export default function Home() {
	const router = useRouter();
	jwtValidHome(router).catch((error) => console.log(error));
}
