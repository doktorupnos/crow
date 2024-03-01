"use client";

import AuthGrid from "@/components/auth/AuthGrid/AuthGrid";

import { validSession } from "@/utils/auth";

import { useState, useEffect } from "react";

const LoginPortal = () => {
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
        <AuthGrid method={1} />
      </>
    )
  );
};

export default LoginPortal;
