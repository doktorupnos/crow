import Image from "next/image";

const IconCrow = () => {
  return (
    <Image
      src="/images/crow/logo.svg"
      alt="crow_logo"
      width={64}
      height={64}
      draggable="false"
    />
  );
};

export default IconCrow;
