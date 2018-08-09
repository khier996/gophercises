package cmd

import (
  "github.com/spf13/cobra"
  "github.com/boltdb/bolt"
  "strings"
  "strconv"
)

var addCmd = &cobra.Command{
  Use:   "add",
  Long: "This command adds a new task to the task list",
  Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    task := strings.Join(args, " ")
    addTask(task)
  },
}

func addTask(task string) {
    db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("Tasks"))

    id, _ := b.NextSequence()
    idStr := strconv.Itoa(int(id))

    err := b.Put([]byte(idStr), []byte(task))
    return err
  })
}



