import Image from "next/image";
import Link from "next/link";

const IconGithub = () => {
  return (
    <Link href="https://github.com/doktorupnos/crow/">
      <Image
        src="/images/nav/github.svg"
        alt="github"
        width={25}
        height={25}
        draggable="false"
      />
    </Link>
  );
};

export default IconGithub;
