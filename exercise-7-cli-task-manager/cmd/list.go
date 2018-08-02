package cmd

import (
  "github.com/spf13/cobra"
  "github.com/boltdb/bolt"
  "fmt"
)

var listCmd = &cobra.Command{
  Use:   "list",
  Run: func(cmd *cobra.Command, args []string) {
    listTasks()
  },
}

func listTasks() {
  db.View(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("Tasks"))
    counter := 0
    b.ForEach(func(k, v []byte) error {
      counter++
      fmt.Printf("%d. %s \n", counter, v)
      return nil
    })
    return nil
  })
}


