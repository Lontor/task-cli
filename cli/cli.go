package cli

import (
	"fmt"
	"lontor/task_tracker/task"
	"strconv"
	"time"
)

const helpText = `Available commands:
  
  add <description>
      Adds a new task with the specified description.

  update <id> <new_description>
      Updates the task with the specified id to the new description.

  delete <id>
      Deletes the task with the specified id.

  mark <id> <status>
      Marks the task with the specified id as one of the following statuses:
      - todo
      - in-progress
      - done

  list [status]
      Lists all tasks. Optionally filter tasks by status:
      - todo
      - in-progress
      - done

  help
      Displays this help text.
  
  Note: All commands are case-sensitive.
`

// Defines a command and calls the required handler
func Execute(args []string, tm *task.TaskManager) error {
	if len(args) <= 1 {
		return fmt.Errorf("no command was specified, try %s help", args[0])
	}
	switch args[1] {
	case "add":
		return handleAdd(args, tm)
	case "update":
		return handleUpdate(args, tm)
	case "delete":
		return handleDelete(args, tm)
	case "mark":
		return handleMark(args, tm)
	case "list":
		return handleList(args, tm)
	case "help":
		return handleHelp()
	default:
		return fmt.Errorf("%s %s: unknown command, try %s help", args[0], args[1], args[0])
	}
}

// Checks the arguments and adds the task
func handleAdd(args []string, tm *task.TaskManager) error {
	if len(args) < 3 {
		return fmt.Errorf("no arguments for add command")
	}
	if len(args) > 3 {
		return fmt.Errorf("too many arguments for add command")
	}
	id, err := tm.AddTask(args[2])
	if err != nil {
		return err
	}
	fmt.Printf("Task added successfully (ID: %d)\n", id)
	return nil
}

// Checks the arguments and updates the task
func handleUpdate(args []string, tm *task.TaskManager) error {
	if len(args) < 4 {
		return fmt.Errorf("no arguments for update command")
	}
	if len(args) > 4 {
		return fmt.Errorf("too many arguments for update command")
	}
	id, err := strconv.Atoi(args[2])
	if err != nil {
		return err
	}
	err = tm.UpdateTask(id, args[3])
	if err != nil {
		return err
	}
	fmt.Printf("Task #%d updated successfully\n", id)
	return nil
}

// Checks the arguments and deletes the task
func handleDelete(args []string, tm *task.TaskManager) error {
	if len(args) < 3 {
		return fmt.Errorf("no arguments for delete command")
	}
	if len(args) > 3 {
		return fmt.Errorf("too many arguments for delete command")
	}
	id, err := strconv.Atoi(args[2])
	if err != nil {
		return err
	}
	err = tm.DeleteTask(id)
	if err != nil {
		return err
	}
	fmt.Printf("Task #%d deleted successfully\n", id)
	return nil
}

// Checks arguments changes task status
func handleMark(args []string, tm *task.TaskManager) error {
	if len(args) < 4 {
		return fmt.Errorf("no arguments for mark command")
	}
	if len(args) > 4 {
		return fmt.Errorf("too many arguments for mark command")
	}
	id, err := strconv.Atoi(args[2])
	if err != nil {
		return err
	}
	err = tm.MarkTask(id, args[3])
	if err != nil {
		return err
	}
	fmt.Printf("Task #%d marked as %s\n", id, args[3])
	return nil
}

// Prints a list of tasks with the ability to filter by status
func handleList(args []string, tm *task.TaskManager) error {
	if len(args) > 3 {
		return fmt.Errorf("too many arguments for list command")
	}
	sortBy := ""
	if len(args) == 3 {
		sortBy = args[2]
	}
	tasks, err := tm.ListTasks(sortBy)
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		fmt.Println("No saved tasks found.")
	}
	for id, task := range tasks {
		fmt.Printf("ID: %d\n", id)
		fmt.Printf("Description: %s\n", task.Description)
		fmt.Printf("Status: %s\n", task.Status)
		fmt.Printf("Created At: %s\n", task.CreatedAt.Format(time.RFC3339))
		fmt.Printf("Updated At: %s\n\n", task.UpdatedAt.Format(time.RFC3339))
	}
	return nil

}

// Prints help
func handleHelp() error {
	fmt.Print(helpText)
	return nil
}
