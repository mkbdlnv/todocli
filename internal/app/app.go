package app

import (
	"fmt"
	"todocli/internal/models"
	"todocli/internal/storage"
)

type App struct {
	storage storage.TaskStorage
	tasks   []models.Task
	running bool
}

func NewTodocliApp(tasksFilepath string) *App {
	fs := &storage.FileStorage{Path: tasksFilepath}
	tasks, err := fs.LoadTasks()
	if err != nil {
		panic("Cannot load task")
	}
	return &App{
		storage: fs,
		tasks:   tasks,
		running: true,
	}
}

func (app *App) Run() {
	app.runMenu()

}

func (app *App) runMenu() {
	fmt.Print(`Welcome to TODOCLI app. Here you can manage your daily tasks.
List of available commands:
1. Create task
2. See upcoming tasks
3. See finished tasks
4. Mark as done
0. Stop program
`)

	for app.running {
		var option uint8
		fmt.Print("Enter command number: ")
		fmt.Scanln(&option)
		switch option {
		case 1:
			err := app.CreateTask()
			if err != nil {
				fmt.Println("cannot create task: ", err)
			} else {
				fmt.Println("Task was succesfully created ")
			}
			app.storage.SaveTasks(app.tasks)
		case 2:
			app.PrintUpcomingTasks()
		case 3:
			app.PrintFinishedTasks()
		case 4:
			app.MarkAsDone()
			app.storage.SaveTasks(app.tasks)
		case 0:
			app.running = false
		default:
			fmt.Print("Unknown command.")
		}
	}

}
