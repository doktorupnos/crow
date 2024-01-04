"use client";

import { useState, useEffect } from "react";

import NavBar from "@/components/nav/NavBar/NavBar";
import ProfileGrid from "@/components/profile/ProfileGrid/ProfileGrid";
import ErrorUser from "@/components/error/ErrorUser/ErrorUser";

import { validSession } from "@/utils/auth";
import { fetchProfile } from "@/utils/profile";

const Profile = () => {
	const query = new URLSearchParams(window.location.search);
	const user = query.get("u");

	const [userData, setUserData] = useState({});
	const [userDataLoad, setUserDataLoad] = useState(null);

	const [session, setSession] = useState(null);

	useEffect(() => {
		const checkSession = async () => {
			try {
				let response = await validSession();
				if (!response) {
					console.error(`Invalid session!`);
					return (window.location.href = "/auth");
				}
				setSession(true);
			} catch (error) {
				console.error(`Invalid session! [${error.message}]`);
				return (window.location.href = "/auth");
			}
		};
		checkSession();
	}, []);

	useEffect(() => {
		const getUserData = async (user) => {
			try {
				let response = await fetchProfile(user);
				if (Object.keys(response).length > 0) {
					setUserData(response);
					setUserDataLoad(true);
				}
			} catch (error) {
				console.error(`Failed to fetch user data! [${error.message}]`);
			}
		};
		getUserData(user);
	}, [user]);

	return (
		session && (
			<>
				<NavBar />
				{!userDataLoad ? (
					<ErrorUser />
				) : (
					<>
						<ProfileGrid userData={userData} />
					</>
				)}
			</>
		)
	);
};

export default Profile;
