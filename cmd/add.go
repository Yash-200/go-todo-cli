package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/Yash-200/go-todo-cli/internal/database"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Adds a new task to the to-do list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")

		statement, err := database.DB.Prepare("INSERT INTO tasks (name) VALUES (?)")
		if err != nil {
			log.Fatalf("Failed to prepare statement: %v", err)
		}
		defer statement.Close()

		_, err = statement.Exec(taskName)
		if err != nil {
			log.Fatalf("Failed to execute statement: %v", err)
		}

		fmt.Printf("Added task: \"%s\"\n", taskName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
