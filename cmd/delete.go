package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Yash-200/go-todo-cli/internal/database"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Deletes a task from the to-do list",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalf("Invalid task ID: %v. Please provide a number.", err)
		}

		rowsAffected, err := database.DeleteTask(taskID)
		if err != nil {
			log.Fatalf("Failed to delete task: %v", err)
		}

		if rowsAffected == 0 {
			fmt.Printf("No task found with ID: %d\n", taskID)
			return
		}

		fmt.Printf("Deleted task #%d.\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
