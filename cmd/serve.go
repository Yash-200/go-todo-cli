package cmd

import (
	"log"
	"net/http"

	"github.com/Yash-200/go-todo-cli/internal/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the API server",
	Run: func(cmd *cobra.Command, args []string) {
		r := chi.NewRouter()

		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)

		r.Get("/tasks", api.ListTasksHandler)
		r.Post("/tasks", api.CreateTaskHandler)
		r.Put("/tasks/{id}/complete", api.UpdateTaskStatusHandler)
		r.Delete("/tasks/{id}", api.DeleteTaskHandler)

		log.Println("Starting server on :3000")
		if err := http.ListenAndServe(":3000", r); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
