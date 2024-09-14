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

## Usage

The application is run from the command line, accepting various commands to perform operations on tasks. Below are the available commands:

1. **Add a task**
    ```bash
    tasktracker add "Task description"
    ```

2. **Update a task**
    ```bash
    tasktracker update <task_id> "Updated description"
    ```

3. **Delete a task**
    ```bash
    tasktracker delete <task_id>
    ```

4. **Mark a task as in progress**
    ```bash
    tasktracker inprogress <task_id>
    ```

5. **Mark a task as done**
    ```bash
    tasktracker done <task_id>
    ```

6. **List all tasks**
    ```bash
    tasktracker list
    ```

7. **List all tasks that are done**
    ```bash
    tasktracker list --done
    ```

8. **List all tasks that are not done**
    ```bash
    tasktracker list --notdone
    ```

9. **List all tasks that are in progress**
    ```bash
    tasktracker list --inprogress
    ```

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

## Example

<!-- ```bash
$ tasktracker add "Finish homework"
Task added: [1] Finish homework

$ tasktracker list
1. [1] Finish homework (todo)

$ tasktracker inprogress 1
Task 1 marked as in-progress.

$ tasktracker list --inprogress
1. [1] Finish homework (in-progress)

$ tasktracker done 1
Task 1 marked as done.

$ tasktracker list --done
1. [1] Finish homework (done) -->

## Conclusion

Task Tracker CLI provides an efficient way to manage tasks through the command line. It enhances your command-line experience while also offering practice with filesystem operations, handling user inputs, and building a small-scale CLI application.

Happy coding!