import Image from "next/image";

export default function AuthCrow() {
	return (
		<Image
			className="mx-auto"
			src="/images/crow.svg"
			height={64}
			width={64}
			alt="crow symbol"
			draggable="false"
		/>
	);
}
