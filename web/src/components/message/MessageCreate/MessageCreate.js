import IconSend from "./_components/IconSend/IconSend";

import { useState, useEffect, useRef } from "react";

import styles from "./MessageCreate.module.scss";

const MessageCreate = ({ ws }) => {
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
      //let response = await messageSend(ws, textValue);
      ws.send(textValue);
      if (response) {
        setTextValue("");
      }
    } catch (error) {
      return console.error(`Failed to send message! [${error.message}] `);
    }
  };

  return (
    <div className={styles.message_create_base}>
      <form onSubmit={handleSubmit} className={styles.post_block}>
        <textarea
          ref={textareaRef}
          className={styles.message_box_text}
          type="text"
          value={textValue}
          onChange={handleChange}
          placeholder="Quoth the raven.."
          rows={1}
          required
        />
        <button className={styles.post_button} type="submit">
          <IconSend />
        </button>
      </form>
    </div>
  );
};

export default MessageCreate;
