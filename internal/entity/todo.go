package entity

import (
	"time"
)

type Todo struct {
	ID          int
	Title       string
	Description string
	DueDate     time.Time
	Priority    int
	Status      string
}
