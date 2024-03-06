package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
  dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable ",
  os.Getenv("HABITS_DB_USER"), os.Getenv("HABITS_DB_PASSWORD"), os.Getenv("HABITS_DB_NAME"))


  return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
