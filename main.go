package main

import (
	"log"

	"github.com/Yash-200/go-todo-cli/cmd"
	"github.com/Yash-200/go-todo-cli/internal/database"
)

func main() {
	database.InitDB()
	defer func() {
		if err := database.DB.Close(); err != nil {
			log.Fatalf("Error closing the database: %v", err)
		}
	}()

	cmd.Execute()
}
