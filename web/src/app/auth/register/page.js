"use client";

import AuthGrid from "@/components/auth/AuthGrid/AuthGrid";

import { validSession } from "@/utils/auth";

import { useState, useEffect } from "react";

const RegisterPortal = () => {
  const [session, setSession] = useState(null);

  useEffect(() => {
    const checkSession = async () => {
      try {
        let response = await validSession();
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
        <AuthGrid method={0} />
      </>
    )
  );
};

export default RegisterPortal;
