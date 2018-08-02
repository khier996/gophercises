package cmd

import (
  "github.com/spf13/cobra"
  "fmt"
  "log"
  "github.com/boltdb/bolt"
)

var rootCmd = &cobra.Command{
  Use:   "task",
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("success main")
  },
}

var db *bolt.DB
var dbErr error

func Execute() {
  openDB()
  defer db.Close()

  rootCmd.AddCommand(addCmd, listCmd)
  rootCmd.Execute()
}

func openDB() {
  db, dbErr = bolt.Open("tasks.db", 0600, nil)
  if dbErr != nil {
    log.Fatal(dbErr)
  }

  db.Update(func(tx *bolt.Tx) error {
    _, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
    return err
  })
}
