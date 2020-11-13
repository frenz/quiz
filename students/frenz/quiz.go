package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problemData struct {
	Question string
	Answer   string
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func readCsvLines(filename string) [][]string {
	csvFile, err := os.Open(filename)
	printError(err)
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	printError(err)
	return csvLines
}
func parseLinesToProblems(csvLines [][]string) []problemData {
	var problems []problemData
	for _, line := range csvLines {
		problem := problemData{
			Question: line[0],
			Answer:   line[1],
		}
		problems = append(problems, problem)
	}
	return problems
}
func readAnswer(answerCh chan string) {
	answer := ""
	fmt.Scanf("%s", &answer)
	answerCh <- answer
}

func askUser(problems []problemData, timer *time.Timer) int {
	correct := 0
	for _, problem := range problems {
		fmt.Printf("%s ?\n", problem.Question)
		answerCh := make(chan string)
		go readAnswer(answerCh)
		select {
		case <-timer.C:
			return correct
		case answer := <-answerCh:
			if answer == problem.Answer {
				correct++
			}
		}
	}
	return correct
}

func main() {
	filename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	csvLines := readCsvLines(*filename)
	problems := parseLinesToProblems(csvLines)
	newtimer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := askUser(problems, newtimer)
	fmt.Printf("%d\n", correct)
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}
