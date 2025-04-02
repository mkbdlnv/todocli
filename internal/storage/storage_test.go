package storage

import (
	"encoding/json"
	"os"
	"testing"
	"time"
	"todocli/internal/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSaveTasks(t *testing.T) {
	tmpFile := "test_tasks.json"
	defer os.Remove(tmpFile)
	fs := &FileStorage{Path: tmpFile}

	uid, _ := uuid.NewUUID()
	title := "Learn_Golang"
	deadline := time.Now()
	done := false

	tasks := []models.Task{
		models.Task{Id: uid, Title: title, Deadline: deadline, Done: done},
	}

	err := fs.SaveTasks(tasks)
	assert.NoError(t, err)
	data, err := os.ReadFile(tmpFile)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)

	loaded := []models.Task{}
	err = json.Unmarshal(data, &loaded)
	assert.NoError(t, err)
	assert.Len(t, loaded, 1)
	assert.Equal(t, title, loaded[0].Title)
	assert.Equal(t, uid, loaded[0].Id)
	assert.WithinDuration(t, deadline, loaded[0].Deadline, time.Second)
	assert.Equal(t, done, loaded[0].Done)

}

func TestLoadTasks(t *testing.T) {
	tmpFile := "test_tasks.json"
	defer os.Remove(tmpFile)
	fs := &FileStorage{Path: tmpFile}
	data := []byte(`[{"id":"2f323da3-0ff9-11f0-8c37-b0227aef31cb", "title":"fail", "time":"2025-04-25T15:00:00Z", "done":false}]`)
	err := os.WriteFile(fs.Path, data, 0644)

	assert.Nil(t, err)
	loaded, err := fs.LoadTasks()
	assert.Nil(t, err)
	assert.Len(t, loaded, 1)
}

func TestLoadTasksWhenFileNotFound(t *testing.T) {
	tmpFile := "non_existing_file.json"
	fs := &FileStorage{Path: tmpFile}

	loaded, err := fs.LoadTasks()
	assert.Nil(t, err)
	assert.Len(t, loaded, 0)
}

func TestLoadTasks_PermissionError(t *testing.T) {
	tmp := "test_perm.json"
	err := os.WriteFile(tmp, []byte(`[]`), 0000) 
	assert.NoError(t, err)
	defer os.Remove(tmp)

	fs := &FileStorage{Path: tmp}
	_, err = fs.LoadTasks()

	assert.Error(t, err)
	assert.True(t, os.IsPermission(err))
}

