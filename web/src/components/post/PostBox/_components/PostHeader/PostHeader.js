import styles from "./PostHeader.module.scss";

const handleVisitProfile = async (user) => {
  return (location.href = `/profile?u=${user}`);
};

const PostHeader = ({ author, date }) => {
  return (
    <header className={styles.post_header}>
      <button onClick={() => handleVisitProfile(author)}>
        <span className={styles.post_author}>@{author}</span>
      </button>
      <span className={styles.post_date}>{date}</span>
    </header>
  );
};

export default PostHeader;
