import IconDelete from "./_components/IconDelete/IconLoad";

import { postDel } from "@/utils/posts";

import styles from "./PostBoxDelete.module.scss";

export default function PostBoxDelete({ id }) {
	const handleDelete = async () => {
		try {
			let response = await postDelete(id);
			if (response) {
				return (window.location.href = "/home");
			}
		} catch (error) {
			return console.error(`Failed to delete post! [${error.message}]`);
		}
	};

	return (
		<button onClick={handleDelete}>
			<IconDelete />
		</button>
	);
}
