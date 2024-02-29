"use client";

import styles from "./AuthForm.module.scss";

import { useState } from "react";
import { useRouter } from "next/navigation";

import { userRegister, userLogin } from "@/utils/auth";

import AuthField from "../AuthField/AuthField";

export default function AuthForm({ method }) {
  const router = useRouter();

  let [post, setPost] = useState({
    name: "",
    password: "",
  });

  function handleInput(event) {
    setPost({ ...post, [event.target.id]: event.target.value });
  }

  async function handleSubmit(event) {
    event.preventDefault();
    try {
      let response;
      if (method) {
        response = await userLogin(post);
      } else {
        response = await userRegister(post);
      }
      if (response) {
        return (location.href = "/home");
      }
    } catch (error) {
      return console.error(`Failed to auth user! [${error.message}]`);
    }
  }

  return (
    <form className={styles.auth_form} onSubmit={handleSubmit}>
      <AuthField type={1} handler={handleInput} />
      <AuthField type={2} handler={handleInput} />
      <input
        className={styles.auth_submit}
        type="submit"
        value={method ? "Enter" : "Join"}
      />
    </form>
  );
}
