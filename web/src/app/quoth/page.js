"use client";

import NavBar from "@/components/nav/NavBar/NavBar";
import MessageGrid from "@/components/message/MessageGrid/MessageGrid";

import { useEffect, useState } from "react";

import { validSession } from "@/utils/auth";

const Quoth = () => {
  const [session, setSession] = useState(null);

  const ws = new WebSocket("http://localhost:8000/api/ws/echo");
  ws.onmessage = (event) => {
    console.log(event.data);
  };

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
        <MessageGrid />
      </>
    )
  );
};

export default Quoth;
