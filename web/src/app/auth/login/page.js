import AuthCrow from "../_components/AuthCrow/AuthCrow";
import AuthTitle from "../_components/AuthTitle/AuthTitle";
import AuthForm from "../_components/AuthForm/AuthForm";

export default function RegisterPage() {
	return (
		<div className="container mx-auto">
			<div className="flex flex-col h-screen justify-center">
				<AuthCrow />
				<AuthTitle title="Welcome Back." />
				<AuthForm method={1} />
			</div>
		</div>
	);
}
