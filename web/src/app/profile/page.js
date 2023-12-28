"use client";

import { useState, useEffect } from "react";
import PostNone from "@/components/post/PostNone/PostNone";
import NavBar from "@/components/nav/NavBar/NavBar";

import ProfileGrid from "@/components/profile/ProfileGrid/ProfileGrid";

import { getProfile } from "@/utils/profile";

const Profile = () => {
	const query = new URLSearchParams(window.location.search);
	const user = query.get("u");

	const [userData, setUserData] = useState({});
	const [posts, setPosts] = useState([]);

	/*
	const fetchUserPosts = async () => {
		try {
			// TODO: Add request to add user posts
			setPosts([]);
		} catch (error) {
			console.error("Error gathering user posts:", error);
			setPosts([]);
		}
	};
	*/

	/*
	useEffect(() => {
		fetchUserPosts();
	}, []);
	*/

	useEffect(() => {
		const fetchUserData = async (user) => {
			try {
				const response = await getProfile(user);
				setUserData(response);
			} catch (error) {
				console.error(`Failed to fetch user data! [${error.message}]`);
			}
		};
		fetchUserData(user);
	}, [user]);

	return (
		<>
			<NavBar />
			{Object.keys(userData).length > 0 ? (
				<>
					<ProfileGrid userData={userData} />
					{posts.length > 0 ? (
						posts.map((post) => <p key={post.id}>Hello</p>)
					) : (
						<PostNone />
					)}
				</>
			) : (
				<p>User does not exist</p>
			)}
		</>
	);
};

export default Profile;
