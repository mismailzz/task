package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
    Use:   "task",
    Short: "task is a CLI task manager",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("This is the root command")
    },
}