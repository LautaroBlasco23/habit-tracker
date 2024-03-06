import React from "react";

const HabitCard = ({habit}) => {
  const {title, description, estado} = habit; 

  return (
    <div className="my-2 w-full border border-neutral-800 flex flex-col p-2 bg-neutral-800 text-neutral-300 rounded-t-lg">

      {/* Title with white bar below */}
      <h2 className="text-lg">{title}</h2>
      <div className="w-full h-[1px] bg-neutral-300"></div>

      {/* Description with white bar below */}
      <p className="my-2">{description}</p>
      <div className="w-full h-[0.5px] bg-neutral-300"></div>

      <div className="relative flex w-full justify-center mt-2">

        {/* Date and Hour info */}
        <div className="text-gray-500 absolute right-2 bottom-0 flex flex-col text-xs">
          <span className="">date: 03/12/23</span>
          <span className="">hour: 13:45</span>
        </div>

        {/* Status Button */}
        <button className="bg-neutral-300 text-neutral-950 border border-neutral-950 hover:bg-neutral-950 hover:text-neutral-300 hover:border-neutral-300 duration-150 m-auto p-1 rounded-b-lg  ">
          Change Status
        </button>

        {/* Status info */}
        <div className="text-gray-500 absolute left-2 bottom-0 flex flex-col text-xs">
          <span className="">status: <span className="text-green-300">{estado}</span></span>
        </div>

      </div>
    </div>
  )
}

export default HabitCard;
