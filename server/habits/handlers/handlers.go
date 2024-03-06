package handlers

import (
	"fmt"

	"habitservice/initializers"
	"habitservice/models"
)

func GetAllHabitsHandler() ([]models.Habit, error) {
  db, err := initializers.ConnectToDB()
  if err != nil {
    fmt.Println("error")
    return []models.Habit{}, err
  }

  listOfHabits := []models.Habit{}
  db.Find(&listOfHabits)

  return listOfHabits, nil
}

func GetHabitByIdHandler(habit_id uint) (models.Habit, error) {
  db, err := initializers.ConnectToDB()
  if err != nil {
    fmt.Println("error")
    return models.Habit{}, fmt.Errorf("internal server error")
  }

  habit := models.Habit{}
  result := db.Find(&habit, habit_id)
  if result.RowsAffected == 0 {
    return models.Habit{}, fmt.Errorf("habit not found")
  }

  return habit, nil

}

func CreateHabitHandler(habit *models.Habit) (uint, error) {
  db, err := initializers.ConnectToDB()
  if err != nil {
    fmt.Println("error", err.Error())
    return 0, err
  }

  db.Create(habit)
  
  return habit.ID, nil
}

func ModifyHabitHandler(habitNewData models.Habit, habitId uint) (models.Habit, error) {
  db, err := initializers.ConnectToDB()
  if err != nil {
    fmt.Println("error", err.Error())
    return models.Habit{}, err
  }

  // Updating habit data:
  
  result := db.Model(&habitNewData).Where("id = ?", habitId).Update("title", habitNewData.Title).Update("text", habitNewData.Text)
  if result.RowsAffected == 0 {
    return models.Habit{}, fmt.Errorf("error modifying habit")
  }
  
  return habitNewData, nil
}

func DeleteHabitHandler(habit_id uint) (uint, error) {
  db, err := initializers.ConnectToDB()
  if err != nil {
    fmt.Println("error", err.Error())
    return 0, err
  }

  // Remove habit from database, if there are 0 rows affected, then, habit didn't exist.
  result := db.Delete(&models.Habit{}, &habit_id)
  if result.RowsAffected == 0 {
    return 0, fmt.Errorf("error removing habit")
  }
  
  return habit_id, nil
}
