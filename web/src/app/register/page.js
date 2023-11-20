"use client"

import { useState } from "react"
import axios from "axios"

const apiendpoint = "//localhost:8000/users"

export default function RegisterBox() {
	return (
		<div>
			<h1>Become A Crow</h1>
			<RegisterForm />
		</div>
	)
}

function RegisterForm() {

	// JSON body values.
	let [post, setPost] = useState({
		name: '',
		password: ''
	})

	// Input field handlers.
	function handleInput(event) {
		setPost({...post, [event.target.name]: event.target.value})
	}

	// Submit button handler.
	function handleSubmit(event) {
		event.preventDefault()
		axios.post({apiendpoint}, post)
		.then(response => console.log(response))
		.catch(err => console.log(err))
	}

	// Html elements.
	return (
		<div>
			<form onSubmit={handleSubmit}>
				<input name="name" type="text" maxlength="20" onChange={handleInput} placeholder="Username" required /><br /><br />
				<input name="password" type="password" maxlength="64" onChange={handleInput} placeholder="Password" required /><br /><br />
				<input type="submit" value="Join" />
			</form>
		</div>
	)

}
