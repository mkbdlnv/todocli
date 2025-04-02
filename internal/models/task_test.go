package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalJSON(t *testing.T) {
	id := "2f323da3-0ff9-11f0-8c37-b0227aef31cb"
	title := "learn_go"
	deadline := "2025-04-25T15:00:00Z"
	done := false
	data := fmt.Sprintf(`{ "id": "%s", "title": "%s", "time": "%s", "done": %t }`, id, title, deadline, done)

	task := &Task{}
	err := task.UnmarshalJSON([]byte(data))
	actualId := task.Id.String()
	actualTitle := task.Title
	actualDeadline := task.Deadline.Format(time.RFC3339)
	actualDone := task.Done

	assert.Nil(t, err)
	assert.Equal(t, id, actualId)
	assert.Equal(t, title, actualTitle)
	assert.Equal(t, actualDeadline, deadline)
	assert.Equal(t, done, actualDone)
}

func TestUnmarshalJSON_InvalidUUID(t *testing.T) {
	badData := `{"id":"not-a-uuid", "title":"fail", "time":"2025-04-25T15:00:00Z", "done":false}`
	task := &Task{}
	err := task.UnmarshalJSON([]byte(badData))
	assert.Error(t, err)
}

func TestUnmarshalJSON_IvalidTime(t *testing.T) {
	badData := `{"id":"2f323da3-0ff9-11f0-8c37-b0227aef31cb", "title":"fail", "time":"25.04.2025", "done":false}`
	task := &Task{}
	err := task.UnmarshalJSON([]byte(badData))
	assert.Error(t, err)
}
