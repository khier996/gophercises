package main

import (
  "encoding/csv"
  "os"
  "log"
  "io"
  "bufio"
  "fmt"
  "strings"
  "flag"
)

func main() {
  filePath := flag.String("file", "answers.csv", "path to answers file")
  flag.Parse()

  file, err := os.Open(*filePath)
  if err != nil {
    log.Fatal(err)
  } else {
    lr := io.MultiReader(file)
    csvReader := csv.NewReader(lr)
    runTest(csvReader)
  }
}

func runTest(csvReader *csv.Reader) {
  correctCount := 0
  totalCount := 0

  inputReader := bufio.NewReader(os.Stdin)
  for {
    if record, err := csvReader.Read(); err != nil {
      if err == io.EOF {
        break
      } else {
        log.Fatal("error reading csv file", err)
      }
    } else {
      totalCount++
      if correct := runAnswer(record, inputReader); correct {
        correctCount++
      }
    }
  }
  fmt.Println("**********************************************")
  fmt.Println("You scored", correctCount, "out of", totalCount)
}

func runAnswer(record []string, inputReader *bufio.Reader) bool {
  var userAnswer string
  var err error

  fmt.Print(record[0], " = ")
  if userAnswer, err = inputReader.ReadString('\n'); err != nil && err != io.EOF {
    log.Fatal("error reading answer", err)
  }

  userAnswer = strings.TrimSpace(userAnswer)
  realAnswer := strings.TrimSpace(record[1])
  return userAnswer == realAnswer
}

