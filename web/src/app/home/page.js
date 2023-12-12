"use client";

import { redirect } from "next/navigation";
import { useState, useEffect, useRef } from "react";
import { userValid } from "@/app/utils/auth";

import NavBar from "@/components/nav/NavBar/NavBar";
import PostBlock from "../_components/PostBlock/PostBlock";

import styles from "./page.module.css";

export default function HomePage() {
	const [session, setSession] = useState(null);
	const [textValue, setTextValue] = useState("");

	const textareaRef = useRef(null);

	const handleChange = (event) => {
		const postText = event.target.value;
		if (postText.trim() !== "") {
			setTextValue(event.target.value);
		} else {
			setTextValue("");
		}
	};

	useEffect(() => {
		const autoResize = () => {
			if (textareaRef.current) {
				textareaRef.current.style.height = "auto";
				textareaRef.current.style.height =
					textareaRef.current.scrollHeight + "px";
			}
		};
		autoResize();
	}, [textValue]);

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
				<NavBar />
				<form className={styles.post_block}>
					<textarea
						ref={textareaRef}
						className={styles.post_text}
						type="text"
						value={textValue}
						onChange={handleChange}
						placeholder="Quoth the raven.."
						required
					/>
					<div>
						<button className={styles.post_button}>Post</button>
					</div>
				</form>
				<PostBlock />
			</>
		);
	} else {
		redirect("/auth");
	}
}
