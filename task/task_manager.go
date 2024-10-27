package task

import (
	"fmt"
	"time"
)

type TaskManager struct {
	tasks map[int]Task
}

// Creates a task manager and returns an error on failure.
func NewTaskManager() (*TaskManager, error) {
	tasks, err := loadTasks()
	if err != nil {
		return nil, err
	}
	return &TaskManager{tasks: tasks}, nil
}

// Returns a copy of the map
func (tm *TaskManager) copyTasks() map[int]Task {
	tasksCopy := make(map[int]Task, len(tm.tasks))
	for key, value := range tm.tasks {
		tasksCopy[key] = value
	}
	return tasksCopy
}

// Adds a task with a unique id to the file
func (tm *TaskManager) AddTask(description string) (int, error) {
	id := 0
	for {
		_, exist := tm.tasks[id]
		if !exist {
			break
		}
		id++
	}
	tm.tasks[id] = Task{
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return id, saveTasks(tm.tasks)
}

// Updates the task if it exists and returns an error if it does not
func (tm *TaskManager) UpdateTask(id int, description string) error {
	task, exist := tm.tasks[id]
	if !exist {
		return fmt.Errorf("task with id %d not found", id)
	}
	task.Description = description
	task.UpdatedAt = time.Now()
	tm.tasks[id] = task
	return saveTasks(tm.tasks)
}

// Deletes the task if it exists and returns an error if it does not
func (tm *TaskManager) DeleteTask(id int) error {
	_, exist := tm.tasks[id]
	if !exist {
		return fmt.Errorf("task with id %d not found", id)
	}
	delete(tm.tasks, id)
	return saveTasks(tm.tasks)
}

// Changes the task status, possible values:
// todo,
// in-progress,
// done
func (tm *TaskManager) MarkTask(id int, status string) error {
	task, exist := tm.tasks[id]
	if !exist {
		return fmt.Errorf("task with id %d not found", id)
	}
	if status != "todo" && status != "in-progress" && status != "done" {
		return fmt.Errorf("invalid status: %s", status)
	}
	task.Status = status
	task.UpdatedAt = time.Now()
	tm.tasks[id] = task
	return saveTasks(tm.tasks)
}

// Returns a task map with the ability to filter by status
func (tm *TaskManager) ListTasks(status string) (map[int]Task, error) {
	if status == "" {
		return tm.copyTasks(), nil
	}
	if status != "todo" && status != "in-progress" && status != "done" {
		return nil, fmt.Errorf("invalid status: %s", status)
	}
	filteredTasks := make(map[int]Task)
	for id, value := range tm.tasks {
		if value.Status == status {
			filteredTasks[id] = value
		}
	}
	return filteredTasks, nil
}
