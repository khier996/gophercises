package cmd

import (
  "github.com/spf13/cobra"
  "strconv"
  "github.com/boltdb/bolt"
  "log"
  "fmt"
)

var doCmd = &cobra.Command{
  Use: "do",
  Long: "This command will delete a task from the task list",
  Args: cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    taskNum, err := strconv.Atoi(args[0])
    if err != nil {
      log.Fatal("Could not convert provided task number into number")
    }
    doTask(taskNum)
  },
}

func doTask(taskNum int) {
  db.Update(func(tx *bolt.Tx) error {
    b := tx.Bucket([]byte("Tasks"))
    c := b.Cursor()

    count := 0
    for k, v := c.First(); k != nil; k, v = c.Next() {
      count++
      if taskNum == count {
        fmt.Printf("You have completed the \"%s\" task.", v)
        b.Delete(k)
        return nil
      }
    }

    fmt.Println("There is no task with the given number")
    return nil
  })
}
