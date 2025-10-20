package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/Yash-200/go-todo-cli/internal/models" // Import the new models package
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Open a connection to the SQLite database. The file will be created if it doesn't exist.
	DB, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to ensure the connection is alive.
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"name" TEXT,
		"status" TEXT DEFAULT 'pending',
		"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	statement, err := DB.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
}

// GetTasks retrieves tasks from the database with optional filtering and sorting
func GetTasks(filterName, filterStatus, sortBy, order string) ([]models.Task, error) {
	query := "SELECT id, name, status, created_at FROM tasks"
	var queryArgs []interface{}
	var whereClauses []string

	if filterName != "" {
		whereClauses = append(whereClauses, "name LIKE ?")
		queryArgs = append(queryArgs, "%"+filterName+"%")
	}

	if filterStatus != "" {
		whereClauses = append(whereClauses, "status = ?")
		queryArgs = append(queryArgs, filterStatus)
	}

	if len(whereClauses) > 0 {
		query += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	if sortBy != "" {
		allowedSortBy := map[string]bool{"id": true, "name": true, "status": true, "created_at": true}
		if !allowedSortBy[sortBy] {
			return nil, fmt.Errorf("invalid sort-by column")
		}
		order = strings.ToUpper(order)
		if order != "ASC" && order != "DESC" {
			return nil, fmt.Errorf("invalid order direction")
		}
		query += fmt.Sprintf(" ORDER BY %s %s", sortBy, order)
	}

	rows, err := DB.Query(query, queryArgs...)
	if err != nil {
		return nil, fmt.Errorf("failed to query tasks: %w", err)
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		// The fix is here: Scan directly into the time.Time field.
		if err := rows.Scan(&task.ID, &task.Name, &task.Status, &task.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func CreateTask(name string) (*models.Task, error) {
	statement, err := DB.Prepare("INSERT INTO tasks (name) VALUES (?)")
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(name)
	if err != nil {
		return nil, fmt.Errorf("failed to execute statement: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	var task models.Task
	// The fix is here: Scan directly into the time.Time field.
	err = DB.QueryRow("SELECT id, name, status, created_at FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Name, &task.Status, &task.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch new task: %w", err)
	}

	return &task, nil
}

func UpdateTaskStatus(id int) (int64, error) {
	statement, err := DB.Prepare("UPDATE tasks SET status = 'completed' WHERE id = ?")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}

	return result.RowsAffected()
}
func DeleteTask(id int) (int64, error) {
	statement, err := DB.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		return 0, fmt.Errorf("failed to execute statement: %w", err)
	}

	return result.RowsAffected()
}
