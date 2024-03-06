package models

import (
	"time"

	"gorm.io/gorm"
)

type Habit struct {
  ID        uint         `gorm:"primaryKey json:"id"`
  CreatedAt time.Time      `json:"-"`
  UpdatedAt time.Time      `json:"-"`
  DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
  Title string             `json:"title"`
  Text string              `json:"text"`
}
