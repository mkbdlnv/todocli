package main

import (
	"os"
	"path/filepath"
	"todocli/internal/app"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			println("Error occured during execution of the program:", r)
			os.Exit(1)
		}
	}()

	cwd, err := os.Getwd()
	const tasksFilename = "tasks.json"

	if err != nil {
		panic("Cannot get working directory")
	}
	tasksPath := filepath.Join(filepath.Dir(cwd), tasksFilename)

	app := app.NewTodocliApp(tasksPath)
	app.Run()
}
