"use client";

import React, { useState, useEffect } from "react";
import axios from "axios";
import PostNone from "@/components/post/PostNone/PostNone";
import NavBar from "@/components/nav/NavBar/NavBar";

import ProfileGrid from "@/components/profile/ProfileGrid/ProfileGrid";

const Profile = () => {
	const [followers, setFollowers] = useState(0);
	const [following, setFollowing] = useState(0);
	const [posts, setPosts] = useState([]);

	const handleFollow = async () => {
		try {
			await axios.post(process.env.followEndpoint, { withCredentials: true });
			await fetchFollowers();
		} catch (error) {
			console.error("Error following user:", error);
		}
	};

	const handleUnfollow = async () => {
		try {
			await axios.post(process.env.unfollowEndpoint, { withCredentials: true });
			await fetchFollowers();
		} catch (error) {
			console.error("Error unfollowing user:", error);
		}
	};

	const fetchFollowersCount = async () => {
		try {
			const response = await axios.get(
				process.env.fetchFollowersCountEndPoint,
				{ withCredentials: true }
			);
			setFollowers(response.data);
		} catch (error) {
			setFollowers(0);
			console.error("Error fetching followers count:", error);
		}
	};

	const fetchFollowingCount = async () => {
		try {
			const response = await axios.get(
				process.env.fetchFollowingCountEndPoint,
				{ withCredentials: true }
			);
			setFollowing(response.data);
		} catch (error) {
			console.error("Error fetching following count:", error);
			setFollowing(0);
		}
	};

	const fetchUserPosts = async () => {
		try {
			// TODO: Add request to add user posts
			setPosts([]);
		} catch (error) {
			console.error("Error gathering user posts:", error);
			setPosts([]);
		}
	};

	useEffect(() => {
		fetchFollowingCount();
		fetchFollowersCount();
		fetchUserPosts();
	}, []);

	return (
		<>
			<NavBar />
			<ProfileGrid />
			<div>
				{posts.length > 0 ? (
					posts.map((post) => <p key={post.id}>Hello</p>)
				) : (
					<PostNone />
				)}
			</div>
		</>
	);
};

export default Profile;
