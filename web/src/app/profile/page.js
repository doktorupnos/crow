"use client";

import NavBar from "@/components/nav/NavBar/NavBar";
import ProfileGrid from "@/components/profile/ProfileGrid/ProfileGrid";
import PostGrid from "@/components/post/PostGrid/PostGrid";
import ErrorUser from "@/components/error/ErrorUser/ErrorUser";

import { useState, useEffect } from "react";

import { validSession } from "@/utils/auth";
import { fetchProfile } from "@/utils/profile";

const Profile = () => {
  const [userData, setUserData] = useState({});
  const [userDataLoad, setUserDataLoad] = useState(null);

  const [session, setSession] = useState(null);

  useEffect(() => {
    const checkSession = async () => {
      try {
        let response = await validSession();
        if (!response) {
          console.error(`Invalid session!`);
          return (location.href = "/auth");
        }
        setSession(true);
      } catch (error) {
        console.error(`Invalid session! [${error.message}]`);
        return (location.href = "/auth");
      }
    };
    if (typeof window !== "undefined") {
      checkSession();
    }
  }, []);

  useEffect(() => {
    const query = new URLSearchParams(window.location.search);
    const user = query.get("u");
    const getUserData = async (user) => {
      try {
        let response = await fetchProfile(user);
        if (Object.keys(response).length > 0) {
          setUserData(response);
          setUserDataLoad(true);
        } else {
          setUserDataLoad(false);
        }
      } catch (error) {
        console.error(`Failed to fetch user data! [${error.message}]`);
        setUserDataLoad(false);
      }
    };
    if (typeof window !== "undefined") {
      getUserData(user);
    }
  }, []);

  return (
    session && (
      <>
        <NavBar />
        {userDataLoad ? (
          <>
            <ProfileGrid userData={userData} />
            <PostGrid user={userData.name} />
          </>
        ) : userDataLoad === false ? (
          <>
            <ErrorUser />
          </>
        ) : null}
      </>
    )
  );
};

export default Profile;
