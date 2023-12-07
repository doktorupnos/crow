"use client";

import { redirect } from "next/navigation";
import { useState, useEffect } from "react";
import { userValid } from "@/app/utils/auth";
import PostBlock from "../_components/PostBlock/PostBlock";
import Navbar from "@/components/Navbar/Navbar";

export default function HomePage() {
	const [session, setSession] = useState(null);

	useEffect(() => {
		const sessionValid = async () => {
			try {
				let response = await userValid();
				setSession(response);
			} catch (error) {
				console.error("Failed to validate session!", error);
				setSession(false);
			}
		};
		sessionValid();
	}, []);

	if (session === null) {
		return null;
	} else if (session) {
		return (
            <>
                <Navbar />
                <div>
                    <div className="mb-5"></div>
                    <PostBlock />
                </div>
            </>
		);
	} else {
		console.error("Invalid session!");
		redirect("/auth");
	}
}
