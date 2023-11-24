import "./global.css";

import { EB_Garamond } from "next/font/google";

export const garamond = EB_Garamond({
	subsets: ["latin"],
	variable: "--font-garamond",
});

export const metadata = {
	title: "CROW",
	description: "null",
};

export default function RootLayout({ children }) {
	return (
		<html lang="en">
			<body className={garamond.variable}>{children}</body>
		</html>
	);
}
