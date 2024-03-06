import React from "react";

const Navbar = () => {
  return (
    <div className="w-screen py-2 fixed z-10 bg-neutral-800 text-neutral-400 flex justify-center">
      <span className="cursor-pointer hover:text-white"><a href="/">Habits</a></span>
      <span className="absolute right-20 cursor-pointer hover:text-white"><a href="/profile">Profile</a></span>
    </div>
  )
}

export default Navbar;
