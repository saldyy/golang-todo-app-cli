/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/saldyy/golang-todo-app-cli/repositories"
)

func getTaskStatusText(status repositories.TaskStatus) string {
	switch status {
		case repositories.Done:
			return "Done" 
		case repositories.Pending:
			return "Pending" 
		case repositories.Cancel:
			return "Cancel" 
		default:
			return ""
	}
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := repositories.GetListTodoTasks()
		for _, task := range(tasks) {
			fmt.Printf("Id: %d, Task: %s, status: %s, last updated at: %s\n", task.Id, task.Title, getTaskStatusText(task.Status), task.UpdatedAt.UTC())
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
