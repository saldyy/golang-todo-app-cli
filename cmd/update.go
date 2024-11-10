/*
Copyright © 2024 NAME HERE <png9981@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/saldyy/golang-todo-app-cli/repositories"
	"github.com/spf13/cobra"
)

var status string
var id int

func getValueFromStatusInput(status string) (repositories.TaskStatus, error) {
	switch status {
	case "pending":
		return repositories.Pending, nil
	case "done":
		return repositories.Done, nil
	case "cancel":
		return repositories.Cancel, nil
	default:
		return -1, fmt.Errorf("Invalid status, must be: pending, done or cancel.")
	}
}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if id == 0 {
			log.Fatal("Invalid Id")
		}

		statusValue, error := getValueFromStatusInput(status)
		if error != nil {
			log.Fatal(error)
		}

		repositories.UpdateTask(id, statusValue)

		fmt.Println("Update task successfully")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVar(&status, "status", "", "Status of the task")
	updateCmd.Flags().IntVar(&id, "id", 0, "")
}
