package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var HelloCmd = &cobra.Command{
    Use:   "hello",
    Short: "say hello",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("hi")
    },
}

func init() {
    RootCmd.AddCommand(HelloCmd)
}
