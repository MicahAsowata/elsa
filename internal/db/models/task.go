package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey;not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	Name      string    `gorm:"not null"`
	Details   string    `gorm:"not null"`
	Start     time.Time `gorm:"not null"`
	End       time.Time `gorm:"not null"`
	Completed bool      `gorm:"not null"`
	Tag       string    `gorm:"type:varchar(255);not null"`
}
