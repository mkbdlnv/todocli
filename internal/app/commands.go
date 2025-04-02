package app

import (
	"fmt"
	"time"
	"todocli/internal/models"

	"github.com/google/uuid"
)
const timeFormat = "02-01-2006 15:04:05"

func (app *App) CreateTask() error {

	var title, dateStr, timeStr string


	fmt.Print("Enter task name: ")
	fmt.Scanln(&title)

	fmt.Print("Enter task deadline (date) in format dd-mm-yyyy: ")
	fmt.Scanln(&dateStr)

	fmt.Print("Enter task deadline (time) in format hh:mm:ss: ")
	fmt.Scanln(&timeStr)

	dateTimeStr := dateStr + " " + timeStr
	deadline, err := time.Parse(timeFormat, dateTimeStr)
	if err != nil {
		return err
	}

	id, _ := uuid.NewUUID()

	task := models.Task{
		Id: id,
		Title: title,
		Deadline: deadline,
		Done: false,
	}
	app.tasks = append(app.tasks, task)
	return err
	
}

func (app *App) PrintUpcomingTasks(){
	for i, task := range app.tasks{
		if !task.Done{
			fmt.Printf("%d. ID: %s, Title: %s, Deadline: %s\n", i, task.Id.String(), task.Title, task.Deadline)
		}	
	}
}

func (app *App) PrintFinishedTasks(){
	for i, task := range app.tasks{
		if task.Done{
			fmt.Printf("%d. ID: %s, Title: %s, Deadline: %s\n", i, task.Id.String(), task.Title, task.Deadline)
		}	
	}
}

func (app *App) MarkAsDone(){
	var id string
	fmt.Print("Enter task ID to makr as done: ")
	
	fmt.Scanln(&id)

	for i := range app.tasks{
		if app.tasks[i].Id.String() == id{
			app.tasks[i].Done = true
			fmt.Println("Task with id " + id + " was successfully mark as done")
			return
		}
	}
	fmt.Println("Task with id " + id + " not found")
}