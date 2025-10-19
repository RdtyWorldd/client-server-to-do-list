package task

import "time"

type Progress string

const (
	TODO       Progress = "todo"
	INPROGRESS Progress = "in-progress"
	DONE       Progress = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Progress  `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
