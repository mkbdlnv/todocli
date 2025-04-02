package main

import (
	"os"
	"path/filepath"
	"todocli/internal/app"
)

func main() {
	//TODO todo cli app
	cwd, err := os.Getwd()
	const tasksFilename = "tasks.json"

	if err != nil {
		panic("Cannot load tasks file")
	}
	tasksPath := filepath.Join(filepath.Dir(cwd), tasksFilename)
	
	app := app.NewTodocliApp(tasksPath)
	app.Run()
}
