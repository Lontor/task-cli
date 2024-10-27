# Task CLI
Task CLI is a simple task manager written in Go. It allows you to add, update, delete, and list tasks via the terminal. All tasks are saved in JSON format.

## Build
1. Install Go (minimum version 1.16 required).
2. Clone this repository:
   ```bash
   git clone https://github.com/Lontor/task-cli.git
   ```
3. Navigate to the project directory:
   ```bash
   cd task-cli
   ```
4. Build the project:
   ```bash
   go build -o task-cli
   ```

## Usage
- Display help text:
  ```bash
  ./task-cli help
  ```
- Add a new task:
  ```bash
  ./task-cli add "my task"
  ```
- List all tasks:
  ```bash
  ./task-cli list
  ```
- Mark a task as done (other statuses: todo, in-progress):
  ```bash
  ./task-cli mark 0 done
  ```
- Update the description of a task:
  ```bash
  ./task-cli update 0 "updated task description"
  ```
- Delete a task:
  ```bash
  ./task-cli delete 0
  ```
