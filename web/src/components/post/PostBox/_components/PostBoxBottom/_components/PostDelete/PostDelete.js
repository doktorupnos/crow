import IconDelete from "./_components/IconDelete/IconDelete";

import { postDelete } from "@/utils/posts";

const PostDelete = ({ id }) => {
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
};

export default PostDelete;
