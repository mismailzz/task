package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("Please provide a task description.")
			return
		}
		taskDescription := strings.Join(args, " ")
		fmt.Printf("Task added: %s\n", taskDescription)
	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
