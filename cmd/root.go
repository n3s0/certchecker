package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "certchecker",
  Short: "Cert Checker checks that TLS certificates are valid.",
  Long: `Check that TLS certificates are valid.`,
  Run: func(cmd *cobra.Command, args []string) {
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
