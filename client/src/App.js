import React from "react";
import {createBrowserRouter, RouterProvider,} from "react-router-dom";
import HabitsPage from "./routes/HabitsPage";
import ProfilePage from "./routes/ProfilePage";

function App() {

  const router = createBrowserRouter([
    {
      path: "/",
      element: <HabitsPage />,
    },
    {
      path: "/profile",
      element: <ProfilePage />
    }
  ]);

  return (
    <div className="App font-mono overflow-x-hidden bg-neutral-600 relative">
      <RouterProvider router={router} />
    </div>
  );
}

export default App;
