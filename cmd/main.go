package main

import (
	"os"
	"todocli/internal/app"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			println("Error occured during execution of the program:", r)
			os.Exit(1)
		}
	}()

	tasksPath := "tasks.json"

	app := app.NewTodocliApp(tasksPath)
	app.Run()
}
