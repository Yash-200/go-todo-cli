# Go To-Do List CLI & API

This is a simple yet powerful to-do list application built with Go. It provides both a command-line interface (CLI) for managing tasks directly in your terminal and a RESTful API for programmatic access.

This project was built with Go version `go1.24.9 linux/amd64`.

## Features

*   **Full CRUD Operations:** Add, list, complete, and delete tasks.
*   **Persistent Storage:** Tasks are stored in a local SQLite database file (`tasks.db`).
*   **Advanced Listing:**
    *   Filter tasks by name (partial match).
    *   Filter tasks by status (`pending`, `completed`).
    *   Sort tasks by `id`, `name`, `status`, or `created_at`.
    *   Specify sort order (`asc`, `desc`).
*   **Dual Interface:**
    *   **CLI:** A robust command-line interface for terminal-based management.
    *   **REST API:** A full-featured API for interacting with the to-do list via HTTP, testable with tools like Postman or curl.

## Prerequisites

*   Go version 1.24.9 or newer.

## Installation & Building

1.  **Clone the repository (or create the files as per the guide):**
    ```bash
    # If you have it in a git repo
    git clone https://github.com/your-username/go-todo-cli.git
    cd go-todo-cli
    ```

2.  **Tidy dependencies:**
    This will download all the necessary libraries (Cobra, Chi, SQLite driver).
    ```bash
    go mod tidy
    ```

3.  **Build the executable:**
    This command compiles the source code into a single executable file named `go-todo-cli`.
    ```bash
    go build
    ```

## CLI Usage

All commands are run from the executable (`./go-todo-cli`).

#### `add`
Adds a new task. The task name must be enclosed in quotes if it contains spaces.

```bash
./go-todo-cli add "My new task"
# Output: Added task: "My new task"
```

#### `list`
Lists all tasks. This command has several flags for filtering and sorting.

```bash
# List all tasks
./go-todo-cli list

# List all pending tasks
./go-todo-cli list --status pending

# List tasks containing "learn", sorted by name
./go-todo-cli list --filter-name "learn" --sort-by name

# List all tasks sorted by ID in descending order
./go-todo-cli list --sort-by id --order desc
```

#### `complete`
Marks a task as completed using its ID.

```bash
./go-todo-cli complete 2
# Output: Task #2 marked as completed.
```

#### `delete`
Deletes a task using its ID.

```bash
./go-todo-cli delete 3
# Output: Deleted task #3.
```

#### `serve`
Starts the REST API server.

```bash
./go-todo-cli serve
# Output: Starting server on :3000
```

## API Usage

Start the server by running `./go-todo-cli serve`. The API will be available at `http://localhost:3000`.

---

### List Tasks
Retrieves a list of tasks. Supports filtering and sorting via query parameters.

*   **Method:** `GET`
*   **Endpoint:** `/tasks`
*   **Query Parameters:**
    *   `filter_name` (string): Filters tasks by a partial name match.
    *   `status` (string): Filters tasks by status (`pending` or `completed`).
    *   `sort_by` (string): Column to sort by (`id`, `name`, `status`, `created_at`).
    *   `order` (string): Sort order (`asc` or `desc`).
*   **Example:** `GET http://localhost:3000/tasks?status=pending&sort_by=id&order=desc`

---

### Create Task
Adds a new task.

*   **Method:** `POST`
*   **Endpoint:** `/tasks`
*   **Body (JSON):**
    ```json
    {
        "name": "Create API documentation"
    }
    ```
*   **Success Response:** `201 Created` with the new task object.

---

### Complete Task
Marks a specific task as completed.

*   **Method:** `PUT`
*   **Endpoint:** `/tasks/{id}/complete`
*   **Example:** `PUT http://localhost:3000/tasks/2/complete`
*   **Success Response:** `204 No Content`.

---

### Delete Task
Deletes a specific task.

*   **Method:** `DELETE`
*   **Endpoint:** `/tasks/{id}`
*   **Example:** `DELETE http://localhost:3000/tasks/2`
*   **Success Response:** `204 No Content`.