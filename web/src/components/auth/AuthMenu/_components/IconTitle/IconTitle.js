import Image from "next/image";

const IconTitle = () => {
  return (
    <Image
      src="/images/crow/title.svg"
      alt="app title"
      height={130}
      width={260}
      draggable="false"
    />
  );
};

export default IconTitle;
