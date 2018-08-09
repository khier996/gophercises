package cmd

import (
  "github.com/spf13/cobra"
  "fmt"
  "log"
  "github.com/boltdb/bolt"
  "github.com/mitchellh/go-homedir"
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

  rootCmd.AddCommand(addCmd, listCmd, doCmd)
  rootCmd.Execute()
}

func openDB() {
  hdir, err := homedir.Dir()
  if err != nil {
    fmt.Println("Error finding home directory")
    return
  }
  if hdir, err = homedir.Expand(hdir); err != nil {
    log.Fatal("Error expanding home directory")
  }

  db, dbErr = bolt.Open(hdir + "/tasks.db", 0600, nil)
  if dbErr != nil {
    log.Fatal(dbErr)
  }

  db.Update(func(tx *bolt.Tx) error {
    _, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
    if err != nil {
      log.Fatal("Could not create bucket Tasks")
    }
    return err
  })
}
