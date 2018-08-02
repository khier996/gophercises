package cmd

import (
  "github.com/spf13/cobra"
)

var doCmd = &cobra.Command({
  Use:   "list",
  Run: func(cmd *cobra.Command, args []string) {
    doTask()
  },
})

func doTask()
