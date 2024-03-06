import React from "react";
import HabitCard from "./HabitCard";

const HabitsContainer = () => {
  const listOfData = []

  const HabitData = {
    "title": "Hacer las compras matutinas lo antes posible",
    "description": "leche, crema, salchichas, fideos, huevos, harina, carne, pollo, mayonesa, papel, servilleta",
    "estado": "Sin hacer"
  }

  for (let i = 0; i < 10; i++) {
    listOfData.push(HabitData);
  }

  return (
    <div className="bg-neutral-600 text-neutral-200 h-full flex flex-col items-center relative w-1/2 m-auto">
      <h1 className="font-bold text-2xl text-neutral-300 py-2">List of Habits</h1>
      {
        listOfData.map((habit) => {
          return (
            <HabitCard habit={habit}/>
          )
        })
      }
    </div>
  )
}

export default HabitsContainer;
