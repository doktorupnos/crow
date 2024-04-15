import Image from "next/image";

const IconLike = ({ likeStatus }) => {
  return (
    <Image
      src={
        likeStatus
          ? "/images/post/like_true.svg"
          : "/images/post/like_false.svg"
      }
      alt="like"
      width={20}
      height={20}
      draggable="false"
    />
  );
};

export default IconLike;
