import Image from "next/image";

const IconDelete = () => {
	return (
		<Image
			src="/images/post/delete.svg"
			alt="delete"
			width={20}
			height={20}
			draggable="false"
		/>
	);
};

export default IconDelete;
