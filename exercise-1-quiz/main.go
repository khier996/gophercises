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
  "time"
  "math/rand"
)

func main() {
  filePath := flag.String("file", "answers.csv", "path to answers file")
  limit := flag.Int("limit", 30, "timer duration")
  flag.Parse()

  file, err := os.Open(*filePath)
  if err != nil {
    log.Fatal(err)
  } else {
    lr := io.MultiReader(file)
    csvReader := csv.NewReader(lr)
    runTest(csvReader, *limit)
  }

}

func runTest(csvReader *csv.Reader, limit int) {
  timer := time.NewTimer(time.Duration(limit) * time.Second)
  correctCount := 0
  inputReader := bufio.NewReader(os.Stdin)
  answerChan := make(chan bool)

  var records [][]string
  var csvErr error
  if records, csvErr = csvReader.ReadAll(); csvErr != nil {
    log.Fatal("error reading csv file", csvErr)
    return
  }

  records = shuffle(records)
  for _, record := range records {
    go testRunAnswer(record, inputReader, answerChan)
    select {
      case answer := <-answerChan:
        if answer { correctCount++ }
      case <-timer.C:
        fmt.Println("\nYou ran out of time")
        printResult(correctCount, len(records))
        return
    }
  }

  printResult(correctCount, len(records))
}

func printResult(correctCount, totalCount int) {
  fmt.Println("**********************************************")
  fmt.Println("You scored", correctCount, "out of", totalCount)
}

func testRunAnswer(record []string, inputReader *bufio.Reader, answerChan chan bool) {
  var userAnswer string
  var err error

  fmt.Print(record[0], " = ")
  if userAnswer, err = inputReader.ReadString('\n'); err != nil && err != io.EOF {
    log.Fatal("error reading answer", err)
  }

  userAnswer = strings.TrimSpace(userAnswer)
  realAnswer := strings.TrimSpace(record[1])
  answerChan <- userAnswer == realAnswer
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

func shuffle(records [][]string) [][]string {
  r := rand.New(rand.NewSource(time.Now().Unix()))
  ret := make([][]string, len(records))
  perm := r.Perm(len(records))
  for i, randIndex := range perm {
    ret[i] = records[randIndex]
  }
  return ret
}
