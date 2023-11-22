"use client";

import { useState } from "react";
import axios from "axios";

const apiendpoint = "//localhost:8000";

export default function RegisterPage() {
	return (
		<div className="container mx-auto">
			<div className="flex flex-col h-screen justify-center">
				<RegisterForm />
			</div>
		</div>
	);
}

function RegisterForm() {
	// JSON body values.
	let [post, setPost] = useState({
		name: "",
		password: "",
	});

	// Input field handlers.
	function handleInput(event) {
		setPost({ ...post, [event.target.id]: event.target.value });
	}

	// Submit button handler.
	function handleSubmit(event) {
		event.preventDefault();
		console.log(post);
		axios
			.post(apiendpoint, post)
			.then((response) => console.log(response))
			.catch((err) => console.log(err));
	}

	// Html elements.
	return (
		<div>
			<form onSubmit={handleSubmit}>
				<LogRegField
					type="text"
					id="name"
					maxlen="20"
					placeholder="Username"
					onchange={handleInput}
				/>
				<br />
				<LogRegField
					type="password"
					id="password"
					maxlen="64"
					placeholder="Password"
					onchange={handleInput}
				/>
				<input type="submit" value="Join" />
			</form>
		</div>
	);
}

function LogRegField({ type, id, maxlen, placeholder, onchange }) {
	return (
		<div className="max-w-sm mx-auto">
			<label
				htmlFor={id}
				className="bg-transparent text-white relative block overflow-hidden rounded-md border border-gray-200 px-3 pt-3 shadow-sm focus-within:border-blue-600 focus-within:ring-1 focus-within:ring-blue-600"
			>
				<input
					type={type}
					id={id}
					maxLength={maxlen}
					placeholder={placeholder}
					className="bg-transparent peer h-8 w-full border-none p-0 placeholder-transparent focus:border-transparent focus:outline-none focus:ring-0 sm:text-sm"
					onChange={onchange}
					required
				/>
				<span className="text-white absolute start-3 top-3 -translate-y-1/2 text-xs text-gray-700 transition-all peer-placeholder-shown:top-1/2 peer-placeholder-shown:text-sm peer-focus:top-3 peer-focus:text-xs">
					{placeholder}
				</span>
			</label>
		</div>
	);
}
