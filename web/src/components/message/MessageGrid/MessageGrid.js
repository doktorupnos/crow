import MessageCreate from "@/components/message/MessageCreate/MessageCreate";
import PostBox from "@/components/post/PostBox/PostBox";

import { useState, useEffect } from "react";

import { fetchPosts } from "@/utils/posts";

import styles from "./MessageGrid.module.scss";

const MessageGrid = ({ user }) => {
  const [postList, setPostList] = useState([]);
  const [postLoad, setPostLoad] = useState(false);
  const [morePosts, setMorePosts] = useState(null);
  const [page, setPage] = useState(0);

  // Create Echo Websocket
  const ws = new WebSocket("http://localhost:8000/api/ws/echo");
  ws.onmessage = (event) => {
    console.log(event.data);
  };

  useEffect(() => {
    const getPosts = async (page) => {
      try {
        let response;
        if (user) {
          response = await fetchPosts(user, page, null);
        } else {
          response = await fetchPosts(null, page, null);
        }
        if (!response.length > 0) {
          setPostLoad(false);
          return setMorePosts(false);
        }
        let newList = response.map((post) => {
          return <PostBox key={post.id} post={post} />;
        });
        setPostList((prevList) => [...prevList, newList]);
        setMorePosts(true);
        setPostLoad(false);
      } catch (error) {
        console.error(`Failed to retrieve posts! [${error.message}]`);
      }
    };
    getPosts(page);
  }, [page, user]);

  useEffect(() => {
    const handleScrollBottom = () => {
      const isScrollAtBottom =
        window.innerHeight + window.scrollY >= document.body.scrollHeight;
      if (isScrollAtBottom && !postLoad && morePosts) {
        setPage((page) => page + 1);
        setPostLoad(true);
      }
    };
    window.addEventListener("scroll", handleScrollBottom);
    return () => {
      window.removeEventListener("scroll", handleScrollBottom);
    };
  }, [postLoad, morePosts]);

  const handleLoad = () => {
    if (!postLoad && morePosts) {
      setPage((page) => page + 1);
      setPostLoad(true);
    }
  };

  const appendNewPost = async () => {
    try {
      let response = await fetchPosts(null, 0, 1);
      if (response) {
        let newPost = response.map((post) => {
          return <PostBox key={post.id} post={post} />;
        });
        setPostList((prevList) => [newPost, ...prevList]);
      }
    } catch (error) {
      console.error(`Failed to load created post! ${error.message}`);
    }
  };

  return (
    <>
      {postList.length > 0 && (
        <>
          <div className={styles.post_grid}>{postList}</div>
          {morePosts && !postLoad && (
            <button className={styles.post_load} onClick={handleLoad}>
              <p>load</p>
            </button>
          )}
          {postLoad && <p>spin</p>}
        </>
      )}
      <div className={styles.message_space}></div>
      <MessageCreate ws={ws} appendMessage={appendNewPost} />
      {morePosts == false && postList.length == 0 && <p>error</p>}
    </>
  );
};

export default MessageGrid;
