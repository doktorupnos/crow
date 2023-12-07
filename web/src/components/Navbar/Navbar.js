"use client";

import './Navbar.css';
import Image from 'next/image';
import { redirect } from "next/navigation";

const Navbar = () => {
    return (
            <>
                <nav>
                    <ul>
                        <li style={{marginLeft: 2 + 'em'}}>
                            <h1>Crow</h1>
                        </li>
                        <li>    
                            <a href="https://github.com/doktorupnos/crow" target="_blank" rel="noopener noreferrer">
                                <Image
                                  src="/images/logo.p"
                                  alt="GitHub"
                                  priority={false}
                                  width={50}
                                  height={50}
                                />
                            </a>
                        </li>
                        <li style={{marginRight: 2 + 'em'}}>
                            <Profile />
                        </li>
                    </ul>
                </nav> 
            </>
    );
};

export default Navbar;

function Profile() {
    return (
        <>
            <button onClick={navigateToProfile}>My Profile</button>
        </>
    )
}

function navigateToProfile() {
	redirect("/profile");
}
