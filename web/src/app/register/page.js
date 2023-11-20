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
	return (
		<div>
			<form action={apiendpoint} method="post">
				<input id="username" type="text" maxlength="20" placeholder="Username" required /><br />
				<input id="password" type="password" maxlength="64" placeholder="Password" required /><br />
				<input type="submit" value="Join" />
			</form>
		</div>
	)
}
