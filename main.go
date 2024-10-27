package main

import (
	"log"
	"lontor/task_tracker/cli"
	"lontor/task_tracker/task"
	"os"
)

func main() {
	tm, err := task.NewTaskManager()
	if err != nil {
		log.Fatal(err)
	}

	err = cli.Execute(os.Args, tm)
	if err != nil {
		log.Fatal(err)
	}
}
