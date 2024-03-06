import React from "react";
import { CiEdit } from "react-icons/ci";

const ProfileCard = () => {
  return (
    <div className="bg-neutral-800 rounded-md p-8 flex flex-col shadow-2xl shadow-black border-2 border-neutral-300">
      <img className="shadow-neutral-600 shrink-0 grow-0 shadow-md mx-auto z-10 h-[150px] w-[150px] rounded-full border-2 border-neutral-800" src={"https://images.ctfassets.net/h6goo9gw1hh6/2sNZtFAWOdP1lmQ33VwRN3/24e953b920a9cd0ff2e1d587742a2472/1-intro-photo-final.jpg?w=1200&h=992&fl=progressive&q=70&fm=jpg"} />
      <div className="mt-4 text-neutral-300 flex items-center">
          <span className="font-bold text-lg mr-4">
            Username: 
          </span>
        <h1>
          LautaroBlasco23
        </h1>
        <CiEdit className="text-xl ml-2 hover:text-white duration-100 cursor-pointer"/>
      </div>

      <div className="h-[1px] w-full bg-white mt-2"></div>

      <div className="text-neutral-300 mt-4 flex-col flex">
        <div className="flex items-center">
          <span className="font-bold text-lg mr-4">
            Email:
          </span>
          <span className="">
          lautaroblasco98@gmail.com
          </span>
          <CiEdit className="text-xl ml-2 hover:text-white duration-100 cursor-pointer"/>
        </div>
          <span className=""><span className="font-bold text-lg mr-4">habits tracked:</span>10000</span>
      </div>

    </div>
  )
}

  export default ProfileCard;
