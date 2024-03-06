package main

import (
	"fmt"
	"habitservice/controllers"
	"habitservice/initializers"
	"habitservice/models"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
  router := mux.NewRouter()

  db, err := initializers.ConnectToDB()
  if err != nil {
    fmt.Println("error initializing database!")
  }
  if err = db.AutoMigrate(&models.Habit{}); err != nil {
    fmt.Println("Error making models migrations to database!")
  }

  fmt.Println("Database ready to use!")

  // Routes
  router.HandleFunc("/habits", controllers.GetAllHabits).Methods("GET")
  router.HandleFunc("/habit/{habitId}", controllers.GetHabitById).Methods("GET")
  router.HandleFunc("/habit/create", controllers.CreateHabit).Methods("POST")
  router.HandleFunc("/habit/modify/{habitId}", controllers.ModifyHabit).Methods("PUT")
  router.HandleFunc("/habit/delete/{habitId}", controllers.DeleteHabit).Methods("DELETE")

  router.PathPrefix("/").HandlerFunc(controllers.InvalidRoute)
  
  fmt.Println("Server running!")

  http.ListenAndServe(":8080", router)
}
