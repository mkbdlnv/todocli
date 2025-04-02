package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Deadline time.Time `json:"time"`
	Done     bool      `json:"done"`
}

func (task *Task) UnmarshalJSON(data []byte) error {
	var raw struct {
		Id       string    `json:"id"`
		Title    string    `json:"title"`
		Deadline time.Time `json:"time"`
		Done     bool      `json:"done"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	uid, err := uuid.Parse(raw.Id)
	if err != nil {
		return err
	}
	task.Id = uid
	task.Title = raw.Title
	task.Deadline = raw.Deadline
	task.Done = raw.Done
	return nil
}
