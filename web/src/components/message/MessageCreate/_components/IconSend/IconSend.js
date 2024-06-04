import Image from "next/image";

const IconSend = () => {
  return (
    <Image
      src="/images/message/send.svg"
      alt="send message"
      width={32}
      height={32}
      draggable="false"
    />
  );
};

export default IconSend;
