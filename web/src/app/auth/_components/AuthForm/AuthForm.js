"use client";

import { useState } from "react";
import axios from "axios";

import AuthField from "../AuthField/AuthField";
import AuthSubmit from "../AuthSubmit/AuthSubmit";

export default function AuthForm({ method }) {
	let [post, setPost] = useState({
		name: "",
		password: "",
	});

	function handleInput(event) {
		setPost({ ...post, [event.target.id]: event.target.value });
	}

	async function handleSubmit(event) {
		event.preventDefault();
		if (method) {
			// Login method.
			await axios
				.post(
					process.env.authLoginEndPoint,
					{},
					{
						auth: {
							username: `${post.name}`,
							password: `${post.password}`,
						},
						withCredentials: true,
					}
				)
				.then((response) => console.log(response))
				.catch((err) => console.log(err));
		} else {
			// Register method.
			await axios
				.post(process.env.authRegEndPoint, post)
				.then((response) => console.log(response))
				.catch((err) => console.log(err));
		}
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
