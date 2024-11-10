/*
Copyright © 2024 NAME HERE <png9981@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/saldyy/golang-todo-app-cli/repositories"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Entering task: ")
		scanner.Scan()
		title := scanner.Text()
		repositories.CreateTask(title)

		fmt.Printf("Create task successfully.")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
