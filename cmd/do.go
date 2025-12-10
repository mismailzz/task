package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DoCmd = &cobra.Command{
	Use:   "do",
	Short: "mark a task as done",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Println("Please provide a task number.")
			return
		}
		taskNumber := args[0]
		fmt.Printf("Task %s marked as done.\n", taskNumber)
	},
}

func init() {
	RootCmd.AddCommand(DoCmd)
}
