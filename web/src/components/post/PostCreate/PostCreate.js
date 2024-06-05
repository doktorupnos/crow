import IconCreate from "./_components/IconCreate";

import { useState, useEffect, useRef } from "react";

import { postCreate } from "@/utils/posts";

import styles from "./PostCreate.module.scss";

const PostCreate = ({ appendNewPost }) => {
  const [textValue, setTextValue] = useState("");

  const textareaRef = useRef(null);

  const handleChange = (event) => {
    const postText = event.target.value;
    if (postText.trim() !== "") {
      const trimmedText = postText.substring(0, 280);
      return setTextValue(trimmedText);
    } else {
      return setTextValue("");
    }
  };

  // Auto Resize Text Input
  useEffect(() => {
    if (textareaRef.current) {
      textareaRef.current.style.height = "auto";
      textareaRef.current.style.height =
        textareaRef.current.scrollHeight + "px";
    }
  }, [textValue]);

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      await postCreate(textValue);
      appendNewPost();
      setTextValue("");
    } catch (error) {
      return console.error(`Failed to create post! [${error.message}] `);
    }
  };

  return (
    <form onSubmit={handleSubmit} className={styles.post_block}>
      <textarea
        ref={textareaRef}
        className={styles.post_create_text}
        type="text"
        value={textValue}
        onChange={handleChange}
        placeholder="Quoth the raven.."
        rows={1}
        required
      />
      <button className={styles.post_button} type="submit">
        <IconCreate />
      </button>
    </form>
  );
};

export default PostCreate;
