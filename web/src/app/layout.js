import "./global.css";

export const metadata = {
	title: "CROW",
	description: "null",
};

export default function RootLayout({ children }) {
	return (
		<html lang="en">
			<body>{children}</body>
		</html>
	);
}
