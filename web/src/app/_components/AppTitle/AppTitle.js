import Image from "next/image";

export default function AppTitle() {
	return (
		<Image
			className="mx-auto"
			src="/images/title.svg"
			width={260}
			height={130}
			alt="app title"
			draggable="false"
		/>
	);
}
