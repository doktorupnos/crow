import Image from "next/image";

const handleLogout = async () => {
  document.cookie = "token=0";
  return (location.href = "/auth");
};

const IconLogout = () => {
  return (
    <button onClick={handleLogout}>
      <Image
        src="/images/nav/logout.svg"
        alt="logout"
        width={25}
        height={25}
        draggable="false"
      />
    </button>
  );
};

export default IconLogout;
