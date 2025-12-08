package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
    Use:   "list",
    Short: "list all tasks",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Listing all tasks...")
    },
}

func init() {
    RootCmd.AddCommand(ListCmd)
}
