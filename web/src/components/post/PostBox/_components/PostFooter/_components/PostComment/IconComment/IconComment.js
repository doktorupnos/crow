import Image from "next/image";

const IconComment = () => {
	return (
		<Image
			src="/images/post/comment.svg"
			alt="comments"
			width={20}
			height={20}
			draggable="false"
		/>
	);
};

export default IconComment;
