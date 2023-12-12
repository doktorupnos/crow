"use client";

import { redirect } from "next/navigation";
import { useState, useEffect, useRef } from "react";
import { useRouter } from "next/navigation";
import { userValid } from "@/app/utils/auth";

import NavBar from "@/components/nav/NavBar/NavBar";
import PostBlock from "../_components/PostBlock/PostBlock";

import Image from "next/image";

import styles from "./page.module.css";

import { createPost } from "@/app/utils/posts";

export default function HomePage() {
	const [session, setSession] = useState(null);
	const [textValue, setTextValue] = useState("");
	const router = useRouter();

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

	const handleSubmit = async (event) => {
		event.preventDefault();
		try {
			let response = await createPost(textValue);
			if (response) window.location.reload(false);
		} catch (error) {
			return console.error("Failed to create post!", error);
		}
	};

	if (session === null) {
		return null;
	} else if (session) {
		return (
			<>
				<NavBar />
				<form onSubmit={handleSubmit} className={styles.post_block}>
					<textarea
						ref={textareaRef}
						className={styles.post_text}
						type="text"
						value={textValue}
						onChange={handleChange}
						placeholder="Quoth the raven.."
						required
					/>
					<button className={styles.post_button} type="submit">
						<Image
							className={styles.post_button_image}
							src="/images/bootstrap/post_create.svg"
							alt="create post"
							width={25}
							height={25}
							draggable="false"
						/>
					</button>
				</form>
				<PostBlock />
			</>
		);
	} else {
		redirect("/auth");
	}
}
