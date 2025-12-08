package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
    Use:   "add",
    Short: "add a new task",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Adding a new task...")
    },
}

func init() {
    RootCmd.AddCommand(AddCmd)
}
