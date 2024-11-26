package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print version for certchecker",
  Long:  `Print version for certchecker.`,
  Run: func(cmd *cobra.Command, args []string) {
    const appVersion = "0.1.0"
    fmt.Printf("certmon v%s", appVersion)
  },
}
