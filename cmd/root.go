package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "Daneshkar server",
    Short: "A simple CLI web server",
    Long:  "This is a simple CLI tool using Cobra to run a web server.",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
