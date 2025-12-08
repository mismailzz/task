package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var DoCmd = &cobra.Command{
    Use:   "do",
    Short: "mark a task as done",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Marking a task as done...")
    },
}

func init() {
    RootCmd.AddCommand(DoCmd)
}
