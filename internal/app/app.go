package app

import (
	"encoding/json"
	"fmt"
	"todocli/internal/models"
	"todocli/utils"
)

type App struct {
	tasksFilepath string
	tasks         []models.Task
	running       bool
}

func NewTodocliApp(tasksFilepath string) *App {
	return &App{
		tasksFilepath: tasksFilepath,
		tasks:         make([]models.Task, 0, 100),
		running:       true,
	}
}

func (app *App) Run() {
	err := app.loadTasks()
	if err != nil {
		panic("Cannot load tasks: " + err.Error())
	}
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
			}else{
				fmt.Println("Task was succesfully created ")
			}
			app.writeTasks()
		case 2:
			app.PrintUpcomingTasks()
		case 3:
			app.PrintFinishedTasks()
		case 4:
			app.MarkAsDone()
			app.writeTasks()
		case 0:
			app.running = false
		default:
			fmt.Print("Unknown command.")
		}
	}

}

func (app *App) loadTasks() error {
	data, err := utils.ReadFile(app.tasksFilepath)
	if err != nil {
		fmt.Println("Cannot read file with path " + app.tasksFilepath + ":" + err.Error())
		return err
	}

	var tasks []models.Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Cannot unmarshal json file" + ":" + err.Error())
		return err
	}

	app.tasks = tasks
	return nil
}

func (app *App) writeTasks() error {
	data, err := json.Marshal(app.tasks)
	if err != nil {
		fmt.Println("Cannot save tasks")
		return err
	}

	err = utils.WriteFile(app.tasksFilepath, data)
	if err != nil {
		fmt.Println("Cannot save tasks")
		return err
	}

	return nil
}
