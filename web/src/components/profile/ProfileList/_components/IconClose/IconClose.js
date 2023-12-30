import Image from "next/image";

const IconClose = () => {
	return (
		<Image
			src="/images/profile/close.svg"
			alt="close"
			height={20}
			width={20}
			draggable="false"
		/>
	);
};

export default IconClose;
