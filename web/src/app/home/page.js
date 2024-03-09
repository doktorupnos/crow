"use client";

import NavBar from "@/components/nav/NavBar/NavBar";
import PostGrid from "@/components/post/PostGrid/PostGrid";

import { useState, useEffect } from "react";

import { validSession } from "@/utils/auth";

const Home = () => {
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
    checkSession();
  }, []);

  return (
    session && (
      <>
        <NavBar />
        <PostGrid />
      </>
    )
  );
};

export default Home;
