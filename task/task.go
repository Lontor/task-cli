package task

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// File name to save tasks
const taskFile = "tasks.json"

// Attempts to load tasks from a file, returns a Task map and error
func loadTasks() (map[int]Task, error) {
	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		return make(map[int]Task, 8), nil
	}

	data, err := os.ReadFile(taskFile)
	if err != nil {
		return make(map[int]Task, 0), fmt.Errorf("error reading task file: %v", err)
	}

	var tasks map[int]Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return make(map[int]Task, 0), fmt.Errorf("error parsing tasks file: %v", err)
	}

	return tasks, nil
}

// Attempts to load tasks from a file, returns an error
func saveTasks(taskMap map[int]Task) error {
	data, err := json.Marshal(taskMap)
	if err != nil {
		return fmt.Errorf("error writing tasks to file: %v", err)
	}

	err = os.WriteFile(taskFile, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing tasks to file: %v", err)
	}

	return nil
}
