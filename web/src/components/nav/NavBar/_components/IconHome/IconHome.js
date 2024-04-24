import Image from "next/image";

const handleLink = async () => {
  return (location.href = "/home");
};

const IconHome = () => {
  return (
    <button onClick={handleLink}>
      <Image
        src="/images/crow/logo.svg"
        alt="home"
        width={30}
        height={30}
        draggable="false"
      />
    </button>
  );
};

export default IconHome;
