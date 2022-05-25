package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"quiz-game/service"
	"time"
)

func main() {

	csvFileName := flag.String("csv", "problems.csv", "csv file")
	quizTime := flag.Int("time", 5, "time for the whole quiz")
	flag.Parse()

	fd, err := os.Open(*csvFileName)
	defer fd.Close()
	if err != nil {
		fmt.Printf(" can't open file %s", err)
		os.Exit(-1)
	}

	reader := csv.NewReader(fd)
	reader.FieldsPerRecord = 2
	all, err := reader.ReadAll()
	if err != nil {
		os.Exit(0)
	}
	parser := service.NewDataParser()
	problems := parser.Parse(all)

	counterPositive := 0
	counterNegative := 0
	answerCh := make(chan string)
	timer := time.NewTimer(time.Duration(*quizTime) * time.Second)

	for _, value := range problems {
		fmt.Printf("%s = ", value.Question)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case answer := <-answerCh:
			if answer != value.Answer {
				counterNegative++
			} else {
				counterPositive++
			}
		case <-timer.C:
			fmt.Println("\ntime is over!\n", counterPositive, " Correct\n", counterNegative, " Incorrect\n ")
			return

		}

	}
	fmt.Printf("correct - %d,incorrect - %d", counterPositive, counterNegative)
}
