"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";

import { userLogin, userRegister } from "../../../utils/auth";

import AuthField from "../AuthField/AuthField";
import AuthSubmit from "../AuthSubmit/AuthSubmit";

export default function AuthForm({ method }) {
	const router = useRouter();

	let [post, setPost] = useState({
		name: "",
		password: "",
	});

	function handleInput(event) {
		setPost({ ...post, [event.target.id]: event.target.value });
	}

	async function handleSubmit(event) {
		let auth;
		event.preventDefault();
		if (method) {
			try {
				auth = await userLogin(post);
			} catch (error) {
				console.error("Could not log-in user!");
			}
		} else {
			try {
				auth = await userRegister(post);
			} catch (error) {
				console.error("Could not register user!", error);
			}
		}
		if (auth) router.push("/home");
	}

	return (
		<div className="flex justify-center">
			<form onSubmit={handleSubmit}>
				<AuthField
					id="name"
					type="text"
					maxlen="20"
					placeholder="Username"
					onchange={handleInput}
				/>
				<AuthField
					id="password"
					type="password"
					maxlen="64"
					placeholder="Password"
					onchange={handleInput}
				/>
				<AuthSubmit value={method ? "Enter" : "Join"} />
			</form>
		</div>
	);
}
