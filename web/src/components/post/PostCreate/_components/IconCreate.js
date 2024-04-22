import Image from "next/image";

const IconCreate = () => {
  return (
    <Image
      src="/images/post/create.svg"
      alt="create post"
      width={25}
      height={25}
      draggable="false"
    />
  );
};

export default IconCreate;
