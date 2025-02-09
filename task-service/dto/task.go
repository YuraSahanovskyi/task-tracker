package dto

import (
	"encoding/json"
	"fmt"
	"time"
)

const DEFAULT_STATUS string = "todo"
const DEFAULT_PRIORITY int = 3
const MIN_PRIORITY int = 1
const MAX_PRIORITY int = 5

type TaskDto struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Status      string    `json:"status"`
	DueTo       time.Time `json:"due_to"`
	//ProjectID
	//AssignedTo
	//CreatedBy
}

func (dto *TaskDto) Validate() error {
	if dto == nil {
		return fmt.Errorf(`["TaskDto is nil"]`)
	}

	var errors []string

	if dto.Name == "" {
		errors = append(errors, "name is mandatory")
	}
	if dto.Priority == 0 {
		dto.Priority = DEFAULT_PRIORITY
	} else if dto.Priority < MIN_PRIORITY || dto.Priority > MAX_PRIORITY {
		errors = append(errors, fmt.Sprintf("priority must be between %d and %d", MIN_PRIORITY, MAX_PRIORITY))
	}
	if dto.Status == "" {
		dto.Status = DEFAULT_STATUS
	}
	if dto.DueTo.Before(time.Now().Truncate(time.Minute)) {
		errors = append(errors, "due_to must be in the future")
	}

	if len(errors) > 0 {
		jsonErrors, _ := json.Marshal(errors)
		return fmt.Errorf("%s", string(jsonErrors))
	}

	return nil
}
