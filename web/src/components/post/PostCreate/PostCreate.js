import Image from "next/image";

import { useState, useEffect, useRef } from "react";

import { postCreate } from "@/utils/posts";

import styles from "./PostCreate.module.scss";

const PostCreate = ({ appendNewPost }) => {
  const [textValue, setTextValue] = useState("");

  const textareaRef = useRef(null);

  const handleChange = (event) => {
    const postText = event.target.value;
    if (postText.trim() !== "") {
      setTextValue(event.target.value);
    } else {
      setTextValue("");
    }
  };

  useEffect(() => {
    const autoResize = () => {
      if (textareaRef.current) {
        textareaRef.current.style.height = "auto";
        textareaRef.current.style.height =
          textareaRef.current.scrollHeight + "px";
      }
    };
    autoResize();
  }, [textValue]);

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      let response = await postCreate(textValue);
      if (response) {
        appendNewPost();
      }
    } catch (error) {
      return console.error(`Failed to create post! [${error.message}] `);
    }
  };

  return (
    <form onSubmit={handleSubmit} className={styles.post_block}>
      <textarea
        ref={textareaRef}
        className={styles.post_text}
        type="text"
        value={textValue}
        onChange={handleChange}
        placeholder="Quoth the raven.."
        required
      />
      <button className={styles.post_button} type="submit">
        <Image
          className={styles.post_button_image}
          src="/images/bootstrap/post_create.svg"
          alt="create post"
          width={25}
          height={25}
          draggable="false"
        />
      </button>
    </form>
  );
};

export default PostCreate;
