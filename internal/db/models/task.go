package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string
	Details   string
	Start     time.Time
	End       time.Time
	Completed bool
	Tags      []string
}
