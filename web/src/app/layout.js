import "./global.css";

export const metadata = {
  title: "CROW",
  description: "null",
};

export default function RootLayout({ children }) {
  return (
    <>
      <title>CROW</title>
      <html lang="en">
        <body>{children}</body>
      </html>
    </>
  );
}
