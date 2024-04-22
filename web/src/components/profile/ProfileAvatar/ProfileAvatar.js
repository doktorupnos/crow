import Image from "next/image";
import { useState, useEffect } from "react";

import { followUser, unfollowUser } from "@/utils/profile";

import styles from "./ProfileAvatar.module.scss";

const ProfileAvatar = ({ uuid, self, following }) => {
  const [followStatus, setFollowStatus] = useState(following);
  const [followIcon, setFollowIcon] = useState();

  useEffect(() => {
    {
      followStatus
        ? setFollowIcon("/images/profile/followed.svg")
        : setFollowIcon("/images/profile/follow.svg");
    }
  }, [followStatus]);

  const handleFollow = async () => {
    try {
      let response = false;
      if (!followStatus) {
        response = await followUser(uuid);
      } else {
        response = await unfollowUser(uuid);
      }
      if (response) {
        setFollowStatus(!followStatus);
      } else {
        console.error("Failed to change follow status!");
      }
    } catch (error) {
      console.error(`Failed to change follow status! [${error.message}]`);
    }
  };

  return (
    <div className={styles.profile_grid}>
      <Image
        src="images/crow_circle.svg"
        alt="avatar"
        height={120}
        width={120}
        draggable="false"
        className={styles.profile_avatar}
      />
      {!self ? (
        <button onClick={handleFollow} className={styles.profile_follow}>
          <Image
            src={followIcon}
            alt="follow status"
            height={26}
            width={26}
            draggable="false"
          />
        </button>
      ) : null}
    </div>
  );
};

export default ProfileAvatar;
