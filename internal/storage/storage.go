package storage

import (
	"encoding/json"
	"os"
	"todocli/internal/models"
)

type TaskStorage interface {
	LoadTasks() ([]models.Task, error)
	SaveTasks([]models.Task) error
}

type FileStorage struct {
	Path string
}

func (fs *FileStorage) LoadTasks() ([]models.Task, error) {
	data, err := os.ReadFile(fs.Path)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		return nil, err
	}

	var tasks []models.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (fs *FileStorage) SaveTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fs.Path, data, 0644)
}
