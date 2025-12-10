package cmd

import (
	"fmt"
	"strings"

	"github.com/mismailzz/task/internal"
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
		db, _ := internal.OpenDB("tasks.db")
		taskText := strings.Join(args, " ")
		err := internal.AddTask(db, taskText)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Added:", taskText)

	},
}

func init() {
	RootCmd.AddCommand(AddCmd)
}
