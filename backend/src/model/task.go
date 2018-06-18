package model

import "fmt"

// Task representa a tarefa para ser gerenciada
type Task struct {
	Description string `json:"description,omitempty"`
	Done        bool   `json:"done"`
}

func (t Task) String() string {
	return fmt.Sprintf("%v - %v", t.Description, t.Done)
}
