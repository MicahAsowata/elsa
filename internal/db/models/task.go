package models

import "time"

type Task struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Details   string
	Completed bool
}
