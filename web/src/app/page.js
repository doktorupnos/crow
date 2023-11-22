import AppTitle from "./_components/AppTitle/AppTitle";
import AuthButton from "./_components/AuthButton/AuthButton";

export default function Home() {
	return (
		<div className="container mx-auto">
			<div className="flex flex-col h-screen justify-center">
				<AppTitle />
				<AuthButton name="Register" link="/auth/register" />
				<AuthButton name="Login" link="/auth/login" />
			</div>
		</div>
	);
}
