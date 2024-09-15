# Task Tracker CLI

Task Tracker is a simple command line interface (CLI) tool used to track and manage tasks. It allows you to add, update, delete, and view tasks, making it easier to manage your to-do list from the terminal.

## Features

- **Add, Update, and Delete tasks**: Easily manage your tasks with simple commands.
- **Task Status Management**: Mark tasks as `todo`, `in-progress`, or `done`.
- **Task Listings**: View all tasks or filter tasks based on their status.
- **Task Persistence**: Tasks are stored in a JSON file, allowing for persistent data across sessions.

## Task Properties

Each task has the following properties:
- `id`: A unique identifier for the task.
- `description`: A short description of the task.
- `status`: The current status of the task (`todo`, `in-progress`, or `done`).
- `createdAt`: The date and time when the task was created.
- `updatedAt`: The date and time when the task was last updated.

## File Descriptions

# main.go

Contains the main function that initializes the application, loads tasks, processes commands, and saves tasks.

# command.go

Defines the command-line flags and the Execute method to handle user commands.

# tracker.go

Contains the Tracker struct and methods to manage tasks, including adding, editing, deleting, toggling status, and printing tasks.

# library.go

Handles saving and loading task data to and from a JSON file.

## Requirements

- The application accepts user actions and inputs as positional arguments.
- Tasks are stored in a JSON file in the current directory.
- The JSON file is created automatically if it does not exist.
- No external libraries or frameworks are required for this project.

### Task Constraints
- Use the native filesystem module of your programming language to read and write tasks to the JSON file.
- Handle errors and edge cases gracefully, such as missing or incorrect inputs.

## Getting Started

1. **Set Up Your Development Environment**
    - Choose your preferred programming language (e.g., Python, JavaScript).
    - Install a code editor or IDE (e.g., VSCode, PyCharm).

2. **Project Initialization**
    - Create a new project directory for your Task Tracker CLI.
    - Initialize version control (e.g., Git) to track changes.

3. **Implementing Features**
    - Begin by setting up a basic CLI to accept user inputs.
    - Implement features step by step (e.g., task addition, listing, updating, etc.), ensuring each is fully functional before moving to the next.

4. **Testing and Debugging**
    - Test each feature to ensure correctness.
    - Verify tasks are stored correctly in the JSON file.
    - Address any bugs or issues encountered during development.

5. **Finalizing the Project**
    - Complete the implementation of all features.
    - Clean up your code and add necessary comments.
    - Write documentation (like this README) to guide users on how to use the Task Tracker CLI.

## Commands

# Add a New Task

Add a new task with a description.
<!-- go run main.go --add "Your task description" -->

# Edit a Task Description

Edit the description of an existing task by ID.
<!-- go run main.go --edit id:new_description -->
Example:
<!-- go run main.go --edit 1:Updated task description -->

# Set a Task Status

Set the status of a task by ID. Allowed statuses are todo, in-progress, and done.
<!-- go run main.go --status id:new_status -->
Example:
<!-- go run main.go --status 1:done -->

# Delete a Task

Delete a task by ID.
<!-- go run main.go --del id -->
Example:
<!-- go run main.go --del 1 -->

# Toggle a Task Status

Toggle a task's status through its lifecycle: todo -> in-progress -> done -> todo.
<!-- go run main.go --toggle id -->
Example:
<!-- go run main.go --toggle 1 -->

# List All Tasks

List all tasks with their details.
<!-- go run main.go --list -->

## Conclusion

Task Tracker CLI provides an efficient way to manage tasks through the command line. It enhances your command-line experience while also offering practice with filesystem operations, handling user inputs, and building a small-scale CLI application.

Happy coding!