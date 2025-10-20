package cmd

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	"github.com/Yash-200/go-todo-cli/internal/database"
	"github.com/spf13/cobra"
)

type Task struct {
	ID        int
	Name      string
	Status    string
	CreatedAt time.Time
}

var sortBy string
var order string
var filterName string
var filterStatus string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := database.GetTasks(filterName, filterStatus, sortBy, order)
		if err != nil {
			log.Fatalf("Error getting tasks: %v", err)
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.FilterHTML)
		fmt.Fprintln(w, "ID\tTASK NAME\tSTATUS\tCREATED AT")
		fmt.Fprintln(w, "--\t---------\t------\t----------")
		for _, task := range tasks {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", task.ID, task.Name, task.Status, task.CreatedAt.Format("2006-01-02 15:04"))
		}
		w.Flush()

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVar(&sortBy, "sort-by", "", "Sort tasks by column (id, name, status, created_at)")
	listCmd.Flags().StringVar(&order, "order", "asc", "Order of sorting (asc, desc)")
	listCmd.Flags().StringVar(&filterName, "filter-name", "", "Filter tasks by name (case-insensitive, partial match)")
	listCmd.Flags().StringVar(&filterStatus, "status", "", "Filter tasks by status (pending, completed)")
}
