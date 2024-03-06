package controllers

import (
	"encoding/json"
	"fmt"
	"habitservice/handlers"
	"habitservice/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func GetAllHabits(w http.ResponseWriter, req *http.Request) {
  // All returns use JSON, so we prepare the encoder.
  w.Header().Set("Content-Type", "application/json")
  json_response := json.NewEncoder(w)

  // Getting all habits from handler.
  listOfHabits, err := handlers.GetAllHabitsHandler()
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json_response.Encode(models.BasicResponse{
      Status: "error",
      Message: err.Error(),
    })
    return
  }

  w.WriteHeader(http.StatusOK)
  json_response.Encode(listOfHabits)
}


func GetHabitById(w http.ResponseWriter, req *http.Request) {
  // All returns use JSON, so we prepare the encoder.
  json_response := json.NewEncoder(w)
  w.Header().Set("Content-Type", "application/json")

  // Getting id from request params, and parsing it to uint.
  habitId, err := strconv.ParseUint(mux.Vars(req)["habitId"], 10, 32)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json_response.Encode(models.BasicResponse{
      Status: "error",
      Message: "invalid id",
    })
    return
  }

  // getting the id from handlers.
  habitData, err := handlers.GetHabitByIdHandler(uint(habitId))
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json_response.Encode(models.BasicResponse{
      Status: "error",
      Message: err.Error(),
    })
    return
  }
  

  // Encoding response and setting it's headers
  w.WriteHeader(http.StatusOK)
  json_response.Encode(habitData)
}


func CreateHabit(w http.ResponseWriter, req *http.Request) {
  // All returns use JSON, so we prepare the encoder.
  json_response := json.NewEncoder(w)
  w.Header().Set("Content-Type", "application/json")

  // Reading json data from request body and saving it into new variable.
  var newHabitData models.Habit
  err := json.NewDecoder(req.Body).Decode(&newHabitData)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json_response.Encode(models.BasicResponse{
      Status: "error",
      Message: err.Error(),
    })
    return
  }

  habitId, err := handlers.CreateHabitHandler(&newHabitData)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json_response.Encode(models.BasicResponse{
      Status: "error",
      Message: err.Error(),
    })
    return
  }

  w.WriteHeader(http.StatusCreated)
  json_response.Encode(habitId)
}


func ModifyHabit(w http.ResponseWriter, req *http.Request) {
  // All returns use JSON, so we prepare the encoder.
  json_response := json.NewEncoder(w)
  w.Header().Set("Content-Type", "application/json")

  habitId, err := strconv.ParseUint(mux.Vars(req)["habitId"], 10, 32)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json_response.Encode(models.BasicResponse{
      Status: "error",
      Message: err.Error(),
    })
    return
  }

  // Getting habit data from request's body, and saving it into variable.
  var habitToModify models.Habit
  err = json.NewDecoder(req.Body).Decode(&habitToModify)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json_response.Encode(models.BasicResponse{
      Status: "error",
      Message: err.Error(),
    })
    return
  }

  fmt.Println(habitToModify)

  // Change habit in database, if error, return error.
  modifiedHabit, err := handlers.ModifyHabitHandler(habitToModify, uint(habitId))
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json_response.Encode(models.BasicResponse{
      Status: "error",
      Message: "internal server error",
    })
    return
  }

  // Success Response
  w.WriteHeader(http.StatusOK)
  json_response.Encode(modifiedHabit)
}


func DeleteHabit(w http.ResponseWriter, req *http.Request) {
  // creating our json response encoder.
  json_response := json.NewEncoder(w)
  w.Header().Set("Content-Type", "application/json")
  
  // Getting habit id from request params.
  habitId, err := strconv.ParseUint(mux.Vars(req)["habitId"], 10, 32)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    json_response.Encode(models.BasicResponse{
      Status: "error",
      Message: "invalid habit id",
    })
    return
  }

  deletedHabitId, err := handlers.DeleteHabitHandler(uint(habitId))
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json_response.Encode(models.BasicResponse{
      Status: "error",
      Message: err.Error(),
    })
    return
  }
  
  w.WriteHeader(http.StatusOK)
  json_response.Encode(models.BasicResponse{
      Status: "success",
      Message: fmt.Sprint("habit with id ", deletedHabitId, " deleted"),
  })
}

func InvalidRoute(w http.ResponseWriter, req *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusBadRequest)
  json.NewEncoder(w).Encode(models.BasicResponse{
    Status: "error",
    Message: "invalid route",
  })
}


