"use client";

import AuthMenu from "@/components/auth/AuthMenu/AuthMenu";

import { useState, useEffect } from "react";

import { validSession } from "@/utils/auth";

const AuthPortal = () => {
  const [session, setSession] = useState(null);

  useEffect(() => {
    const checkSession = async () => {
      try {
        let response = await validSession();
        console.log(response);
        if (response) {
          return (location.href = "/home");
        }
        setSession(true);
      } catch (error) {
        setSession(true);
      }
    };
    checkSession();
  }, []);

  return (
    session && (
      <>
        <AuthMenu />
      </>
    )
  );
};

export default AuthPortal;
